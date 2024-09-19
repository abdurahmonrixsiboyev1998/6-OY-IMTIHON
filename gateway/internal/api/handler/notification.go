package handler

import (
	"gateway/proto/notification"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetNotification(c *gin.Context) {
	notifyList, err := h.Notification.GetNotfication(c.Request.Context(), &notification.Empty{})
	if err != nil {
		log.Println("Error get notification")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to get notification"})
		return
	}
	c.IndentedJSON(200, notifyList)
}

func (h *Handler) GetUnreadNotifications(c *gin.Context) {
	notifyList, err := h.Notification.GetUnreadNotfications(c.Request.Context(), &notification.Empty{})
	if err != nil {
		log.Println("Error get notification")
		c.AbortWithStatusJSON(500, gin.H{"error": "Unable to get notification"})
		return
	}
	c.IndentedJSON(200, notifyList)
}

