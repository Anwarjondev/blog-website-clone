package storage

import (
	"github.com/Anwarjondev/blog-website-clone/storage/postgres"
	"github.com/Anwarjondev/blog-website-clone/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserStorageI
}

type storagePg struct {
	userRepo repo.UserStorageI
}

func NewStorage(db *sqlx.DB) StorageI {
	return &storagePg{
		userRepo: postgres.NewUserStorage(db),
	}
}
func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}
