package main

import (
	"encoding/json"
	"go-lms-of-pupilfirst/cmd/api/handlers"
	"go-lms-of-pupilfirst/configs"
	"go-lms-of-pupilfirst/migrations"
	"go-lms-of-pupilfirst/pkg/auth"
	"go-lms-of-pupilfirst/pkg/database"
	"go-lms-of-pupilfirst/pkg/flag"
	"log"
	"os"
	"time"

	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/kelseyhightower/envconfig"
)

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
	database.New(dbConfig.Storage)
	authenticator, _ := auth.NewAuthenticatorFile("", time.Now().UTC(), configs.CFG.Auth.KeyExpiration)

	migrations.Migrate()

	app := gin.Default()
	handlers.ApplyRoutes(app, authenticator)
	app.Run(configs.CFG.Server.Host)
}
