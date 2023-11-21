package router

import (
	"github.com/ljcbaby/plan/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter
func SetupRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Create controller
	Controller := &controller.Controller{}

	// Index
	r.GET("/", Controller.Index)

	// v1
	v1 := r.Group("/v1")
	{
		// Course
		v1.POST("/courses", Controller.Course.CreateCourse)
		v1.DELETE("/courses/:id", Controller.Course.DeleteCourse)
		v1.PATCH("/courses/:id", Controller.Course.UpdateCourse)
		v1.GET("/courses", Controller.Course.GetCourseList)
		v1.POST("/courses/files", Controller.Course.UploadCourseFile)
		v1.GET("/courses/files", Controller.Course.GetCourseFileList)

		// Program
		v1.POST("/programs", Controller.Program.CreateProgram)
		v1.DELETE("/programs/:id", Controller.Program.DeleteProgram)
		v1.PATCH("/programs/:id", Controller.Program.UpdateProgram)
		v1.GET("/programs/:id", Controller.Program.GetProgram)
		v1.GET("/programs/:id/calculate", Controller.Program.CalculateProgram)
		v1.GET("/programs", Controller.Program.GetProgramList)
		v1.GET("/programs/:id/files", Controller.Program.GetProgramFile)

		// Public
		v1.GET("/download/:filename", Controller.DownloadFile)
	}

	return r
}
