// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package extended_properties

import "github.com/unidoc/unioffice"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_Properties", NewCT_Properties)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_VectorVariant", NewCT_VectorVariant)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_VectorLpstr", NewCT_VectorLpstr)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_DigSigBlob", NewCT_DigSigBlob)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "Properties", NewProperties)
}