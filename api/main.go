package main

//  Require the packages
import (
	"context"
	"fmt"
	"log"

	"api/config"
	"api/controllers"
	"api/routes"
	"api/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Create required variables that we'll re-assign later
var (
	userService             services.UserService
	UserController          controllers.UserController
	UserRouteController     routes.UserRouteController
	categoryCollection      *mongo.Collection
	categoryService         services.CategoryService
	CategoryController      controllers.CategoryController
	CategoryRouteController routes.CategoryRouteController
	productCollection       *mongo.Collection
	productService          services.ProductService
	ProductController       controllers.ProductController
	ProductRouteController  routes.ProductRouteController
	authCollection          *mongo.Collection
	authService             services.AuthService
	AuthController          controllers.AuthController
	AuthRouteController     routes.AuthRouteController
	server                  *gin.Engine
	ctx                     context.Context
	mongoclient             *mongo.Client
)

// Init function that will run before the "main" function

func init() {

	// Load env
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	ctx = context.TODO()

	// Connect to MongoDB

	mongoconn := options.Client().ApplyURI(config.DBUri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// Collections
	authCollection = mongoclient.Database("golang_mongodb").Collection("users")
	productCollection = mongoclient.Database("golang_mongodb").Collection("Products")
	categoryCollection = mongoclient.Database("golang_mongodb").Collection("Categories")
	categoryService = services.NewCategoryService(categoryCollection, ctx)
	CategoryController = controllers.NewCategoryController(categoryService)
	CategoryRouteController = routes.NewRouteCategoryController(CategoryController)
	userService = services.NewUserServiceImpl(authCollection, ctx)
	authService = services.NewAuthService(authCollection, ctx)
	productService = services.NewProductService(productCollection, ctx)
	ProductController = controllers.NewProductController(productService)
	ProductRouteController = routes.NewRouteProductController(ProductController)
	AuthController = controllers.NewAuthController(authService, userService)
	AuthRouteController = routes.NewAuthRouteController(AuthController)
	UserController = controllers.NewUserController(userService)
	UserRouteController = routes.NewRouteUserController(UserController)

	model := mongo.IndexModel{Keys: bson.D{{"name", "text"}}}
	name, err := mongoclient.Database("golang_mongodb").Collection("Products").Indexes().CreateOne(context.TODO(), model)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Name of index created", name)
	//  Create the Gin Engine instance
	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}
	defer mongoclient.Disconnect(ctx)
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowOrigins = []string{"http://localhost:4200"}
	server.Use(cors.New(corsConfig))
	router := server.Group("/api")

	AuthRouteController.AuthRoute(router, userService)
	UserRouteController.UserRoute(router, userService)
	ProductRouteController.ProductRoute(router, productService)
	CategoryRouteController.CategoryRoute(router, categoryService)
	log.Fatal(server.Run(":" + config.Port))
}
