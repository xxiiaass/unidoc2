// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package chart

import (
	"github.com/xxiiaass/unidoc2"
	"github.com/xxiiaass/unidoc2/color"
	"github.com/xxiiaass/unidoc2/drawing"
	"github.com/xxiiaass/unidoc2/schema/soo/dml"
	crt "github.com/xxiiaass/unidoc2/schema/soo/dml/chart"
)

// PieOfPieChart is a Pie chart with an extra Pie chart.
type PieOfPieChart struct {
	chartBase
	x *crt.CT_OfPieChart
}

// X returns the inner wrapped XML type.
func (c PieOfPieChart) X() *crt.CT_OfPieChart {
	return c.x
}

// InitializeDefaults the bar chart to its defaults
func (c PieOfPieChart) InitializeDefaults() {
	c.x.VaryColors = crt.NewCT_Boolean()
	c.x.VaryColors.ValAttr = unioffice.Bool(true)
	c.SetType(crt.ST_OfPieTypePie)
	c.x.SecondPieSize = crt.NewCT_SecondPieSize()
	c.x.SecondPieSize.ValAttr = &crt.ST_SecondPieSize{}
	c.x.SecondPieSize.ValAttr.ST_SecondPieSizeUShort = unioffice.Uint16(75)
	cl := crt.NewCT_ChartLines()
	cl.SpPr = dml.NewCT_ShapeProperties()
	sp := drawing.MakeShapeProperties(cl.SpPr)
	sp.LineProperties().SetSolidFill(color.Auto)
	c.x.SerLines = append(c.x.SerLines, cl)
}

// SetType sets the type the secone pie to either pie or bar
func (c PieOfPieChart) SetType(t crt.ST_OfPieType) {
	c.x.OfPieType.ValAttr = t
}

// AddSeries adds a default series to an Pie chart.
func (c PieOfPieChart) AddSeries() PieChartSeries {
	ser := crt.NewCT_PieSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	bs := PieChartSeries{ser}
	bs.InitializeDefaults()
	return bs
}
