package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/karankumarshreds/go-blog-api/internal/core"
	"github.com/karankumarshreds/go-blog-api/internal/middlewares"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogHandlers struct {
	blogService BlogServicePort
}

type BlogServicePort interface {
	Create(payload core.CreateBlogDto) (*primitive.ObjectID, error)
}

func NewBlogHandlers(blogService BlogServicePort) *BlogHandlers {
	return &BlogHandlers{blogService}
}

func (b *BlogHandlers) Create(c *fiber.Ctx) error {
	blog := new(core.CreateBlogDto)
	if parseError := c.BodyParser(&blog); parseError != nil {
		log.Println("Body parser error")
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": parseError.Error(),
		})
	}
	if validateError := middlewares.Validator.Struct(blog); validateError != nil {
		log.Println("Body validation error")
		return c.Status(fiber.ErrBadRequest.Code).JSON(validateError.Error())
	}

	if id, createError := b.blogService.Create(*blog); createError != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(createError.Error())
	} else {
		return c.JSON(id)
	}
}
