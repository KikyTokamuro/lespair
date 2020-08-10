package main

// Shedule struct
type Shedule struct {
	Group1 Group1 `json:"group1"`
	Group2 Group2 `json:"group2"`
	Time   Time   `json:"time"`
}

// Group1 struct
type Group1 struct {
	Monday    Monday    `json:"monday"`
	Tuesday   Tuesday   `json:"tuesday"`
	Wednesday Wednesday `json:"wednesday"`
	Thursday  Thursday  `json:"thursday"`
	Friday    Friday    `json:"friday"`
	Saturday  Saturday  `json:"saturday"`
	Sunday    Sunday    `json:"Sunday"`
}

// Group2 struct
type Group2 struct {
	Monday    Monday    `json:"monday"`
	Tuesday   Tuesday   `json:"tuesday"`
	Wednesday Wednesday `json:"wednesday"`
	Thursday  Thursday  `json:"thursday"`
	Friday    Friday    `json:"friday"`
	Saturday  Saturday  `json:"saturday"`
	Sunday    Sunday    `json:"Sunday"`
}

// Monday struct
type Monday struct {
	S1 S1 `json:"S1"`
	S2 S2 `json:"S2"`
}

// Tuesday struct
type Tuesday struct {
	S1 S1 `json:"S1"`
	S2 S2 `json:"S2"`
}

// Wednesday struct
type Wednesday struct {
	S1 S1 `json:"S1"`
	S2 S2 `json:"S2"`
}

// Thursday struct
type Thursday struct {
	S1 S1 `json:"S1"`
	S2 S2 `json:"S2"`
}

// Friday struct
type Friday struct {
	S1 S1 `json:"S1"`
	S2 S2 `json:"S2"`
}

// Saturday struct
type Saturday struct {
	S1 S1 `json:"S1"`
	S2 S2 `json:"S2"`
}

// Sunday struct
type Sunday struct {
	S1 S1 `json:"S1"`
	S2 S2 `json:"S2"`
}

// S1 struct
type S1 struct {
	P1 string `json:"1"`
	P2 string `json:"2"`
	P3 string `json:"3"`
	P4 string `json:"4"`
	P5 string `json:"5"`
}

// S2 struct
type S2 struct {
	P1 string `json:"1"`
	P2 string `json:"2"`
	P3 string `json:"3"`
	P4 string `json:"4"`
	P5 string `json:"5"`
}

// Time struct
type Time struct {
	P1 string `json:"1"`
	P2 string `json:"2"`
	P3 string `json:"3"`
	P4 string `json:"4"`
	P5 string `json:"5"`
}
