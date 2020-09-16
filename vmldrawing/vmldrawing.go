//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package vmldrawing

import (
	_b "encoding/xml"
	_ff "fmt"
	_c "github.com/stasomega1/mylib"
	_ba "github.com/stasomega1/mylib/schema/soo/ofc/sharedTypes"
	_d "github.com/stasomega1/mylib/schema/urn/schemas_microsoft_com/office/excel"
	_a "github.com/stasomega1/mylib/schema/urn/schemas_microsoft_com/vml"
)

// NewCommentDrawing constructs a new comment drawing.
func NewCommentDrawing() *Container {
	_e := NewContainer()
	_e.Layout = _a.NewOfcShapelayout()
	_e.Layout.ExtAttr = _a.ST_ExtEdit
	_e.Layout.Idmap = _a.NewOfcCT_IdMap()
	_e.Layout.Idmap.DataAttr = _c.String("\u0031")
	_e.Layout.Idmap.ExtAttr = _a.ST_ExtEdit
	_e.ShapeType = _a.NewShapetype()
	_e.ShapeType.IdAttr = _c.String("_\u0078\u0030\u0030\u0030\u0030\u005f\u0074\u0032\u0030\u0032")
	_e.ShapeType.CoordsizeAttr = _c.String("2\u0031\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030")
	_e.ShapeType.SptAttr = _c.Float32(202)
	_e.ShapeType.PathAttr = _c.String("\u006d\u0030\u002c0l\u0030\u002c\u0032\u0031\u0036\u0030\u0030\u002c\u00321\u00360\u0030,\u00321\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030\u002c\u0030\u0078\u0065")
	_g := _a.NewEG_ShapeElements()
	_e.ShapeType.EG_ShapeElements = append(_e.ShapeType.EG_ShapeElements, _g)
	_g.Path = _a.NewPath()
	_g.Path.GradientshapeokAttr = _ba.ST_TrueFalseT
	_g.Path.ConnecttypeAttr = _a.OfcST_ConnectTypeRect
	return _e
}

type Container struct {
	Layout    *_a.OfcShapelayout
	ShapeType *_a.Shapetype
	Shape     []*_a.Shape
}

func (_af *Container) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0076"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d:v\u006d\u006c"})
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u006f"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006di\u0063\u0072\u006f\u0073\u006f\u0066t\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u006ff\u0066\u0069\u0063\u0065"})
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c"})
	start.Name.Local = "\u0078\u006d\u006c"
	e.EncodeToken(start)
	if _af.Layout != nil {
		_bdd := _b.StartElement{Name: _b.Name{Local: "\u006f\u003a\u0073\u0068\u0061\u0070\u0065\u006c\u0061\u0079\u006f\u0075\u0074"}}
		e.EncodeElement(_af.Layout, _bdd)
	}
	if _af.ShapeType != nil {
		_cb := _b.StartElement{Name: _b.Name{Local: "v\u003a\u0073\u0068\u0061\u0070\u0065\u0074\u0079\u0070\u0065"}}
		e.EncodeElement(_af.ShapeType, _cb)
	}
	for _, _da := range _af.Shape {
		_cdg := _b.StartElement{Name: _b.Name{Local: "\u0076:\u0073\u0068\u0061\u0070\u0065"}}
		e.EncodeElement(_da, _cdg)
	}
	return e.EncodeToken(_b.EndElement{Name: start.Name})
}
func NewContainer() *Container { return &Container{} }
func (_be *Container) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	_be.Shape = nil
_bge:
	for {
		_ac, _daf := d.Token()
		if _daf != nil {
			return _daf
		}
		switch _dd := _ac.(type) {
		case _b.StartElement:
			switch _dd.Name.Local {
			case "s\u0068\u0061\u0070\u0065\u006c\u0061\u0079\u006f\u0075\u0074":
				_be.Layout = _a.NewOfcShapelayout()
				if _bdg := d.DecodeElement(_be.Layout, &_dd); _bdg != nil {
					return _bdg
				}
			case "\u0073h\u0061\u0070\u0065\u0074\u0079\u0070e":
				_be.ShapeType = _a.NewShapetype()
				if _fa := d.DecodeElement(_be.ShapeType, &_dd); _fa != nil {
					return _fa
				}
			case "\u0073\u0068\u0061p\u0065":
				_gb := _a.NewShape()
				if _gf := d.DecodeElement(_gb, &_dd); _gf != nil {
					return _gf
				}
				_be.Shape = append(_be.Shape, _gb)
			}
		case _b.EndElement:
			break _bge
		}
	}
	return nil
}

