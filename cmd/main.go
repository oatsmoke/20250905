package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/oatsmoke/20250905/docs"
	"github.com/oatsmoke/20250905/internal/handler"
	"github.com/oatsmoke/20250905/internal/lib/env"
	"github.com/oatsmoke/20250905/internal/lib/http_server"
	"github.com/oatsmoke/20250905/internal/lib/logger"
	"github.com/oatsmoke/20250905/internal/lib/postgres_db"
	"github.com/oatsmoke/20250905/internal/repository"
	"github.com/oatsmoke/20250905/internal/service"
)

// @title Users online subscriptions
// @version 1.0
func main() {
	logger.New()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	postgresDB := postgres_db.Connect(ctx, env.GetPostgresDsn())
	defer postgresDB.Close()

	newR := repository.New(postgresDB)
	newS := service.New(newR)
	newH := handler.New(newS)

	httpPort := env.GetHttpPort()
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost%s", httpPort)
	httpS := http_server.New(httpPort, newH)
	httpS.Run()
	defer httpS.Stop(ctx)

	<-ctx.Done()
}
