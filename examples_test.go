package notion_test

import (
	"fmt"
	"os"

	"github.com/tmc/notion"
)

var (
	testPageSimple = "aa8fc12667704e83ad6c3968dcfc9b82"
)

func ExampleNewClient() {
	c, err := notion.NewClient()
	if err != nil {
		fmt.Println(err)
	}
	page, err := c.GetPage(testPageSimple)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(page.RecordMap.Blocks[testPageSimple].Value.Type)
	// output:
	// page
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
