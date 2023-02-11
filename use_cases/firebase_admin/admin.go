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
	// [START list_all_users_golang]
	// Note, behind the scenes, the Users() iterator will retrive 1000 Users at a time through the API
	iter := client.Users(ctx, "")
	for {
		user, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("error listing users: %s\n", err)
		}
		log.Printf("read user user: %v\n", user)
	}

	// Iterating by pages 100 users at a time.
	// Note that using both the Next() function on an iterator and the NextPage()
	// on a Pager wrapping that same iterator will result in an error.
	pager := iterator.NewPager(client.Users(ctx, ""), 100, "")
	for {
		var users []*auth.ExportedUserRecord
		nextPageToken, err := pager.NextPage(&users)
		if err != nil {
			log.Fatalf("paging error %v\n", err)
		}
		for _, u := range users {
			log.Printf("read user user: %v\n", u)
		}
		if nextPageToken == "" {
			break
		}
	}
	// [END list_all_users_golang]
}
