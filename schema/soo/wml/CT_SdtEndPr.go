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

	"github.com/unidoc/unioffice"
)

type CT_SdtEndPr struct {
	// Structured Document Tag End Character Run Properties
	RPr []*CT_RPr
}

func NewCT_SdtEndPr() *CT_SdtEndPr {
	ret := &CT_SdtEndPr{}
	return ret
}

func (m *CT_SdtEndPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		for _, c := range m.RPr {
			e.EncodeElement(c, serPr)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SdtEndPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SdtEndPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "rPr"}:
				tmp := NewCT_RPr()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.RPr = append(m.RPr, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_SdtEndPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SdtEndPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SdtEndPr and its children
func (m *CT_SdtEndPr) Validate() error {
	return m.ValidateWithPath("CT_SdtEndPr")
}

// ValidateWithPath validates the CT_SdtEndPr and its children, prefixing error messages with path
func (m *CT_SdtEndPr) ValidateWithPath(path string) error {
	for i, v := range m.RPr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/RPr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
