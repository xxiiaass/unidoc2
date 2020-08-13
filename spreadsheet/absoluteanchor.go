// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package spreadsheet

import (
	"github.com/xxiiaass/unidoc2"
	"github.com/xxiiaass/unidoc2/measurement"
	sd "github.com/xxiiaass/unidoc2/schema/soo/dml/spreadsheetDrawing"
)

// AbsoluteAnchor has a fixed top-left corner in distance units as well as a
// fixed height/width.
type AbsoluteAnchor struct {
	x *sd.CT_AbsoluteAnchor
}

// SetColOffset sets the column offset of the top-left of the image in fixed units.
func (a AbsoluteAnchor) SetColOffset(m measurement.Distance) {
	a.x.Pos.XAttr.ST_CoordinateUnqualified = unioffice.Int64(int64(m / measurement.EMU))
}

// SetRowOffset sets the row offset of the top-left of the image in fixed units.
func (a AbsoluteAnchor) SetRowOffset(m measurement.Distance) {
	a.x.Pos.YAttr.ST_CoordinateUnqualified = unioffice.Int64(int64(m / measurement.EMU))
}

// SetHeight sets the height of the anchored object.
func (a AbsoluteAnchor) SetHeight(h measurement.Distance) {
	a.x.Ext.CyAttr = int64(h / measurement.EMU)
}

// SetWidth sets the width of the anchored object.
func (a AbsoluteAnchor) SetWidth(w measurement.Distance) {
	a.x.Ext.CxAttr = int64(w / measurement.EMU)
}

// SetHeightCells is a no-op.
func (a AbsoluteAnchor) SetHeightCells(int32) {
}

// TopLeft is a no-op.
func (a AbsoluteAnchor) TopLeft() CellMarker {
	return CellMarker{}
}

// BottomRight is a no-op.
func (a AbsoluteAnchor) BottomRight() CellMarker {
	return CellMarker{}
}

// MoveTo is a no-op.
func (a AbsoluteAnchor) MoveTo(x, y int32) {
}

// SetWidthCells is a no-op.
func (a AbsoluteAnchor) SetWidthCells(int32) {
}

// Type returns the type of anchor
func (a AbsoluteAnchor) Type() AnchorType {
	return AnchorTypeAbsolute
}
