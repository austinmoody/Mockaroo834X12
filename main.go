package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	//"strconv"
)

type InterchangeControlHeader struct {
	Isa ISA `json:ISA`
}

type ISA struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
	De06 string `json:"06"`
	De07 string `json:"07"`
	De08 string `json:"08"`
	De09 string `json:"09"`
	De10 string `json:"10"`
	De11 string `json:"11"`
	De12 string `json:"12"`
	De13 string `json:"13"`
	De14 string `json:"14"`
	De15 string `json:"15"`
	De16 string `json:"16"`
	Gs   []GS   `json:"GS"`
}

type GS struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
	De06 string `json:"06"`
	De07 string `json:"07"`
	De08 string `json:"08"`
	St   []ST   `json:"ST"`
}

type ST struct {
	De01      string    `json:"01"`
	De02      string    `json:"02"`
	De03      string    `json:"03"`
	Bgn       BGN       `json:"BGN"`
	Ref       REF       `json:"REF"`
	Dtp       DTP       `json:"DTP"`
	Loop1000A Loop1000A `json:"1000A"`
	Loop1000B Loop1000B `json:"1000B"`
	Ins       []INS     `json:"INS"`
}

type Loop1000A struct {
	N1 N1 `json:"N1"`
}

type Loop1000B struct {
	N1 N1 `json:"N1"`
}

type BGN struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
	De06 string `json:"06"`
	De07 string `json:"07"`
	De08 string `json:"08"`
}

type REF struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
}

type DTP struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
}

type N1 struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
}

type INS struct {
	De01      string    `json:"01"`
	De02      string    `json:"02"`
	De03      string    `json:"03"`
	De04      string    `json:"04"`
	De05      string    `json:"05"`
	De06      string    `json:"06"`
	De07      string    `json:"07"`
	De08      string    `json:"08"`
	De09      string    `json:"09"`
	De10      string    `json:"10"`
	Ref       []REF     `json:"REF"`
	Dtp       []DTP     `json:"DTP"`
	Loop2100A Loop2100A `json:"2100A"`
	Loop2100C Loop2100C `json:"2100C"`
	Loop2100F Loop2100F `json:"2100F"`
	Loop2100G Loop2100G `json:"2100G"`
	// Add Loop2300
	Loop2310 []Loop2310 `json:"2310"`
}

type Loop2100A struct {
	NM1 NM1 `json:"NM1"`
	PER PER `json:"PER"`
	N3  N3  `json:"N3"`
	N4  N4  `json:"N4"`
	DMG DMG `json:"DMG"`
	LUI LUI `json:"LUI"`
}

type Loop2100C struct {
	NM1 NM1 `json:"NM1"`
	N3  N3  `json:"N3"`
	N4  N4  `json:"N4"`
}

type Loop2100F struct {
	NM1 NM1 `json:"NM1"`
}

type Loop2100G struct {
	NM1 NM1 `json:"NM1"`
	PER PER `json:"PER"`
}

type Loop2310 struct {
	LX  LX  `json:"LX"`
	NM1 NM1 `json:"NM1"`
	N3  N3  `json:"N3"`
	N4  N4  `json:"N4"`
	PLA PLA `json:"PLA"`
}

type NM1 struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
	De06 string `json:"06"`
	De07 string `json:"07"`
	De08 string `json:"08"`
	De09 string `json:"09"`
	De10 string `json:"10"`
}

type PER struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
	De06 string `json:"06"`
	De07 string `json:"07"`
	De08 string `json:"08"`
}

type N3 struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
}

type N4 struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
	De06 string `json:"06"`
}

type DMG struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
}

type LUI struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
}

type LX struct {
	De01 string `json:"01"`
}

type PLA struct {
	De01 string `json:"01"`
	De02 string `json:"02"`
	De03 string `json:"03"`
	De04 string `json:"04"`
	De05 string `json:"05"`
}

func main() {
	jsonFile, err := os.Open("/Users/amoody/Downloads/X12-834-JSON.json")

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var x12 InterchangeControlHeader
	json.Unmarshal(byteValue, &x12)

	fmt.Printf("ISA*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s*%s~\r\n",
		x12.Isa.De01,
		x12.Isa.De02,
		x12.Isa.De03,
		x12.Isa.De04,
		x12.Isa.De05,
		x12.Isa.De06,
		x12.Isa.De07,
		x12.Isa.De08,
		x12.Isa.De09,
		x12.Isa.De10,
		x12.Isa.De11,
		x12.Isa.De12,
		x12.Isa.De13,
		x12.Isa.De14,
		x12.Isa.De15,
		x12.Isa.De16)

	for _, gs := range x12.Isa.Gs {
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

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
}
