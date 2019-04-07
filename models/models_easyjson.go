// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	types "github.com/studtool/common/types"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD2b7633eDecodeGithubComStudtoolUsersServiceModels(in *jlexer.Lexer, out *UserMap) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
	} else {
		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make(UserMap)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v1 interface{}
			if m, ok := v1.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := v1.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				v1 = in.Interface()
			}
			(*out)[key] = v1
			in.WantComma()
		}
		in.Delim('}')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComStudtoolUsersServiceModels(out *jwriter.Writer, in UserMap) {
	if in == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v2First := true
		for v2Name, v2Value := range in {
			if v2First {
				v2First = false
			} else {
				out.RawByte(',')
			}
			out.String(string(v2Name))
			out.RawByte(':')
			if m, ok := v2Value.(easyjson.Marshaler); ok {
				m.MarshalEasyJSON(out)
			} else if m, ok := v2Value.(json.Marshaler); ok {
				out.Raw(m.MarshalJSON())
			} else {
				out.Raw(json.Marshal(v2Value))
			}
		}
		out.RawByte('}')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v UserMap) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComStudtoolUsersServiceModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserMap) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComStudtoolUsersServiceModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserMap) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComStudtoolUsersServiceModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserMap) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComStudtoolUsersServiceModels(l, v)
}
func easyjsonD2b7633eDecodeGithubComStudtoolUsersServiceModels1(in *jlexer.Lexer, out *UserInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "userId":
			out.Id = string(in.String())
		case "username":
			out.Username = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComStudtoolUsersServiceModels1(out *jwriter.Writer, in UserInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComStudtoolUsersServiceModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComStudtoolUsersServiceModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComStudtoolUsersServiceModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComStudtoolUsersServiceModels1(l, v)
}
func easyjsonD2b7633eDecodeGithubComStudtoolUsersServiceModels2(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "userId":
			out.Id = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "fullName":
			if in.IsNull() {
				in.Skip()
				out.FullName = nil
			} else {
				if out.FullName == nil {
					out.FullName = new(string)
				}
				*out.FullName = string(in.String())
			}
		case "dateOfBirth":
			if in.IsNull() {
				in.Skip()
				out.DateOfBirth = nil
			} else {
				if out.DateOfBirth == nil {
					out.DateOfBirth = new(types.Date)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.DateOfBirth).UnmarshalJSON(data))
				}
			}
		case "locationInfo":
			if in.IsNull() {
				in.Skip()
				out.Location = nil
			} else {
				if out.Location == nil {
					out.Location = new(LocationInfo)
				}
				(*out.Location).UnmarshalEasyJSON(in)
			}
		case "universityInfo":
			if in.IsNull() {
				in.Skip()
				out.University = nil
			} else {
				if out.University == nil {
					out.University = new(UniversityInfo)
				}
				(*out.University).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComStudtoolUsersServiceModels2(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"fullName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.FullName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.FullName))
		}
	}
	{
		const prefix string = ",\"dateOfBirth\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.DateOfBirth == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.DateOfBirth).MarshalJSON())
		}
	}
	{
		const prefix string = ",\"locationInfo\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Location == nil {
			out.RawString("null")
		} else {
			(*in.Location).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"universityInfo\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.University == nil {
			out.RawString("null")
		} else {
			(*in.University).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComStudtoolUsersServiceModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComStudtoolUsersServiceModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComStudtoolUsersServiceModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComStudtoolUsersServiceModels2(l, v)
}
func easyjsonD2b7633eDecodeGithubComStudtoolUsersServiceModels3(in *jlexer.Lexer, out *UniversityInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "department":
			if in.IsNull() {
				in.Skip()
				out.Department = nil
			} else {
				if out.Department == nil {
					out.Department = new(string)
				}
				*out.Department = string(in.String())
			}
		case "speciality":
			if in.IsNull() {
				in.Skip()
				out.Speciality = nil
			} else {
				if out.Speciality == nil {
					out.Speciality = new(string)
				}
				*out.Speciality = string(in.String())
			}
		case "admissionYear":
			if in.IsNull() {
				in.Skip()
				out.AdmissionYear = nil
			} else {
				if out.AdmissionYear == nil {
					out.AdmissionYear = new(int)
				}
				*out.AdmissionYear = int(in.Int())
			}
		case "graduationYear":
			if in.IsNull() {
				in.Skip()
				out.GraduationYear = nil
			} else {
				if out.GraduationYear == nil {
					out.GraduationYear = new(int)
				}
				*out.GraduationYear = int(in.Int())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComStudtoolUsersServiceModels3(out *jwriter.Writer, in UniversityInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"department\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Department == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Department))
		}
	}
	{
		const prefix string = ",\"speciality\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Speciality == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Speciality))
		}
	}
	{
		const prefix string = ",\"admissionYear\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.AdmissionYear == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.AdmissionYear))
		}
	}
	{
		const prefix string = ",\"graduationYear\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.GraduationYear == nil {
			out.RawString("null")
		} else {
			out.Int(int(*in.GraduationYear))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UniversityInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComStudtoolUsersServiceModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UniversityInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComStudtoolUsersServiceModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UniversityInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComStudtoolUsersServiceModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UniversityInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComStudtoolUsersServiceModels3(l, v)
}
func easyjsonD2b7633eDecodeGithubComStudtoolUsersServiceModels4(in *jlexer.Lexer, out *LocationInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "country":
			out.Country = string(in.String())
		case "city":
			if in.IsNull() {
				in.Skip()
				out.City = nil
			} else {
				if out.City == nil {
					out.City = new(string)
				}
				*out.City = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComStudtoolUsersServiceModels4(out *jwriter.Writer, in LocationInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"country\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Country))
	}
	{
		const prefix string = ",\"city\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.City == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.City))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LocationInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComStudtoolUsersServiceModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LocationInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComStudtoolUsersServiceModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LocationInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComStudtoolUsersServiceModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LocationInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComStudtoolUsersServiceModels4(l, v)
}