package v1

import (
	"database/sql"
	"errors"
	"github.com/Anwarjondev/blog-website-clone/server/model"
	"github.com/Anwarjondev/blog-website-clone/storage/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func (h *handlerV1) CreateUser(ctx *gin.Context) {
	var req model.CreateUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "We have internal error",
		})
		return
	}
	user, err := h.strg.User().Create(ctx, &repo.User{
		ID:        id.String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	})
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "We have internal error",
		})
		return
	}
	ctx.JSON(http.StatusCreated, model.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	})
}
func (h *handlerV1) UpdateUser(ctx *gin.Context) {
	var req model.UpdateUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	id := ctx.Param("id")

	err := h.strg.User().Update(ctx, &repo.Updateuser{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "We have internal error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated",
	})
}

func (h *handlerV1) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := h.strg.User().Get(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "We have internal error",
		})
		return
	}
	ctx.JSON(http.StatusOK, model.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	})
}

func (h *handlerV1) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.strg.User().Delete(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "We have internal error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})

}
