package adapters

import (
	"encoding/json"
	"io/ioutil"
	"laches1sm/demo-parrot-service/infrastructure"
	"log"
	"net/http"
)

type ParrotHTTPAdapter struct {
	logger log.Logger
	infra  infrastructure.ParrotInfra
}

// GibeParrot accepts only GET HTTP method. You get a parrot, you get a parrot, everyone gets a parrot!
func (adapter *ParrotHTTPAdapter) GibeParrot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		adapter.logger.Printf(`nope`)
		_ = marshalAndWriteErrorResponse(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		adapter.logger.Printf(`whoops there's an error while reading request body`)
		_ = marshalAndWriteErrorResponse(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// read resp body...
	resp, err := adapter.infra.GetParrot(body)
	if err != nil {
		return
	}
	parrotGet, err := json.Marshal(resp)
	if err != nil {
		return
	}
	writeSuccessResponse(w, parrotGet)
}

func (adapter *ParrotHTTPAdapter) AddParrot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		adapter.logger.Printf(`lol_nope`)
		_ = marshalAndWriteErrorResponse(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		adapter.logger.Printf(`whoops there's an error while reading request body`)
		_ = marshalAndWriteErrorResponse(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	resp, err := adapter.infra.AddParrot(body)
	if err != nil {
		return
	}
	parrotAdd, err := json.Marshal(resp)
	if err != nil {
		return
	}
	writeSuccessResponse(w, parrotAdd)
}
