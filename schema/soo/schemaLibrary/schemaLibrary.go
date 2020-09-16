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

package schemaLibrary

import (
	_d "encoding/xml"
	_fe "fmt"
	_a "github.com/stasomega1/mylib"
)

type CT_Schema struct {
	UriAttr              *string
	ManifestLocationAttr *string
	SchemaLocationAttr   *string
	SchemaLanguageAttr   *string
}

func (_cfg *SchemaLibrary) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u006d\u0061"}, Value: "\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u006d\u0061:\u0073\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079"
	return _cfg.CT_SchemaLibrary.MarshalXML(e, start)
}
func (_e *CT_Schema) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _de := range start.Attr {
		if _de.Name.Local == "\u0075\u0072\u0069" {
			_dea, _dg := _de.Value, error(nil)
			if _dg != nil {
				return _dg
			}
			_e.UriAttr = &_dea
			continue
		}
		if _de.Name.Local == "\u006d\u0061n\u0069\u0066\u0065s\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e" {
			_df, _cb := _de.Value, error(nil)
			if _cb != nil {
				return _cb
			}
			_e.ManifestLocationAttr = &_df
			continue
		}
		if _de.Name.Local == "\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e" {
			_fb, _b := _de.Value, error(nil)
			if _b != nil {
				return _b
			}
			_e.SchemaLocationAttr = &_fb
			continue
		}
		if _de.Name.Local == "\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065" {
			_bb, _fa := _de.Value, error(nil)
			if _fa != nil {
				return _fa
			}
			_e.SchemaLanguageAttr = &_bb
			continue
		}
	}
	for {
		_aeb, _cbc := d.Token()
		if _cbc != nil {
			return _fe.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0053\u0063\u0068e\u006d\u0061\u003a\u0020\u0025\u0073", _cbc)
		}
		if _fc, _eg := _aeb.(_d.EndElement); _eg && _fc.Name == start.Name {
			break
		}
	}
	return nil
}
func NewCT_SchemaLibrary() *CT_SchemaLibrary { _ce := &CT_SchemaLibrary{}; return _ce }

// ValidateWithPath validates the SchemaLibrary and its children, prefixing error messages with path
func (_cd *SchemaLibrary) ValidateWithPath(path string) error {
	if _gc := _cd.CT_SchemaLibrary.ValidateWithPath(path); _gc != nil {
		return _gc
	}
	return nil
}

type SchemaLibrary struct{ CT_SchemaLibrary }

// ValidateWithPath validates the CT_Schema and its children, prefixing error messages with path
func (_g *CT_Schema) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the CT_SchemaLibrary and its children, prefixing error messages with path
func (_cgc *CT_SchemaLibrary) ValidateWithPath(path string) error {
	for _fbc, _eb := range _cgc.Schema {
		if _gg := _eb.ValidateWithPath(_fe.Sprintf("\u0025\u0073\u002f\u0053\u0063\u0068\u0065\u006d\u0061\u005b\u0025\u0064\u005d", path, _fbc)); _gg != nil {
			return _gg
		}
	}
	return nil
}

// Validate validates the CT_SchemaLibrary and its children
func (_be *CT_SchemaLibrary) Validate() error {
	return _be.ValidateWithPath("\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079")
}
func (_ae *CT_Schema) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _ae.UriAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006d\u0061\u003a\u0075\u0072\u0069"}, Value: _fe.Sprintf("\u0025\u0076", *_ae.UriAttr)})
	}
	if _ae.ManifestLocationAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006d\u0061\u003a\u006dan\u0069\u0066\u0065\u0073\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"}, Value: _fe.Sprintf("\u0025\u0076", *_ae.ManifestLocationAttr)})
	}
	if _ae.SchemaLocationAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"}, Value: _fe.Sprintf("\u0025\u0076", *_ae.SchemaLocationAttr)})
	}
	if _ae.SchemaLanguageAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"}, Value: _fe.Sprintf("\u0025\u0076", *_ae.SchemaLanguageAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

// Validate validates the SchemaLibrary and its children
func (_beb *SchemaLibrary) Validate() error {
	return _beb.ValidateWithPath("\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079")
}
func NewSchemaLibrary() *SchemaLibrary {
	_ba := &SchemaLibrary{}
	_ba.CT_SchemaLibrary = *NewCT_SchemaLibrary()
	return _ba
}

type CT_SchemaLibrary struct{ Schema []*CT_Schema }

// Validate validates the CT_Schema and its children
func (_ea *CT_Schema) Validate() error {
	return _ea.ValidateWithPath("\u0043T\u005f\u0053\u0063\u0068\u0065\u006da")
}
func (_dge *CT_SchemaLibrary) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
_ga:
	for {
		_fca, _gae := d.Token()
		if _gae != nil {
			return _gae
		}
		switch _dgb := _fca.(type) {
		case _d.StartElement:
			switch _dgb.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", Local: "\u0073\u0063\u0068\u0065\u006d\u0061"}:
				_cf := NewCT_Schema()
				if _af := d.DecodeElement(_cf, &_dgb); _af != nil {
					return _af
				}
				_dge.Schema = append(_dge.Schema, _cf)
			default:
				_a.Log("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u0020\u0025v", _dgb.Name)
				if _fed := d.Skip(); _fed != nil {
					return _fed
				}
			}
		case _d.EndElement:
			break _ga
		case _d.CharData:
		}
	}
	return nil
}
func (_ec *SchemaLibrary) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ec.CT_SchemaLibrary = *NewCT_SchemaLibrary()
_ed:
	for {
		_db, _egda := d.Token()
		if _egda != nil {
			return _egda
		}
		switch _ecf := _db.(type) {
		case _d.StartElement:
			switch _ecf.Name {
			case _d.Name{Space: "\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", Local: "\u0073\u0063\u0068\u0065\u006d\u0061"}:
				_ca := NewCT_Schema()
				if _dbg := d.DecodeElement(_ca, &_ecf); _dbg != nil {
					return _dbg
				}
				_ec.Schema = append(_ec.Schema, _ca)
			default:
				_a.Log("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0053\u0063\u0068\u0065m\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079 \u0025\u0076", _ecf.Name)
				if _eae := d.Skip(); _eae != nil {
					return _eae
				}
			}
		case _d.EndElement:
			break _ed
		case _d.CharData:
		}
	}
	return nil
}
func (_cg *CT_SchemaLibrary) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	if _cg.Schema != nil {
		_egd := _d.StartElement{Name: _d.Name{Local: "\u006da\u003a\u0073\u0063\u0068\u0065\u006da"}}
		for _, _ac := range _cg.Schema {
			e.EncodeElement(_ac, _egd)
		}
	}
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func NewCT_Schema() *CT_Schema { _fd := &CT_Schema{}; return _fd }
func init() {
	_a.RegisterConstructor("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", "\u0043T\u005f\u0053\u0063\u0068\u0065\u006da", NewCT_Schema)
	_a.RegisterConstructor("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", "\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079", NewCT_SchemaLibrary)
	_a.RegisterConstructor("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", "\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079", NewSchemaLibrary)
}
