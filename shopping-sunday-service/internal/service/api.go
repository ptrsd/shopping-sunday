package service

import (
	"encoding/json"
	"fmt"
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
	router.HandleFunc("/sunday", CalculatorHandler)
	return http.ListenAndServe(":"+server.Config.Port, router)
}

func CalculatorHandler(w http.ResponseWriter, r *http.Request) {
	var parsedDate time.Time
	var parseErr error

	if date, ok := mux.Vars(r)["date"]; ok {
		parsedDate, parseErr = time.Parse(sunday.ShoppingSundayFormat, date)
	} else {
		parsedDate = time.Now()
	}

	if parseErr != nil {
		errorResponse := ErrorResponse{fmt.Sprintf("selected date has invalid format, expected format %s", sunday.ShoppingSundayFormat)}
		errorBody, err := json.Marshal(errorResponse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(errorBody)
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
