package messageService

type MessageService struct {
	repo MessageRepository
}

func NewService(repo MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(message Message) (Message, error) {
	return s.repo.CreateMessage(message)
}

func (s *MessageService) GetAllMessages() ([]Message, error) {
	return s.repo.GetAllMessages()
}

func (s *MessageService) DeleteMessageByID(m Message) error {
	return s.repo.DeleteMessageByID(m)

}

func (s *MessageService) UpdateMessageByID(message Message) (Message, error) {

	return s.repo.UpdateMessageByID(message)
}
