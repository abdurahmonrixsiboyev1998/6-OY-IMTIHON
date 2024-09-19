package handler

import (
	"gateway/proto/income"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *Handler) CreateTransaction(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var createIncome income.TransactionInfo
	if err := protojson.Unmarshal(bytes, &createIncome); err != nil {
		log.Println("Error unmarshaling: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error":"Invalid request body"})
		return
	}
	trID, err := h.Income.CreateTransaction(c.Request.Context(), &createIncome)
	if err != nil {
		log.Println("Error creating income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to create income"})
		return
	}

	c.IndentedJSON(201, trID)
}

func (h *Handler) UpdateTransactionByID(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var TransactionWithID income.TransactionWithID
	if err := protojson.Unmarshal(bytes, &TransactionWithID); err != nil {
		log.Println("Error unmarshaling: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error":"Invalid request body"})
		return
	}

	tres, err := h.Income.UpdateTransactionByID(c.Request.Context(), &TransactionWithID)
	if err != nil {
		log.Println("Error updating income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to update income"})
		return
	}

	c.IndentedJSON(200, tres)
}

func (h *Handler) DeleteTransactionByID(c *gin.Context) {
	id := c.Param("id")
	var trId income.TransactionID
	trId.Id = id
	tres, err := h.Income.DeleteTransactionByID(c.Request.Context(), &trId)
	if err != nil {
		log.Println("Error deleted income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to delete income"})
		return
	}

	c.IndentedJSON(200, tres)
}

func (h *Handler) GetTransactionByID(c *gin.Context) {
	id := c.Param("id")
	var trId income.TransactionID
	trId.Id = id
	transactionWithID, err := h.Income.GetTransactionByID(c.Request.Context(), &trId)
	if err != nil {
		log.Println("Error get income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to delete income"})
		return
	}

	c.IndentedJSON(200, transactionWithID)
}

func (h *Handler) GetTransactionsByCategory(c *gin.Context) {
	category := c.Param("categroy")
	var trCategory income.TransactionCategory
	trCategory.Category = category
	list, err := h.Income.GetTransactionsByCategory(c.Request.Context(), &trCategory)
	if err != nil {
		log.Println("Error get income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to get income"})
		return
	}

	c.IndentedJSON(200, list)
}

func (h *Handler) GetTransactionByDate(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var trDate income.TransactionDate
	if err := protojson.Unmarshal(bytes, &trDate); err != nil {
		log.Println("Error unmarshaling: ", err)
		c.AbortWithStatusJSON(400, gin.H{"error":"Invalid request body"})
		return
	}
	lst, err := h.Income.GetTransactionByDate(c.Request.Context(), &trDate)
	if err != nil {
		log.Println("Error get income")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to delete income"})
		return
	}

	c.IndentedJSON(200, lst)
}
