package services

var (
	sumimasenService SumimasenService
)

func init() {
	sumimasenService = newSumimasenService()
}

func newSumimasenService() SumimasenService {
	return &SumimasenServiceImpl{}
}

func GetSumimasenService() SumimasenService {
	return sumimasenService
}