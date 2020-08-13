// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package sml

import (
	"encoding/xml"

	"github.com/xxiiaass/unidoc2"
)

type Chartsheet struct {
	CT_Chartsheet
}

func NewChartsheet() *Chartsheet {
	ret := &Chartsheet{}
	ret.CT_Chartsheet = *NewCT_Chartsheet()
	return ret
}

func (m *Chartsheet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "ma:chartsheet"
	return m.CT_Chartsheet.MarshalXML(e, start)
}

func (m *Chartsheet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Chartsheet = *NewCT_Chartsheet()
lChartsheet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "sheetPr"}:
				m.SheetPr = NewCT_ChartsheetPr()
				if err := d.DecodeElement(m.SheetPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetViews"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "sheetViews"}:
				if err := d.DecodeElement(m.SheetViews, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetProtection"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "sheetProtection"}:
				m.SheetProtection = NewCT_ChartsheetProtection()
				if err := d.DecodeElement(m.SheetProtection, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "customSheetViews"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "customSheetViews"}:
				m.CustomSheetViews = NewCT_CustomChartsheetViews()
				if err := d.DecodeElement(m.CustomSheetViews, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pageMargins"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "pageMargins"}:
				m.PageMargins = NewCT_PageMargins()
				if err := d.DecodeElement(m.PageMargins, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pageSetup"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "pageSetup"}:
				m.PageSetup = NewCT_CsPageSetup()
				if err := d.DecodeElement(m.PageSetup, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "headerFooter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "headerFooter"}:
				m.HeaderFooter = NewCT_HeaderFooter()
				if err := d.DecodeElement(m.HeaderFooter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "drawing"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "drawing"}:
				if err := d.DecodeElement(m.Drawing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "legacyDrawing"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "legacyDrawing"}:
				m.LegacyDrawing = NewCT_LegacyDrawing()
				if err := d.DecodeElement(m.LegacyDrawing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "legacyDrawingHF"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "legacyDrawingHF"}:
				m.LegacyDrawingHF = NewCT_LegacyDrawing()
				if err := d.DecodeElement(m.LegacyDrawingHF, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "drawingHF"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "drawingHF"}:
				m.DrawingHF = NewCT_DrawingHF()
				if err := d.DecodeElement(m.DrawingHF, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "picture"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "picture"}:
				m.Picture = NewCT_SheetBackgroundPicture()
				if err := d.DecodeElement(m.Picture, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "webPublishItems"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "webPublishItems"}:
				m.WebPublishItems = NewCT_WebPublishItems()
				if err := d.DecodeElement(m.WebPublishItems, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on Chartsheet %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lChartsheet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Chartsheet and its children
func (m *Chartsheet) Validate() error {
	return m.ValidateWithPath("Chartsheet")
}

// ValidateWithPath validates the Chartsheet and its children, prefixing error messages with path
func (m *Chartsheet) ValidateWithPath(path string) error {
	if err := m.CT_Chartsheet.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
