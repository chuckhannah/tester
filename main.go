package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	token64, err := ioutil.ReadFile("token")
	token_bytes, _ := base64.StdEncoding.DecodeString(string(token64))
	token := strings.TrimSpace(string(token_bytes))

	check(err)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repos, _, err := client.Repositories.List(ctx, "", nil)
	check(err)
	fmt.Print(repos)
}
