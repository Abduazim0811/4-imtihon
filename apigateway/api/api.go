package api

import (
	devicehandler "Api-GateWay/api/device_handler"
	userhandler "Api-GateWay/api/user_handler"
	"Api-GateWay/clients/device"
	"Api-GateWay/clients/user"
	"Api-GateWay/internal/jwt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes() {
	// User service
	userclient := user.DialGRPCUser()
	userHandler := &userhandler.UserHandler{ClientUser: userclient}

	// Device service
	deviceclient := device.DialGrpcDevice()
	deviceHandler := &devicehandler.DeviceHandler{ClientDevice: deviceclient}

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/users", userHandler.CreateUser)
	router.POST("/verifycode", userHandler.VerifyCode)
	router.POST("/login", userHandler.Login)
	router.GET("/users/:id", jwt.Protected(), userHandler.GetByIdUser)
	router.GET("/users", userHandler.GetUsers)
	router.PUT("/users", jwt.Protected(), userHandler.UpdateUser)
	router.PUT("/users/password", jwt.Protected(), userHandler.UpdatePassword)
	router.PUT("/users/email", jwt.Protected(), userHandler.UpdateEmail)
	router.DELETE("/users", jwt.Protected(), userHandler.DeleteUser)

	router.POST("/device", jwt.Protected(), deviceHandler.CreateDevice)
	router.GET("/device/:id", jwt.Protected(), deviceHandler.GetDeviceById)
	router.PUT("/device", jwt.Protected(), deviceHandler.UpdateDevice)
	router.DELETE("/device", jwt.Protected(), deviceHandler.DeleteDevice)

	router.Run(":7777")
}
