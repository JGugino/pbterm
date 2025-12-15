package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JGugino/pbterm/pb"
	"github.com/joho/godotenv"
)

var authToken string = ""

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pbURL := "http://127.0.0.1:8090"

	pbAuth := pb.PBAuth{
		BaseURL: pbURL,
	}

	pbCollection := pb.PBCollection{
		BaseURL: pbURL,
	}

	email := os.Getenv("TESTING_EMAIL")
	password := os.Getenv("TESTING_PASSWORD")

	success, err := pbAuth.AuthWithPasswordForCollection("_superusers", "", "", email, password)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	authToken = success.Token
	fmt.Printf("User logged in (%s) \n", email)

	// options := pb.CollectionOptions{
	// 	Name:   "testing_collection",
	// 	Type:   pb.BaseCollection,
	// 	System: true,
	// 	Fields: []map[string]any{
	// 		{
	// 			"name":     "title",
	// 			"type":     "text",
	// 			"required": true,
	// 		},
	// 	},
	// }

	// collection, err := pbCollection.CreateNewCollection(authToken, options)

	_, err = pbCollection.DeleteCollection(authToken, "testing_collection")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
