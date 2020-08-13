// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package picture

import "github.com/xxiiaass/unidoc2"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/picture", "CT_PictureNonVisual", NewCT_PictureNonVisual)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/picture", "CT_Picture", NewCT_Picture)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/picture", "pic", NewPic)
}
