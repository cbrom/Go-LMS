package database

import (
	"fmt"
	"go-lms-of-pupilfirst/configs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //
)

// Initialize gets the config and returns a database pointer
func Initialize(conf configs.Storage) (*gorm.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.Dbuser, conf.Dbpassword, conf.Database)

	db, err := gorm.Open("postgres", url)
	return db, err
}

// InjectDB injects database to gin server
func InjectDB(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

// USER
//to0k the database users id,name , age , email , phone  //how to take users form the database?
// func getUsers(c *gin.Context){  // in c place we will use the database data
// 	c.IndentedJSON(http.StatusOK, Users)  //this reder our data
// }
//write some functions to get the user
//get from ("/user" , function)
// func main(){
// 	router:=gin.Default()  //this is the defualt gin api labrary
// 	router.GET("/users",getUsers) // this resourse called the function runs
// }

// took the database books id,author,quantity ,
//write some functions to get the book
//jsonfunc handle by json
//get from ("/book" , jsonfunc)
