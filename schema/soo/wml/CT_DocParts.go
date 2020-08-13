// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/xxiiaass/unidoc2"
)

type CT_DocParts struct {
	// Glossary Document Entry
	DocPart []*CT_DocPart
}

func NewCT_DocParts() *CT_DocParts {
	ret := &CT_DocParts{}
	return ret
}

func (m *CT_DocParts) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.DocPart != nil {
		sedocPart := xml.StartElement{Name: xml.Name{Local: "w:docPart"}}
		for _, c := range m.DocPart {
			e.EncodeElement(c, sedocPart)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocParts) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DocParts:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "docPart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "docPart"}:
				tmp := NewCT_DocPart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DocPart = append(m.DocPart, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_DocParts %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DocParts
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DocParts and its children
func (m *CT_DocParts) Validate() error {
	return m.ValidateWithPath("CT_DocParts")
}

// ValidateWithPath validates the CT_DocParts and its children, prefixing error messages with path
func (m *CT_DocParts) ValidateWithPath(path string) error {
	for i, v := range m.DocPart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DocPart[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
