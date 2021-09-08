package service

type Service struct {
	ConvServ *ConversationService
}

func NewService(token string) *Service {
	return &Service{
		ConvServ: NewConversationService(token),
	}
}
