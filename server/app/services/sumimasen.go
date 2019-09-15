package services

import "golang.org/x/xerrors"

type SumimasenService interface {
	Sumimasen() (*Sumimasen, error)
}

type SumimasenServiceImpl struct {
}

func (svc *SumimasenServiceImpl) Sumimasen() (*Sumimasen, error) {
	return newSumimasen("すみません。。。")
}

type Sumimasen struct {
	Message string
}

func newSumimasen(message string) (*Sumimasen, error) {
	if len(message) > 100 {
		return nil, xerrors.New("too long message")
	}
	return &Sumimasen{Message: message}, nil
}
