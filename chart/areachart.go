// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package chart

import crt "github.com/unidoc/unioffice/schema/soo/dml/chart"

// AreaChart is an area chart that has a shaded area underneath a curve.
type AreaChart struct {
	chartBase
	x *crt.CT_AreaChart
}

// X returns the inner wrapped XML type.
func (c AreaChart) X() *crt.CT_AreaChart {
	return c.x
}

// InitializeDefaults the bar chart to its defaults
func (c AreaChart) InitializeDefaults() {
}

// AddSeries adds a default series to an area chart.
func (c AreaChart) AddSeries() AreaChartSeries {
	clr := c.nextColor(len(c.x.Ser))
	ser := crt.NewCT_AreaSer()
	c.x.Ser = append(c.x.Ser, ser)
	ser.Idx.ValAttr = uint32(len(c.x.Ser) - 1)
	ser.Order.ValAttr = uint32(len(c.x.Ser) - 1)

	bs := AreaChartSeries{ser}
	bs.InitializeDefaults()
	bs.Properties().SetSolidFill(clr)
	return bs
}

func (c AreaChart) AddAxis(axis Axis) {
	axisID := crt.NewCT_UnsignedInt()
	axisID.ValAttr = axis.AxisID()
	c.x.AxId = append(c.x.AxId, axisID)
}