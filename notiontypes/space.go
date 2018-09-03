package notiontypes

// SpaceWithRole holds a user's role associated with a space and a space.
type SpaceWithRole struct {
	Role  string `json:"role,omitempty"`
	Value *Space `json:"value,omitempty"`
}

// Space is a notion.so workspace.
type Space struct {
	ID          string        `json:"id"`
	Version     float64       `json:"version"`
	Name        string        `json:"name"`
	BetaEnabled bool          `json:"beta_enabled"`
	Permissions *[]Permission `json:"permissions,omitempty"`
	Pages       []string      `json:"pages,omitempty"`
}
