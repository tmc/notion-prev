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

const (
	// PermissionTypeUser describes permissions for a user
	PermissionTypeUser = "user_permission"
	// PermissionTypePublic describes permissions for public
	PermissionTypePublic = "public_permission"
)

const (
	// BlockPage is a notion Page
	BlockPage = "page"
	// BlockText is a text block
	BlockText = "text"
	// BlockBookmark is a bookmark block
	BlockBookmark = "bookmark"
	// BlockGist is a gist block
	BlockGist = "gist"
	// BlockBulletedList is a bulleted list block
	BlockBulletedList = "bulleted_list"
	// BlockNumberedList is a numbered list block
	BlockNumberedList = "numbered_list"
	// BlockToggle is a toggle block
	BlockToggle = "toggle"
	// BlockTodo is a todo block
	BlockTodo = "to_do"
	// BlockDivider is a divider block
	BlockDivider = "divider"
	// BlockImage is an image block
	BlockImage = "image"
	// BlockHeader is a header block
	BlockHeader = "header"
	// BlockSubHeader is a header block
	BlockSubHeader = "sub_header"
	// BlockQuote is a quote block
	BlockQuote = "quote"
	// BlockComment is a comment block
	BlockComment = "comment"
	// BlockCode is a code block
	BlockCode = "code"
	// BlockColumnList is for multi-column. Number of columns is
	// number of content blocks of type TypeColumn
	BlockColumnList = "column_list"
	// BlockColumn is a child of TypeColumnList
	BlockColumn = "column"
	// BlockTable is a table block
	BlockTable = "table"
	// BlockCollectionView is a collection view block
	BlockCollectionView = "collection_view"
	// BlockVideo is youtube video embed
	BlockVideo = "video"
	// BlockFile is an embedded file
	BlockFile = "file"
)

// for CollectionColumnInfo.Type
const (
	// ColumnMultiSelect is multi-select column
	ColumnMultiSelect = "multi_select"
	ColumnTypeNumber  = "number"
	ColumnTypeTitle   = "title"
	// TODO: text, select, date, person, Files&Media, checkbox, URL, Email, phone
	// formula, relaion, created time, created by, last edited time, last edited by
)

const (
	// TableSpace represents a Notion workspace
	TableSpace = "space"
	// TableBlock represents a Notion block
	TableBlock = "block"
)

const (
	// RoleReader represents a reader
	RoleReader = "reader"
	// RoleEditor represents an editor
	RoleEditor = "editor"
)

const (
	// DateTypeDate represents a date in Date.Type
	DateTypeDate = "date"
	// DateTypeDateTime represents a datetime in Date.Type
	DateTypeDateTime = "datetime"
)
