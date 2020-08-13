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
	"fmt"

	"github.com/unidoc/unioffice"
)

type EG_OfficeArtExtensionList struct {
	Ext []*CT_OfficeArtExtension
}

func NewEG_OfficeArtExtensionList() *EG_OfficeArtExtensionList {
	ret := &EG_OfficeArtExtensionList{}
	return ret
}

func (m *EG_OfficeArtExtensionList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Ext != nil {
		seext := xml.StartElement{Name: xml.Name{Local: "a:ext"}}
		for _, c := range m.Ext {
			e.EncodeElement(c, seext)
		}
	}
	return nil
}

func (m *EG_OfficeArtExtensionList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_OfficeArtExtensionList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "ext"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "ext"}:
				tmp := NewCT_OfficeArtExtension()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ext = append(m.Ext, tmp)
			default:
				unioffice.Log("skipping unsupported element on EG_OfficeArtExtensionList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_OfficeArtExtensionList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_OfficeArtExtensionList and its children
func (m *EG_OfficeArtExtensionList) Validate() error {
	return m.ValidateWithPath("EG_OfficeArtExtensionList")
}

// ValidateWithPath validates the EG_OfficeArtExtensionList and its children, prefixing error messages with path
func (m *EG_OfficeArtExtensionList) ValidateWithPath(path string) error {
	for i, v := range m.Ext {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ext[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
