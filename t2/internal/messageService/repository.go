package messageService

import "gorm.io/gorm"

type MessageRepository interface {
	// CreateMessage - Передаем в функцию message типа Message из orm.go
	// возвращаем созданный Message и ошибку
	CreateMessage(message Message) (Message, error)
	// GetAllMessages - Возвращаем массив из всех писем в БД и ошибку
	GetAllMessages() ([]Message, error)
	// UpdateMessageByID - Передаем id и Message, возвращаем обновленный Message
	// и ошибку
	UpdateMessageByID(Message) (Message, error)
	// DeleteMessageByID - Передаем id для удаления, возвращаем только ошибку
	DeleteMessageByID(message Message) error
}

type messageRepository struct {
	db *gorm.DB
}

func (r *messageRepository) UpdateMessageByID(message Message) (Message, error) {
	result := r.db.Model(&message).Updates(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

func (r *messageRepository) DeleteMessageByID(message Message) error {
	result := r.db.Delete(&message)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

// (r *messageRepository) привязывает данную функцию к нашему репозиторию
func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}
