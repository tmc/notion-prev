// BSD 2-Clause License
//
// Copyright (c) 2018, Krzysztof Kowalczyk
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package notiontypes

// RecordMap contains a collections of blocks, a space, users, and collections.
type RecordMap struct {
	Blocks          map[string]*BlockWithRole          `json:"block"`
	Space           map[string]interface{}             `json:"space"` // TODO: figure out the type
	Users           map[string]*notionUserInfo         `json:"notion_user"`
	Collections     map[string]*CollectionWithRole     `json:"collection"`
	CollectionViews map[string]*CollectionViewWithRole `json:"collection_view"`
}

// CollectionViewWithRole describes a role and a collection view
type CollectionViewWithRole struct {
	Role  string          `json:"role"`
	Value *CollectionView `json:"value"`
}

// CollectionView describes a collection
type CollectionView struct {
	ID          string                `json:"id"`
	Alive       bool                  `json:"alive"`
	Format      *CollectionViewFormat `json:"format"`
	Name        string                `json:"name"`
	PageSort    []string              `json:"page_sort"`
	ParentID    string                `json:"parent_id"`
	ParentTable string                `json:"parent_table"`
	Query       *CollectionViewQuery  `json:"query"`
	Type        string                `json:"type"`
	Version     int                   `json:"version"`
}

// CollectionViewFormat describes a fomrat of a collection view
type CollectionViewFormat struct {
	TableProperties []*TableProperty `json:"table_properties"`
	TableWrap       bool             `json:"table_wrap"`
}

// CollectionViewQuery describes a query
type CollectionViewQuery struct {
	Aggregate []*AggregateQuery `json:"aggregate"`
}

// AggregateQuery describes an aggregate query
type AggregateQuery struct {
	AggregationType string `json:"aggregation_type"`
	ID              string `json:"id"`
	Property        string `json:"property"`
	Type            string `json:"type"`
	ViewType        string `json:"view_type"`
}

// CollectionWithRole describes a collection
type CollectionWithRole struct {
	Role  string      `json:"role"`
	Value *Collection `json:"value"`
}

// Collection describes a collection
type Collection struct {
	Alive            bool                             `json:"alive"`
	Format           *CollectionFormat                `json:"format"`
	ID               string                           `json:"id"`
	Name             [][]string                       `json:"name"`
	ParentID         string                           `json:"parent_id"`
	ParentTable      string                           `json:"parent_table"`
	CollectionSchema map[string]*CollectionColumnInfo `json:"schema"`
	Version          int                              `json:"version"`
}

// CollectionFormat describes format of a collection
type CollectionFormat struct {
	CollectionPageProperties []*CollectionPageProperty `json:"collection_page_properties"`
}

// CollectionPageProperty describes properties of a collection
type CollectionPageProperty struct {
	Property string `json:"property"`
	Visible  bool   `json:"visible"`
}

// CollectionColumnInfo describes a info of a collection column
type CollectionColumnInfo struct {
	Name    string                    `json:"name"`
	Options []*CollectionColumnOption `json:"options"`
	Type    string                    `json:"type"`
}

// CollectionColumnOption describes options for a collection column
type CollectionColumnOption struct {
	Color string `json:"color"`
	ID    string `json:"id"`
	Value string `json:"value"`
}

type notionUserInfo struct {
	Role  string `json:"role"`
	Value *User  `json:"value"`
}

// User describes a user
type User struct {
	Email                     string `json:"email"`
	FamilyName                string `json:"family_name"`
	GivenName                 string `json:"given_name"`
	ID                        string `json:"id"`
	Locale                    string `json:"locale"`
	MobileOnboardingCompleted bool   `json:"mobile_onboarding_completed"`
	OnboardingCompleted       bool   `json:"onboarding_completed"`
	ProfilePhoto              string `json:"profile_photo"`
	TimeZone                  string `json:"time_zone"`
	Version                   int    `json:"version"`
}

// Date describes a date
type Date struct {
	// "MMM DD, YYYY", "MM/DD/YYYY", "DD/MM/YYYY", "YYYY/MM/DD", "relative"
	DateFormat string    `json:"date_format"`
	Reminder   *Reminder `json:"reminder,omitempty"`
	// "2018-07-12"
	StartDate string `json:"start_date"`
	// "09:00"
	StartTime *string `json:"start_time,omitempty"`
	// "America/Los_Angeles"
	TimeZone *string `json:"time_zone,omitempty"`
	// "H:mm" for 24hr, not given for 12hr
	TimeFormat *string `json:"time_format,omitempty"`
	// "date", "datetime"
	Type string `json:"type"`
}

// Reminder describes date reminder
type Reminder struct {
	Time  string `json:"time"` // e.g. "09:00"
	Unit  string `json:"unit"` // e.g. "day"
	Value int64  `json:"value"`
}
