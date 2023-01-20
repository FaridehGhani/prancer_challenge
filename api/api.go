package api

import "github.com/gin-gonic/gin"

var api apiHandler

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/deliver_point", api.DeliverPoint)

	return router
}
