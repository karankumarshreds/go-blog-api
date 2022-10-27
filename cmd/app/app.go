package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/karankumarshreds/go-blog-api/internal/handlers"
	"github.com/karankumarshreds/go-blog-api/internal/repos"
	"github.com/karankumarshreds/go-blog-api/internal/services"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	app *fiber.App
}

func NewApp() *App {
	return &App{}
}

const CONNECTION_TIMEOUT = 10

func (a *App) Init() {
	CheckEnvs()
	db := a.MongoConnect()
	a.app = fiber.New()

	// Dependency injection
	blogRepo := repos.NewBlogRepo(db)
	blogService := services.NewBlogService(blogRepo)
	blogHandlers := handlers.NewBlogHandlers(blogService)

	a.app.Post(createPath("/blog"), blogHandlers.Create)
	a.app.Get(createPath("/blog/:id"), blogHandlers.Get)
	a.app.Put(createPath("/blog/:id"), blogHandlers.Update)
	a.app.Delete(createPath("/blog/:id"), blogHandlers.Delete)
}

func (a *App) Start() {
	log.Panic(a.app.Listen(os.Getenv("PORT")))
}

func (a *App) MongoConnect() *mongo.Database {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(os.Getenv("CONN_STRING")).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), CONNECTION_TIMEOUT*time.Second)
	defer cancel()

	mongoClient, mongoErr := mongo.Connect(ctx, clientOptions)
	if mongoErr != nil {
		log.Panic("Mongo connection error", mongoErr)
	} else {
		log.Println("Mongo connection successfull")
	}

	return mongoClient.Database("go-blog-api")
}

func CheckEnvs() {
	if err := godotenv.Load(); err != nil {
		log.Panic("Error loading .env file")
	}
	envs := []string{
		"CONN_STRING",
		"PORT",
		"API_PREFIX",
	}
	for _, e := range envs {
		if os.Getenv(e) == "" {
			log.Panic(fmt.Sprintf("%v env not defined", e))
		}
	}
}

func createPath(path string) string {
	return fmt.Sprintf("%v%v", os.Getenv("API_PREFIX"), path)
}
