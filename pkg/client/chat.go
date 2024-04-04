package auth

import (
	"api/pkg/client/interfaces"
	"api/pkg/config"
	"api/pkg/models"
	"api/pkg/pb/chat"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ChatClient struct{
	Server chat.ChatServiceClient
}

func InitChatClient(c *config.Config) (chat.ChatServiceClient,error) {
	cc,err:=grpc.Dial(c.ChatService,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		return nil,err
	}
	return chat.NewChatServiceClient(cc),nil
}

func NewChatServiceClient(server chat.ChatServiceClient) interfaces.ChatClient{
	return &ChatClient{
		Server: server,
	}
}

func (a *ChatClient) ChatHistory(ctx context.Context,body models.ChathistoryRequest) (*chat.ChatHistoryResponse,error) {
	res,err:=a.Server.ChatHistory(ctx,&chat.ChatHistoryRequest{
		Username1: body.Username1,
		Username2: body.Username2,
		FromTs: body.Fromts,
		ToTs: body.Tots,
	})
	if err != nil{
		return nil,err
	}
	return res,nil
}

func (a *ChatClient) ContactList(ctx context.Context,username string) (*chat.ContactListResponse,error) {
	res,err:=a.Server.ContactList(ctx,&chat.ContactListRequest{
		Username: username,
	})
	if err != nil{
		return nil,err
	}
	return res,nil
}

func(a *ChatClient) VerifyContact(ctx context.Context,username string) (*chat.VerifyContactResponse,error) {
	res,err:=a.Server.VerifyContact(ctx,&chat.VerifyContactRequest{
		Username: username,
	})
	if err != nil{
		return nil,err
	}
	return res,nil
}
