package main

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	type args struct {
		encodedstring string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"A", args{"A"}, "B"},
		{"B", args{"B"}, "C"},
		{"C", args{"C"}, "A"},
		{"AA", args{"AA"}, "BC"},
		{"D", args{"D"}, "E"},
		{"DD", args{"DD"}, "EF"},
		{"EE", args{"EE"}, "FD"},
		{"FF", args{"FF"}, "DE"},
		{"H", args{"H"}, "I"},
		{"a", args{"a"}, "b"},
		{"AZ", args{"AZ"}, "BX"},
		{"MNO", args{"MNO"}, "NMO"},
		{"TUV", args{"TUV"}, "UTV"},
		{"VVVV", args{"VVVV"}, "TUVT"},
		{"az", args{"az"}, "bx"},
		{"lover second", args{"IhatesMS"}, "GgaudqNS"},
		{"First", args{"ABCD"}, "BACE"},
		{"Second", args{"IhateSMS"}, "GgaudQNS"},
		{"Third", args{"ccd"}, "abd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeString(tt.args.encodedstring); got != tt.want {
				t.Errorf("DecodeString(%v) = %v, want %v", tt.args.encodedstring, got, tt.want)
			}
		})
	}
}

func TestEncodeString(t *testing.T) {
	type args struct {
		encodedstring string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"A", args{"A"}, "B"},
		{"B", args{"B"}, "C"},
		{"C", args{"C"}, "A"},
		{"AA", args{"AA"}, "BC"},
		{"D", args{"D"}, "E"},
		{"DD", args{"DD"}, "EF"},
		{"EE", args{"EE"}, "FD"},
		{"FF", args{"FF"}, "DE"},
		{"H", args{"H"}, "I"},
		{"a", args{"a"}, "b"},
		{"AZ", args{"AZ"}, "BX"},
		{"MNO", args{"MNO"}, "NMO"},
		{"TUV", args{"TUV"}, "UTV"},
		{"VVVV", args{"VVVV"}, "TUVT"},
		{"az", args{"az"}, "bx"},
		{"lover second", args{"IhatesMS"}, "GgaudqNS"},
		{"First", args{"ABCD"}, "BACE"},
		{"Second", args{"IhateSMS"}, "GgaudQNS"},
		{"Third", args{"ccd"}, "abd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeString(tt.want); got != tt.args.encodedstring {
				t.Errorf("EncodeString(%v) = %v, want %v", tt.args.encodedstring, got, tt.want)
			}
		})
	}
}
