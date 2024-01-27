package controller

import (
	"net/http"
	"payeasy/config"
	"payeasy/delivery/middleware"
	"payeasy/entity"
	"payeasy/usecase"

	"github.com/gin-gonic/gin"
)



type HistoryController struct {
	usecase    usecase.HistoryUsecase
	rg         *gin.RouterGroup
	middleware middleware.AuthMiddleware
}

func NewHistoryController(historyUC usecase.HistoryUsecase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *HistoryController {
	return &HistoryController{
		usecase:        historyUC,
		rg:             rg,
		middleware: authMiddleware,
	}
}


func (h *HistoryController) HistoryList(c *gin.Context) {
	histories, err := h.usecase.ReadHistoryTransaction()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": histories})
}

func (h *HistoryController) HistoryGetByIdUsers(c *gin.Context) {
	idUser := c.Param("user")
	histories, err := h.usecase.GetByIdUsers(idUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": histories})
}

func (h *HistoryController) HistoryGetByIdMerchant(c *gin.Context) {
	idMerchant := c.Param("merchant")
	histories, err := h.usecase.GetByIdMerchant(idMerchant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": histories})
}

func (h *HistoryController) HistoryGetBalance(c *gin.Context) {
	idUser := c.Param("user")
	balance := h.usecase.GetBalanceByUserId(idUser)
	if balance == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": 0})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": balance})
}


func (h *HistoryController) HistoryCreate(c *gin.Context) {
	history := entity.History{}

	if err := c.ShouldBindJSON(&history); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.RegisterNewTransaction(&history)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": history})
}

func (h *HistoryController) Route() {
	h.rg.POST(config.HistoryCreate, h.middleware.RequireToken("customer", "merchant", "admin"), h.HistoryCreate)                   // Membuat transaksi baru
	h.rg.GET(config.HistoryCreate,h.middleware.RequireToken("customer", "merchant", "admin"), h.HistoryList)                      // Menampilkan daftar transaksi
	h.rg.GET(config.HistoryGetByIdUsers, h.middleware.RequireToken("customer", "merchant", "admin"), h.HistoryGetByIdUsers)        // Menampilkan transaksi berdasarkan ID pengguna
	h.rg.GET(config.HistoryGetByIdMerchant, h.middleware.RequireToken("customer", "merchant", "admin"), h.HistoryGetByIdMerchant) // Menampilkan transaksi berdasarkan ID merchant
	h.rg.GET(config.HistoryGetBalance, h.middleware.RequireToken("customer", "merchant", "admin"), h.HistoryGetBalance) // Menampilkan transaksi berdasarkan ID merchant

	
}
