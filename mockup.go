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
	"bufio"
	"flag"
	"fmt"
	"os"
)

func EncodeString(encodedstring string) string {
	output := ""
	//fmt.Println(encodedstring)

	for i, v := range encodedstring {
		if int(v) < 0x41 || int(v) > 0x5a &&
			int(v) < 0x61 || int(v) > 0x7a {
			continue
		}
		mod := 0x41
		if int(v) >= 0x60 {
			mod = 0x61
		}
		mv := 3
		if (int(v) >= 0x50 && int(v) <= 0x53) ||
			(int(v) >= 0x70 && int(v) <= 0x73) {
			mv = 4
		}
		if (int(v) >= 0x54 && int(v) <= 0x56) ||
			(int(v) >= 0x74 && int(v) <= 0x76) {
			mod += 0x0d
		}
		if (int(v) >= 0x57 && int(v) <= 0x5A) ||
			(int(v) >= 0x77 && int(v) <= 0x7A) {
			mv = 4
			mod += 0x16
		}

		val := (int(v) - mod)

		idx := i + 1
		m := ((val + (mv - idx)) % mv)
		d := val / mv

		enc := string(m + (d * mv) + mod)
		output += enc
	}

	return output
}

func DecodeString(encodedstring string) string {
	output := ""
	//fmt.Println(encodedstring)

	for i, v := range encodedstring {
		if int(v) < 0x41 || int(v) > 0x5a &&
			int(v) < 0x61 || int(v) > 0x7a {
			continue
		}
		mod := 0x41
		if int(v) >= 0x60 {
			mod = 0x61
		}
		mv := 3
		if (int(v) >= 0x50 && int(v) <= 0x53) ||
			(int(v) >= 0x70 && int(v) <= 0x73) {
			mv = 4
		}
		if (int(v) >= 0x54 && int(v) <= 0x56) ||
			(int(v) >= 0x74 && int(v) <= 0x76) {
			mod += 0x0d
		}
		if (int(v) >= 0x57 && int(v) <= 0x5A) ||
			(int(v) >= 0x77 && int(v) <= 0x7A) {
			mv = 4
			mod += 0x16
		}

		val := (int(v) - mod)

		idx := i + 1
		m := (val + idx) % mv
		d := val / mv

		enc := string(m + (d * mv) + mod)
		output += enc
	}

	return output
}

func main() {
	encdec := flag.Bool("encode", false, "Encode the values from the input file")
	flag.Parse()
	args := flag.Args()

	fmt.Println("encdec", *encdec)
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "requires at 1 argument")
		return
	}
	f, err := os.Open(args[0])
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if !*encdec {
			fmt.Println(DecodeString(scanner.Text()))
		} else {
			fmt.Println(EncodeString(scanner.Text()))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