// NewCommentShape creates a new comment shape for a given cell index.  The
// indices here are zero based.
func NewCommentShape(col, row int64) *_a.Shape {
	_cd := _a.NewShape()
	_cd.IdAttr = _c.String(_ff.Sprintf("\u0063\u0073\u005f\u0025\u0064\u005f\u0025\u0064", col, row))
	_cd.TypeAttr = _c.String("\u0023\u005f\u00780\u0030\u0030\u0030\u005f\u0074\u0032\u0030\u0032")
	_cd.StyleAttr = _c.String("\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u003a\u0061\u0062\u0073\u006f\u006cu\u0074\u0065\u003b\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065\u0066\u0074:\u0038\u0030\u0070\u0074;\u006d\u0061\u0072\u0067\u0069n-\u0074o\u0070\u003a\u0032pt\u003b\u0077\u0069\u0064\u0074\u0068\u003a1\u0030\u0034\u0070\u0074\u003b\u0068\u0065\u0069\u0067\u0068\u0074\u003a\u0037\u0036\u0070\u0074\u003b\u007a\u002d\u0069\u006e\u0064\u0065x\u003a\u0031\u003bv\u0069\u0073\u0069\u0062\u0069\u006c\u0069t\u0079\u003a\u0068\u0069\u0064\u0064\u0065\u006e")
	_cd.FillcolorAttr = _c.String("\u0023f\u0062\u0066\u0036\u0064\u0036")
	_cd.StrokecolorAttr = _c.String("\u0023e\u0064\u0065\u0061\u0061\u0031")
	_ea := _a.NewEG_ShapeElements()
	_ea.Fill = _a.NewFill()
	_ea.Fill.Color2Attr = _c.String("\u0023f\u0062\u0066\u0065\u0038\u0032")
	_ea.Fill.AngleAttr = _c.Float64(-180)
	_ea.Fill.TypeAttr = _a.ST_FillTypeGradient
	_ea.Fill.Fill = _a.NewOfcFill()
	_ea.Fill.Fill.ExtAttr = _a.ST_ExtView
	_ea.Fill.Fill.TypeAttr = _a.OfcST_FillTypeGradientUnscaled
	_cd.EG_ShapeElements = append(_cd.EG_ShapeElements, _ea)
	_bg := _a.NewEG_ShapeElements()
	_bg.Shadow = _a.NewShadow()
	_bg.Shadow.OnAttr = _ba.ST_TrueFalseT
	_bg.Shadow.ObscuredAttr = _ba.ST_TrueFalseT
	_cd.EG_ShapeElements = append(_cd.EG_ShapeElements, _bg)
	_bd := _a.NewEG_ShapeElements()
	_bd.Path = _a.NewPath()
	_bd.Path.ConnecttypeAttr = _a.OfcST_ConnectTypeNone
	_cd.EG_ShapeElements = append(_cd.EG_ShapeElements, _bd)
	_cdf := _a.NewEG_ShapeElements()
	_cdf.Textbox = _a.NewTextbox()
	_cdf.Textbox.StyleAttr = _c.String("\u006d\u0073\u006f\u002ddi\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u002d\u0061\u006c\u0074\u003a\u0061\u0075t\u006f")
	_cd.EG_ShapeElements = append(_cd.EG_ShapeElements, _cdf)
	_dg := _a.NewEG_ShapeElements()
	_dg.ClientData = _d.NewClientData()
	_dg.ClientData.ObjectTypeAttr = _d.ST_ObjectTypeNote
	_dg.ClientData.MoveWithCells = _ba.ST_TrueFalseBlankT
	_dg.ClientData.SizeWithCells = _ba.ST_TrueFalseBlankT
	_dg.ClientData.Anchor = _c.String("\u0031,\u0020\u0031\u0035\u002c\u0020\u0030\u002c\u0020\u0032\u002c\u00202\u002c\u0020\u0035\u0034\u002c\u0020\u0035\u002c\u0020\u0033")
	_dg.ClientData.AutoFill = _ba.ST_TrueFalseBlankFalse
	_dg.ClientData.Row = _c.Int64(row)
	_dg.ClientData.Column = _c.Int64(col)
	_cd.EG_ShapeElements = append(_cd.EG_ShapeElements, _dg)
	return _cd
}
