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
)

type CT_MdxKPI struct {
	// Member Unique Name Index
	NAttr uint32
	// KPI Index
	NpAttr uint32
	// KPI Property
	PAttr ST_MdxKPIProperty
}

func NewCT_MdxKPI() *CT_MdxKPI {
	ret := &CT_MdxKPI{}
	ret.PAttr = ST_MdxKPIProperty(1)
	return ret
}

func (m *CT_MdxKPI) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "n"},
		Value: fmt.Sprintf("%v", m.NAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "np"},
		Value: fmt.Sprintf("%v", m.NpAttr)})
	attr, err := m.PAttr.MarshalXMLAttr(xml.Name{Local: "p"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MdxKPI) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PAttr = ST_MdxKPIProperty(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "n" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.NAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "np" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.NpAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "p" {
			m.PAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_MdxKPI: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_MdxKPI and its children
func (m *CT_MdxKPI) Validate() error {
	return m.ValidateWithPath("CT_MdxKPI")
}

// ValidateWithPath validates the CT_MdxKPI and its children, prefixing error messages with path
func (m *CT_MdxKPI) ValidateWithPath(path string) error {
	if m.PAttr == ST_MdxKPIPropertyUnset {
		return fmt.Errorf("%s/PAttr is a mandatory field", path)
	}
	if err := m.PAttr.ValidateWithPath(path + "/PAttr"); err != nil {
		return err
	}
	return nil
}