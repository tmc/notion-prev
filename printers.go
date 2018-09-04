package notion

import (
	"bytes"
	"fmt"

	"github.com/tmc/notion/notiontypes"
)

type vimPrinter struct {
	buf      *bytes.Buffer
	indent   string
	indentBy string
}

func (v *vimPrinter) W(s string) {
	v.w(s)
	v.w("\n")
}

func (v *vimPrinter) w(s string) {
	v.buf.WriteString(v.indent)
	v.buf.WriteString(s)
}

func (v *vimPrinter) print(block *notiontypes.Block) error {
	v.W(fmt.Sprintf("%v %v %v {{{", block.Title, block.Type, block.ID))
	v.incIndent()
	//spew.Dump(block)
	for _, b := range block.InlineContent {
		v.W(b.Text)
	}
	for _, b := range block.Content {
		v.print(b)
	}
	v.decIndent()
	v.W("}}}")
	return nil
}

func (v *vimPrinter) incIndent() {
	v.indent += v.indentBy
}

func (v *vimPrinter) decIndent() {
	v.indent = string(v.indent[:len(v.indent)-len(v.indentBy)])
}

// PrintAsVim renders a notion block as a vim block.
func PrintAsVim(block *notiontypes.Block, indent string) ([]byte, error) {
	v := &vimPrinter{buf: new(bytes.Buffer), indentBy: indent}
	err := v.print(block)
	return v.buf.Bytes(), err
}
