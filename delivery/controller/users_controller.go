package controller

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"payeasy/config"
	"payeasy/delivery/middleware"
	"payeasy/entity"
	"payeasy/shared/common"
	"payeasy/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	usersUC        usecase.UsersUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (u *UsersController) createHandler(ctx *gin.Context) {

	var payload entity.Users

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	users, err := u.usersUC.RegisterNewUsers(payload)

	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return

	}
	common.SendCreateResponse(ctx, users, "Created")
}

func (u *UsersController) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := u.usersUC.DeleteUsers(id)

	log.Println(err)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			common.SendErrorResponse(ctx, http.StatusNotFound, "Users with ID "+id+" not found. Delete data failed")
		} else {
			common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		}

		return
	}

	common.SendDeletedResponse(ctx, "Delete user successfully")
}

// read by
func (u *UsersController) getByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	users, err := u.usersUC.FindUsersByID(id)
	if err != nil {

		common.SendErrorResponse(ctx, http.StatusNotFound, "Users with ID "+id+" not found")
		return
	}
	common.SendSingleResponse(ctx, users, "Ok")
}
func (u *UsersController) getByEmailHandler(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := u.usersUC.FindUsersByEmail(email)
	if err != nil {

		common.SendErrorResponse(ctx, http.StatusNotFound, "Employee with Username "+email+" not found")
		return
	}
	common.SendSingleResponse(ctx, user, "Ok")
}

// update

func (u *UsersController) putHandler(ctx *gin.Context) {
	var payload entity.Users
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Failed to bind data")
		return
	}
	users, err := u.usersUC.UpdateUsers(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	common.SendSingleResponse(ctx, users, "Updated Successfully")

}

// pagination
func (u *UsersController) ListHandler(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "5"))

	users, paging, err := u.usersUC.ListAll(page, size)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var response []interface{}

	for _, v := range users {
		response = append(response, v)
	}
	common.SendPagedResponse(ctx, response, paging, "Ok")
}

// route
func (u *UsersController) Route() {
	u.rg.GET(config.UserGetById, u.authMiddleware.RequireToken("customer", "merchant", "admin"), u.getByIdHandler)
	u.rg.GET(config.UserGetByEmail, u.authMiddleware.RequireToken("customer", "merchant", "admin"), u.getByEmailHandler)
	u.rg.POST(config.UserCreate, u.createHandler)
	u.rg.PUT(config.UserUpdate, u.authMiddleware.RequireToken("customer", "merchant", "admin"), u.putHandler)
	u.rg.GET(config.UserList, u.authMiddleware.RequireToken("customer", "merchant", "admin"), u.ListHandler)
	u.rg.DELETE(config.UserDelete, u.authMiddleware.RequireToken("customer", "merchant", "admin"), u.deleteHandler)
}

func NewUsersController(usersUC usecase.UsersUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *UsersController {
	return &UsersController{
		usersUC:        usersUC,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}
