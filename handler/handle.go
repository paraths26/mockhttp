package handler

import (
	"log"
	"net/http"

	"github.com/paraths26/mockhttp/service"
)

//Mock :
type Mock struct {
	HttpClient service.Client
	CurrSrvURL string
}

//Handle :
func (m *Mock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	curr, err := service.GetCurrency(m.HttpClient, m.CurrSrvURL, "GET")
	if err != nil {
		log.Println("Error fetching currency service response :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(curr)
	return
}
