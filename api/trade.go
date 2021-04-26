package api

import (
	"encoding/json"
	"log"
	"net/http"
	"smallcase/config/database"
	"smallcase/dto"
	"smallcase/logic"
	"smallcase/utils"
	"strings"
)

const (
	AddTradeAPITag     = "AddTradeAPI"
	UpdateTradeAPITag  = "UpdateTradeAPI"
	RemoveTradeAPITag  = "RemoverTradeAPI"
	FetchTradesAPITag  = "FetchTradesAPI"
	GetPortfolioAPITag = "GetPortfolioAPI"
	GetReturnsAPITag   = "GetReturnsAPI"
)

type TradeAPI struct {
	Trader logic.Trader
}

func AddTradeHandler(w http.ResponseWriter, req *http.Request) {
	defer GenericRecovery(w, AddTradeAPITag)
	ctx := req.Context()
	var trade *dto.AddTradeRequest
	err := json.NewDecoder(req.Body).Decode(&trade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("AddTrade Request: ", utils.Marshal(trade))
	tx := database.GetConnection().Begin()
	t := TradeAPI{&logic.TradeAPI{}}
	response, err := t.Trader.AddTrade(ctx, trade, tx)
	GenerateResponse(w, response, err, http.MethodPost, false)

}

func UpdateTradeHandler(w http.ResponseWriter, req *http.Request) {
	defer GenericRecovery(w, UpdateTradeAPITag)
	ctx := req.Context()
	var trade *dto.UpdateTradeRequest
	err := json.NewDecoder(req.Body).Decode(&trade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Update Trade Request: ", utils.Marshal(trade))
	tx := database.GetConnection().Begin()
	t := TradeAPI{&logic.TradeAPI{}}
	response, err := t.Trader.UpdateTrade(ctx, trade, tx)
	GenerateResponse(w, response, err, http.MethodPut, true)
}

func RemoveTradeHandler(w http.ResponseWriter, req *http.Request) {
	defer GenericRecovery(w, RemoveTradeAPITag)
	ctx := req.Context()
	tradeId := strings.Split(req.URL.Path, "/")[4]
	log.Println("Remove Trade Request: ", utils.Marshal(tradeId))
	tx := database.GetConnection().Begin()
	t := TradeAPI{&logic.TradeAPI{}}
	response, err := t.Trader.RemoveTrade(ctx, &tradeId, tx)
	GenerateResponse(w, response, err, http.MethodPut, true)
}

func FetchTradeHandler(w http.ResponseWriter, req *http.Request) {
	defer GenericRecovery(w, FetchTradesAPITag)
	ctx := req.Context()
	tx := database.GetConnection().Begin()
	t := TradeAPI{&logic.TradeAPI{}}
	response, err := t.Trader.FetchTrades(ctx, tx)
	GenerateResponse(w, response, err, http.MethodGet, false)
}

func GetPortfolioHandler(w http.ResponseWriter, req *http.Request) {
	defer GenericRecovery(w, GetPortfolioAPITag)
	ctx := req.Context()
	tx := database.GetConnection().Begin()
	t := TradeAPI{&logic.TradeAPI{}}
	response, err := t.Trader.GetPortfolio(ctx, tx)
	GenerateResponse(w, response, err, http.MethodGet, false)
}

func GetReturnsHandler(w http.ResponseWriter, req *http.Request) {
	defer GenericRecovery(w, GetReturnsAPITag)
	ctx := req.Context()
	tx := database.GetConnection().Begin()
	t := TradeAPI{&logic.TradeAPI{}}
	response, err := t.Trader.GetReturns(ctx, tx)
	GenerateResponse(w, response, err, http.MethodGet, false)
}
