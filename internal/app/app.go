package app

import (
	"context"
	"flag"

	"github.com/valek177/chat-client/grpc/pkg/chat_v1"
	"github.com/valek177/chat-client/internal/config"
)

var configPath = "/home/valek/microservices_course/chat-client/local.env"

var app *App

// App contains application object
type App struct {
	serviceProvider *serviceProvider
}

// NewApp creates new App object
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	flag.Parse()

	err := config.Load(configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func ConnectChat(ctx context.Context, chatname, username string) (chat_v1.ChatV1_ConnectChatClient,
	error,
) {
	srv, err := app.serviceProvider.CommandService(ctx)
	if err != nil {
		return nil, err
	}
	return srv.ConnectChat(ctx, chatname, username)
}

// func CreateChat(ctx context.Context, chatname, username string) (chat_v1.ChatV1_ConnectChatClient,
// 	error,
// ) {
// 	srv, err := app.serviceProvider.CommandService(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return srv.CreateChat(ctx, chatname, username)
// }

// func DeleteChat(ctx context.Context, chatname, username string) (chat_v1.ChatV1_ConnectChatClient,
// 	error,
// ) {
// 	srv, err := app.serviceProvider.CommandService(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return srv.DeleteChat(ctx, chatname, username)
// }
