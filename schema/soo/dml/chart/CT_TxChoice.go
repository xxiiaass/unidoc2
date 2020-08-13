// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package chart

import (
	"encoding/xml"

	"github.com/xxiiaass/unidoc2"
	"github.com/xxiiaass/unidoc2/schema/soo/dml"
)

type CT_TxChoice struct {
	StrRef *CT_StrRef
	Rich   *dml.CT_TextBody
}

func NewCT_TxChoice() *CT_TxChoice {
	ret := &CT_TxChoice{}
	return ret
}

func (m *CT_TxChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.StrRef != nil {
		sestrRef := xml.StartElement{Name: xml.Name{Local: "c:strRef"}}
		e.EncodeElement(m.StrRef, sestrRef)
	}
	if m.Rich != nil {
		serich := xml.StartElement{Name: xml.Name{Local: "c:rich"}}
		e.EncodeElement(m.Rich, serich)
	}
	return nil
}

func (m *CT_TxChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TxChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "strRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "strRef"}:
				m.StrRef = NewCT_StrRef()
				if err := d.DecodeElement(m.StrRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "rich"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "rich"}:
				m.Rich = dml.NewCT_TextBody()
				if err := d.DecodeElement(m.Rich, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TxChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TxChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TxChoice and its children
func (m *CT_TxChoice) Validate() error {
	return m.ValidateWithPath("CT_TxChoice")
}

// ValidateWithPath validates the CT_TxChoice and its children, prefixing error messages with path
func (m *CT_TxChoice) ValidateWithPath(path string) error {
	if m.StrRef != nil {
		if err := m.StrRef.ValidateWithPath(path + "/StrRef"); err != nil {
			return err
		}
	}
	if m.Rich != nil {
		if err := m.Rich.ValidateWithPath(path + "/Rich"); err != nil {
			return err
		}
	}
	return nil
}
