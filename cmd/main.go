package main

import (
	"fmt"
	"github.com/Anwarjondev/blog-website-clone/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
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
	_ = db
}
