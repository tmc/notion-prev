package notion

import "github.com/tmc/notion/notiontypes"

// Page is a notion.so page.
type Page struct {
	RecordMap notiontypes.RecordMap `json:"recordMap"`
	Cursor    Cursor                `json:"cursor"`
}

// Cursor is used for pagination of entities.
type Cursor struct {
	Stack []interface{} `json:"stack"`
}
