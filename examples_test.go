package notion_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tmc/notion"
)

func ExampleNewClient() {
	c, err := notion.NewClient(os.Getenv("NOTION_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}
	// do something with projects, don't ignore errors.
	page, err := c.GetPage(os.Getenv("NOTION_PAGE_ID"))
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(os.Stdout).Encode(page)
	// output:
	// .
}
