package routes

import (
	"github.com/alexguidi/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/:name", controllers.Greetings)
	r.POST("/students", controllers.CreateNewStudent)
	r.GET("/students/:id", controllers.FindStudentByID)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/students/cpf/:cpf", controllers.FindStudentByCPF)
	r.Run()
}
