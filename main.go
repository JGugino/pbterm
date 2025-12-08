package main

import (
	"fmt"

	"github.com/JGugino/pbterm/pb"
)

var authToken string = ""

func main() {

	pbURL := "http://127.0.0.1:8090"

	pbAuth := pb.PBAuth{
		BaseURL: pbURL,
	}

	pbRecord := pb.PBRecord{
		BaseURL: pbURL,
	}

	email := "gugino.inquires@gmail.com"
	password := "password123"

	success, err := pbAuth.AuthWithPasswordForCollection("_superusers", "", "", email, password)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	authToken = success.Token
	fmt.Printf("User logged in (%s) \n", email)

	record, err := pbRecord.ViewRecord("posts", "otm5rz3l1jz0nwo", authToken)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Found Record")
	fmt.Println(record)
}
