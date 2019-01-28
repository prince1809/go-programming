package main

import (
	"fmt"
	"github.com/prince1809/go-programming/ch4/github"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.555\n", item.Number, item.User.Login, item.Title)
	}
}
