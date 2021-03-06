// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package presentation

import (
	"github.com/xxiiaass/unidoc2/drawing"
	"github.com/xxiiaass/unidoc2/schema/soo/dml"
	"github.com/xxiiaass/unidoc2/schema/soo/pml"
)

// TextBox is a text box within a slide.
type TextBox struct {
	x *pml.CT_Shape
}

// AddParagraph adds a paragraph to the text box
func (t TextBox) AddParagraph() drawing.Paragraph {
	p := dml.NewCT_TextParagraph()
	t.x.TxBody.P = append(t.x.TxBody.P, p)
	return drawing.MakeParagraph(p)
}

// Properties returns the properties of the TextBox.
func (t TextBox) Properties() drawing.ShapeProperties {
	if t.x.SpPr == nil {
		t.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(t.x.SpPr)
}

// SetTextAnchor controls the text anchoring
func (t TextBox) SetTextAnchor(a dml.ST_TextAnchoringType) {
	t.x.TxBody.BodyPr = dml.NewCT_TextBodyProperties()
	t.x.TxBody.BodyPr.AnchorAttr = a
}
