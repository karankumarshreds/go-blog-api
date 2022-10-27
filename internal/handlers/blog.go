package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/karankumarshreds/go-blog-api/custom_errors"
	"github.com/karankumarshreds/go-blog-api/internal/core"
	"github.com/karankumarshreds/go-blog-api/internal/middlewares"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogHandlers struct {
	blogService BlogServicePort
}

type BlogServicePort interface {
	Create(payload core.CreateBlogDto) (*primitive.ObjectID, *custom_errors.CustomError)
	Get(id string) (*core.Blog, *custom_errors.CustomError)
	Update(id string, payload core.CreateBlogDto) (*core.Blog, *custom_errors.CustomError)
	Delete(id string) *custom_errors.CustomError
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
		return c.Status(fiber.StatusInternalServerError).JSON(createError.Message)
	} else {
		return c.JSON(id)
	}
}

func (b *BlogHandlers) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := b.blogService.Get(id)

	if err != nil {
		return c.Status(err.Status).JSON(err.Message)
	} else {
		return c.JSON(res)
	}
}

func (b *BlogHandlers) Update(c *fiber.Ctx) error {
	id := c.Params("id")
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

	if res, createError := b.blogService.Update(id, *blog); createError != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(createError.Message)
	} else {
		return c.JSON(res)
	}
}

func (b *BlogHandlers) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := b.blogService.Delete(id)
	if err != nil {
		return c.Status(err.Status).JSON(err.Message)
	} else {
		return c.SendStatus(fiber.StatusAccepted)
	}
}
