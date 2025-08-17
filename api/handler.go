package api

import (
	"net/http"

	"github.com/digitcodestudio/accounting-module/models"
	"github.com/digitcodestudio/accounting-module/service"
	"github.com/gin-gonic/gin"
)

type ApiHandler struct {
    PayableService    *service.PayableService
    ReceivableService *service.ReceivableService
}

// Constructor
func NewApiHandler(payable *service.PayableService, receivable *service.ReceivableService) *ApiHandler {
    return &ApiHandler{
        PayableService:    payable,
        ReceivableService: receivable,
    }
}

// Register all routes
func (h *ApiHandler) RegisterRoutes(r *gin.Engine) {
    r.POST("/payables", h.CreatePayable)
    r.POST("/receivables", h.CreateReceivable)
    r.GET("/payables", h.ListPayables)
    r.GET("/receivables", h.ListReceivables)
}

// Create Payable
func (h *ApiHandler) CreatePayable(c *gin.Context) {
    var p models.Payable
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.PayableService.AddPayable(&p); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, p)
}

// Create Receivable
func (h *ApiHandler) CreateReceivable(c *gin.Context) {
    var r models.Receivable
    if err := c.ShouldBindJSON(&r); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.ReceivableService.AddReceivable(&r); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, r)
}

// List Payables
func (h *ApiHandler) ListPayables(c *gin.Context) {
    payables, err := h.PayableService.List()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, payables)
}

// List Receivables
func (h *ApiHandler) ListReceivables(c *gin.Context) {
    receivables, err := h.ReceivableService.List()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, receivables)
}
