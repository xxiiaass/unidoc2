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

	"github.com/unidoc/unioffice"
)

type CT_Picture struct {
	// Embedded Video
	Movie *CT_Rel
	// Floating Embedded Control
	Control *CT_Control
}

func NewCT_Picture() *CT_Picture {
	ret := &CT_Picture{}
	return ret
}

func (m *CT_Picture) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Movie != nil {
		semovie := xml.StartElement{Name: xml.Name{Local: "w:movie"}}
		e.EncodeElement(m.Movie, semovie)
	}
	if m.Control != nil {
		secontrol := xml.StartElement{Name: xml.Name{Local: "w:control"}}
		e.EncodeElement(m.Control, secontrol)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Picture) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Picture:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "movie"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "movie"}:
				m.Movie = NewCT_Rel()
				if err := d.DecodeElement(m.Movie, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "control"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "control"}:
				m.Control = NewCT_Control()
				if err := d.DecodeElement(m.Control, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Picture %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Picture
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Picture and its children
func (m *CT_Picture) Validate() error {
	return m.ValidateWithPath("CT_Picture")
}

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (m *CT_Picture) ValidateWithPath(path string) error {
	if m.Movie != nil {
		if err := m.Movie.ValidateWithPath(path + "/Movie"); err != nil {
			return err
		}
	}
	if m.Control != nil {
		if err := m.Control.ValidateWithPath(path + "/Control"); err != nil {
			return err
		}
	}
	return nil
}
