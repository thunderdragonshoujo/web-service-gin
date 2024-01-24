package main

import (
    "encoding/json"
	"fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    

)

// Contact represents data about a record Contact.
type contact struct {
    Name   string  `json:"name"`
    Address  string  `json:"address"`
    Date string  `json:"date"`
    OrderNumber  string `json:"ordernumber"`
}

var contacts = []contact{
	{Name: "clint", Address: "1752 N Bus", Date: "now", OrderNumber: "12345"},
}

func main() {
    router := gin.Default()
    router.GET("/getcontacts", getContacts)
    router.POST("/contacts", postContacts)
    // same as
    // config := cors.DefaultConfig()
    // config.AllowAllOrigins = true
    // router.Use(cors.New(config))
    router.Use(cors.Default())
    router.Run("localhost:8080")

    
  }

// getContacts responds with the list of all contacts as JSON.
func getContacts(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, contacts)
}

// postContacts adds an contact from JSON received in the request body.
func postContacts(c *gin.Context) {
    var newContact contact

    // Call BindJSON to bind the received JSON to
    // newContact.
    if err := c.BindJSON(&newContact); err != nil {
        return
    }

    // Add the new album to the slice.
    contacts = append(contacts, newContact)
    c.IndentedJSON(http.StatusCreated, newContact)
    fmt.Println("hello")
    jsonStr, err := json.Marshal(newContact)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
    fmt.Println(string(jsonStr))
}