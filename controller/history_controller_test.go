package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyHistory = []entity.History{
	{
		Id:               1,
		Member_Username:  "member dummy 1",
		Transaction_Type: "transaction dummy 1",
		Kredit:           1000,
		Debit:            0,
		Date_Time:        time.Now(),
		Status:           "success",
		Transaction_Code: "code dummy 1",
	},
	{
		Id:               2,
		Member_Username:  "member dummy 2",
		Transaction_Type: "transaction dummy 2",
		Kredit:           2000,
		Debit:            0,
		Date_Time:        time.Now(),
		Status:           "success",
		Transaction_Code: "code dummy 2",
	},
}

type historyUseCaseMock struct {
	mock.Mock
}

func (u historyUseCaseMock) GetAll(id int) ([]entity.History, error) {
	args := u.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.History), nil
}

func (u historyUseCaseMock) GetAllPayment(id int) ([]entity.History, error) {
	args := u.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.History), nil
}

func (u historyUseCaseMock) GetAllTransfer(id int) ([]entity.History, error) {
	//TODO implement me
	panic("implement me")
}

func (u historyUseCaseMock) GetAllDeposit(id int) ([]entity.History, error) {
	//TODO implement me
	panic("implement me")
}

func (u historyUseCaseMock) ExportPDF(histories []entity.History) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

type HistoryApiTestSuite struct {
	suite.Suite
	useCaseTest     historyUseCaseMock
	routerTest      *gin.Engine
	routerGroupTest *gin.RouterGroup
}

func (suite *HistoryApiTestSuite) SetupTest() {
	suite.useCaseTest = historyUseCaseMock{}
	suite.routerTest = gin.Default()
	suite.routerGroupTest = suite.routerTest.Group("/v1")
}

func (suite *HistoryApiTestSuite) Test_GetAllAPI_Success() {
	suite.useCaseTest.On("GetAll", mock.AnythingOfType("int")).Return(dummyHistory, nil)
	history := NewHistoryController(suite.routerGroupTest, suite.useCaseTest)
	handler := history.GetAll
	suite.routerTest.GET("/v1/history", handler)
	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/history", nil)
	request.Header.Set("Content-Type", "application/json")

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, 200)
}

// func (suite *HistoryApiTestSuite) Test_GetAllPaymentAPI_Success() {
// 	// hapus pendaftaran handler untuk path '/v1/history/payment' jika sudah ada
// 	suite.routerTest.Handlers.Delete("GET", "/v1/history/payment")

// 	suite.useCaseTest.On("GetAll", mock.AnythingOfType("int")).Return(dummyHistory, nil)
// 	history := NewHistoryController(suite.routerGroupTest, suite.useCaseTest)
// 	handler := history.GetAllPayment
// 	suite.routerTest.GET("/v1/history/payment", handler)
// 	rr := httptest.NewRecorder()
// 	request, _ := http.NewRequest(http.MethodGet, "/v1/history/payment", nil)
// 	request.Header.Set("Content-Type", "application/json")

// 	suite.routerTest.ServeHTTP(rr, request)
// 	assert.Equal(suite.T(), rr.Code, 200)
// }

func TestHistoryApiTestSuite(t *testing.T) {
	suite.Run(t, new(HistoryApiTestSuite))
}
