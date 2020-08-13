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

	"github.com/xxiiaass/unidoc2"
)

type TblStyleLst struct {
	CT_TableStyleList
}

func NewTblStyleLst() *TblStyleLst {
	ret := &TblStyleLst{}
	ret.CT_TableStyleList = *NewCT_TableStyleList()
	return ret
}

func (m *TblStyleLst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "a:tblStyleLst"
	return m.CT_TableStyleList.MarshalXML(e, start)
}

func (m *TblStyleLst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_TableStyleList = *NewCT_TableStyleList()
	for _, attr := range start.Attr {
		if attr.Name.Local == "def" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DefAttr = parsed
			continue
		}
	}
lTblStyleLst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tblStyle"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "tblStyle"}:
				tmp := NewCT_TableStyle()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TblStyle = append(m.TblStyle, tmp)
			default:
				unioffice.Log("skipping unsupported element on TblStyleLst %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lTblStyleLst
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the TblStyleLst and its children
func (m *TblStyleLst) Validate() error {
	return m.ValidateWithPath("TblStyleLst")
}

// ValidateWithPath validates the TblStyleLst and its children, prefixing error messages with path
func (m *TblStyleLst) ValidateWithPath(path string) error {
	if err := m.CT_TableStyleList.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
