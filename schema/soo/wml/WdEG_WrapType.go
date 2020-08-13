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

type WdEG_WrapType struct {
	Choice *WdEG_WrapTypeChoice
}

func NewWdEG_WrapType() *WdEG_WrapType {
	ret := &WdEG_WrapType{}
	return ret
}

func (m *WdEG_WrapType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	return nil
}

func (m *WdEG_WrapType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lWdEG_WrapType:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapNone"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wrapNone"}:
				m.Choice = NewWdEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapNone, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapSquare"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wrapSquare"}:
				m.Choice = NewWdEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapSquare, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapTight"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wrapTight"}:
				m.Choice = NewWdEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapTight, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapThrough"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wrapThrough"}:
				m.Choice = NewWdEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapThrough, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapTopAndBottom"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wrapTopAndBottom"}:
				m.Choice = NewWdEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapTopAndBottom, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on WdEG_WrapType %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdEG_WrapType
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdEG_WrapType and its children
func (m *WdEG_WrapType) Validate() error {
	return m.ValidateWithPath("WdEG_WrapType")
}

// ValidateWithPath validates the WdEG_WrapType and its children, prefixing error messages with path
func (m *WdEG_WrapType) ValidateWithPath(path string) error {
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	return nil
}