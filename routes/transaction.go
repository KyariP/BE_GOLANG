package routes

import (
	"dumbmerch/handler"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	TransactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handler.HandlerTransaction(TransactionRepository)
	e.GET("/transaction", middleware.Auth(h.FindTransaction))
	e.GET("/transactions", middleware.Auth(h.GetTransByUser))
	e.GET("/transaction/:id", middleware.Auth(h.FindTransactionId))
	e.DELETE("/transaction/:id", middleware.Auth(h.DeleteTransaction))
	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	// e.PATCH("/transaction/:id", middleware.Auth(h.Notification))
	e.POST("/notification", h.Notification)
}
