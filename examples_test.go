package notion_test

import (
	"fmt"
	"os"

	"github.com/tmc/notion"
)

var (
	testPageSimple = "aa8fc126-6770-4e83-ad6c-3968dcfc9b82"
)

func ExampleNewClient() {
	c, err := notion.NewClient()
	if err != nil {
		fmt.Println(err)
	}
	page, err := c.GetPage(testPageSimple)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(page.Type)
	fmt.Println(page.Title)
	// output:
	// page
	// le-title
}

func ExampleNewClient_authenticated() {
	token := os.Getenv("NOTION_TOKEN")
	client, err := notion.NewClient(notion.WithToken(token), notion.WithDebugLogging())

	_, _ = client, err
}

func ExampleWithDebugLogging() {
	client, err := notion.NewClient(notion.WithDebugLogging())

	_, _ = client, err
}
