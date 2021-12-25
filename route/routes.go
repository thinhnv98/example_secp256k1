package route

import (
	"database/sql"

	"example_secp256k1/controller"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Db     *sql.DB
	Server *gin.Engine
}

func (_self Route) Register() {
	messageController := controller.MessageHandler{}
	messageRoutes := _self.Server.Group("")
	{
		messageRoutes.GET("/generate-keys", messageController.GenKeys)
		messageRoutes.POST("/someone-create-mess", messageController.SomeoneCreate)
		messageRoutes.POST("/i-create-mess", messageController.ICreate)
		messageRoutes.POST("/decrypt-mess", messageController.DecryptMessage)
		messageRoutes.POST("/verify-mess", messageController.VerifyMessage)
	}
}
