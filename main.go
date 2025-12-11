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

	options := pb.PocketBaseListOptions{
		Page:    1,
		PerPage: 10,
		Sort:    "-created",
		Filter:  "",
	}

	records, err := pbRecord.ListRecords("posts", authToken, options)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(records)

	// recordData := map[string]any{
	// 	"title":   "Example Post",
	// 	"content": "This is some example content",
	// }

	// record, err := pbRecord.CreateNewRecord("posts", authToken, recordData)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// updatedData := map[string]any{
	// 	"content": "This is a modified content for the example post",
	// }

	// record, err = pbRecord.UpdateRecord("posts", record["id"].(string), authToken, updatedData)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Printf("Record Updated (%s)\n", record["id"])
}
