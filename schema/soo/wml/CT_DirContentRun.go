// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/xxiiaass/unidoc2"
	"github.com/xxiiaass/unidoc2/schema/soo/ofc/math"
)

type CT_DirContentRun struct {
	// Direction of Embedding
	ValAttr ST_Direction
	// Simple Field
	FldSimple []*CT_SimpleField
	// Hyperlink
	Hyperlink *CT_Hyperlink
	// Anchor for Subdocument Location
	SubDoc               *CT_Rel
	EG_ContentRunContent []*EG_ContentRunContent
}

func NewCT_DirContentRun() *CT_DirContentRun {
	ret := &CT_DirContentRun{}
	return ret
}

func (m *CT_DirContentRun) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAttr != ST_DirectionUnset {
		attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.FldSimple != nil {
		sefldSimple := xml.StartElement{Name: xml.Name{Local: "w:fldSimple"}}
		for _, c := range m.FldSimple {
			e.EncodeElement(c, sefldSimple)
		}
	}
	if m.Hyperlink != nil {
		sehyperlink := xml.StartElement{Name: xml.Name{Local: "w:hyperlink"}}
		e.EncodeElement(m.Hyperlink, sehyperlink)
	}
	if m.SubDoc != nil {
		sesubDoc := xml.StartElement{Name: xml.Name{Local: "w:subDoc"}}
		e.EncodeElement(m.SubDoc, sesubDoc)
	}
	if m.EG_ContentRunContent != nil {
		for _, c := range m.EG_ContentRunContent {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DirContentRun) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_DirContentRun:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "fldSimple"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "fldSimple"}:
				tmp := NewCT_SimpleField()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.FldSimple = append(m.FldSimple, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hyperlink"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "hyperlink"}:
				m.Hyperlink = NewCT_Hyperlink()
				if err := d.DecodeElement(m.Hyperlink, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "subDoc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "subDoc"}:
				m.SubDoc = NewCT_Rel()
				if err := d.DecodeElement(m.SubDoc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXml"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXml"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.CustomXml = NewCT_CustomXmlRun()
				if err := d.DecodeElement(tmpcontentruncontent.CustomXml, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "smartTag"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "smartTag"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.SmartTag = NewCT_SmartTagRun()
				if err := d.DecodeElement(tmpcontentruncontent.SmartTag, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sdt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "sdt"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.Sdt = NewCT_SdtRun()
				if err := d.DecodeElement(tmpcontentruncontent.Sdt, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "dir"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "dir"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.Dir = NewCT_DirContentRun()
				if err := d.DecodeElement(tmpcontentruncontent.Dir, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bdo"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bdo"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.Bdo = NewCT_BdoContentRun()
				if err := d.DecodeElement(tmpcontentruncontent.Bdo, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "r"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "r"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmpcontentruncontent.R = NewCT_R()
				if err := d.DecodeElement(tmpcontentruncontent.R, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "proofErr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "proofErr"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.ProofErr = NewCT_ProofErr()
				if err := d.DecodeElement(tmprunlevelelts.ProofErr, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "permStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "permStart"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermStart = NewCT_PermStart()
				if err := d.DecodeElement(tmprunlevelelts.PermStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "permEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "permEnd"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermEnd = NewCT_Perm()
				if err := d.DecodeElement(tmprunlevelelts.PermEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ins"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "ins"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Ins = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Ins, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "del"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "del"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Del = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Del, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFrom"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveFrom"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveFrom = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveFrom, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveTo"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveTo"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveTo = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveTo, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookmarkStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bookmarkStart"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkStart = NewCT_Bookmark()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookmarkEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bookmarkEnd"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFromRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveFromRangeStart"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFromRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveFromRangeEnd"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveToRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveToRangeStart"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveToRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveToRangeEnd"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "commentRangeStart"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeStart = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "commentRangeEnd"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlInsRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlInsRangeStart"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlInsRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlInsRangeEnd"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlDelRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlDelRangeStart"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlDelRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlDelRangeEnd"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveFromRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlMoveFromRangeStart"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveFromRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlMoveFromRangeEnd"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveToRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlMoveToRangeStart"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveToRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlMoveToRangeEnd"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "oMathPara"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "oMathPara"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMathPara = math.NewOMathPara()
				if err := d.DecodeElement(tmpmathcontent.OMathPara, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "oMath"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "oMath"}:
				tmpcontentruncontent := NewEG_ContentRunContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMath = math.NewOMath()
				if err := d.DecodeElement(tmpmathcontent.OMath, &el); err != nil {
					return err
				}
				m.EG_ContentRunContent = append(m.EG_ContentRunContent, tmpcontentruncontent)
				tmpcontentruncontent.EG_RunLevelElts = append(tmpcontentruncontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			default:
				unioffice.Log("skipping unsupported element on CT_DirContentRun %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DirContentRun
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DirContentRun and its children
func (m *CT_DirContentRun) Validate() error {
	return m.ValidateWithPath("CT_DirContentRun")
}

// ValidateWithPath validates the CT_DirContentRun and its children, prefixing error messages with path
func (m *CT_DirContentRun) ValidateWithPath(path string) error {
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	for i, v := range m.FldSimple {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/FldSimple[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.Hyperlink != nil {
		if err := m.Hyperlink.ValidateWithPath(path + "/Hyperlink"); err != nil {
			return err
		}
	}
	if m.SubDoc != nil {
		if err := m.SubDoc.ValidateWithPath(path + "/SubDoc"); err != nil {
			return err
		}
	}
	for i, v := range m.EG_ContentRunContent {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_ContentRunContent[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
