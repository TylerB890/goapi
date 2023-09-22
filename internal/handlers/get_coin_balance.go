package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tylerb890/goapi/api"
	"github.com/tylerb890/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	 var params = api.CoinBalanceParams{}
	 var decoder *schema.Decoder = schema.NewDecoder()
	 var err error

	 err = decoder.Decode(&params, r.URL.Query())
	 if err != nil {
	 	log.Error(err)
	 	api.InternalErrorHandler(w)
	 	return
	 }

	 var database *tools.DatabaseInterface
	 database, err = tools.NewDatabase()
	 if err != nil {
	 	log.Error(err)
	 	api.InternalErrorHandler(w)
	 	return
	 }

	 var coinDetails *tools.CoinDetails
	 coinDetails = (*database).GetUserCoins(params.Username)
	 if coinDetails == nil {
	 	log.Error(err)
	 	api.InternalErrorHandler(w)
	 	return
	 }

	 var response = api.CoinBalanceResponse{
	 	Balance: (*coinDetails).Coins,
	 	Code: http.StatusOK,
	 }

	 w.Header().Set("Content-Type", "application/json")
	 err = json.NewEncoder(w).Encode(response)
	 if err != nil {
	 	log.Error(err)
	 	api.InternalErrorHandler(w)
	 	return
	 }
}