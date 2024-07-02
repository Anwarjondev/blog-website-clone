package main

import (
	"fmt"
	"log"

	"github.com/Anwarjondev/blog-website-clone/config"
	"github.com/Anwarjondev/blog-website-clone/server"
	"github.com/Anwarjondev/blog-website-clone/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load(".")
	psqUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	db, err := sqlx.Connect("postgres", psqUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	}

	log.Println("Successfully connected to the database!")
	strg := storage.NewStorage(db)
	router := server.NewServer(&server.Options{
		Strg: strg,
	})
	if err = router.Run(cfg.Port); err != nil {
		log.Fatalf("Failed to connect to server[%v]: %v", cfg.Port, err)
	}
}
