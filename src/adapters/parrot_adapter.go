package adapters

import (
	"io/ioutil"
	"log"
	"net/http"
	"errors"
)

type ParrotHTTPAdapter struct {
	logger log.Logger
}

// GibeParrot accepts only GET HTTP method. You get a parrot, you get a parrot, everyone gets a parrot!
func (adapter *ParrotHTTPAdapter) GibeParrot(w http.ResponseWriter, r *http.Request) *models.Parrot, error {
	if r.Method != http.MethodGet {
		adapter.logger.Printf(`nope`)
		_ = marshalAndWriteErrorResponse(w, http.StatusText(http.StatusMethodNotAllowed), r.Context(), nil)
		return nil, errors.New(`status not allowed`)
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil{
	    adapter.logger.Printf(`whoops there's an error while reading request body`)
	    _ = marshalAndWriteErrorResponse(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	    return nil, err
	}
	// read resp body...
	parrot := &domain.Parrot{}
	// add infra here
}

func (adapter *ParrotHTTPAdapter) AddParrot (w http.ResponseWriter, r *http.Request){
    if r.Method != http.MethodPost{
	adapter.logger.Printf(`lol_nope`)
	_ = marshalAndWriteErrorResponse(w, http.StatusText(http.StatusMethodNotAllowed), r.Context(), nil)
	return
    }
    body, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil{
	adapter.logger.Printf(`whoops there's an error while reading request body`)
	_ = marshalAndWriteErrorResponse(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
    }
    parrot := &domain.Parrot{}
    // add infra here
}
