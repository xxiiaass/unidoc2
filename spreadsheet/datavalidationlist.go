// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package spreadsheet

import (
	"strings"

	"github.com/xxiiaass/unidoc2"
	"github.com/xxiiaass/unidoc2/schema/soo/sml"
)

// DataValidationList is just a view on a DataValidation configured as a list.
// It presents a drop-down combo box for spreadsheet users to select values. The
// contents of the dropdown can either pull from a rang eof cells (SetRange) or
// specified directly (SetValues).
type DataValidationList struct {
	x *sml.CT_DataValidation
}

// SetRange sets the range that contains the possible values. This is incompatible with SetValues.
func (d DataValidationList) SetRange(cellRange string) {
	d.x.Formula1 = unioffice.String(cellRange)
	d.x.Formula2 = unioffice.String("0")
}

// SetValues sets the possible values. This is incompatible with SetRange.
func (d DataValidationList) SetValues(values []string) {
	d.x.Formula1 = unioffice.String("\"" + strings.Join(values, ",") + "\"")
	d.x.Formula2 = unioffice.String("0")
}
