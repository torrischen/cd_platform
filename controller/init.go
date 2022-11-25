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
	pc.GET("/:project/repolist", projectController.GetRepoList)
	pc.GET("/:project/:repo/list", projectController.GetImageList)
	pc.POST("/application/create", projectController.CreateApplication)
	pc.POST("/application/destroy", projectController.DestroyApplication)
	pc.POST("/application/deploy", projectController.DeployApplication)
	pc.GET("/:project/application/all", projectController.GetApplicationList)
	pc.GET("/:project/application/:application/pod/all", projectController.GetApplicationDetailsByApplication)
	pc.GET("/:project/detail", projectController.GetApplicationDetailsByProject)
	pc.POST("/application/ingress/insert", projectController.InsertApplicationIngressPath)
	pc.GET("/:project/application/:application/ingress", projectController.GetApplicationIngress)
	pc.GET("/:project/pod/:podname/log", projectController.GetPodLog)
	pc.POST("/application/ingress/delete", projectController.DeleteSpecifiedIngressRule)
	pc.POST("/application/replica/patch", projectController.PatchApplicationReplica)
	pc.POST("/application/restart", projectController.RestartApplication)
	pc.GET("/:project/application/:application/yaml", projectController.GetApplicationYaml)

	return engine
}
