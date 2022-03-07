package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	delimiters := X12Delimiters{}
	delimiters.Segment = "~\r\n"
	delimiters.Element = "*"
	delimiters.SubElement = "/"

	var x12 = GetX12FromStdin()

	/*
		var err error
		jsonFile, err := os.Open("/Users/amoody/Downloads/X12-834-JSON.json")

		if err != nil {
			fmt.Println(err)
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
	*/
	fmt.Print(x12.ISA.String(delimiters))

	for _, gs := range x12.ISA.Gs {
		groupControlNumber := gs.De02

		fmt.Print(gs.String(delimiters))

		setHeaderCount := len(gs.St)
		for _, st := range gs.St {

			fmt.Print(st.String(delimiters))
			fmt.Print(st.Bgn.String(delimiters))
			fmt.Print(st.Ref.String(delimiters))
			fmt.Print(st.Dtp.String(delimiters))
			fmt.Print(st.Loop1000A.N1.String(delimiters))
			fmt.Print(st.Loop1000B.N1.String(delimiters))

			for _, ins := range st.Ins {

				fmt.Print(ins.String(delimiters))

				for _, ref := range ins.Ref {
					fmt.Print(ref.String(delimiters))
				}

				for _, dtp := range ins.Dtp {
					fmt.Print(dtp.String(delimiters))
				}

				// Loop2100A
				fmt.Print(ins.Loop2100A.NM1.String(delimiters))
				fmt.Print(ins.Loop2100A.PER.String(delimiters))
				fmt.Print(ins.Loop2100A.N3.String(delimiters))
				fmt.Print(ins.Loop2100A.N4.String(delimiters))
				fmt.Print(ins.Loop2100A.DMG.String(delimiters))
				fmt.Print(ins.Loop2100A.LUI.String(delimiters))

				// Loop2100C
				fmt.Print(ins.Loop2100C.NM1.String(delimiters))
				fmt.Print(ins.Loop2100C.N3.String(delimiters))
				fmt.Print(ins.Loop2100C.N4.String(delimiters))

				// Loop2100F
				if ins.Loop2100F.NM1.De01 != "" {
					fmt.Print(ins.Loop2100F.NM1.String(delimiters))
				}

				// Loop2100G
				if ins.Loop2100G.NM1.De01 != "" {
					fmt.Print(ins.Loop2100G.NM1.String(delimiters))

					if ins.Loop2100G.PER.De01 != "" {
						fmt.Print(ins.Loop2100G.PER.String(delimiters))
					}
				}

				// TODO - Add Loop 2300 in code & Mockaroo

				for _, l2310 := range ins.Loop2310 {
					fmt.Print(l2310.LX.String(delimiters))
					fmt.Print(l2310.NM1.String(delimiters))
					fmt.Print(l2310.N3.String(delimiters))
					fmt.Print(l2310.N4.String(delimiters))
					fmt.Print(l2310.PLA.String(delimiters))
				}
			}
		}

		// End of Function Group Header
		fmt.Printf("GE*%d*%s~\r\n", setHeaderCount, groupControlNumber)
	}

	fmt.Printf("IEA*%d*%s~\r\n", len(x12.ISA.Gs), x12.ISA.De13)

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
