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

// BubbleChartSeries is a series to be used on a Bubble chart.
type BubbleChartSeries struct {
	x *crt.CT_BubbleSer
}

// X returns the inner wrapped XML type.
func (c BubbleChartSeries) X() *crt.CT_BubbleSer {
	return c.x
}

// InitializeDefaults initializes a Bubble chart series to the default values.
func (c BubbleChartSeries) InitializeDefaults() {
}

// SetText sets the series text.
func (c BubbleChartSeries) SetText(s string) {
	c.x.Tx = crt.NewCT_SerTx()
	c.x.Tx.Choice.V = &s
}

// CategoryAxis returns the category data source.
func (c BubbleChartSeries) CategoryAxis() CategoryAxisDataSource {
	if c.x.XVal == nil {
		c.x.XVal = crt.NewCT_AxDataSource()
	}
	return MakeAxisDataSource(c.x.XVal)
}

// Values returns the value data source.
func (c BubbleChartSeries) Values() NumberDataSource {
	if c.x.YVal == nil {
		c.x.YVal = crt.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(c.x.YVal)
}

// Values returns the bubble size data source.
func (c BubbleChartSeries) BubbleSizes() NumberDataSource {
	if c.x.BubbleSize == nil {
		c.x.BubbleSize = crt.NewCT_NumDataSource()
	}
	return MakeNumberDataSource(c.x.BubbleSize)
}

// Properties returns the Bubble chart series shape properties.
func (c BubbleChartSeries) Properties() drawing.ShapeProperties {
	if c.x.SpPr == nil {
		c.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(c.x.SpPr)
}
