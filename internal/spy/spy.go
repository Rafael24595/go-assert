package spy

import "fmt"

type T struct {
	HasFailed bool
	HasFatal  bool
	IsHelper  bool
	Message   string
}

func (s *T) Helper() {
	s.IsHelper = true
}

func (s *T) Errorf(format string, args ...any) {
	s.HasFailed = true
	s.Message = fmt.Sprintf(format, args...)
}

func (s *T) Fatalf(format string, args ...any) {
	s.HasFatal = true
	s.Message = fmt.Sprintf(format, args...)
}

type Report struct {
	Calls   int
	Message string
}

func (r Report) Called() bool {
	return r.Calls > 0
}

func (r *Report) Func() func(format string, args ...any) {
	return func(format string, args ...any) {
		r.Calls++
		r.Message = fmt.Sprintf(format, args...)
	}
}
