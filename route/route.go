package route

import (
	"sesi_8/handler"
	"sesi_8/service"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app service.ServiceInterface) {
	handler := handler.NewHttpServer(app)
	api := r.Group("/book")
	{
		api.GET("/", handler.GetAllBook)
		api.GET("/:id", handler.GetBookById)
		api.POST("/add", handler.AddBook)
		api.PUT("/update/:id", handler.UpdateBook)
		api.DELETE("/delete/:id", handler.DeleteBook)
	}
}
