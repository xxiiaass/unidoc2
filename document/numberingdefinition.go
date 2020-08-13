// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package document

import "github.com/xxiiaass/unidoc2/schema/soo/wml"

// NumberingDefinition defines a numbering definition for a list of pragraphs.
type NumberingDefinition struct {
	x *wml.CT_AbstractNum
}

// X returns the inner wrapped XML type.
func (n NumberingDefinition) X() *wml.CT_AbstractNum {
	return n.x
}

// AbstractNumberID returns the ID that is unique within all numbering
// definitions that is used to assign the definition to a paragraph.
func (n NumberingDefinition) AbstractNumberID() int64 {
	return n.x.AbstractNumIdAttr
}

// Levels returns all of the numbering levels defined in the definition.
func (n NumberingDefinition) Levels() []NumberingLevel {
	ret := []NumberingLevel{}
	for _, nl := range n.x.Lvl {
		ret = append(ret, NumberingLevel{nl})
	}
	return ret
}

// AddLevel adds a new numbering level to a NumberingDefinition.
func (n NumberingDefinition) AddLevel() NumberingLevel {
	nl := wml.NewCT_Lvl()
	nl.Start = &wml.CT_DecimalNumber{ValAttr: 1}
	nl.IlvlAttr = int64(len(n.x.Lvl))
	n.x.Lvl = append(n.x.Lvl, nl)
	return NumberingLevel{nl}
}

// MultiLevelType returns the multilevel type, or ST_MultiLevelTypeUnset if not set.
func (n NumberingDefinition) MultiLevelType() wml.ST_MultiLevelType {
	if n.x.MultiLevelType != nil {
		return n.x.MultiLevelType.ValAttr
	} else {
		return wml.ST_MultiLevelTypeUnset
	}
}

// SetMultiLevelType sets the multilevel type.
func (n NumberingDefinition) SetMultiLevelType(t wml.ST_MultiLevelType) {
	if t == wml.ST_MultiLevelTypeUnset {
		n.x.MultiLevelType = nil
	} else {
		n.x.MultiLevelType = wml.NewCT_MultiLevelType()
		n.x.MultiLevelType.ValAttr = t
	}
}
