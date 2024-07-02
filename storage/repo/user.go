package repo

import (
	"context"
	"time"
)

type UserStorageI interface {
	Create(ctx context.Context, req *User) (*User, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, req *Updateuser) error
}

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
}
type Updateuser struct {
	ID        string
	FirstName string
	LastName  string
}
