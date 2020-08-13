// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package chart

import (
	"github.com/xxiiaass/unidoc2/drawing"
	"github.com/xxiiaass/unidoc2/schema/soo/dml"
	crt "github.com/xxiiaass/unidoc2/schema/soo/dml/chart"
)

// RadarChartSeries is a series to be used on an Radar chart.
type RadarChartSeries struct {
	x *crt.CT_RadarSer
}

// X returns the inner wrapped XML type.
func (c RadarChartSeries) X() *crt.CT_RadarSer {
	return c.x
}

// InitializeDefaults initializes an Radar series to the default values.
func (c RadarChartSeries) InitializeDefaults() {

}

// SetText sets the series text.
func (c RadarChartSeries) SetText(s string) {
	c.x.Tx = crt.NewCT_SerTx()
	c.x.Tx.Choice.V = &s
}

// CategoryAxis returns the category data source.
func (c RadarChartSeries) CategoryAxis() CategoryAxisDataSource {
	if c.x.Cat == nil {
		c.x.Cat = crt.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(c.x.Cat)
}

// Values returns the value data source.
func (c RadarChartSeries) Values() NumberDataSource {
	if c.x.Val == nil {
		c.x.Val = crt.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(c.x.Val)
}

// Properties returns the bar chart series shape properties.
func (c RadarChartSeries) Properties() drawing.ShapeProperties {
	if c.x.SpPr == nil {
		c.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(c.x.SpPr)
}
