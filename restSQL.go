package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"	
	"fmt"
)

var db *gorm.DB
var err error

type User struct{
	gorm.Model `json:"model"`	
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

type Msg struct{
	Message string `json:"message"`
}

func initiaizeDB(){
	db, err = gorm.Open("sqlite3","users.db")
	if err != nil {
		fmt.Println("Could not create a database!",err.Error())
	} 
	db.AutoMigrate(&User{})
	// defer db.Close()
}

func listUsers(c *gin.Context){
	var users []User	 
	err := db.Find(&users)
	if err.Error != nil {
	    fmt.Println(err.Error)
	}
    c.JSON(200, users)	
}

func getUser(c *gin.Context,){
	var users []User
	id := c.Param("id")
	err := db.Find(&users, "id = ?", id)
	if err.Error != nil {
	    fmt.Println(err.Error)
	}	
	if len(users) == 0 {
	    c.AbortWithStatus(404)		
	} else {		
	    c.JSON(200, users)	
	}
}

func deleteUser(c *gin.Context,){
	var users []User
	id := c.Param("id")
	err := db.Delete(&users, id)
	if err.Error != nil {
	    fmt.Println(err.Error)
	}	
	msg := Msg{"User deleted successfully!"}
    c.JSON(200, msg)		
}

func createUser(c *gin.Context){
	var user User
	c.BindJSON(&user)
	db.Create(&user)
	c.JSON(200, user)
}

func updateUser(c *gin.Context){
	var user User	
	id := c.Param("id")
	c.BindJSON(&user)	
	db.Model(&user).Where("id = ?", id).Update(&user)	
	c.JSON(200, user)
}

func main(){
	router := gin.Default()
	initiaizeDB()
	router.GET("/users/",listUsers)
	router.POST("/users/",createUser)
	router.PUT("/users/:id",updateUser)
	router.GET("/users/:id",getUser)
	router.DELETE("/users/:id",deleteUser)
	router.Run()	
}