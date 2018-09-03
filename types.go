package notion

import "github.com/tmc/notion/notiontypes"

// Page is a notion.so page.
type Page struct {
	*notiontypes.Block
}

// StackPosition refers to a position within a list of entities (usually blocks).
type StackPosition struct {
	ID    string  `json:"id,omitempty"`
	Index float64 `json:"index,omitempty"`
	Table string  `json:"table,omitempty"`
}

// Cursor is used for pagination of entities.
type Cursor struct {
	Stack [][]StackPosition `json:"stack"`
}
