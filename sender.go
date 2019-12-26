// go run sender.go -i input file -o output file	- encode
// go run sender.go -d -i output file -o input file	- decode

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func find(s string, NATO []string) int {
	for i, v := range NATO {
		if v == s {
			return i
		}
	}
	return -1
}

type BintoNATO struct {
	in  io.Reader
	out io.Writer
}

func (btn BintoNATO) read(p []string) (int, error) {
	inBuffer := make([]byte, len(p))
	n, err := btn.in.Read(inBuffer)
	if err != nil {
		return n, err
	}

	c := 0
	for i := 0; i < n; i++ {
		p[i] = NATO[inBuffer[i]]
		c++
	}
	return c, nil
}

func (btn BintoNATO) decode() {
	scanner := bufio.NewScanner(btn.in)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		v := find(scanner.Text(), NATO)
		if v < 0 {
			panic("word not found at dict")
		}
		btn.out.Write([]byte{byte(v)})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func (btn BintoNATO) encode() {
       
        final := []string{}
	for {
		encoded := make([]string, 1024)
		n, err := btn.read(encoded)
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
	for i := 0; i < len(final); i += 5{
		if i+5 < len(final) {
			// win32 ln
			fmt.Fprintf(btn.out, "%s\r\n", strings.Join(final[i:(i+5)], " "))
		} else {
			fmt.Fprintf(btn.out, "%s", strings.Join(final[i:len(final)], " "))
		}
	}
}

func main() {
	inFP := flag.String("i", "", "input file")
	outFP := flag.String("o", "", "output file")
	action := flag.Bool("d", false, "use -d to specify decode")
	flag.Parse()

	if *inFP == "" || *outFP == "" {
		fmt.Printf("Usage: sender [-d] -i input file -o output file")
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
		BintoNATO := BintoNATO{inFD, outFD}
		BintoNATO.encode()
	} else {
		BintoNATO := BintoNATO{inFD, outFD}
		BintoNATO.decode()
	}

}

// NATO is our dict to convert binary to string
var NATO = []string{
        "Zero-Zero......",
	"Zero-One.......",
	"Zero-Two.......",
	"Zero-Three.....",
	"Zero-Four......",
	"Zero-Five......",
	"Zero-Six.......",
	"Zero-Seven.....",
	"Zero-Eight.....",
	"Zero-Nine......",
	"Zero-Alfa......",
	"Zero-Bravo.....",
	"Zero-Charlie...",
	"Zero-Delta.....",
	"Zero-Echo......",
	"Zero-Foxtrot...",
	"One-Zero.......",
	"One-One........",
	"One-Two........",
	"One-Three......",
	"One-Four.......",
	"One-Five.......",
	"One-Six........",
	"One-Seven......",
	"One-Eight......",
	"One-Nine.......",
	"One-Alfa.......",
	"One-Bravo......",
	"One-Charlie....",
	"One-Delta......",
	"One-Echo.......",
	"One-Foxtrot....",
	"Two-Zero.......",
	"Two-One........",
	"Two-Two........",
	"Two-Three......",
	"Two-Four.......",
	"Two-Five.......",
	"Two-Six........",
	"Two-Seven......",
	"Two-Eight......",
	"Two-Nine.......",
	"Two-Alfa.......",
	"Two-Bravo......",
	"Two-Charlie....",
	"Two-Delta......",
	"Two-Echo.......",
	"Two-Foxtrot....",
	"Three-Zero.....",
	"Three-One......",
	"Three-Two......",
	"Three-Three....",
	"Three-Four.....",
	"Three-Five.....",
	"Three-Six......",
	"Three-Seven....",
	"Three-Eight....",
	"Three-Nine.....",
	"Three-Alfa.....",
	"Three-Bravo....",
	"Three-Charlie..",
	"Three-Delta....",
	"Three-Echo.....",
	"Three-Foxtrot..",
	"Four-Zero......",
	"Four-One.......",
	"Four-Two.......",
	"Four-Three.....",
	"Four-Four......",
	"Four-Five......",
	"Four-Six.......",
	"Four-Seven.....",
	"Four-Eight.....",
	"Four-Nine......",
	"Four-Alfa......",
	"Four-Bravo.....",
	"Four-Charlie...",
	"Four-Delta.....",
	"Four-Echo......",
	"Four-Foxtrot...",
	"Five-Zero......",
	"Five-One.......",
	"Five-Two.......",
	"Five-Three.....",
	"Five-Four......",
	"Five-Five......",
	"Five-Six.......",
	"Five-Seven.....",
	"Five-Eight.....",
	"Five-Nine......",
	"Five-Alfa......",
	"Five-Bravo.....",
	"Five-Charlie...",
	"Five-Delta.....",
	"Five-Echo......",
	"Five-Foxtrot...",
	"Six-Zero.......",
	"Six-One........",
	"Six-Two........",
	"Six-Three......",
	"Six-Four.......",
	"Six-Five.......",
	"Six-Six........",
	"Six-Seven......",
	"Six-Eight......",
	"Six-Nine.......",
	"Six-Alfa.......",
	"Six-Bravo......",
	"Six-Charlie....",
	"Six-Delta......",
	"Six-Echo.......",
	"Six-Foxtrot....",
	"Seven-Zero.....",
	"Seven-One......",
	"Seven-Two......",
	"Seven-Three....",
	"Seven-Four.....",
	"Seven-Five.....",
	"Seven-Six......",
	"Seven-Seven....",
	"Seven-Eight....",
	"Seven-Nine.....",
	"Seven-Alfa.....",
	"Seven-Bravo....",
	"Seven-Charlie..",
	"Seven-Delta....",
	"Seven-Echo.....",
	"Seven-Foxtrot..",
	"Eight-Zero.....",
	"Eight-One......",
	"Eight-Two......",
	"Eight-Three....",
	"Eight-Four.....",
	"Eight-Five.....",
	"Eight-Six......",
	"Eight-Seven....",
	"Eight-Eight....",
	"Eight-Nine.....",
	"Eight-Alfa.....",
	"Eight-Bravo....",
	"Eight-Charlie..",
	"Eight-Delta....",
	"Eight-Echo.....",
	"Eight-Foxtrot..",
	"Nine-Zero......",
	"Nine-One.......",
	"Nine-Two.......",
	"Nine-Three.....",
	"Nine-Four......",
	"Nine-Five......",
	"Nine-Six.......",
	"Nine-Seven.....",
	"Nine-Eight.....",
	"Nine-Nine......",
	"Nine-Alpha.....",
	"Nine-Bravo.....",
	"Nine-Charlie...",
	"Nine-Delta.....",
	"Nine-Echo......",
	"Nine-Foxtrot...",
	"Alfa-Zero......",
	"Alfa-One.......",
	"Alfa-Two.......",
	"Alfa-Three.....",
	"Alfa-Four......",
	"Alfa-Five......",
	"Alfa-Six.......",
	"Alfa-Seven.....",
	"Alfa-Eight.....",
	"Alfa-Nine......",
	"Alfa-Alfa......",
	"Alfa-Bravo.....",
	"Alfa-Charlie...",
	"Alfa-Delta.....",
	"Alfa-Echo......",
	"Alfa-Foxtrot...",
	"Bravo-Zero.....",
	"Bravo-One......",
	"Bravo-Two......",
	"Bravo-Three....",
	"Bravo-Four.....",
	"Bravo-Five.....",
	"Bravo-Six......",
	"Bravo-Seven....",
	"Bravo-Eight....",
	"Bravo-Nine.....",
	"Bravo-Alfa.....",
	"Bravo-Bravo....",
	"Bravo-Charlie..",
	"Bravo-Delta....",
	"Bravo-Echo.....",
	"Bravo-Foxtrot..",
	"Charlie-Zero...",
	"Charlie-One....",
	"Charlie-Two....",
	"Charlie-Three..",
	"Charlie-Four...",
	"Charlie-Five...",
	"Charlie-Six....",
	"Charlie-Seven..",
	"Charlie-Eight..",
	"Charlie-Nine...",
	"Charlie-Alfa...",
	"Charlie-Bravo..",
	"Charlie-Charlie",
	"Charlie-Delta..",
	"Charlie-Echo...",
	"Charlie-Foxtrot",
	"Delta-Zero.....",
	"Delta-One......",
	"Delta-Two......",
	"Delta-Three....",
	"Delta-Four.....",
	"Delta-Five.....",
	"Delta-Six......",
	"Delta-Seven....",
	"Delta-Eight....",
	"Delta-Nine.....",
	"Delta-Alfa.....",
	"Delta-Bravo....",
	"Delta-Charlie..",
	"Delta-Delta....",
	"Delta-Echo.....",
	"Delta-Foxtrot..",
	"Echo-Zero......",
	"Echo-One.......",
	"Echo-Two.......",
	"Echo-Three.....",
	"Echo-Four......",
	"Echo-Five......",
	"Echo-Six.......",
	"Echo-Seven.....",
	"Echo-Eight.....",
	"Echo-Nine......",
	"Echo-Alfa......",
	"Echo-Bravo.....",
	"Echo-Charlie...",
	"Echo-Delta.....",
	"Echo-Echo......",
	"Echo-Foxtrot...",
	"Foxtrot-Zero...",
	"Foxtrot-One....",
	"Foxtrot-Two....",
	"Foxtrot-Three..",
	"Foxtrot-Four...",
	"Foxtrot-Five...",
	"Foxtrot-Six....",
	"Foxtrot-Seven..",
	"Foxtrot-Eight..",
	"Foxtrot-Nine...",
	"Foxtrot-Alfa...",
	"Foxtrot-Bravo..",
	"Foxtrot-Charlie",
	"Foxtrot-Delta..",
	"Foxtrot-Echo...",
	"Foxtrot-Foxtrot",
}
