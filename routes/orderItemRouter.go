package routes

import (
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderitem/", controller.GetFoods())
	incomingRoutes.GET("/orderitem/:food_id", controller.GetFood())
	incomingRoutes.POST("/orderitem", controller.CreateFood())
	incomingRoutes.PATCH("/orderitem/:food_id", controller.UpdateFood())
}
