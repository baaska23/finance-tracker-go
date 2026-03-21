package transactions

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service *TransactionService
}

func NewTransactionHandler(service *TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (handler *TransactionHandler) Create(c *gin.Context) {
	var req Transaction
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	transaction, err := handler.service.CreateTransaction(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

func (handler *TransactionHandler) ListAll(c *gin.Context) {
	transactions, err := handler.service.ListTransaction()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (handler *TransactionHandler) GetById(c *gin.Context) {
	idParam := c.Query("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	transaction, err := handler.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (handler *TransactionHandler) Update(c *gin.Context) {
	idParam := c.Query("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction"})
		return
	}

	var req TransactionPatch
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	req.TransactionId = &id

	transaction, err := handler.service.UpdateTransaction(req)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (handler *TransactionHandler) Delete(c *gin.Context){
	idParam := c.Query("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction"})
		return
	}

	var req Transaction
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	req.TransactionId = id

	transaction, err := handler.service.DeleteTransaction(req)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}
