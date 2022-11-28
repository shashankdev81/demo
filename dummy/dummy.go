package dummy

type Stringer interface {
	String() string
}

type StringerImpl struct {
}

func (s *StringerImpl) String() string {
	return ""
}
