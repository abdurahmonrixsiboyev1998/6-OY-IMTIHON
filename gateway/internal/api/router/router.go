package router

import (
	"gateway/client"
	"gateway/internal/api/handler"
	"gateway/storage/redisdb"
	"os"


	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func Routes() {
	r := gin.Default()

	userClient := client.DialUsersClient(os.Getenv("users_url"))
	budgetClient := client.DialBudgetClient(os.Getenv("budget_url"))
	redis := redisdb.ConnectRedis(os.Getenv("redis_url"))
	incomeClient := client.DialIncomeClient(os.Getenv("income_url"))
	notificationClient := client.DialNotifyClient("notification_url")
	Handler := handler.NewHandler(userClient, redis, budgetClient, incomeClient, notificationClient)

	//USERS
	r.POST("/api/users/register", Handler.RegisterUser)
	r.POST("/api/users/verification", Handler.VerifyCode)
	r.POST("/api/users/login", Handler.UserLogin)
	r.POST("/api/users/forgot-password", Handler.ForgetPassword)
	r.PUT("/api/users/update-password", Handler.UpdatePassword)
	r.DELETE("/api/users/logout", Handler.UserLogout)
	r.GET("/api/users/:id", Handler.GetUserByID)

	//BUDGETS
	r.POST("/api/budgets/create", Handler.CreateBudget)
	r.PUT("/api/budgets/update", Handler.UpdateBudgetAmount)
	r.GET("/api/budgets", Handler.GetBudgets)
	r.DELETE("/api/budgets/delete/:id", Handler.DeleteBudgetByID)

	//Income
	r.POST("/api/income/create", Handler.CreateTransaction)
	r.PUT("api/income/update", Handler.UpdateTransactionByID)
	r.DELETE("api/incomde/delete/:id", Handler.DeleteTransactionByID)
	r.GET("api/incomeID", Handler.GetTransactionByID)
	r.GET("api/incodeByDate", Handler.GetTransactionByDate)

	//Notify
	r.GET("api/notify/lst", Handler.GetNotification)
	r.GET("api/unreadnotify/lst", Handler.GetUnreadNotifications)
	
	r.Run(os.Getenv("gateway_url"))

}
