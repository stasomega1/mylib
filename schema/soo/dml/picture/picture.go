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

package picture

import (
	_b "encoding/xml"
	_a "github.com/unidoc/unioffice"
	_c "github.com/unidoc/unioffice/schema/soo/dml"
)

func (_bdg *CT_PictureNonVisual) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	e.EncodeToken(start)
	_ba := _b.StartElement{Name: _b.Name{Local: "\u0070i\u0063\u003a\u0063\u004e\u0076\u0050r"}}
	e.EncodeElement(_bdg.CNvPr, _ba)
	_beb := _b.StartElement{Name: _b.Name{Local: "\u0070\u0069\u0063:\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}}
	e.EncodeElement(_bdg.CNvPicPr, _beb)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_PictureNonVisual and its children
func (_cfc *CT_PictureNonVisual) Validate() error {
	return _cfc.ValidateWithPath("\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c")
}

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (_dd *CT_Picture) ValidateWithPath(path string) error {
	if _bd := _dd.NvPicPr.ValidateWithPath(path + "\u002f\u004e\u0076\u0050\u0069\u0063\u0050\u0072"); _bd != nil {
		return _bd
	}
	if _eb := _dd.BlipFill.ValidateWithPath(path + "\u002fB\u006c\u0069\u0070\u0046\u0069\u006cl"); _eb != nil {
		return _eb
	}
	if _de := _dd.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _de != nil {
		return _de
	}
	return nil
}

type CT_Picture struct {
	NvPicPr  *CT_PictureNonVisual
	BlipFill *_c.CT_BlipFillProperties
	SpPr     *_c.CT_ShapeProperties
}

func (_ag *CT_Picture) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	e.EncodeToken(start)
	_ea := _b.StartElement{Name: _b.Name{Local: "p\u0069\u0063\u003a\u006e\u0076\u0050\u0069\u0063\u0050\u0072"}}
	e.EncodeElement(_ag.NvPicPr, _ea)
	_agb := _b.StartElement{Name: _b.Name{Local: "\u0070\u0069\u0063:\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}}
	e.EncodeElement(_ag.BlipFill, _agb)
	_g := _b.StartElement{Name: _b.Name{Local: "\u0070\u0069\u0063\u003a\u0073\u0070\u0050\u0072"}}
	e.EncodeElement(_ag.SpPr, _g)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_Picture and its children
func (_df *CT_Picture) Validate() error {
	return _df.ValidateWithPath("\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065")
}
func NewCT_Picture() *CT_Picture {
	_e := &CT_Picture{}
	_e.NvPicPr = NewCT_PictureNonVisual()
	_e.BlipFill = _c.NewCT_BlipFillProperties()
	_e.SpPr = _c.NewCT_ShapeProperties()
	return _e
}

type CT_PictureNonVisual struct {
	CNvPr    *_c.CT_NonVisualDrawingProps
	CNvPicPr *_c.CT_NonVisualPictureProperties
}

func (_cba *Pic) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	_cba.CT_Picture = *NewCT_Picture()
_dfc:
	for {
		_gde, _fb := d.Token()
		if _fb != nil {
			return _fb
		}
		switch _dea := _gde.(type) {
		case _b.StartElement:
			switch _dea.Name {
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}, _b.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}:
				if _eg := d.DecodeElement(_cba.NvPicPr, &_dea); _eg != nil {
					return _eg
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}, _b.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}:
				if _ab := d.DecodeElement(_cba.BlipFill, &_dea); _ab != nil {
					return _ab
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0073\u0070\u0050\u0072"}, _b.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0073\u0070\u0050\u0072"}:
				if _ge := d.DecodeElement(_cba.SpPr, &_dea); _ge != nil {
					return _ge
				}
			default:
				_a.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006fn\u0020\u0050i\u0063\u0020\u0025\u0076", _dea.Name)
				if _gea := d.Skip(); _gea != nil {
					return _gea
				}
			}
		case _b.EndElement:
			break _dfc
		case _b.CharData:
		}
	}
	return nil
}
func NewCT_PictureNonVisual() *CT_PictureNonVisual {
	_f := &CT_PictureNonVisual{}
	_f.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_f.CNvPicPr = _c.NewCT_NonVisualPictureProperties()
	return _f
}

