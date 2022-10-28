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

// Create godoc
// @Summary Create a blog
// @Description Create a blog
// @Accept  json
// @Produce  json
// @Tags Item
// @Param blog body core.CreateBlogDto true "Blog"
// @Success 200 {object} core.Blog
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /api/blog/ [post]
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

// Get godoc
// @Summary Get a Blog
// @Description Get a Blog by its ID
// @ID get-item-by-id
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path string true "Blog ID"
// @Success 200 {object} core.Blog
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /api/blog/{id} [get]
func (b *BlogHandlers) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := b.blogService.Get(id)

	if err != nil {
		return c.Status(err.Status).JSON(err.Message)
	} else {
		return c.JSON(res)
	}
}

// Update godoc
// @Summary Update a blog
// @Description Update a blog by ID
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path string true "Blog ID"
// @Param blog body core.CreateBlogDto true "Blog"
// @Success 200 {object} core.Blog
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /api/blog/{id} [put]
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

// Delete godoc
// @Summary Delete a blog
// @Description Delete a blog by its ID
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path string true "Blog ID"
// @Success 200 {object} core.Blog
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /api/blog/{id} [delete]
func (b *BlogHandlers) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := b.blogService.Delete(id)
	if err != nil {
		return c.Status(err.Status).JSON(err.Message)
	} else {
		return c.SendStatus(fiber.StatusAccepted)
	}
}
