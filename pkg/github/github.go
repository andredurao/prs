package github

import (
	"context"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"log"
	"os"
)

type PullRequest struct {
	URL         string
	HeadRefName string
	BaseRefName string
}

var query struct {
	Search struct {
		Nodes []struct {
			PullRequest PullRequest `graphql:"... on PullRequest"`
		}
	} `graphql:"search(query: \"is:open is:pr repo:my_repo\", type: ISSUE, first: 90)"`
}

func githubToken() string {
	return os.Getenv("GITHUB_AUTH_TOKEN")
}

func PullRequests() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken()},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		log.Panicln(err)
	}
}
