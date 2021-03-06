// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package chart

import (
	"github.com/xxiiaass/unidoc2"
	"github.com/xxiiaass/unidoc2/drawing"
	"github.com/xxiiaass/unidoc2/schema/soo/dml"
	crt "github.com/xxiiaass/unidoc2/schema/soo/dml/chart"
)

type Legend struct {
	x *crt.CT_Legend
}

func MakeLegend(l *crt.CT_Legend) Legend {
	return Legend{l}
}

// X returns the inner wrapped XML type.
func (l Legend) X() *crt.CT_Legend {
	return l.x
}
func (l Legend) InitializeDefaults() {
	l.SetPosition(crt.ST_LegendPosR)
	l.SetOverlay(false)
	l.Properties().SetNoFill()
	l.Properties().LineProperties().SetNoFill()
}

func (l Legend) SetPosition(p crt.ST_LegendPos) {
	if p == crt.ST_LegendPosUnset {
		l.x.LegendPos = nil
	} else {
		l.x.LegendPos = crt.NewCT_LegendPos()
		l.x.LegendPos.ValAttr = p
	}
}

func (l Legend) SetOverlay(b bool) {
	l.x.Overlay = crt.NewCT_Boolean()
	l.x.Overlay.ValAttr = unioffice.Bool(b)
}

func (l Legend) Properties() drawing.ShapeProperties {
	if l.x.SpPr == nil {
		l.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(l.x.SpPr)
}
