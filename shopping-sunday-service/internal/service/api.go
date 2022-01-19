package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"shopping-sunday-service/pkg/sunday"
	"time"
)

type RestApi struct {
	Config Config
}

func (server *RestApi) Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/sunday/{date}", CalculatorHandler)
	return http.ListenAndServe(":"+server.Config.Port, router)
}

func CalculatorHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]
	parsedDate, err := time.Parse(sunday.ShoppingSundayFormat, date)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var reasonsList []Reason
	shopping, reasons := sunday.IsShopping(parsedDate)
	for _, reason := range reasons {
		reasonsList = append(reasonsList, Reason{reason.Message, reason.Id})
	}

	response := ShoppingSundayResponse{IsShoppingSunday: shopping, Reasons: reasonsList}
	body, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}
