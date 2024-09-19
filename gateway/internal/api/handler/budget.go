package handler

import (
	"gateway/proto/budget"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *Handler) CreateBudget(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var createBudget budget.BudgetInfo
	if err := protojson.Unmarshal(bytes, &createBudget); err != nil {
		log.Println("Error unmarshaling:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	budgetID, err := h.Budget.CreateBudget(c.Request.Context(), &createBudget)
	if err != nil {
		log.Println("Error creating budget:", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to create budget"})
		return
	}

	c.IndentedJSON(201, budgetID)
}

func (h *Handler) UpdateBudgetAmount(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var updateBudget budget.BudgetUpdate
	if err := protojson.Unmarshal(bytes, &updateBudget); err != nil {
		log.Println("Error unmarshaling:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	budgetResponse, err := h.Budget.UpdateBudgetAmount(c.Request.Context(), &updateBudget)
	if err != nil {
		log.Println("Error updating budget:", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to update budget"})
		return
	}

	c.IndentedJSON(200, budgetResponse)
}

func (h *Handler) GetBudgets(c *gin.Context) {
	var budgetEmpty budget.Empty
	budgetWithID, err := h.Budget.GetBudgets(c.Request.Context(), &budgetEmpty)
	if err != nil {
		log.Println("Error getting budgets:", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to retrieve budgets"})
		return
	}

	c.IndentedJSON(200, budgetWithID)
}

func (h *Handler) DeleteBudgetByID(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var budgetID budget.BudgetID
	if err := protojson.Unmarshal(bytes, &budgetID); err != nil {
		log.Println("Error unmarshaling:", err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	budgetResponse, err := h.Budget.DeleteBudgetByID(c.Request.Context(), &budgetID)
	if err != nil {
		log.Println("Error deleting budget:", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to delete budget"})
		return
	}

	c.IndentedJSON(200, budgetResponse)
}
