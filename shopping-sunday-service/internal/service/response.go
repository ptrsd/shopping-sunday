package service

type ShoppingSundayResponse struct {
	IsShoppingSunday bool     `json:"isShoppingSunday"`
	Reasons          []Reason `json:"reasons"`
}

type Reason struct {
	Message string `json:"message"`
	Id      int    `json:"id"`
}
