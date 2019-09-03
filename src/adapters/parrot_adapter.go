package adapters

import (
	"encoding/json"
	"io/ioutil"
	"laches1sm/demo-parrot-service/infrastructure"
	"log"
	"net/http"
)

type ParrotHTTPAdapter struct {
	Logger log.Logger
	Infra  infrastructure.ParrotInfra
}

func NewParrotHTTPAdapter(logger log.Logger, infra infrastructure.ParrotInfra) *ParrotHTTPAdapter {
	return &ParrotHTTPAdapter{
		Logger: logger,
		Infra:  infra,
	}
}

// GibeParrot accepts only GET HTTP method. You get a parrot, you get a parrot, everyone gets a parrot!
func (adapter *ParrotHTTPAdapter) GibeParrot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		adapter.Logger.Printf(`nope`)
		_ = marshalAndWriteErrorResponse(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		adapter.Logger.Printf(`whoops there's an error while reading request body`)
		_ = marshalAndWriteErrorResponse(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// read resp body...
	resp, err := adapter.Infra.GetParrot(body)
	if err != nil {
		return
	}
	parrotGet, err := json.Marshal(resp)
	if err != nil {
		return
	}
	writeResponse(w, parrotGet, http.StatusOK)
}

func (adapter *ParrotHTTPAdapter) AddParrot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		adapter.Logger.Printf(`lol_nope`)
		_ = marshalAndWriteErrorResponse(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		adapter.Logger.Printf(`whoops there's an error while reading request body`)
		_ = marshalAndWriteErrorResponse(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	resp, err := adapter.Infra.AddParrot(body)
	if err != nil {
		return
	}
	parrotAdd, err := json.Marshal(resp)
	if err != nil {
		return
	}
	writeResponse(w, parrotAdd, http.StatusCreated)
}
