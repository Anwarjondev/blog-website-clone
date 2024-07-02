package server

import (
	v1 "github.com/Anwarjondev/blog-website-clone/server/v1"
	"github.com/Anwarjondev/blog-website-clone/storage"
	"github.com/gin-gonic/gin"
)

type Options struct {
	Strg storage.StorageI
}

func NewServer(opt *Options) *gin.Engine {
	router := gin.New()
	handler := v1.New(&v1.HandlerV1{
		Strg: opt.Strg,
	})
	router.POST("/v1/users", handler.CreateUser)
	router.GET("/v1/users/:id", handler.GetUser)
	router.DELETE("/v1/users/:id", handler.DeleteUser)
	router.PUT("/v1/users/:id", handler.UpdateUser)
	return router
}
