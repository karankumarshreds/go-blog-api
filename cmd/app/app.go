package app

import (
	"context"
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
	if err := godotenv.Load(); err != nil {
		log.Panic("Error loading .env file")
	}
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

	db := mongoClient.Database("go-blog-api")

	a.app = fiber.New()
	blogRepo := repos.NewBlogRepo(db)
	blogService := services.NewBlogService(blogRepo)
	blogHandlers := handlers.NewBlogHandlers(blogService)

	a.app.Post("/", blogHandlers.Create)
}

func (a *App) Start() {
	log.Panic(a.app.Listen(os.Getenv("PORT")))
}

func (a *App) MongoConnect() {

}
