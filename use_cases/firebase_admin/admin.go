package main

import (
	"context"
	"fmt"
	"log"
	initializer "starter_pack/initializer/initialize"

	"firebase.google.com/go/auth"
	"google.golang.org/api/iterator"
)

func main() {
	fmt.Println("This is main function")
	app := initializer.Initialize()
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	listUsers(ctx, client)
}

func listUsers(ctx context.Context, client *auth.Client) {
	pager := iterator.NewPager(client.Users(ctx, ""), 100, "")
	for {
		var users []*auth.ExportedUserRecord
		nextPageToken, err := pager.NextPage(&users)
		if err != nil {
			log.Fatalf("paging error %v\n", err)
		}
		for _, u := range users {
			log.Printf("pager, read user user: %v\n", string(u.UserRecord.UserInfo.DisplayName))
		}
		if nextPageToken == "" {
			break
		}
	}
}
