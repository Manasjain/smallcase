package logic

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"reflect"
	"regexp"
	"smallcase/db/dao"
	"smallcase/db/models"
	"smallcase/dto"
	"strings"
	"testing"
)

func TestAddTrade(t *testing.T) {

	var TestTicker string = "BSE"
	var TestSellTradeType string = "BUY"
	var TestTradingUnitsTen int64 = 10
	var TestUnitPriceNinety float64 = 90
	var RandomUUID string = "abbdk-829nNjkA89h-Imsa"
	tx := &gorm.DB{}
	type Params struct {
		context context.Context
		request *dto.AddTradeRequest
	}
	tests := []struct {
		name   string
		logic  *TradeAPI
		params *Params

		expectedResponse *dto.AddTradeResponse
		expectedErr      error

		dbFindAllMockResponse []*models.Trades
		dbFindAllError        error

		dbSaveMockError error
	}{
		{
			name:  "Test_AddTrade_When_Trade_Is_Valid_Return_Success",
			logic: &TradeAPI{},
			params: &Params{
				context: context.Background(),
				request: &dto.AddTradeRequest{
					RequestBody: dto.AddTradeRequestBody{
						Ticker:       &TestTicker,
						TradingUnits: &TestTradingUnitsTen,
						TradeType:    &TestSellTradeType,
						UnitPrice:    &TestUnitPriceNinety,
					},
				},
			},
			expectedResponse: &dto.AddTradeResponse{
				TradeId: RandomUUID,
			},
			expectedErr: nil,

			dbFindAllMockResponse: []*models.Trades{
				{},
			},
			dbFindAllError: nil,

			dbSaveMockError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockController := gomock.NewController(t)
			defer mockController.Finish()

			tradeDatabaseClient := dao.NewMockDBClientTrader(mockController)
			tradeDatabaseClient.
				EXPECT().
				FindAll(gomock.Any(), gomock.Any()).
				Return(test.dbFindAllMockResponse, test.dbFindAllError).
				MaxTimes(1)
			tradeDatabaseClient.
				EXPECT().
				Save(gomock.Any(), gomock.Any()).
				Return(nil).
				MaxTimes(1)
			dao.DbClientTraderDao = tradeDatabaseClient
		})
		gotResponse, gotError := test.logic.AddTrade(test.params.context, test.params.request, tx)

		if strings.Contains(test.name, "Success") {
			if !IsValidUUID(gotResponse.TradeId) {
				t.Errorf("response got=%+v, Error: TradeId must be uuid", gotResponse)
			}
			if !reflect.DeepEqual(gotError, test.expectedErr) {
				t.Errorf("error got=%+v, want=%+v", gotError, test.expectedErr)
			}
		}
	}
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
