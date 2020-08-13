// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package chart

import (
	"github.com/xxiiaass/unidoc2"
	crt "github.com/xxiiaass/unidoc2/schema/soo/dml/chart"
)

// StockChart is a 2D Stock chart.
type StockChart struct {
	chartBase
	x *crt.CT_StockChart
}

// X returns the inner wrapped XML type.
func (c StockChart) X() *crt.CT_StockChart {
	return c.x
}

// InitializeDefaults the Stock chart to its defaults
func (c StockChart) InitializeDefaults() {
	c.x.HiLowLines = crt.NewCT_ChartLines()
	c.x.UpDownBars = crt.NewCT_UpDownBars()
	c.x.UpDownBars.GapWidth = crt.NewCT_GapAmount()
	c.x.UpDownBars.GapWidth.ValAttr = &crt.ST_GapAmount{}
	c.x.UpDownBars.GapWidth.ValAttr.ST_GapAmountUShort = unioffice.Uint16(150)
	c.x.UpDownBars.UpBars = crt.NewCT_UpDownBar()
	c.x.UpDownBars.DownBars = crt.NewCT_UpDownBar()
}

// AddSeries adds a default series to a Stock chart.
func (c StockChart) AddSeries() LineChartSeries {
	ser := crt.NewCT_LineSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	bs := LineChartSeries{ser}
	bs.Values().CreateEmptyNumberCache()
	// don't use defaults as the stock chart needs special
	// formatting
	bs.Properties().LineProperties().SetNoFill()
	return bs
}

func (c StockChart) AddAxis(axis Axis) {
	axisID := crt.NewCT_UnsignedInt()
	axisID.ValAttr = axis.AxisID()
	c.x.AxId = append(c.x.AxId, axisID)
}