type Pic struct{ CT_Picture }

// ValidateWithPath validates the Pic and its children, prefixing error messages with path
func (_ga *Pic) ValidateWithPath(path string) error {
	if _fc := _ga.CT_Picture.ValidateWithPath(path); _fc != nil {
		return _fc
	}
	return nil
}
func (_fad *Pic) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"})
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0070\u0069c"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"})
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0070i\u0063\u003a\u0070\u0069\u0063"
	return _fad.CT_Picture.MarshalXML(e, start)
}
func (_ca *CT_Picture) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	_ca.NvPicPr = NewCT_PictureNonVisual()
	_ca.BlipFill = _c.NewCT_BlipFillProperties()
	_ca.SpPr = _c.NewCT_ShapeProperties()
_gf:
	for {
		_cf, _be := d.Token()
		if _be != nil {
			return _be
		}
		switch _dc := _cf.(type) {
		case _b.StartElement:
			switch _dc.Name {
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}, _b.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}:
				if _ad := d.DecodeElement(_ca.NvPicPr, &_dc); _ad != nil {
					return _ad
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}, _b.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}:
				if _gd := d.DecodeElement(_ca.BlipFill, &_dc); _gd != nil {
					return _gd
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0073\u0070\u0050\u0072"}, _b.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0073\u0070\u0050\u0072"}:
				if _bb := d.DecodeElement(_ca.SpPr, &_dc); _bb != nil {
					return _bb
				}
			default:
				_a.Log("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fP\u0069\u0063\u0074\u0075\u0072\u0065\u0020\u0025\u0076", _dc.Name)
				if _cd := d.Skip(); _cd != nil {
					return _cd
				}
			}
		case _b.EndElement:
			break _gf
		case _b.CharData:
		}
	}
	return nil
}
func NewPic() *Pic { _ebf := &Pic{}; _ebf.CT_Picture = *NewCT_Picture(); return _ebf }

// ValidateWithPath validates the CT_PictureNonVisual and its children, prefixing error messages with path
func (_cg *CT_PictureNonVisual) ValidateWithPath(path string) error {
	if _fa := _cg.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _fa != nil {
		return _fa
	}
	if _bbd := _cg.CNvPicPr.ValidateWithPath(path + "\u002fC\u004e\u0076\u0050\u0069\u0063\u0050r"); _bbd != nil {
		return _bbd
	}
	return nil
}
func (_ee *CT_PictureNonVisual) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	_ee.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_ee.CNvPicPr = _c.NewCT_NonVisualPictureProperties()
_adf:
	for {
		_add, _dcd := d.Token()
		if _dcd != nil {
			return _dcd
		}
		switch _baa := _add.(type) {
		case _b.StartElement:
			switch _baa.Name {
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0063\u004e\u0076P\u0072"}, _b.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0063\u004e\u0076P\u0072"}:
				if _bf := d.DecodeElement(_ee.CNvPr, &_baa); _bf != nil {
					return _bf
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}, _b.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}:
				if _ebc := d.DecodeElement(_ee.CNvPicPr, &_baa); _ebc != nil {
					return _ebc
				}
			default:
				_a.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065No\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _baa.Name)
				if _cb := d.Skip(); _cb != nil {
					return _cb
				}
			}
		case _b.EndElement:
			break _adf
		case _b.CharData:
		}
	}
	return nil
}

// Validate validates the Pic and its children
func (_faf *Pic) Validate() error { return _faf.ValidateWithPath("\u0050\u0069\u0063") }
func init() {
	_a.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", "\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c", NewCT_PictureNonVisual)
	_a.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", "\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065", NewCT_Picture)
	_a.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", "\u0070\u0069\u0063", NewPic)
}
