package service

type ISay interface {
	GetHello() string
	GetMonming() string
}

type Say struct{}

func NewSay() ISay {
	return &Say{}
}

func (s *Say) GetHello() string {
	return "hi"
}

func (s *Say) GetMonming() string {
	return "Good monming"
}
