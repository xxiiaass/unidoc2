// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package dml

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type CT_GvmlShapeNonVisual struct {
	CNvPr   *CT_NonVisualDrawingProps
	CNvSpPr *CT_NonVisualDrawingShapeProps
}

func NewCT_GvmlShapeNonVisual() *CT_GvmlShapeNonVisual {
	ret := &CT_GvmlShapeNonVisual{}
	ret.CNvPr = NewCT_NonVisualDrawingProps()
	ret.CNvSpPr = NewCT_NonVisualDrawingShapeProps()
	return ret
}

func (m *CT_GvmlShapeNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "a:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvSpPr := xml.StartElement{Name: xml.Name{Local: "a:cNvSpPr"}}
	e.EncodeElement(m.CNvSpPr, secNvSpPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GvmlShapeNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = NewCT_NonVisualDrawingProps()
	m.CNvSpPr = NewCT_NonVisualDrawingShapeProps()
lCT_GvmlShapeNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cNvPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cNvPr"}:
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cNvSpPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cNvSpPr"}:
				if err := d.DecodeElement(m.CNvSpPr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_GvmlShapeNonVisual %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GvmlShapeNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GvmlShapeNonVisual and its children
func (m *CT_GvmlShapeNonVisual) Validate() error {
	return m.ValidateWithPath("CT_GvmlShapeNonVisual")
}

// ValidateWithPath validates the CT_GvmlShapeNonVisual and its children, prefixing error messages with path
func (m *CT_GvmlShapeNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvSpPr.ValidateWithPath(path + "/CNvSpPr"); err != nil {
		return err
	}
	return nil
}