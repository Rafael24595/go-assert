package require

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

func exported(t *testing.T, dir string) map[string]bool {
	t.Helper()

	files, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		t.Fatal(err)
	}

	names := map[string]bool{}

	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}

		file, err := parser.ParseFile(token.NewFileSet(), f, nil, parser.ParseComments)
		if err != nil {
			t.Fatal(err)
		}

		for _, d := range file.Decls {
			fn, ok := d.(*ast.FuncDecl)
			if !ok || fn.Recv != nil || !fn.Name.IsExported() {
				continue
			}

			if fn.Doc != nil && strings.Contains(fn.Doc.Text(), "Deprecated:") {
				continue
			}

			names[fn.Name.Name] = true
		}
	}

	return names
}

func diff(a, b map[string]bool) []string {
	var out []string
	for name := range a {
		if !b[name] {
			out = append(out, name)
		}
	}

	sort.Strings(out)
	return out
}

func TestSameAssertionsAsAssert(t *testing.T) {
	assertFns := exported(t, "../assert")
	requireFns := exported(t, "../require")

	if missing := diff(assertFns, requireFns); len(missing) > 0 {
		t.Errorf("missing in require: %s", strings.Join(missing, ", "))
	}

	if extra := diff(requireFns, assertFns); len(extra) > 0 {
		t.Errorf("missing in assert: %s", strings.Join(extra, ", "))
	}
}
