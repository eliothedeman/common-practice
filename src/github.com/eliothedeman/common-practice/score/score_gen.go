package score

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Movement) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "index":
			z.Index, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "parts":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Parts) >= int(xsz) {
				z.Parts = z.Parts[:xsz]
			} else {
				z.Parts = make([]Part, xsz)
			}
			for xvk := range z.Parts {
				var isz uint32
				isz, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for isz > 0 {
					isz--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "insturment":
						z.Parts[xvk].Insturment, err = dc.ReadString()
						if err != nil {
							return
						}
					case "measures":
						var xsz uint32
						xsz, err = dc.ReadArrayHeader()
						if err != nil {
							return
						}
						if cap(z.Parts[xvk].Measures) >= int(xsz) {
							z.Parts[xvk].Measures = z.Parts[xvk].Measures[:xsz]
						} else {
							z.Parts[xvk].Measures = make([]Measure, xsz)
						}
						for bzg := range z.Parts[xvk].Measures {
							err = z.Parts[xvk].Measures[bzg].DecodeMsg(dc)
							if err != nil {
								return
							}
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Movement) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(3)
	if err != nil {
		return
	}
	err = en.WriteString("name")
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	err = en.WriteString("index")
	if err != nil {
		return
	}
	err = en.WriteInt(z.Index)
	if err != nil {
		return
	}
	err = en.WriteString("parts")
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Parts)))
	if err != nil {
		return
	}
	for xvk := range z.Parts {
		err = en.WriteMapHeader(2)
		if err != nil {
			return
		}
		err = en.WriteString("insturment")
		if err != nil {
			return
		}
		err = en.WriteString(z.Parts[xvk].Insturment)
		if err != nil {
			return
		}
		err = en.WriteString("measures")
		if err != nil {
			return
		}
		err = en.WriteArrayHeader(uint32(len(z.Parts[xvk].Measures)))
		if err != nil {
			return
		}
		for bzg := range z.Parts[xvk].Measures {
			err = z.Parts[xvk].Measures[bzg].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Movement) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 3)
	o = msgp.AppendString(o, "name")
	o = msgp.AppendString(o, z.Name)
	o = msgp.AppendString(o, "index")
	o = msgp.AppendInt(o, z.Index)
	o = msgp.AppendString(o, "parts")
	o = msgp.AppendArrayHeader(o, uint32(len(z.Parts)))
	for xvk := range z.Parts {
		o = msgp.AppendMapHeader(o, 2)
		o = msgp.AppendString(o, "insturment")
		o = msgp.AppendString(o, z.Parts[xvk].Insturment)
		o = msgp.AppendString(o, "measures")
		o = msgp.AppendArrayHeader(o, uint32(len(z.Parts[xvk].Measures)))
		for bzg := range z.Parts[xvk].Measures {
			o, err = z.Parts[xvk].Measures[bzg].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Movement) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "index":
			z.Index, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "parts":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Parts) >= int(xsz) {
				z.Parts = z.Parts[:xsz]
			} else {
				z.Parts = make([]Part, xsz)
			}
			for xvk := range z.Parts {
				var isz uint32
				isz, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for isz > 0 {
					isz--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "insturment":
						z.Parts[xvk].Insturment, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "measures":
						var xsz uint32
						xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.Parts[xvk].Measures) >= int(xsz) {
							z.Parts[xvk].Measures = z.Parts[xvk].Measures[:xsz]
						} else {
							z.Parts[xvk].Measures = make([]Measure, xsz)
						}
						for bzg := range z.Parts[xvk].Measures {
							bts, err = z.Parts[xvk].Measures[bzg].UnmarshalMsg(bts)
							if err != nil {
								return
							}
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Movement) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 4 + msgp.StringPrefixSize + len(z.Name) + msgp.StringPrefixSize + 5 + msgp.IntSize + msgp.StringPrefixSize + 5 + msgp.ArrayHeaderSize
	for xvk := range z.Parts {
		s += msgp.MapHeaderSize + msgp.StringPrefixSize + 10 + msgp.StringPrefixSize + len(z.Parts[xvk].Insturment) + msgp.StringPrefixSize + 8 + msgp.ArrayHeaderSize
		for bzg := range z.Parts[xvk].Measures {
			s += z.Parts[xvk].Measures[bzg].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Part) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "insturment":
			z.Insturment, err = dc.ReadString()
			if err != nil {
				return
			}
		case "measures":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Measures) >= int(xsz) {
				z.Measures = z.Measures[:xsz]
			} else {
				z.Measures = make([]Measure, xsz)
			}
			for bai := range z.Measures {
				err = z.Measures[bai].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Part) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(2)
	if err != nil {
		return
	}
	err = en.WriteString("insturment")
	if err != nil {
		return
	}
	err = en.WriteString(z.Insturment)
	if err != nil {
		return
	}
	err = en.WriteString("measures")
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Measures)))
	if err != nil {
		return
	}
	for bai := range z.Measures {
		err = z.Measures[bai].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Part) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 2)
	o = msgp.AppendString(o, "insturment")
	o = msgp.AppendString(o, z.Insturment)
	o = msgp.AppendString(o, "measures")
	o = msgp.AppendArrayHeader(o, uint32(len(z.Measures)))
	for bai := range z.Measures {
		o, err = z.Measures[bai].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Part) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "insturment":
			z.Insturment, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "measures":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Measures) >= int(xsz) {
				z.Measures = z.Measures[:xsz]
			} else {
				z.Measures = make([]Measure, xsz)
			}
			for bai := range z.Measures {
				bts, err = z.Measures[bai].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Part) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 10 + msgp.StringPrefixSize + len(z.Insturment) + msgp.StringPrefixSize + 8 + msgp.ArrayHeaderSize
	for bai := range z.Measures {
		s += z.Measures[bai].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TimeSignature) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "beat":
			z.Beat, err = dc.ReadUint8()
			if err != nil {
				return
			}
		case "value":
			z.Value, err = dc.ReadUint8()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z TimeSignature) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(2)
	if err != nil {
		return
	}
	err = en.WriteString("beat")
	if err != nil {
		return
	}
	err = en.WriteUint8(z.Beat)
	if err != nil {
		return
	}
	err = en.WriteString("value")
	if err != nil {
		return
	}
	err = en.WriteUint8(z.Value)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z TimeSignature) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 2)
	o = msgp.AppendString(o, "beat")
	o = msgp.AppendUint8(o, z.Beat)
	o = msgp.AppendString(o, "value")
	o = msgp.AppendUint8(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TimeSignature) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "beat":
			z.Beat, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "value":
			z.Value, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z TimeSignature) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 4 + msgp.Uint8Size + msgp.StringPrefixSize + 5 + msgp.Uint8Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Measure) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "time_signature":
			var isz uint32
			isz, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for isz > 0 {
				isz--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "beat":
					z.TimeSignature.Beat, err = dc.ReadUint8()
					if err != nil {
						return
					}
				case "value":
					z.TimeSignature.Value, err = dc.ReadUint8()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "index":
			z.Index, err = dc.ReadUint32()
			if err != nil {
				return
			}
		case "notes":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Notes) >= int(xsz) {
				z.Notes = z.Notes[:xsz]
			} else {
				z.Notes = make([]Note, xsz)
			}
			for cmr := range z.Notes {
				err = z.Notes[cmr].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Measure) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(3)
	if err != nil {
		return
	}
	err = en.WriteString("time_signature")
	if err != nil {
		return
	}
	err = en.WriteMapHeader(2)
	if err != nil {
		return
	}
	err = en.WriteString("beat")
	if err != nil {
		return
	}
	err = en.WriteUint8(z.TimeSignature.Beat)
	if err != nil {
		return
	}
	err = en.WriteString("value")
	if err != nil {
		return
	}
	err = en.WriteUint8(z.TimeSignature.Value)
	if err != nil {
		return
	}
	err = en.WriteString("index")
	if err != nil {
		return
	}
	err = en.WriteUint32(z.Index)
	if err != nil {
		return
	}
	err = en.WriteString("notes")
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Notes)))
	if err != nil {
		return
	}
	for cmr := range z.Notes {
		err = z.Notes[cmr].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Measure) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 3)
	o = msgp.AppendString(o, "time_signature")
	o = msgp.AppendMapHeader(o, 2)
	o = msgp.AppendString(o, "beat")
	o = msgp.AppendUint8(o, z.TimeSignature.Beat)
	o = msgp.AppendString(o, "value")
	o = msgp.AppendUint8(o, z.TimeSignature.Value)
	o = msgp.AppendString(o, "index")
	o = msgp.AppendUint32(o, z.Index)
	o = msgp.AppendString(o, "notes")
	o = msgp.AppendArrayHeader(o, uint32(len(z.Notes)))
	for cmr := range z.Notes {
		o, err = z.Notes[cmr].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Measure) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "time_signature":
			var isz uint32
			isz, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for isz > 0 {
				isz--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "beat":
					z.TimeSignature.Beat, bts, err = msgp.ReadUint8Bytes(bts)
					if err != nil {
						return
					}
				case "value":
					z.TimeSignature.Value, bts, err = msgp.ReadUint8Bytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "index":
			z.Index, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "notes":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Notes) >= int(xsz) {
				z.Notes = z.Notes[:xsz]
			} else {
				z.Notes = make([]Note, xsz)
			}
			for cmr := range z.Notes {
				bts, err = z.Notes[cmr].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Measure) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 14 + msgp.MapHeaderSize + msgp.StringPrefixSize + 4 + msgp.Uint8Size + msgp.StringPrefixSize + 5 + msgp.Uint8Size + msgp.StringPrefixSize + 5 + msgp.Uint32Size + msgp.StringPrefixSize + 5 + msgp.ArrayHeaderSize
	for cmr := range z.Notes {
		s += z.Notes[cmr].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Note) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "step":
			z.Step, err = dc.ReadUint8()
			if err != nil {
				return
			}
		case "octave":
			z.Octave, err = dc.ReadUint8()
			if err != nil {
				return
			}
		case "duration":
			z.Duration, err = dc.ReadUint16()
			if err != nil {
				return
			}
		case "placement":
			z.Placement, err = dc.ReadUint16()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Note) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(4)
	if err != nil {
		return
	}
	err = en.WriteString("step")
	if err != nil {
		return
	}
	err = en.WriteUint8(z.Step)
	if err != nil {
		return
	}
	err = en.WriteString("octave")
	if err != nil {
		return
	}
	err = en.WriteUint8(z.Octave)
	if err != nil {
		return
	}
	err = en.WriteString("duration")
	if err != nil {
		return
	}
	err = en.WriteUint16(z.Duration)
	if err != nil {
		return
	}
	err = en.WriteString("placement")
	if err != nil {
		return
	}
	err = en.WriteUint16(z.Placement)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Note) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 4)
	o = msgp.AppendString(o, "step")
	o = msgp.AppendUint8(o, z.Step)
	o = msgp.AppendString(o, "octave")
	o = msgp.AppendUint8(o, z.Octave)
	o = msgp.AppendString(o, "duration")
	o = msgp.AppendUint16(o, z.Duration)
	o = msgp.AppendString(o, "placement")
	o = msgp.AppendUint16(o, z.Placement)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Note) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "step":
			z.Step, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "octave":
			z.Octave, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "duration":
			z.Duration, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		case "placement":
			z.Placement, bts, err = msgp.ReadUint16Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Note) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 4 + msgp.Uint8Size + msgp.StringPrefixSize + 6 + msgp.Uint8Size + msgp.StringPrefixSize + 8 + msgp.Uint16Size + msgp.StringPrefixSize + 9 + msgp.Uint16Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Score) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "composer":
			z.Composer, err = dc.ReadString()
			if err != nil {
				return
			}
		case "licence":
			z.Licence, err = dc.ReadString()
			if err != nil {
				return
			}
		case "title":
			z.Title, err = dc.ReadString()
			if err != nil {
				return
			}
		case "movements":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Movements) >= int(xsz) {
				z.Movements = z.Movements[:xsz]
			} else {
				z.Movements = make([]Movement, xsz)
			}
			for ajw := range z.Movements {
				err = z.Movements[ajw].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Score) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteMapHeader(4)
	if err != nil {
		return
	}
	err = en.WriteString("composer")
	if err != nil {
		return
	}
	err = en.WriteString(z.Composer)
	if err != nil {
		return
	}
	err = en.WriteString("licence")
	if err != nil {
		return
	}
	err = en.WriteString(z.Licence)
	if err != nil {
		return
	}
	err = en.WriteString("title")
	if err != nil {
		return
	}
	err = en.WriteString(z.Title)
	if err != nil {
		return
	}
	err = en.WriteString("movements")
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Movements)))
	if err != nil {
		return
	}
	for ajw := range z.Movements {
		err = z.Movements[ajw].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Score) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendMapHeader(o, 4)
	o = msgp.AppendString(o, "composer")
	o = msgp.AppendString(o, z.Composer)
	o = msgp.AppendString(o, "licence")
	o = msgp.AppendString(o, z.Licence)
	o = msgp.AppendString(o, "title")
	o = msgp.AppendString(o, z.Title)
	o = msgp.AppendString(o, "movements")
	o = msgp.AppendArrayHeader(o, uint32(len(z.Movements)))
	for ajw := range z.Movements {
		o, err = z.Movements[ajw].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Score) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "composer":
			z.Composer, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "licence":
			z.Licence, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "title":
			z.Title, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "movements":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Movements) >= int(xsz) {
				z.Movements = z.Movements[:xsz]
			} else {
				z.Movements = make([]Movement, xsz)
			}
			for ajw := range z.Movements {
				bts, err = z.Movements[ajw].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Score) Msgsize() (s int) {
	s = msgp.MapHeaderSize + msgp.StringPrefixSize + 8 + msgp.StringPrefixSize + len(z.Composer) + msgp.StringPrefixSize + 7 + msgp.StringPrefixSize + len(z.Licence) + msgp.StringPrefixSize + 5 + msgp.StringPrefixSize + len(z.Title) + msgp.StringPrefixSize + 9 + msgp.ArrayHeaderSize
	for ajw := range z.Movements {
		s += z.Movements[ajw].Msgsize()
	}
	return
}
