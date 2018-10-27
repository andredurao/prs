package github

import (
	"context"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"log"
	"os"
)

func githubToken() string {
	return os.Getenv("GITHUB_AUTH_TOKEN")
}

func PullRequests() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken()},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	err := client.Query(context.Background(), &base.Query, nil)
	if err != nil {
		log.Panicln(err)
	}
}
