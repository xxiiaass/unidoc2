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

type CT_GvmlPicture struct {
	NvPicPr  *CT_GvmlPictureNonVisual
	BlipFill *CT_BlipFillProperties
	SpPr     *CT_ShapeProperties
	Style    *CT_ShapeStyle
	ExtLst   *CT_OfficeArtExtensionList
}

func NewCT_GvmlPicture() *CT_GvmlPicture {
	ret := &CT_GvmlPicture{}
	ret.NvPicPr = NewCT_GvmlPictureNonVisual()
	ret.BlipFill = NewCT_BlipFillProperties()
	ret.SpPr = NewCT_ShapeProperties()
	return ret
}

func (m *CT_GvmlPicture) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	senvPicPr := xml.StartElement{Name: xml.Name{Local: "a:nvPicPr"}}
	e.EncodeElement(m.NvPicPr, senvPicPr)
	seblipFill := xml.StartElement{Name: xml.Name{Local: "a:blipFill"}}
	e.EncodeElement(m.BlipFill, seblipFill)
	sespPr := xml.StartElement{Name: xml.Name{Local: "a:spPr"}}
	e.EncodeElement(m.SpPr, sespPr)
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "a:style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GvmlPicture) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvPicPr = NewCT_GvmlPictureNonVisual()
	m.BlipFill = NewCT_BlipFillProperties()
	m.SpPr = NewCT_ShapeProperties()
lCT_GvmlPicture:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "nvPicPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "nvPicPr"}:
				if err := d.DecodeElement(m.NvPicPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "blipFill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "blipFill"}:
				if err := d.DecodeElement(m.BlipFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "spPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "spPr"}:
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "style"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "style"}:
				m.Style = NewCT_ShapeStyle()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_GvmlPicture %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GvmlPicture
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GvmlPicture and its children
func (m *CT_GvmlPicture) Validate() error {
	return m.ValidateWithPath("CT_GvmlPicture")
}

// ValidateWithPath validates the CT_GvmlPicture and its children, prefixing error messages with path
func (m *CT_GvmlPicture) ValidateWithPath(path string) error {
	if err := m.NvPicPr.ValidateWithPath(path + "/NvPicPr"); err != nil {
		return err
	}
	if err := m.BlipFill.ValidateWithPath(path + "/BlipFill"); err != nil {
		return err
	}
	if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
		return err
	}
	if m.Style != nil {
		if err := m.Style.ValidateWithPath(path + "/Style"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}