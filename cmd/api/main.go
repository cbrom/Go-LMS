package main

import (
	"encoding/json"
	"go-lms-of-pupilfirst/cmd/api/handlers"
	"go-lms-of-pupilfirst/cmd/api/routes"
	"go-lms-of-pupilfirst/configs"

	_ "go-lms-of-pupilfirst/cmd/api/docs"
	"go-lms-of-pupilfirst/migrations"
	"go-lms-of-pupilfirst/pkg/auth"
	"go-lms-of-pupilfirst/pkg/database"
	"go-lms-of-pupilfirst/pkg/flag"

	"log"
	"os"
	"time"

	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/kelseyhightower/envconfig"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GO-LMS service API
// @version 1.0
// @description This is GO-LMS server.
// @termsOfService GO-LMS.com

// @contact.name API Support
// @contact.url http://GO-LMS.com/support
// @contact.email GO-LMS@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3001
// @BasePath /v1

// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetPrefix("go-lms-api" + " : ")
	log := log.New(os.Stdout, log.Prefix(), log.Flags())

	if err := envconfig.Process("go-lms-api", &configs.CFG); err != nil {
		log.Fatalf("main : Error Parsing Config file: %+v", err)
	}

	log.Println("main : Initialize Redis")
	redisClient := redistrace.NewClient(&redis.Options{
		Addr:        configs.CFG.Redis.Host,
		DB:          configs.CFG.Redis.DB,
		DialTimeout: configs.CFG.Redis.DialTimeout,
	})

	defer redisClient.Close()

	if err := flag.Process(&configs.CFG); err != nil {
		if err != flag.ErrHelp {
			log.Fatalf("main : Error Parsing Command Line : %+v", err)
		}
		// else provide help here
		return
	}

	// Print the config
	{
		cfgJSON, err := json.MarshalIndent(configs.CFG, "", "")
		if err != nil {
			log.Fatalf("main : Error marshaling config to JSON : %+v", err)
		}
		log.Printf("main : Config : %v\n", string(cfgJSON))
	}

	dbConfig, err := configs.LoadConfig()
	if err != nil {
		log.Printf("main : Error loading database %+v", err)
	}
	log.Printf("%+v", dbConfig)
	db, err := database.Initialize(dbConfig.Storage)
	defer db.Close()

	if err != nil {
		log.Fatalf("main: Error initializing database %+v", err)
	}
	authenticator, _ := auth.NewAuthenticatorFile("", time.Now().UTC(), configs.CFG.Auth.KeyExpiration)

	migrations.Migrate(db)

	config, err := configs.LoadConfig()
	app := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	app.Use(cors.New(corsConfig))
	routes.ApplyRoutes(app, authenticator, db, &handlers.UserController{})
	app.Use(database.InjectDB(db))
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.Run(configs.CFG.Server.Host)
}

// go install github.com/swaggo/swag/cmd/swag@latest
