package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	jsonFile, err := os.Open("/Users/amoody/Downloads/X12-834-JSON.json")

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var x12 InterchangeControlHeader
	json.Unmarshal(byteValue, &x12)
	defer jsonFile.Close()

	fmt.Printf("ISA*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s~\r\n",
		x12.ISA.De01,
		x12.ISA.De02,
		x12.ISA.De03,
		x12.ISA.De04,
		x12.ISA.De05,
		x12.ISA.De06,
		x12.ISA.De07,
		x12.ISA.De08,
		x12.ISA.De09,
		x12.ISA.De10,
		x12.ISA.De11,
		x12.ISA.De12,
		x12.ISA.De13,
		x12.ISA.De14,
		x12.ISA.De15,
		x12.ISA.De16)

	for _, gs := range x12.ISA.Gs {
		fmt.Printf("GS*%s*%s*%s*%s*%s*%s*%s*%s~\r\n",
			gs.De01,
			gs.De02,
			gs.De03,
			gs.De04,
			gs.De05,
			gs.De06,
			gs.De07,
			gs.De08)

		groupControlNumber := gs.De02

		for _, st := range gs.St {
			fmt.Printf("ST*%s*%s*%s~\r\n", st.De01, st.De02, st.De03)

			fmt.Printf("BGN*%s*%s*%s*%s*%s*%s*%s*%s~\r\n",
				st.Bgn.De01,
				st.Bgn.De02,
				st.Bgn.De03,
				st.Bgn.De04,
				st.Bgn.De05,
				st.Bgn.De06,
				st.Bgn.De07,
				st.Bgn.De08)

			fmt.Printf("REF*%s*%s~\r\n", st.Ref.De01, st.Ref.De02)

			fmt.Printf("DTP*%s*%s*%s~\r\n", st.Dtp.De01, st.Dtp.De02, st.Dtp.De03)

			fmt.Printf("N1*%s*%s*%s*%s~\r\n", st.Loop1000A.N1.De01, st.Loop1000A.N1.De02, st.Loop1000A.N1.De03, st.Loop1000A.N1.De04)

			fmt.Printf("N1*%s*%s*%s*%s~\r\n", st.Loop1000B.N1.De01, st.Loop1000B.N1.De02, st.Loop1000B.N1.De03, st.Loop1000B.N1.De04)

			for _, ins := range st.Ins {

				fmt.Printf("INS*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s~\r\n",
					ins.De01,
					ins.De02,
					ins.De03,
					ins.De04,
					ins.De05,
					ins.De06,
					ins.De07,
					ins.De08,
					ins.De09,
					ins.De10)

				for _, ref := range ins.Ref {
					fmt.Printf("REF*%s*%s~\r\n", ref.De01, ref.De02)
				}

				for _, dtp := range ins.Dtp {
					fmt.Printf("DTP*%s*%s*%s~\r\n", dtp.De01, dtp.De02, dtp.De03)
				}

				// Loop2100A
				fmt.Printf("NM1*%s*%s*%s*%s*%s*%s*%s~\r\n", ins.Loop2100A.NM1.De01,
					ins.Loop2100A.NM1.De02,
					ins.Loop2100A.NM1.De03,
					ins.Loop2100A.NM1.De04,
					ins.Loop2100A.NM1.De05,
					ins.Loop2100A.NM1.De06,
					ins.Loop2100A.NM1.De07)

				fmt.Printf("PER*%s*%s*%s*%s*%s*%s*%s*%s~\r\n", ins.Loop2100A.PER.De01,
					ins.Loop2100A.PER.De02,
					ins.Loop2100A.PER.De03,
					ins.Loop2100A.PER.De04,
					ins.Loop2100A.PER.De05,
					ins.Loop2100A.PER.De06,
					ins.Loop2100A.PER.De07,
					ins.Loop2100A.PER.De08)

				fmt.Printf("N3*%s*%s~\r\n", ins.Loop2100A.N3.De01, ins.Loop2100A.N3.De02)

				fmt.Printf("N4*%s*%s*%s*%s*%s*%s~\r\n", ins.Loop2100A.N4.De01,
					ins.Loop2100A.N4.De02,
					ins.Loop2100A.N4.De03,
					ins.Loop2100A.N4.De04,
					ins.Loop2100A.N4.De05,
					ins.Loop2100A.N4.De06)

				fmt.Printf("DMG*%s*%s*%s*%s*%s~\r\n", ins.Loop2100A.DMG.De01,
					ins.Loop2100A.DMG.De02,
					ins.Loop2100A.DMG.De03,
					ins.Loop2100A.DMG.De04,
					ins.Loop2100A.DMG.De05)

				fmt.Printf("LUI*%s*%s*%s*%s~\r\n", ins.Loop2100A.LUI.De01,
					ins.Loop2100A.LUI.De02,
					ins.Loop2100A.LUI.De03,
					ins.Loop2100A.LUI.De04)

				// Loop2100C
				fmt.Printf("NM1*%s*%s~\r\n",
					ins.Loop2100C.NM1.De01,
					ins.Loop2100C.NM1.De02)

				fmt.Printf("N3*%s*%s~\r\n",
					ins.Loop2100C.N3.De01,
					ins.Loop2100C.N3.De02)

				fmt.Printf("N4*%s*%s*%s~\r\n",
					ins.Loop2100C.N4.De01,
					ins.Loop2100C.N4.De02,
					ins.Loop2100C.N4.De03)

				// Loop2100F
				if ins.Loop2100F.NM1.De01 != "" {
					fmt.Printf("NM1*%s*%s*%s*%s*%s*%s*%s~\r\n",
						ins.Loop2100F.NM1.De01,
						ins.Loop2100F.NM1.De02,
						ins.Loop2100F.NM1.De03,
						ins.Loop2100F.NM1.De04,
						ins.Loop2100F.NM1.De05,
						ins.Loop2100F.NM1.De06,
						ins.Loop2100F.NM1.De07)
				}

				// Loop2100G
				if ins.Loop2100G.NM1.De01 != "" {
					fmt.Printf("NM1*%s*%s*%s*%s*%s*%s*%s~\r\n",
						ins.Loop2100G.NM1.De01,
						ins.Loop2100G.NM1.De02,
						ins.Loop2100G.NM1.De03,
						ins.Loop2100G.NM1.De04,
						ins.Loop2100G.NM1.De05,
						ins.Loop2100G.NM1.De06,
						ins.Loop2100G.NM1.De07)

					if ins.Loop2100G.PER.De01 != "" {
						fmt.Printf("PER*%s*%s*%s*%s*%s*%s*%s*%s~\r\n",
							ins.Loop2100G.PER.De01,
							ins.Loop2100G.PER.De02,
							ins.Loop2100G.PER.De03,
							ins.Loop2100G.PER.De04,
							ins.Loop2100G.PER.De05,
							ins.Loop2100G.PER.De06,
							ins.Loop2100G.PER.De07,
							ins.Loop2100G.PER.De08)
					}
				}

				// TODO - Add Loop 2300 in code & Mockaroo

				for _, l2310 := range ins.Loop2310 {
					fmt.Printf("LX*%s~\r\n", l2310.LX.De01)
					fmt.Printf("NM1*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s~\r\n",
						l2310.NM1.De01,
						l2310.NM1.De02,
						l2310.NM1.De03,
						l2310.NM1.De04,
						l2310.NM1.De05,
						l2310.NM1.De06,
						l2310.NM1.De07,
						l2310.NM1.De08,
						l2310.NM1.De09,
						l2310.NM1.De10)
					fmt.Printf("N3*%s~\r\n", l2310.N3.De01)
					fmt.Printf("N4*%s*%s*%s~\r\n",
						l2310.N4.De01,
						l2310.N4.De02,
						l2310.N4.De03)
					fmt.Printf("PLA*%s*%s*%s*%s*%s~\r\n",
						l2310.PLA.De01,
						l2310.PLA.De02,
						l2310.PLA.De03,
						l2310.PLA.De04,
						l2310.PLA.De05)

				}
			}
		}

		// End of Function Group Header
		fmt.Printf("GE*%s*%s~\r\n", "?", groupControlNumber)
	}
}
