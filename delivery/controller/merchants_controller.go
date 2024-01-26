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

type MerchantController struct {
	merchantUC     usecase.MerchantUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (m *MerchantController) createHandler(ctx *gin.Context) {

	var payload entity.Merchant

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	merchant, err := m.merchantUC.RegisterNewMerchant(payload)

	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return

	}
	common.SendCreateResponse(ctx, merchant, "Created")
}

func (m *MerchantController) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := m.merchantUC.DeleteMerchant(id)

	log.Println(err)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			common.SendErrorResponse(ctx, http.StatusNotFound, "Merchant with ID "+id+" not found. Delete data failed")
		} else {
			common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		}

		return
	}

	common.SendDeletedResponse(ctx, "Delete merchant successfully")
}

// read by
func (m *MerchantController) getByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	merchant, err := m.merchantUC.FindMerchantByID(id)
	if err != nil {

		common.SendErrorResponse(ctx, http.StatusNotFound, "Users with ID "+id+" not found")
		return
	}
	common.SendSingleResponse(ctx, merchant, "Ok")
}

// update

func (m *MerchantController) putHandler(ctx *gin.Context) {
	var payload entity.Merchant
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Failed to bind data")
		return
	}
	merchant, err := m.merchantUC.UpdateMerchant(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	common.SendSingleResponse(ctx, merchant, "Updated Successfully")

}

// pagination
func (m *MerchantController) ListHandler(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "5"))

	merchant, paging, err := m.merchantUC.ListAll(page, size)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var response []interface{}

	for _, v := range merchant {
		response = append(response, v)
	}
	common.SendPagedResponse(ctx, response, paging, "Ok")
}

// route
func (m *MerchantController) Route() {
	m.rg.GET(config.MerchantGetById, m.authMiddleware.RequireToken("customer", "merchant", "admin"), m.getByIdHandler)
	m.rg.POST(config.MerchantCreate, m.authMiddleware.RequireToken("customer", "merchant", "admin"), m.createHandler)
	m.rg.PUT(config.MerchantUpdate, m.authMiddleware.RequireToken("customer", "merchant", "admin"), m.putHandler)
	m.rg.GET(config.MerchantList, m.authMiddleware.RequireToken("customer", "merchant", "admin"), m.ListHandler)
	m.rg.DELETE(config.MerchantDelete, m.authMiddleware.RequireToken("customer", "merchant", "admin"), m.deleteHandler)
}

func NewMerchantController(merchantUC usecase.MerchantUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *MerchantController {
	return &MerchantController{
		merchantUC:     merchantUC,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}
