package service

import (
	"servicesdemo/accountservice/dbclient"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// DBClient is a instance of accounts db client
var DBClient dbclient.IBoltClient

// GetAccount returns a account by id
func GetAccount(w http.ResponseWriter, r *http.Request) {

	var accountID = mux.Vars(r)["accountId"]

	account, err := DBClient.QueryAccount(accountID)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, jsonErr := json.Marshal(account)

	if jsonErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("error in json: %v", jsonErr)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
