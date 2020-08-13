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
	"strconv"

	"github.com/xxiiaass/unidoc2"
)

type CT_Groups struct {
	// Level Group Count
	CountAttr *uint32
	// OLAP Group
	Group []*CT_LevelGroup
}

func NewCT_Groups() *CT_Groups {
	ret := &CT_Groups{}
	return ret
}

func (m *CT_Groups) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	segroup := xml.StartElement{Name: xml.Name{Local: "ma:group"}}
	for _, c := range m.Group {
		e.EncodeElement(c, segroup)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Groups) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_Groups:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "group"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "group"}:
				tmp := NewCT_LevelGroup()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Group = append(m.Group, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_Groups %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Groups
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Groups and its children
func (m *CT_Groups) Validate() error {
	return m.ValidateWithPath("CT_Groups")
}

// ValidateWithPath validates the CT_Groups and its children, prefixing error messages with path
func (m *CT_Groups) ValidateWithPath(path string) error {
	for i, v := range m.Group {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Group[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
