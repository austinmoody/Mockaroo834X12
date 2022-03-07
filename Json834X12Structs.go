package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type X12Delimiters struct {
	Element    string
	SubElement string
	Segment    string
}

type X12n834 struct {
	ISA ISA `json:"ISA"`
}

type ISA struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
	De06 string `json:"06"`
	De07 string `json:"07"`
	De08 string `json:"08"`
	De09 string `json:"09"`
	De10 string `json:"10"`
	De11 string `json:"11"`
	De12 string `json:"12"`
	De13 string `json:"13"`
	De14 string `json:"14"`
	De15 string `json:"15"`
	De16 string `json:"16"`
	Gs   []GS   `json:"GS"`
}

func (segment ISA) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

type GS struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
	De06 string `json:"06"`
	De07 string `json:"07"`
	De08 string `json:"08"`
	St   []ST   `json:"ST"`
}

func (segment GS) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

type ST struct {
	De01      string    `json:"01"`
	De02      string    `json:"02"`
	De03      string    `json:"03"`
	Bgn       BGN       `json:"BGN"`
	Ref       REF       `json:"REF"`
	Dtp       DTP       `json:"DTP"`
	Loop1000A Loop1000A `json:"1000A"`
	Loop1000B Loop1000B `json:"1000B"`
	Ins       []INS     `json:"INS"`
}

func (segment ST) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

type Loop1000A struct {
	N1 N1 `json:"N1"`
}

type Loop1000B struct {
	N1 N1 `json:"N1"`
}

type BGN struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
	De06 string `json:"06"`
	De07 string `json:"07"`
	De08 string `json:"08"`
}

func (segment BGN) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

type REF struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
}

func (segment REF) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

type DTP struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
}

func (segment DTP) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

type N1 struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
}

func (segment N1) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

type INS struct {
	De01      string    `json:"01"`
	De02      string    `json:"02"`
	De03      string    `json:"03"`
	De04      string    `json:"04"`
	De05      string    `json:"05"`
	De06      string    `json:"06"`
	De07      string    `json:"07"`
	De08      string    `json:"08"`
	De09      string    `json:"09"`
	De10      string    `json:"10"`
	Ref       []REF     `json:"REF"`
	Dtp       []DTP     `json:"DTP"`
	Loop2100A Loop2100A `json:"2100A"`
	Loop2100C Loop2100C `json:"2100C"`
	Loop2100F Loop2100F `json:"2100F"`
	Loop2100G Loop2100G `json:"2100G"`
	// Add Loop2300
	Loop2310 []Loop2310 `json:"2310"`
}

func (segment INS) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

type Loop2100A struct {
	NM1 NM1 `json:"NM1"`
	PER PER `json:"PER"`
	N3  N3  `json:"N3"`
	N4  N4  `json:"N4"`
	DMG DMG `json:"DMG"`
	LUI LUI `json:"LUI"`
}

type Loop2100C struct {
	NM1 NM1 `json:"NM1"`
	N3  N3  `json:"N3"`
	N4  N4  `json:"N4"`
}

type Loop2100F struct {
	NM1 NM1 `json:"NM1"`
}

type Loop2100G struct {
	NM1 NM1 `json:"NM1"`
	PER PER `json:"PER"`
}

type Loop2310 struct {
	LX  LX  `json:"LX"`
	NM1 NM1 `json:"NM1"`
	N3  N3  `json:"N3"`
	N4  N4  `json:"N4"`
	PLA PLA `json:"PLA"`
}

type NM1 struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
	De06 string `json:"06"`
	De07 string `json:"07"`
	De08 string `json:"08"`
	De09 string `json:"09"`
	De10 string `json:"10"`
}

func (segment NM1) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

type PER struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
	De06 string `json:"06"`
	De07 string `json:"07"`
	De08 string `json:"08"`
}

func (segment PER) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

type N3 struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
}

func (segment N3) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

type N4 struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
	De06 string `json:"06"`
}

func (segment N4) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

type DMG struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
}

func (segment DMG) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

type LUI struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
}

func (segment LUI) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

type LX struct {
	De01 string `json:"01"`
}

func (segment LX) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

type PLA struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
}

func (segment PLA) String(x12delimiters X12Delimiters, x12options ...map[string]string) string {
	delimiters := x12delimiters
	options := GetOptions(x12options)
	return GetSegmentAsString(reflect.ValueOf(segment), GetType(segment), delimiters, options)
}

func GetSegmentAsString(v reflect.Value, segmentName string, delimiters X12Delimiters, options map[string]string) string {
	// TODO - how to handle sub-element?
	var fieldSlice []string

	// Set max # of elements to print using MaxElements from options
	// if passed in.  This is useful for segments like NM1 where
	// some Loops print different #'s of elements.
	maxElements := v.NumField()
	if options["MaxElements"] != "" {
		maxElements, _ = strconv.Atoi(options["MaxElements"])
	}

	// Seed slice w/ passed in segment name
	fieldSlice = append(fieldSlice, segmentName)
	for i := 0; i < maxElements; i++ {
		value := v.Field(i)
		switch value.Kind() {
		case reflect.String:
			fieldSlice = append(fieldSlice, v.Field(i).String())
		case reflect.Int, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int64:
			numberValue := strconv.FormatInt(v.Field(i).Int(), 10)
			fieldSlice = append(fieldSlice, numberValue)
		case reflect.Float64, reflect.Float32:
			numberValue := strconv.FormatFloat(v.Field(i).Float(), 'f', 5, 32)
			fieldSlice = append(fieldSlice, numberValue)
		case reflect.Bool:
			boolValue := strconv.FormatBool(v.Field(i).Bool())
			fieldSlice = append(fieldSlice, boolValue)
		}
	}

	return fmt.Sprintf("%s%s", strings.Join(fieldSlice, delimiters.Element), delimiters.Segment)
}

func GetOptions(x12options []map[string]string) map[string]string {
	var options map[string]string

	if len(x12options) == 1 {
		options = x12options[0]
	} else {
		options = map[string]string{}
	}

	return options
}

func GetType(object interface{}) string {
	t := reflect.TypeOf(object)
	if t != nil && t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
