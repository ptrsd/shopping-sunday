package service

type ShoppingSundayResponse struct {
	IsShoppingSunday bool     `json:"isShoppingSunday"`
	Reasons          []string `json:"reasons"`
}
