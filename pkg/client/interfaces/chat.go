package interfaces

import (
	"api/pkg/models"
	"api/pkg/pb/chat"
	"context"
)

type ChatClient interface {
	ChatHistory(context.Context, models.ChathistoryRequest) (*chat.ChatHistoryResponse, error)
	ContactList(context.Context, string) (*chat.ContactListResponse, error)
	VerifyContact( context.Context, string) (*chat.VerifyContactResponse,error)
}
