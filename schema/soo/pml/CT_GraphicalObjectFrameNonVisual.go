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

	"github.com/xxiiaass/unidoc2"
	"github.com/xxiiaass/unidoc2/schema/soo/dml"
)

type CT_GraphicalObjectFrameNonVisual struct {
	// Non-Visual Drawing Properties
	CNvPr *dml.CT_NonVisualDrawingProps
	// Non-Visual Graphic Frame Drawing Properties
	CNvGraphicFramePr *dml.CT_NonVisualGraphicFrameProperties
	// Application Non-Visual Drawing Properties
	NvPr *CT_ApplicationNonVisualDrawingProps
}

func NewCT_GraphicalObjectFrameNonVisual() *CT_GraphicalObjectFrameNonVisual {
	ret := &CT_GraphicalObjectFrameNonVisual{}
	ret.CNvPr = dml.NewCT_NonVisualDrawingProps()
	ret.CNvGraphicFramePr = dml.NewCT_NonVisualGraphicFrameProperties()
	ret.NvPr = NewCT_ApplicationNonVisualDrawingProps()
	return ret
}

func (m *CT_GraphicalObjectFrameNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "p:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvGraphicFramePr := xml.StartElement{Name: xml.Name{Local: "p:cNvGraphicFramePr"}}
	e.EncodeElement(m.CNvGraphicFramePr, secNvGraphicFramePr)
	senvPr := xml.StartElement{Name: xml.Name{Local: "p:nvPr"}}
	e.EncodeElement(m.NvPr, senvPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GraphicalObjectFrameNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = dml.NewCT_NonVisualDrawingProps()
	m.CNvGraphicFramePr = dml.NewCT_NonVisualGraphicFrameProperties()
	m.NvPr = NewCT_ApplicationNonVisualDrawingProps()
lCT_GraphicalObjectFrameNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cNvPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cNvPr"}:
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cNvGraphicFramePr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cNvGraphicFramePr"}:
				if err := d.DecodeElement(m.CNvGraphicFramePr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "nvPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "nvPr"}:
				if err := d.DecodeElement(m.NvPr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_GraphicalObjectFrameNonVisual %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GraphicalObjectFrameNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GraphicalObjectFrameNonVisual and its children
func (m *CT_GraphicalObjectFrameNonVisual) Validate() error {
	return m.ValidateWithPath("CT_GraphicalObjectFrameNonVisual")
}

// ValidateWithPath validates the CT_GraphicalObjectFrameNonVisual and its children, prefixing error messages with path
func (m *CT_GraphicalObjectFrameNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvGraphicFramePr.ValidateWithPath(path + "/CNvGraphicFramePr"); err != nil {
		return err
	}
	if err := m.NvPr.ValidateWithPath(path + "/NvPr"); err != nil {
		return err
	}
	return nil
}
