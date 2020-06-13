package main

import (
	"go-lms-of-pupilfirst/cmd/graphql/resolvers"
	"go-lms-of-pupilfirst/configs"
	"go-lms-of-pupilfirst/migrations"
	"go-lms-of-pupilfirst/pkg/auth"
	"go-lms-of-pupilfirst/pkg/database"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetPrefix("go-lms-graphql" + " : ")
	log := log.New(os.Stdout, log.Prefix(), log.Flags())

	if err := envconfig.Process("go-lms-api", &configs.CFG); err != nil {
		log.Fatalf("main : Error Parsing Config file: %+v", err)
	}

	dbConfig, err := configs.LoadConfig()
	if err != nil {
		log.Printf("main : Error loading database %+v", err)
	}
	log.Printf("%+v", dbConfig)
	db, err := database.Initialize(dbConfig.Storage)
	defer db.Close()

	migrations.Migrate(db)

	authenticator, _ := auth.NewAuthenticatorFile("", time.Now().UTC(), configs.CFG.Auth.KeyExpiration)

	r := gin.Default()
	resolvers.ApplyResolvers(r, db, authenticator)

	log.Println("Running @http://" + configs.CFG.Server.Graphql + "/graphql")
	log.Fatalln(r.Run(configs.CFG.Server.Graphql))

}
