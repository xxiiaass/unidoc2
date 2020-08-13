// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package chart

import (
	"encoding/xml"
	"fmt"
)

// ST_LblOffset is a union type
type ST_LblOffset struct {
	ST_LblOffsetPercent *string
	ST_LblOffsetUShort  *uint16
}

func (m *ST_LblOffset) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_LblOffset) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_LblOffsetPercent != nil {
		e.EncodeToken(xml.CharData(*m.ST_LblOffsetPercent))
	}
	if m.ST_LblOffsetUShort != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_LblOffsetUShort)))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_LblOffset) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_LblOffsetPercent != nil {
		mems = append(mems, "ST_LblOffsetPercent")
	}
	if m.ST_LblOffsetUShort != nil {
		mems = append(mems, "ST_LblOffsetUShort")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_LblOffset) String() string {
	if m.ST_LblOffsetPercent != nil {
		return fmt.Sprintf("%v", *m.ST_LblOffsetPercent)
	}
	if m.ST_LblOffsetUShort != nil {
		return fmt.Sprintf("%v", *m.ST_LblOffsetUShort)
	}
	return ""
}