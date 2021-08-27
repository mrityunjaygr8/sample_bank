package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/mrityunjaygr8/sample_bank/db/sqlc"
	"github.com/mrityunjaygr8/sample_bank/utils"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type createUserResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"fullName"`
	CreatedAt         time.Time `json:"createdAt"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
	Email             string    `json:"email"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, erroresponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, erroresponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pq, ok := err.(*pq.Error); ok {
			fmt.Println(pq.Code.Name())
			switch pq.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, erroresponse(err))
				return

			}
		}
		ctx.JSON(http.StatusInternalServerError, erroresponse(err))
		return
	}

	response := createUserResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		CreatedAt:         user.CreatedAt,
		PasswordChangedAt: user.PasswordChangedAt,
		Email:             user.Email,
	}
	ctx.JSON(http.StatusOK, response)
}
