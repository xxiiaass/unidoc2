// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/xxiiaass/unidoc2"
)

type CT_TLTemplate struct {
	// Level
	LvlAttr *uint32
	// Time Node List
	TnLst *CT_TimeNodeList
}

func NewCT_TLTemplate() *CT_TLTemplate {
	ret := &CT_TLTemplate{}
	ret.TnLst = NewCT_TimeNodeList()
	return ret
}

func (m *CT_TLTemplate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LvlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lvl"},
			Value: fmt.Sprintf("%v", *m.LvlAttr)})
	}
	e.EncodeToken(start)
	setnLst := xml.StartElement{Name: xml.Name{Local: "p:tnLst"}}
	e.EncodeElement(m.TnLst, setnLst)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLTemplate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TnLst = NewCT_TimeNodeList()
	for _, attr := range start.Attr {
		if attr.Name.Local == "lvl" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.LvlAttr = &pt
			continue
		}
	}
lCT_TLTemplate:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "tnLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "tnLst"}:
				if err := d.DecodeElement(m.TnLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TLTemplate %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLTemplate
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLTemplate and its children
func (m *CT_TLTemplate) Validate() error {
	return m.ValidateWithPath("CT_TLTemplate")
}

// ValidateWithPath validates the CT_TLTemplate and its children, prefixing error messages with path
func (m *CT_TLTemplate) ValidateWithPath(path string) error {
	if err := m.TnLst.ValidateWithPath(path + "/TnLst"); err != nil {
		return err
	}
	return nil
}
