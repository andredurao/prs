package main

import (
	"fmt"
	"github.com/andredurao/prs/pkg/gui"
	_ "github.com/shurcooL/githubv4"
	_ "golang.org/x/oauth2"
)

func main() {
	fmt.Println("initializing...")
	gui.NewGui()
}
