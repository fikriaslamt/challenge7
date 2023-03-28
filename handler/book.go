package handler

import (
	"sesi_8/helper"
	"sesi_8/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) GetAllBook(c *gin.Context) {
	res, err := h.service.GetAllBook()
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}
	if len(res) == 0 {
		helper.NoContent(c)
		return
	}

	helper.Ok(c, res)

}
func (h HttpServer) GetBookById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.BadRequest(c, "id Must be Integer")
		return
	}

	res, err := h.service.GetBookById(id)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}
func (h HttpServer) AddBook(c *gin.Context) {
	var book model.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.service.AddBook(book)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, res)

}
func (h HttpServer) UpdateBook(c *gin.Context) {
	idParam := c.Param("id")
	book := model.Book{}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.BadRequest(c, "id Must be Integer")
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		helper.BadRequest(c, "Invalid Format data, Try again")
		return
	}

	if err := h.service.UpdateBook(id, book); err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Data Updated")

}
func (h HttpServer) DeleteBook(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.BadRequest(c, "id Must be Integer")
		return
	}

	if err := h.service.DeleteBook(id); err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.OkWithMessage(c, "Data Deleted")
}
