package controller

import "github.com/gin-gonic/gin"

func InitController() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	_ = NewBaseController()
	projectController := NewProjectController()

	pc := engine.Group("/project")
	pc.POST("/init", projectController.InitProject)

	return engine
}
