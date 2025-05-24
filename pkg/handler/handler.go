package handler

import (
	"net/http"
	"todoapp/pkg/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}






func (h *Handler) InitRoute() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Разрешить только этот источник
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Разрешенные методы
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{ 
		api.GET("/conn",h.conn)   
		api.GET("/orders",h.getOrders)
		api.POST("/orders",h.createOrder)
		api.GET("orders/by-phone",h.getOrdersByPhoneNumber)
		api.PUT("/orders/:id/accept",h.acceptOrder)
		api.PUT("/orders/:id/complete",h.completeOrder)
		api.PUT("/orders/:id/cancel",h.cancleOrder)
		api.GET("/orders/active",h.activeOrders)

		executors := api.Group("/executors")
		{
			executors.GET("/history",h.getExecutorsHistory)
		}

	}


	return router 
} 

func (h *Handler) conn(c * gin.Context) {
	id, _ := c.Get(userCtx)
	res := h.services.Conn.Conn()
	c.JSON(http.StatusOK,map[string]interface{}{
		"connection": res,
		"userId": id,
	})
}
