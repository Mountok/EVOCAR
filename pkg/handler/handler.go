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
		auth.POST("/sign-in", h.signUp)
	}

	api := router.Group("/api")
	{ 
		api.GET("/conn",h.conn)   
	}  

	return router 
} 

func (h *Handler) conn(c * gin.Context) {
	res := h.services.Conn.Conn()
	c.JSON(http.StatusOK,map[string]interface{}{
		"connection": res,
	})
}