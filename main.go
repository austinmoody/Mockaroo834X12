package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	elementDelimiter    = flag.String("element", "*", "Element Delimiter")
	subElementDelimiter = flag.String("subelement", "/", "Sub-Element Delimiter")
	segmentDelimiter    = flag.String("segment", "~\r\n", "Segment Delimiter")
	inputFile           = flag.String("input", "", "Input File")
)

func main() {

	flag.Parse()

	delimiters := X12Delimiters{}
	delimiters.Segment = *segmentDelimiter
	delimiters.Element = *elementDelimiter
	delimiters.SubElement = *subElementDelimiter

	var x12 X12n834
	if *inputFile == "" {
		x12 = GetX12FromStdin()
	} else {
		x12 = GetX12FromFile(*inputFile)
	}

	fmt.Print(x12.ISA.String(delimiters))

	for _, gs := range x12.ISA.Gs {
		groupControlNumber := gs.De02

		fmt.Print(gs.String(delimiters))

		setHeaderCount := len(gs.St)
		transactionSetSegmentCount := 0
		for _, st := range gs.St {

			fmt.Print(st.String(delimiters))
			transactionSetSegmentCount++

			fmt.Print(st.Bgn.String(delimiters))
			transactionSetSegmentCount++

			fmt.Print(st.Ref.String(delimiters))
			transactionSetSegmentCount++

			fmt.Print(st.Dtp.String(delimiters))
			transactionSetSegmentCount++

			fmt.Print(st.Loop1000A.String(delimiters))
			transactionSetSegmentCount++

			fmt.Print(st.Loop1000B.String(delimiters))
			transactionSetSegmentCount++

			for _, ins := range st.Ins {

				fmt.Print(ins.String(delimiters))
				transactionSetSegmentCount++

				for _, ref := range ins.Ref {
					fmt.Print(ref.String(delimiters))
					transactionSetSegmentCount++
				}

				for _, dtp := range ins.Dtp {
					fmt.Print(dtp.String(delimiters))
					transactionSetSegmentCount++
				}

				// Loop2100A
				fmt.Print(ins.Loop2100A.NM1.String(delimiters, map[string]string{"MaxElements": "7"}))
				transactionSetSegmentCount++
				fmt.Print(ins.Loop2100A.PER.String(delimiters))
				transactionSetSegmentCount++
				fmt.Print(ins.Loop2100A.N3.String(delimiters))
				transactionSetSegmentCount++
				fmt.Print(ins.Loop2100A.N4.String(delimiters))
				transactionSetSegmentCount++
				fmt.Print(ins.Loop2100A.DMG.String(delimiters))
				transactionSetSegmentCount++
				fmt.Print(ins.Loop2100A.LUI.String(delimiters))
				transactionSetSegmentCount++

				// Loop2100C
				fmt.Print(ins.Loop2100C.NM1.String(delimiters, map[string]string{"MaxElements": "2"}))
				transactionSetSegmentCount++
				fmt.Print(ins.Loop2100C.N3.String(delimiters))
				transactionSetSegmentCount++
				fmt.Print(ins.Loop2100C.N4.String(delimiters, map[string]string{"MaxElements": "3"}))
				transactionSetSegmentCount++

				// Loop2100F
				if ins.Loop2100F.NM1.De01 != "" {
					fmt.Print(ins.Loop2100F.NM1.String(delimiters, map[string]string{"MaxElements": "7"}))
					transactionSetSegmentCount++
				}

				// Loop2100G
				if ins.Loop2100G.NM1.De01 != "" {
					fmt.Print(ins.Loop2100G.NM1.String(delimiters, map[string]string{"MaxElements": "7"}))
					transactionSetSegmentCount++

					if ins.Loop2100G.PER.De01 != "" {
						fmt.Print(ins.Loop2100G.PER.String(delimiters))
						transactionSetSegmentCount++
					}
				}

				// TODO - Add Loop 2300 in code & Mockaroo

				for _, l2310 := range ins.Loop2310 {
					fmt.Print(l2310.LX.String(delimiters))
					transactionSetSegmentCount++
					fmt.Print(l2310.NM1.String(delimiters))
					transactionSetSegmentCount++
					fmt.Print(l2310.N3.String(delimiters, map[string]string{"MaxElements": "1"}))
					transactionSetSegmentCount++
					fmt.Print(l2310.N4.String(delimiters, map[string]string{"MaxElements": "3"}))
					transactionSetSegmentCount++
					fmt.Print(l2310.PLA.String(delimiters))
					transactionSetSegmentCount++
				}
			}

			// TODO - I hate this copy/paste of transactionSetSegmentCount++ has to be a better way.
			fmt.Printf(
				"%s%s",
				strings.Join([]string{"SE", strconv.FormatInt(int64(transactionSetSegmentCount), 10), st.De02},
					delimiters.Element),
				delimiters.Segment,
			)
		}

		// End of Function Group Header
		fmt.Printf("%s%s", strings.Join(
			[]string{"GE", strconv.FormatInt(int64(setHeaderCount), 10), groupControlNumber},
			delimiters.Element,
		), delimiters.Segment)
	}

	fmt.Printf("%s%s", strings.Join(
		[]string{"IEA", strconv.FormatInt(int64(len(x12.ISA.Gs)), 10), x12.ISA.De13},
		delimiters.Element,
	), delimiters.Segment)

}

func GetX12FromFile(fileName string) X12n834 {
	var err error
	jsonFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err.Error())
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var x12 X12n834
	err = json.Unmarshal(byteValue, &x12)
	if err != nil {
		log.Fatal("Error unmarshalling JSON: " + err.Error())
	}

	defer func(jsonFile *os.File) {
		err = jsonFile.Close()
		if err != nil {
			log.Fatal("Error closing file: " + err.Error())
		}
	}(jsonFile)

	return x12
}

func GetX12FromStdin() X12n834 {
	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)

	var jsonBuffer bytes.Buffer

	for {
		n, err := reader.Read(buf[:cap(buf)])
		jsonBuffer.Write(buf[:n])

		if n == 0 {

			if err == nil {
				continue
			}

			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}

	var x12 X12n834
	err := json.Unmarshal(jsonBuffer.Bytes(), &x12)
	if err != nil {
		log.Fatal(err)
	}

	return x12
}
