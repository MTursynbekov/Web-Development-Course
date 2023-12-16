package main

import (
	"log"
	"twitter/internal/app"
	"twitter/internal/service"
	"twitter/internal/store"
	"twitter/pkg/config"
	"twitter/pkg/db"
)

func init() {
	config.ParseEnv()
}

func main() {
	db, err := db.Connect(config.Get().DB)
	if err != nil {
		log.Fatalf("failed to connect db: %s", err)
	}

	s := store.NewStore(db)
	userService := service.NewUserService(s)
	migrationService := service.NewMigrationService(s)
	postService := service.NewPostsService(s)

	server := app.NewServer(userService, migrationService, postService)
	if err = server.Start(config.Get().Port); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
