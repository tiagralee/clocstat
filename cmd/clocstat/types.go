package main

type ClocResult struct {
	Sum Inner `json:"SUM"`
	Go  Inner `json:"Go"`
}

type Inner struct {
	Code    int `json:"code"`
	Blank   int `json:"blank"`
	Comment int `json:"comment"`
	NFiles  int `json:"nFiles"`
}
