package main

type InterchangeControlHeader struct {
	ISA ISA `json:"ISA""`
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
