package notion_test

import (
	"encoding/json"
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
	json.NewEncoder(os.Stdout).Encode(page)
	// output:
	// .
}

func ExampleNewClient_authenticated() {
	token := os.Getenv("NOTION_TOKEN")
	c, err := notion.NewClient(notion.WithToken(token), notion.WithDebugLogging)
	if err != nil {
		fmt.Println(err)
	}
	_, err = c.GetPage(os.Getenv("NOTION_PAGE_ID"))
	fmt.Println(err)
	// output:
	// <nil>
}
