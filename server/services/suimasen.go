package services

type SumimasenService interface {
	Sumimasen() *SumimasenResponse
}

type SumimasenServiceImpl struct{}

type SumimasenResponse struct {
	Message string
}

func (svc *SumimasenServiceImpl) Sumimasen() *SumimasenResponse {
	return &SumimasenResponse{
		Message: "sorry...",
	}
}
