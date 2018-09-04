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

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// ResolveBlock populates a block.
func ResolveBlock(block *Block, idToBlock map[string]*Block) error {
	err := parseProperties(block)
	if err != nil {
		return err
	}
	err = parseFormat(block)
	if err != nil {
		return err
	}

	if block.Content != nil || len(block.ContentIDs) == 0 {
		return nil
	}
	n := len(block.ContentIDs)
	block.Content = make([]*Block, n, n)
	notResolved := []int{}
	for i, id := range block.ContentIDs {
		resolved := idToBlock[id]
		if resolved == nil {
			// This can happen e.g. for page fa3fc358e5644f39b89c57f13d426d54
			notResolved = append(notResolved, i)
			//return fmt.Errorf("Couldn't resolve block with id '%s'", id)
			continue
		}
		block.Content[i] = resolved
		ResolveBlock(resolved, idToBlock)
	}
	// remove blocks that are not resolved
	for idx, toRemove := range notResolved {
		i := toRemove - idx
		{
			a := block.ContentIDs
			block.ContentIDs = append(a[:i], a[i+1:]...)
		}
		{
			a := block.Content
			block.Content = append(a[:i], a[i+1:]...)
		}
	}
	return nil
}

func getFirstInline(inline []*InlineBlock) string {
	if len(inline) == 0 {
		return ""
	}
	return inline[0].Text
}

func getFirstInlineBlock(v interface{}) (string, error) {
	inline, err := parseInlineBlocks(v)
	if err != nil {
		return "", err
	}
	return getFirstInline(inline), nil
}

func getProp(block *Block, name string, toSet *string) bool {
	v, ok := block.Properties[name]
	if !ok {
		return false
	}
	s, err := getFirstInlineBlock(v)
	if err != nil {
		return false
	}
	*toSet = s
	return true
}

func parseProperties(block *Block) error {
	var err error
	props := block.Properties

	if title, ok := props["title"]; ok {
		if block.Type == BlockPage {
			block.Title, err = getFirstInlineBlock(title)
		} else if block.Type == BlockCode {
			block.Code, err = getFirstInlineBlock(title)
		} else {
			block.InlineContent, err = parseInlineBlocks(title)
		}
		if err != nil {
			return err
		}
	}

	if BlockTodo == block.Type {
		if checked, ok := props["checked"]; ok {
			s, _ := getFirstInlineBlock(checked)
			// fmt.Printf("checked: '%s'\n", s)
			block.IsChecked = strings.EqualFold(s, "Yes")
		}
	}

	// for BlockBookmark
	getProp(block, "description", &block.Description)
	// for BlockBookmark
	getProp(block, "link", &block.Link)

	// for BlockBookmark, BlockImage, BlockGist, BlockFile
	// don't over-write if was already set from "source" json field
	if block.Source != "" {
		getProp(block, "source", &block.Source)
	}

	if block.Source != "" && block.IsImage() {
		block.ImageURL = makeImageURL(block.Source)
	}

	// for BlockCode
	getProp(block, "language", &block.CodeLanguage)

	// for BlockFile
	if block.Type == BlockFile {
		getProp(block, "size", &block.FileSize)
	}

	return nil
}

// sometimes image url in "source" is not accessible but can
// be accessed when proxied via notion server as
// www.notion.so/image/${source}
// This also allows resizing via ?width=${n} arguments
//
// from: /images/page-cover/met_vincent_van_gogh_cradle.jpg
// =>
// https://www.notion.so/image/https%3A%2F%2Fwww.notion.so%2Fimages%2Fpage-cover%2Fmet_vincent_van_gogh_cradle.jpg?width=3290
func makeImageURL(uri string) string {
	if uri == "" || strings.Contains(uri, "//www.notion.so/image/") {
		return uri
	}
	// if the url has https://, it's already in s3.
	// If not, it's only a relative URL (like those for built-in
	// cover pages)
	if !strings.HasPrefix(uri, "https://") {
		uri = "https://www.notion.so" + uri
	}
	return "https://www.notion.so/image/" + url.PathEscape(uri)
}

func parseFormat(block *Block) error {
	if len(block.FormatRaw) == 0 {
		// TODO: maybe if BlockPage, set to default &FormatPage{}
		return nil
	}
	var err error
	switch block.Type {
	case BlockPage:
		var format FormatPage
		err = json.Unmarshal(block.FormatRaw, &format)
		if err == nil {
			format.PageCoverURL = makeImageURL(format.PageCover)
			block.FormatPage = &format
		}
	case BlockBookmark:
		var format FormatBookmark
		err = json.Unmarshal(block.FormatRaw, &format)
		if err == nil {
			block.FormatBookmark = &format
		}
	case BlockImage:
		var format FormatImage
		err = json.Unmarshal(block.FormatRaw, &format)
		if err == nil {
			format.ImageURL = makeImageURL(format.DisplaySource)
			block.FormatImage = &format
		}
	case BlockColumn:
		var format FormatColumn
		err = json.Unmarshal(block.FormatRaw, &format)
		if err == nil {
			block.FormatColumn = &format
		}
	case BlockTable:
		var format FormatTable
		err = json.Unmarshal(block.FormatRaw, &format)
		if err == nil {
			block.FormatTable = &format
		}
	case BlockText:
		var format FormatText
		err = json.Unmarshal(block.FormatRaw, &format)
		if err == nil {
			block.FormatText = &format
		}
	case BlockVideo:
		var format FormatVideo
		err = json.Unmarshal(block.FormatRaw, &format)
		if err == nil {
			block.FormatVideo = &format
		}
	}

	if err != nil {
		fmt.Printf("parseFormat: json.Unamrshal() failed with '%s', format: '%s'\n", err, string(block.FormatRaw))
		return err
	}
	return nil
}
