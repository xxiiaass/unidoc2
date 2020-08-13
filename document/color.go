// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package document

import (
	"fmt"

	"github.com/xxiiaass/unidoc2/color"
	"github.com/xxiiaass/unidoc2/schema/soo/wml"
)

// Color controls the run or styles color.
type Color struct {
	x *wml.CT_Color
}

// X returns the inner wrapped XML type.
func (c Color) X() *wml.CT_Color {
	return c.x
}

// SetColor sets a specific color or auto.
func (c Color) SetColor(v color.Color) {
	if v.IsAuto() {
		c.x.ValAttr.ST_HexColorAuto = wml.ST_HexColorAutoAuto
		c.x.ValAttr.ST_HexColorRGB = nil
	} else {
		c.x.ValAttr.ST_HexColorAuto = wml.ST_HexColorAutoUnset
		c.x.ValAttr.ST_HexColorRGB = v.AsRGBString()
	}
}

// SetThemeColor sets the color from the theme.
func (c Color) SetThemeColor(t wml.ST_ThemeColor) {
	c.x.ThemeColorAttr = t
}

// SetThemeShade sets the shade based off the theme color.
func (c Color) SetThemeShade(s uint8) {
	shd := fmt.Sprintf("%02x", s)
	c.x.ThemeShadeAttr = &shd
}
