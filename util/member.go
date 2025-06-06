package util

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

// MemberType defines a member config node used for sending and receiving XML
// from PAN-OS.
type MemberType struct {
	Members []Member `xml:"member"`
}

// Member defines a member config node used for sending and receiving XML
// from PANOS.
type Member struct {
	XMLName xml.Name `xml:"member"`
	Value   string   `xml:",chardata"`
}

// MemToStr normalizes a MemberType pointer into a list of strings.
func MemToStr(e *MemberType) []string {
	if e == nil {
		return nil
	}

	ans := make([]string, len(e.Members))
	for i := range e.Members {
		ans[i] = e.Members[i].Value
	}

	return ans
}

var ErrEmptyMemberList = fmt.Errorf("a list had no member elements")

// MemToInt64 normalizes a MemberType pointer into a list of int64 values.
func MemToInt64(e *MemberType) ([]int64, error) {
	if e == nil {
		return nil, ErrEmptyMemberList
	}

	ans := make([]int64, len(e.Members))
	for idx, elt := range e.Members {
		if converted, err := strconv.ParseInt(elt.Value, 10, 64); err != nil {
			return nil, err
		} else {
			ans[idx] = converted
		}
	}

	return ans, nil
}

// StrToMem converts a list of strings into a MemberType pointer.
func StrToMem(e []string) *MemberType {
	if e == nil {
		return nil
	}

	ans := make([]Member, len(e))
	for i := range e {
		ans[i] = Member{Value: e[i]}
	}

	return &MemberType{ans}
}

// Int64ToMem converts a list of int64 values into a MemberType pointer
func Int64ToMem(e []int64) *MemberType {
	if e == nil {
		return nil
	}

	ans := make([]Member, len(e))
	for idx, elt := range e {
		ans[idx].Value = strconv.FormatInt(elt, 10)
	}

	return &MemberType{ans}
}

// MemToOneStr normalizes a MemberType pointer for a max_items=1 XML node
// into a string.
func MemToOneStr(e *MemberType) string {
	if e == nil || len(e.Members) == 0 {
		return ""
	}

	return e.Members[0].Value
}

// OneStrToMem converts a string into a MemberType pointer for a max_items=1
// XML node.
func OneStrToMem(e string) *MemberType {
	if e == "" {
		return nil
	}

	return &MemberType{[]Member{
		{Value: e},
	}}
}
