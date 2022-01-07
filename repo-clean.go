package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	//START CONFIG variables section
	ctx := context.Background()
	opt := &github.RepositoryListByOrgOptions{Type: "private"}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_API")},
	)

	github_org := os.Getenv("GITHUB_ORG")
	timeIntervalMonths, _ := strconv.Atoi(os.Getenv("REPO_MONTHS"))
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	var allRepos []*github.Repository
	timeToday := time.Now().UTC()
	expireDate := timeToday.AddDate(0, -timeIntervalMonths, 0)

	type Repo struct {
		name string
		timeValue string
	}
	var expiredRepoStruct []Repo
	var csvList [][]string
    //END CONFIG variables section


	//the loop below returns all the Private repos for the given org
	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, github_org, opt)
		if err != nil {
			fmt.Println(err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	//The section below finds which repos are old based upon the interval
	fmt.Println("**** THE FOLLOWING REPOS ARE OLDER THAN YOUR DECLARED INTERVAL")
	for _, repo := range allRepos {
		lastCommit := *repo.PushedAt
		archived := *repo.Archived
		expired := lastCommit.Before(expireDate)
		if expired && !archived {
			//now why did I make a struct of structs?...for future use with the auto archive future
			expiredRepoStruct = append(expiredRepoStruct, Repo{*repo.Name,lastCommit.String()})
			csvList = append(csvList,[]string{*repo.Name,lastCommit.String()})
			fmt.Printf("Repo: %s\n", *repo.Name)
			fmt.Printf("Last Push Date: %s\n", lastCommit)
		}
	}

	//outputs the number of repos old based upon the total private ones
	arrayLength := strconv.Itoa(len(expiredRepoStruct))
	fmt.Println(arrayLength + " out of " + strconv.Itoa(len(allRepos)) + " have not been committed to in " + strconv.Itoa(timeIntervalMonths) + " months")



	//Writing to the CSV file below
	csvFile, err := os.Create("expired_repos.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvWrite := csv.NewWriter(csvFile)
	for _,item := range csvList {
		csvWrite.Write(item)
	}
	csvWrite.Flush()
	fmt.Println("CSV file with repos and dates written to 'expired_repos.csv' in root of this project")

}
