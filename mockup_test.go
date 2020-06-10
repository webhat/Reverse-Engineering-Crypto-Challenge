/*
   Copyright (C) 2020 Daniel W. Crompton, Special Brands Holding <test10@specialbrands.net>

   This program is free software; you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation; either version 2 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program; if not, write to the Free Software
   Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA  02111-1307  USA
*/
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
