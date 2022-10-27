package controller

import "github.com/gin-gonic/gin"

func InitController() *gin.Engine {
	engine := gin.Default()

	_ = NewBaseController()
	watchController := NewWatchController()

	dep := engine.Group("/deployment")
	dep.GET("/get/:namespace/:name", watchController.GetDeploymentByName)
	dep.POST("/getByLabel", watchController.GetDeploymentByLabel)

	return engine
}
