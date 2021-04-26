// SmallCase API Assignment Documentation
//
// This is the detailed overview of all the REST API
//
// Terms Of Service:
//
//     Schemes: http
//     Host: localhost:3000
//     Version: 1.0.0
//     Contact: Manas Jain <manas.jain@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"smallcase/api"
	"smallcase/config"
	"smallcase/config/database"
	"time"
)


func StartServer() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// swagger:operation POST /api/v1/trade/add AddTradeHandler
	// ---
	// summary: Returns a Unique Id (TradeId) for each Trade
	// description: If the trade is successful, a TradeId will be returned else an error response will be sent
	// Consumes:
	// - application/json
	// responses:
	//   "201":
	//     "$ref": "#/responses/AddTradeResponse"
	//   "400":
	//     "$ref": "#/responses/Error"
	//   "500":
	//     "$ref": "#/responses/Error"
	myRouter.HandleFunc("/api/v1/trade/add", api.AddTradeHandler).Methods(http.MethodPost)
	// swagger:operation PUT /api/v1/trade/update UpdateTradeHandler
	// ---
	// summary: Returns no response
	// description: if the trade is successfully updated, http.StatusNoContent = 204 will be returned
	// responses:
	//   "201":
	//     "$ref": "#/responses/UpdateTradeResponse"
	//   "400":
	//     "$ref": "#/responses/Error"
	//   "500":
	//     "$ref": "#/responses/Error"
	myRouter.HandleFunc("/api/v1/trade/update", api.UpdateTradeHandler).Methods(http.MethodPut)
	// swagger:operation PUT /api/v1/trade/{tradeId}/remove RemoveTradeHandler
	// ---
	// summary: Returns no response
	// description: If the trade is successfully deleted, http.StatusNoContent = 204 will be returned
	// parameters:
	// - name: tradeId
	//   in: path
	//   description: trade id generated when a new trade is added
	//   type: string
	//   required: true
	// responses:
	//   "204":
	//     "$ref": "#/responses/RemoveTradeResponse"
	//   "400":
	//     "$ref": "#/responses/Error"
	//   "500":
	//     "$ref": "#/responses/Error"
	myRouter.HandleFunc("/api/v1/trade/{tradeId}/remove", api.RemoveTradeHandler).Methods(http.MethodPut)
	// swagger:operation GET /api/v1/trades/fetch FetchTradeHandler
	// ---
	// summary: Returns all the Securities & Corresponding Trades
	// responses:
	//   "204":
	//     "$ref": "#/responses/FetchTradeResponse"
	//   "400":
	//     "$ref": "#/responses/Error"
	//   "500":
	//     "$ref": "#/responses/Error"
	myRouter.HandleFunc("/api/v1/trades/fetch", api.FetchTradeHandler).Methods(http.MethodGet)
	// swagger:operation GET /api/v1/trades/portfolio GetPortfolioHandler
	// ---
	// summary: Returns aggregated view of all the Securities
	// responses:
	//   "200":
	//     "$ref": "#/responses/GetPortfolioResponse"
	//   "400":
	//     "$ref": "#/responses/Error"
	//   "500":
	//     "$ref": "#/responses/Error"
	myRouter.HandleFunc("/api/v1/trades/portfolio", api.GetPortfolioHandler).Methods(http.MethodGet)
	// swagger:operation GET /api/v1/trades/returns GetReturnsHandler
	// ---
	// summary: Returns cumulative returns on all the securities
	// responses:
	//   "200":
	//     "$ref": "#/responses/GetReturnsResponse"
	//   "400":
	//     "$ref": "#/responses/Error"
	//   "500":
	//     "$ref": "#/responses/Error"
	myRouter.HandleFunc("/api/v1/trades/returns", api.GetReturnsHandler).Methods(http.MethodGet)

	fs := http.FileServer(http.Dir("../swaggerui"))
	myRouter.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, myRouter),
		Addr:         config.Config.HostAndPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("starting server at host:port = ", config.Config.HostAndPort)
	log.Fatal(srv.ListenAndServe())
}

func main() {
	defer doAPIRecover()
	defer database.CloseDatabaseConnection()
	StartServer()
}

func doAPIRecover() {
	if err := recover(); err != nil {
		log.Println("Failed to recover the server")
	}
}