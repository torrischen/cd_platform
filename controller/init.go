package controller

import "github.com/gin-gonic/gin"

func InitController() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	_ = NewBaseController()
	watchController := NewWatchController()
	execController := NewExecController()
	sitController := NewSitController()

	dft := engine.Group("/api")

	ns := dft.Group("/namespace")
	ns.GET("/get/:name", watchController.GetNamespaceByName)
	ns.POST("/create", execController.CreateNamespace)

	dep := dft.Group("/deployment")
	dep.GET("/get/:namespace/:name", watchController.GetDeploymentByName)
	dep.POST("/getByLabel", watchController.GetDeploymentByLabel)
	dep.POST("/create", execController.CreateDeployment)

	po := dft.Group("/pod")
	po.GET("/get/:namespace/:name", watchController.GetPodByName)
	po.POST("/getByLabel", watchController.GetPodByLabel)

	sts := dft.Group("/statefulset")
	sts.GET("/get/:namespace/:name", watchController.GetStatefulsetByName)
	sts.POST("/getByLabel", watchController.GetStatefulsetByLabel)

	svc := dft.Group("/service")
	svc.GET("/get/:namespace/:name", watchController.GetServiceByName)
	svc.POST("/getByLabel", watchController.GetServiceByLabel)

	igs := dft.Group("/ingress")
	igs.GET("/get/:namespace/:name", watchController.GetIngressByName)
	igs.POST("/getByLabel", watchController.GetIngressByLabel)

	sit := dft.Group("/sit")
	sit.POST("/deployment/create", sitController.CreateDeployment)

	return engine
}
