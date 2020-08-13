// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice"
)

type CT_Rst struct {
	// Text
	T *string
	// Rich Text Run
	R []*CT_RElt
	// Phonetic Run
	RPh []*CT_PhoneticRun
	// Phonetic Properties
	PhoneticPr *CT_PhoneticPr
}

func NewCT_Rst() *CT_Rst {
	ret := &CT_Rst{}
	return ret
}

func (m *CT_Rst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.T != nil {
		set := xml.StartElement{Name: xml.Name{Local: "ma:t"}}
		unioffice.AddPreserveSpaceAttr(&set, *m.T)
		e.EncodeElement(m.T, set)
	}
	if m.R != nil {
		ser := xml.StartElement{Name: xml.Name{Local: "ma:r"}}
		for _, c := range m.R {
			e.EncodeElement(c, ser)
		}
	}
	if m.RPh != nil {
		serPh := xml.StartElement{Name: xml.Name{Local: "ma:rPh"}}
		for _, c := range m.RPh {
			e.EncodeElement(c, serPh)
		}
	}
	if m.PhoneticPr != nil {
		sephoneticPr := xml.StartElement{Name: xml.Name{Local: "ma:phoneticPr"}}
		e.EncodeElement(m.PhoneticPr, sephoneticPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Rst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Rst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "t"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "t"}:
				m.T = new(string)
				if err := d.DecodeElement(m.T, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "r"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "r"}:
				tmp := NewCT_RElt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.R = append(m.R, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rPh"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "rPh"}:
				tmp := NewCT_PhoneticRun()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.RPh = append(m.RPh, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "phoneticPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "phoneticPr"}:
				m.PhoneticPr = NewCT_PhoneticPr()
				if err := d.DecodeElement(m.PhoneticPr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Rst %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Rst
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Rst and its children
func (m *CT_Rst) Validate() error {
	return m.ValidateWithPath("CT_Rst")
}

// ValidateWithPath validates the CT_Rst and its children, prefixing error messages with path
func (m *CT_Rst) ValidateWithPath(path string) error {
	for i, v := range m.R {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/R[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.RPh {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/RPh[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.PhoneticPr != nil {
		if err := m.PhoneticPr.ValidateWithPath(path + "/PhoneticPr"); err != nil {
			return err
		}
	}
	return nil
}
