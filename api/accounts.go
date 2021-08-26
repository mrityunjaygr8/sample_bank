package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/mrityunjaygr8/sample_bank/db/sqlc"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD CAD"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, erroresponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, erroresponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, erroresponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, erroresponse(err))
			return

		}
		ctx.JSON(http.StatusInternalServerError, erroresponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountsRequest struct {
	PageID   int32 `form:"page_id,default=1" binding:"min=1"`
	PageSize int32 `form:"page_size,default=10" binding:"min=5,max=100"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, erroresponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, erroresponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, erroresponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type updateAccountRequest struct {
	Balance int64 `json:"balance" binding:"required,min=1"`
}

func (server *Server) updateAccount(ctx *gin.Context) {
	var getAccountReq getAccountRequest
	var updateAccountReq updateAccountRequest

	if err := ctx.ShouldBindJSON(&updateAccountReq); err != nil {
		ctx.JSON(http.StatusBadRequest, erroresponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&getAccountReq); err != nil {
		ctx.JSON(http.StatusBadRequest, erroresponse(err))
		return
	}

	arg := db.UpdateAccountParams{
		ID:      getAccountReq.ID,
		Balance: updateAccountReq.Balance,
	}

	account, err := server.store.UpdateAccount(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, erroresponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, erroresponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) deleteAccount(ctx *gin.Context) {
	var getAccountReq getAccountRequest
	if err := ctx.ShouldBindUri(&getAccountReq); err != nil {
		ctx.JSON(http.StatusBadRequest, erroresponse(err))
		return
	}

	err := server.store.DeleteAccount(ctx, getAccountReq.ID)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, erroresponse(err))
		return
	}
	ctx.JSON(http.StatusNoContent, struct{}{})
}
