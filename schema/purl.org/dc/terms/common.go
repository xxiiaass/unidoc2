// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package terms

import "github.com/xxiiaass/unidoc2"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "LCSH", NewLCSH)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "MESH", NewMESH)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "DDC", NewDDC)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "LCC", NewLCC)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "UDC", NewUDC)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "Period", NewPeriod)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "W3CDTF", NewW3CDTF)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "DCMIType", NewDCMIType)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "IMT", NewIMT)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "URI", NewURI)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "ISO639-2", NewISO639_2)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "RFC1766", NewRFC1766)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "RFC3066", NewRFC3066)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "Point", NewPoint)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "ISO3166", NewISO3166)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "Box", NewBox)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "TGN", NewTGN)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "elementOrRefinementContainer", NewElementOrRefinementContainer)
	unioffice.RegisterConstructor("http://purl.org/dc/terms/", "elementsAndRefinementsGroup", NewElementsAndRefinementsGroup)
}
