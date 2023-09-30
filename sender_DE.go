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
	for i := 0; i < len(final); i += 4{
		if i+4 < len(final) {
			// win32 ln
			fmt.Fprintf(btn.out, "%s\r\n", strings.Join(final[i:(i+4)], " "))
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
        "Null-Null..........",
	"Null-Eins..........",
	"Null-Zwei..........",
	"Null-Drei..........",
	"Null-Vier..........",
	"Null-Fünf..........",
	"Null-Sechs.........",
	"Null-Sieben........",
	"Null-Acht..........",
	"Null-Neun..........",
	"Null-Anton.........",
	"Null-Berta.........",
	"Null-Cäsar.........",
	"Null-Dora..........",
	"Null-Emil..........",
	"Null-Friedrich.....",
	"Eins-Null..........",
	"Eins-Eins..........",
	"Eins-Zwei..........",
	"Eins-Drei..........",
	"Eins-Vier..........",
	"Eins-Fünf..........",
	"Eins-Sechs.........",
	"Eins-Sieben........",
	"Eins-Acht..........",
	"Eins-Neun..........",
	"Eins-Anton.........",
	"Eins-Berta.........",
	"Eins-Cäsar.........",
	"Eins-Dora..........",
	"Eins-Emil..........",
	"Eins-Friedrich.....",
	"Zwei-Null..........",
	"Zwei-Eins..........",
	"Zwei-Zwei..........",
	"Zwei-Drei..........",
	"Zwei-Vier..........",
	"Zwei-Fünf..........",
	"Zwei-Sechs.........",
	"Zwei-Sieben........",
	"Zwei-Acht..........",
	"Zwei-Neun..........",
	"Zwei-Anton.........",
	"Zwei-Berta.........",
	"Zwei-Cäsar.........",
	"Zwei-Dora..........",
	"Zwei-Emil..........",
	"Zwei-Friedrich.....",
	"Drei-Null..........",
	"Drei-Eins..........",
	"Drei-Zwei..........",
	"Drei-Drei..........",
	"Drei-Vier..........",
	"Drei-Fünf..........",
	"Drei-Sechs.........",
	"Drei-Sieben........",
	"Drei-Acht..........",
	"Drei-Neun..........",
	"Drei-Anton.........",
	"Drei-Berta.........",
	"Drei-Cäsar.........",
	"Drei-Dora..........",
	"Drei-Emil..........",
	"Drei-Friedrich.....",
	"Vier-Null..........",
	"Vier-Eins..........",
	"Vier-Zwei..........",
	"Vier-Drei..........",
	"Vier-Vier..........",
	"Vier-Fünf..........",
	"Vier-Sechs.........",
	"Vier-Sieben........",
	"Vier-Acht..........",
	"Vier-Neun..........",
	"Vier-Anton.........",
	"Vier-Berta.........",
	"Vier-Cäsar.........",
	"Vier-Dora..........",
	"Vier-Emil..........",
	"Vier-Friedrich.....",
	"Fünf-Null..........",
	"Fünf-Eins..........",
	"Fünf-Zwei..........",
	"Fünf-Drei..........",
	"Fünf-Vier..........",
	"Fünf-Fünf..........",
	"Fünf-Sechs.........",
	"Fünf-Sieben........",
	"Fünf-Acht..........",
	"Fünf-Neun..........",
	"Fünf-Anton.........",
	"Fünf-Berta.........",
	"Fünf-Cäsar.........",
	"Fünf-Dora..........",
	"Fünf-Emil..........",
	"Fünf-Friedrich.....",
	"Sechs-Null.........",
	"Sechs-Eins.........",
	"Sechs-Zwei.........",
	"Sechs-Drei.........",
	"Sechs-Vier.........",
	"Sechs-Fünf.........",
	"Sechs-Sechs........",
	"Sechs-Sieben.......",
	"Sechs-Acht.........",
	"Sechs-Neun.........",
	"Sechs-Anton........",
	"Sechs-Berta........",
	"Sechs-Cäsar........",
	"Sechs-Dora.........",
	"Sechs-Emil.........",
	"Sechs-Friedrich....",
	"Sieben-Null........",
	"Sieben-Eins........",
	"Sieben-Zwei........",
	"Sieben-Drei........",
	"Sieben-Vier........",
	"Sieben-Fünf........",
	"Sieben-Sechs.......",
	"Sieben-Sieben......",
	"Sieben-Acht........",
	"Sieben-Neun........",
	"Sieben-Anton.......",
	"Sieben-Berta.......",
	"Sieben-Cäsar.......",
	"Sieben-Dora........",
	"Sieben-Emil........",
	"Sieben-Friedrich...",
	"Acht-Null..........",
	"Acht-Eins..........",
	"Acht-Zwei..........",
	"Acht-Drei..........",
	"Acht-Vier..........",
	"Acht-Fünf..........",
	"Acht-Sechs.........",
	"Acht-Sieben........",
	"Acht-Acht..........",
	"Acht-Neun..........",
	"Acht-Anton.........",
	"Acht-Berta.........",
	"Acht-Cäsar.........",
	"Acht-Dora..........",
	"Acht-Emil..........",
	"Acht-Friedrich.....",
	"Neun-Null..........",
	"Neun-Eins..........",
	"Neun-Zwei..........",
	"Neun-Drei..........",
	"Neun-Vier..........",
	"Neun-Fünf..........",
	"Neun-Sechs.........",
	"Neun-Sieben........",
	"Nein-Acht..........",
	"Neun-Neun..........",
	"Neun-Anton----.....",
	"Neun-Berta.........",
	"Neun-Cäsar.........",
	"Neun-Dora..........",
	"Neun-Emil..........",
	"Neun-Friedrich.....",
	"Anton-Null.........",
	"Anton-Eins.........",
	"Anton-Zwei.........",
	"Anton-Drei.........",
	"Anton-Vier.........",
	"Anton-Fünf.........",
	"Anton-Sechs........",
	"Anton-Sieben.......",
	"Anton-Acht.........",
	"Anton-Neun.........",
	"Anton-Anton........",
	"Anton-Berta........",
	"Anton-Cäsar........",
	"Anton-Dora.........",
	"Anton-Emil.........",
	"Anton-Friedrich....",
	"Berta-Null.........",
	"Berta-Eins.........",
	"Berta-Zwei.........",
	"Berta-Drei.........",
	"Berta-Vier.........",
	"Berta-Fünf.........",
	"Berta-Sechs........",
	"Berta-Sieben.......",
	"Berta-Acht.........",
	"Berta-Neun.........",
	"Berta-Anton........",
	"Berta-Berta........",
	"Berta-Cäsar........",
	"Berta-Dora.........",
	"Berta-Emil.........",
	"Berta-Friedrich....",
	"Cäsar-Null.........",
	"Cäsar-Eins.........",
	"Cäsar-Zwei.........",
	"Cäsar-Drei.........",
	"Cäsar-Vier.........",
	"Cäsar-Fünf.........",
	"Cäsar-Sechs........",
	"Cäsar-Sieben.......",
	"Cäsar-Acht.........",
	"Cäsar-Neun.........",
	"Cäsar-Anton........",
	"Cäsar-Berta........",
	"Cäsar-Cäsar........",
	"Cäsar-Dora.........",
	"Cäsar-Emil.........",
	"Cäsar-Friedrich....",
	"Dora-Null..........",
	"Dora-Eins..........",
	"Dora-Zwei..........",
	"Dora-Drei..........",
	"Dora-Vier..........",
	"Dora-Fünf..........",
	"Dora-Sechs.........",
	"Dora-Sieben........",
	"Dora-Acht..........",
	"Dora-Neun..........",
	"Dora-Anton.........",
	"Dora-Berta.........",
	"Dora-Cäsar.........",
	"Dora-Dora..........",
	"Dora-Emil..........",
	"Dora-Friedrich.....",
	"Emil-Null..........",
	"Emil-Eins..........",
	"Emil-Zwei..........",
	"Emil-Drei..........",
	"Emil-Vier..........",
	"Emil-Fünf..........",
	"Emil-Sechs.........",
	"Emil-Sieben........",
	"Emil-Acht..........",
	"Emil-Neun..........",
	"Emil-Anton.........",
	"Emil-Berta.........",
	"Emil-Cäsar.........",
	"Emil-Dora..........",
	"Emil-Emil..........",
	"Emil-Friedrich.....",
	"Friedrich-Null.....",
	"Friedrich-Eins.....",
	"Friedrich-Zwei.....",
	"Friedrich-Drei.....",
	"Friedrich-Vier.....",
	"Friedrich-Fünf.....",
	"Friedrich-Sechs....",
	"Friedrich-Sieben...",
	"Friedrich-Acht.....",
	"Friedrich-Neun.....",
	"Friedrich-Anton....",
	"Friedrich-Berta....",
	"Friedrich-Cäsar....",
	"Friedrich-Dora.....",
	"Friedrich-Emil.....",
	"Friedrich-Friedrich",
}