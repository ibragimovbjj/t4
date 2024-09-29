package handlers

import (
	"context"
	messageService2 "t1/internal/messageService"
	"t1/internal/web/message"
)

type Handler struct {
	Service *messageService2.MessageService
}

func (h *Handler) DeleteMessages(ctx context.Context, request message.DeleteMessagesRequestObject) (message.DeleteMessagesResponseObject, error) {
	messageRequest := request.Body
	messageToDelete := messageService2.Message{Text: *messageRequest.Message}
	err := h.Service.DeleteMessageByID(messageToDelete)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *Handler) PatchMessages(ctx context.Context, request message.PatchMessagesRequestObject) (message.PatchMessagesResponseObject, error) {

	messageRequest := request.Body
	messageToUpdate := messageService2.Message{Text: *messageRequest.Message}
	updatedMessage, err := h.Service.UpdateMessageByID(messageToUpdate)
	if err != nil {
		return nil, err
	}

	response := message.PatchMessages200JSONResponse{
		Message: &updatedMessage.Text,
	}
	return response, nil
}

func (h *Handler) GetMessages(ctx context.Context, request message.GetMessagesRequestObject) (message.GetMessagesResponseObject, error) {

	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	//Просто возвращаем респонс!
	response := message.GetMessages200JSONResponse{}
	for _, msg := range allMessages {
		msgs := message.Message{
			Message: &msg.Text,
		}
		response = append(response, msgs)
	}
	return response, nil
}

func (h *Handler) PostMessages(ctx context.Context, request message.PostMessagesRequestObject) (message.PostMessagesResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	messageRequest := request.Body
	// Обращаемся к сервису и создаем сообщение
	messageToCreate := messageService2.Message{Text: *messageRequest.Message}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := message.PostMessages201JSONResponse{
		Message: &createdMessage.Text,
	}
	// Просто возвращаем респонс!
	return response, nil
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *messageService2.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}
