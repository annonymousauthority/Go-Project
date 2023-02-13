package main

import (
	"context"
	"fmt"
	"log"
	initializer "starter_pack/initializer/initialize"
	"strings"

	"tryporpra/repo/models"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/iterator"
)

var iDs []models.IDs
var userData models.Users

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func listUsers(ctx context.Context, client *auth.Client) {
	pager := iterator.NewPager(client.Users(ctx, ""), 500, "")

	for {
		var users []*auth.ExportedUserRecord
		nextPageToken, err := pager.NextPage(&users)
		if err != nil {
			log.Fatalf("paging error %v\n", err)
		}
		for _, u := range users {
			iDs = append(iDs, models.IDs{u.UserRecord.UserInfo.UID})
			// log.Printf(" user: %v\n", string(u.UserRecord.UserInfo.DisplayName))
		}
		if nextPageToken == "" {
			break
		}
	}
	// fmt.Println(iDs)
}

func getUserDataByID(ctx context.Context, app *firebase.App, id models.IDs) models.Users {
	dbClient, dbErr := app.Firestore(ctx)
	defer dbClient.Close()
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	dataSnap, err := dbClient.Collection("USERS").Doc(id.Id).Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return models.Users{id.Id, dataSnap.Data()["generatedResume"].(string), dataSnap.Data()["fullname"].(string), dataSnap.Data()["email"].(string)}
}

func breakDownResume(resume string) (string, []string, []string, []string) {
	sections := strings.Split(resume, "\n")
	sections = delete_empty(sections)
	var objective string
	var skills []string
	var workExperience []string
	var education []string

	/*
		the resume is thought to traverse through objective -> skills -> work experience -> education
		or objective ->work experience -> skills -> education

		We know both the start and the end of the algorithm traversal.

		So the algorithm knows the prefixes to check for.

		The resume has to be cleaned to remove any empty strings from the array.
		if prefix is Objective then the algorithm should extract the next index as objective on the assumption that that index is not empty

		Approach one.
		-> Get the index for all target prefixes. Objectives -> Skills -> Work Experience -> Education
		-> Search through and extract every information between as information for each prefix.

		Question: Can concurrency work in extracting this value?
	*/

	// Approach one

	for i, information := range sections {
		if information == "" {
			continue
		} else if strings.HasPrefix(information, "Objective") {
			objective = sections[i+1]
			continue
		} else if strings.HasPrefix(information, "Skills") {
			for j := i + 1; j < len(sections); j++ {
				if !strings.HasPrefix(sections[j], "Work Experience") && !strings.HasPrefix(sections[j], "Education") {
					skills = append(skills, sections[j])
				} else {
					break
				}
			}
		} else if strings.HasPrefix(information, "Work Experience") {
			for j := 1; j <= len(sections); j++ {
				if !strings.HasPrefix(sections[i+j], "Skills") && !strings.HasPrefix(sections[i+j], "Education") {
					workExperience = append(workExperience, sections[i+j])
				} else {
					fmt.Println("Algorithm Stopped")
					break
				}
			}
		} else if strings.HasPrefix(information, "Education") {
			for j := i + 1; j < len(sections); j++ {
				if !strings.HasPrefix(sections[j], "Work Experience") && !strings.HasPrefix(sections[j], "Skills") {
					education = append(education, sections[j])
				} else {
					break
				}
			}
		}
	}

	fmt.Println("Objective")
	fmt.Println(objective)
	fmt.Println("Work Experience")
	fmt.Println(strings.Join(workExperience, "\n"))
	fmt.Println("Skills")
	fmt.Println(strings.Join(skills, "\n"))
	fmt.Println("Education")
	fmt.Println(strings.Join(education, "\n"))

	return objective, skills, workExperience, education
}

func main() {
	fmt.Println("This is main function")
	ctx := context.Background()
	app := initializer.Initialize()
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatal(err)
	}

	listUsers(ctx, client)

	userData = getUserDataByID(ctx, app, iDs[len(iDs)-37])
	_, _, _, _ = breakDownResume(userData.Resume)
}
