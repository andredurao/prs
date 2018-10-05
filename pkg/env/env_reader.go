package envReader

import (
	"os"
)

func GithubToken() string {
	return os.Getenv("GITHUB_AUTH_TOKEN")
}
