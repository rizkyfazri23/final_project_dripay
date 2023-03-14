package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyGateway = []entity.Gateway{
	{
		Gateway_Id:   1,
		Gateway_Name: "gateway dummy 1",
		Status:       true,
	},
	{
		Gateway_Id:   2,
		Gateway_Name: "gateway dummy 2",
		Status:       true,
	},
}

type gatewayUseCaseMock struct {
	mock.Mock
}

func (u gatewayUseCaseMock) GetAll() ([]entity.Gateway, error) {
	args := u.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Gateway), nil
}

func (u gatewayUseCaseMock) GetOne(id int) (entity.Gateway, error) {
	//TODO implement me
	panic("implement me")
}

func (u gatewayUseCaseMock) Add(newUser *entity.Gateway) (entity.Gateway, error) {
	//TODO implement me
	panic("implement me")
}

func (u gatewayUseCaseMock) Edit(user *entity.Gateway) (entity.Gateway, error) {
	//TODO implement me
	panic("implement me")
}

func (u gatewayUseCaseMock) Remove(id int) error {
	//TODO implement me
	panic("implement me")
}

type GatewayApiTestSuite struct {
	suite.Suite
	useCaseTest     gatewayUseCaseMock
	routerTest      *gin.Engine
	routerGroupTest *gin.RouterGroup
}

func (suite *GatewayApiTestSuite) SetupTest() {
	suite.useCaseTest = gatewayUseCaseMock{}
	suite.routerTest = gin.Default()
	suite.routerGroupTest = suite.routerTest.Group("/v1")
}

func (suite *GatewayApiTestSuite) Test_GetAllAPI_Success() {
	suite.useCaseTest.On("GetAll").Return(dummyGateway, nil)
	user := NewGatewayController(suite.routerGroupTest, suite.useCaseTest)
	handler := user.GetAll
	suite.routerTest.GET("/v1/gateway", handler)
	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/gateway", nil)
	request.Header.Set("Content-Type", "application/json")

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, 200)
}
func TestGatewayApiTestSuite(t *testing.T) {
	suite.Run(t, new(GatewayApiTestSuite))
}
