package controller

import "github.com/gin-gonic/gin"

func InitController() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	_ = NewBaseController()
	projectController := NewProjectController()

	pc := engine.Group("/api/dev/project")
	pc.POST("/init", projectController.InitProject)
	pc.GET("/list", projectController.GetProjectList)
	pc.GET("/:project/:repo/list", projectController.GetRepoList)
	pc.POST("/application/create", projectController.CreateApplication)
	pc.POST("/application/destroy", projectController.DestroyApplication)
	pc.POST("/application/deploy", projectController.DeployApplication)

	return engine
}
