// go run receiver.go -i input file -o output file	- encode
// go run receiver.go -d -i output file -o input file	- decode

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func find(s string, HEX []string) int {
	for i, v := range HEX {
		if v == s {
			return i
		}
	}
	return -1
}

type Bintohex struct {
	in  io.Reader
	out io.Writer
}

func (bth Bintohex) read(p []string) (int, error) {
	inBuffer := make([]byte, len(p))
	n, err := bth.in.Read(inBuffer)
	if err != nil {
		return n, err
	}

	c := 0
	for i := 0; i < n; i++ {
		p[i] = HEX[inBuffer[i]]
		c++
	}
	return c, nil
}

func (bth Bintohex) decode() {
	scanner := bufio.NewScanner(bth.in)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		v := find(scanner.Text(), HEX)
		if v < 0 {
			panic("word not found at dict")
		}
		bth.out.Write([]byte{byte(v)})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func (bth Bintohex) encode() {
	final := []string{}
	for {
		encoded := make([]string, 1024)
		n, err := bth.read(encoded)
		if err != nil {
			final = append(final, encoded[0:n]...)
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}

		final = append(final, encoded[0:n]...)
	}
	for i := 0; i < len(final); i += 16{
		if i+16 < len(final) {
			// win32 ln
			fmt.Fprintf(bth.out, "%s\r\n", strings.Join(final[i:(i+16)], " "))
		} else {
			fmt.Fprintf(bth.out, "%s", strings.Join(final[i:len(final)], " "))
		}
	}
}

func main() {
	inFP := flag.String("i", "", "input file")
	outFP := flag.String("o", "", "output file")
	action := flag.Bool("d", false, "use -d to specify decode")
	flag.Parse()

	if *inFP == "" || *outFP == "" {
		fmt.Printf("Usage: receiver [-d] -i input file -o output file")
		os.Exit(1)
	}

	inFD, err := os.Open(*inFP)
	if err != nil {
		panic(err)
	}
	defer inFD.Close()

	outFD, err := os.Create(*outFP)
	if err != nil {
		panic(err)
	}
	defer outFD.Close()

	if *action == false {
		Bintohex := Bintohex{inFD, outFD}
		Bintohex.encode()
	} else {
		Bintohex := Bintohex{inFD, outFD}
		Bintohex.decode()
	}

}

// HEX is our dict to convert binary to string
var HEX = []string{
        "00",
	"01",
	"02",
	"03",
	"04",
	"05",
	"06",
	"07",
	"08",
	"09",
	"0A",
	"0B",
	"0C",
	"0D",
	"0E",
	"0F",
	"10",
	"11",
	"12",
	"13",
	"14",
	"15",
	"16",
	"17",
	"18",
	"19",
	"1A",
	"1B",
	"1C",
	"1D",
	"1E",
	"1F",
	"20",
	"21",
	"22",
	"23",
	"24",
	"25",
	"26",
	"27",
	"28",
	"29",
	"2A",
	"2B",
	"2C",
	"2D",
	"2E",
	"2F",
	"30",
	"31",
	"32",
	"33",
	"34",
	"35",
	"36",
	"37",
	"38",
	"39",
	"3A",
	"3B",
	"3C",
	"3D",
	"3E",
	"3F",
	"40",
	"41",
	"42",
	"43",
	"44",
	"45",
	"46",
	"47",
	"48",
	"49",
	"4A",
	"4B",
	"4C",
	"4D",
	"4E",
	"4F",
	"50",
	"51",
	"52",
	"53",
	"54",
	"55",
	"56",
	"57",
	"58",
	"59",
	"5A",
	"5B",
	"5C",
	"5D",
	"5E",
	"5F",
	"60",
	"61",
	"62",
	"63",
	"64",
	"65",
	"66",
	"67",
	"68",
	"69",
	"6A",
	"6B",
	"6C",
	"6D",
	"6E",
	"6F",
	"70",
	"71",
	"72",
	"73",
	"74",
	"75",
	"76",
	"77",
	"78",
	"79",
	"7A",
	"7B",
	"7C",
	"7D",
	"7E",
	"7F",
	"80",
	"81",
	"82",
	"83",
	"84",
	"85",
	"86",
	"87",
	"88",
	"89",
	"8A",
	"8B",
	"8C",
	"8D",
	"8E",
	"8F",
	"90",
	"91",
	"92",
	"93",
	"94",
	"95",
	"96",
	"97",
	"98",
	"99",
	"9A",
	"9B",
	"9C",
	"9D",
	"9E",
	"9F",
	"A0",
	"A1",
	"A2",
	"A3",
	"A4",
	"A5",
	"A6",
	"A7",
	"A8",
	"A9",
	"AA",
	"AB",
	"AC",
	"AD",
	"AE",
	"AF",
	"B0",
	"B1",
	"B2",
	"B3",
	"B4",
	"B5",
	"B6",
	"B7",
	"B8",
	"B9",
	"BA",
	"BB",
	"BC",
	"BD",
	"BE",
	"BF",
	"C0",
	"C1",
	"C2",
	"C3",
	"C4",
	"C5",
	"C6",
	"C7",
	"C8",
	"C9",
	"CA",
	"CB",
	"CC",
	"CD",
	"CE",
	"CF",
	"D0",
	"D1",
	"D2",
	"D3",
	"D4",
	"D5",
	"D6",
	"D7",
	"D8",
	"D9",
	"DA",
	"DB",
	"DC",
	"DD",
	"DE",
	"DF",
	"E0",
	"E1",
	"E2",
	"E3",
	"E4",
	"E5",
	"E6",
	"E7",
	"E8",
	"E9",
	"EA",
	"EB",
	"EC",
	"ED",
	"EE",
	"EF",
	"F0",
	"F1",
	"F2",
	"F3",
	"F4",
	"F5",
	"F6",
	"F7",
	"F8",
	"F9",
	"FA",
	"FB",
	"FC",
	"FD",
	"FE",
	"FF",
}
