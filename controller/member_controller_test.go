package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	// "github.com/rizkyfazri23/dripay/model/app_error"
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyMember = []entity.Member{
	{
		Member_Id:      1,
		Username:       "member dummy 1",
		Password:       "password dummy 1",
		Email_Address:  "dummy1@home.com",
		Contact_Number: "12345678",
		Wallet_Amount:  1000,
		Status:         true,
	},
	{
		Member_Id:      2,
		Username:       "member dummy 2",
		Password:       "password dummy 2",
		Email_Address:  "dummy2@home.com",
		Contact_Number: "123456789",
		Wallet_Amount:  2000,
		Status:         true,
	},
}

type memberUseCaseMock struct {
	mock.Mock
}

func (m memberUseCaseMock) GetAll() ([]entity.Member, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Member), nil
}

func (m memberUseCaseMock) GetOne(id int) (entity.Member, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return entity.Member{}, args.Error(1)
	}
	return args.Get(0).(entity.Member), nil
}

func (m memberUseCaseMock) Add(newUser *entity.Member) (entity.Member, error) {
	if newUser.Username == "existing_username" {
		return entity.Member{}, fmt.Errorf("username %s already exists", newUser.Username)
	}

	newUser.Member_Id = 999 // assign a new member_id to the added member
	return *newUser, nil
}

func (m memberUseCaseMock) Edit(user *entity.Member) (entity.Member, error) {
	//TODO implement me
	panic("implement me")
}

func (m memberUseCaseMock) Remove(id int) error {
	//TODO implement me
	panic("implement me")
}

func (m memberUseCaseMock) LoginCheck(username string, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

// func (m memberUseCaseMock) LoginCheck(username string, password string) (string, error) {
// 	args := m.Called(username, password)
// 	if args.Get(0) == nil {
// 		return "", args.Error(1)
// 	}
// 	return args.Get(0).(string), nil
// }

type MemberApiTestSuite struct {
	suite.Suite
	useCaseTest     memberUseCaseMock
	routerTest      *gin.Engine
	routerGroupTest *gin.RouterGroup
}

func (suite *MemberApiTestSuite) SetupTest() {
	suite.useCaseTest = memberUseCaseMock{}
	suite.routerTest = gin.Default()
	suite.routerGroupTest = suite.routerTest.Group("/v1")
}
func (suite *MemberApiTestSuite) Test_GetAllAPI_Success() {
	suite.useCaseTest.On("GetAll").Return(dummyMember, nil)
	member := NewMemberController(suite.routerGroupTest, suite.useCaseTest)
	handler := member.GetAll
	suite.routerTest.GET("/v1/member", handler)
	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/member", nil)
	request.Header.Set("Content-Type", "application/json")

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, 200)
}

// func (suite *MemberApiTestSuite) Test_Add_Success() {
// 	newMember := &entity.Member{
// 		Username:       "johndoe",
// 		Password:       "password",
// 		Email_Address:  "johndoe@example.com",
// 		Contact_Number: "+6281234567890",
// 		Wallet_Amount:  0,
// 		Status:         true,
// 	}

// 	expectedResult := &entity.Member{
// 		Member_Id:      1,
// 		Username:       "johndoe",
// 		Password:       "password",
// 		Email_Address:  "johndoe@example.com",
// 		Contact_Number: "+6281234567890",
// 		Wallet_Amount:  0,
// 		Status:         true,
// 	}

// 	suite.useCaseTest.On("Add", newMember).Return(*expectedResult, nil)

// 	member := NewMemberController(suite.routerGroupTest, suite.useCaseTest)
// 	handler := member.Add
// 	suite.routerTest.POST("", handler)

// 	reqBody, _ := json.Marshal(newMember)
// 	req, _ := http.NewRequest(http.MethodPost, "/v1/member", bytes.NewBuffer(reqBody))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp := httptest.NewRecorder()

// 	suite.routerTest.ServeHTTP(resp, req)

// 	assert.Equal(suite.T(), http.StatusCreated, resp.Code)

// 	response := entity.Member{}
// 	err := json.Unmarshal(resp.Body.Bytes(), &response)
// 	assert.Nil(suite.T(), err)

// 	assert.Equal(suite.T(), expectedResult.Member_Id, response.Member_Id)
// 	assert.Equal(suite.T(), expectedResult.Username, response.Username)
// 	assert.Equal(suite.T(), expectedResult.Password, response.Password)
// 	assert.Equal(suite.T(), expectedResult.Email_Address, response.Email_Address)
// 	assert.Equal(suite.T(), expectedResult.Contact_Number, response.Contact_Number)
// 	assert.Equal(suite.T(), expectedResult.Wallet_Amount, response.Wallet_Amount)
// 	assert.Equal(suite.T(), expectedResult.Status, response.Status)
// }

// func (suite *MemberApiTestSuite) Test_Add_InvalidRequest() {
// 	newMember := &entity.Member{
// 		Username:       "",
// 		Password:       "password",
// 		Email_Address:  "johndoe@example.com",
// 		Contact_Number: "+6281234567890",
// 		Wallet_Amount:  0,
// 		Status:         true,
// 	}

// 	member := NewMemberController(suite.routerGroupTest, suite.useCaseTest)
// 	handler := member.Add
// 	suite.routerTest.POST("", handler)

// 	reqBody, _ := json.Marshal(newMember)
// 	req, _ := http.NewRequest(http.MethodPost, "/v1/member", bytes.NewBuffer(reqBody))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp := httptest.NewRecorder()

// 	suite.routerTest.ServeHTTP(resp, req)

// 	assert.Equal(suite.T(), http.StatusBadRequest, resp.Code)

// 	response := app_error.ErrorResponse{}
// 	err := json.Unmarshal(resp.Body.Bytes(), &response)
// 	assert.Nil(suite.T(), err)

// 	assert.Equal(suite.T(), "X01", response.Code)
// 	assert.Equal(suite.T(), "invalid request", response.Message)
// }

// func (suite *MemberApiTestSuite) Test_Add_Error() {
// 	newUser := &entity.Member{
// 		Username:       "johndoe",
// 		Password:       "johndoepass",
// 		Email_Address:  "johndoe@example.com",
// 		Contact_Number: "08123456789",
// 		Wallet_Amount:  100000,
// 		Status:         1,
// 	}

// 	expectedError := errors.New("error while adding new member")
// 	suite.memberRepoMock.On("Add", newUser).Return(entity.Member{}, expectedError)

// 	result, err := suite.memberUseCase.Add(newUser)

// 	assert.Error(suite.T(), err)
// 	assert.Equal(suite.T(), entity.Member{}, result)
// 	assert.Equal(suite.T(), expectedError, err)
// }

// func (suite *MemberApiTestSuite) Test_GetOneAPI_Success() {
// 	expectedMember := dummyMember[0]

// 	suite.useCaseTest.On("GetOne", expectedMember.Member_Id).Return(expectedMember, nil)
// 	member := NewMemberController(suite.routerGroupTest, suite.useCaseTest)
// 	handler := member.GetOne
// 	suite.routerTest.GET("/:id", handler)

// 	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/v1/member/%d", expectedMember.Member_Id), nil)
// 	resp := httptest.NewRecorder()

// 	suite.routerTest.ServeHTTP(resp, req)

// 	assert.Equal(suite.T(), http.StatusOK, resp.Code)

// 	response := entity.Member{}
// 	err := json.Unmarshal(resp.Body.Bytes(), &response)
// 	assert.Nil(suite.T(), err)

// 	fmt.Println(response)
// 	fmt.Println("masmamdlwmld")

// 	assert.Equal(suite.T(), expectedMember.Member_Id, response.Member_Id)
// 	assert.Equal(suite.T(), expectedMember.Username, response.Username)
// 	assert.Equal(suite.T(), expectedMember.Password, response.Password)
// 	assert.Equal(suite.T(), expectedMember.Email_Address, response.Email_Address)
// 	assert.Equal(suite.T(), expectedMember.Contact_Number, response.Contact_Number)
// 	assert.Equal(suite.T(), expectedMember.Wallet_Amount, response.Wallet_Amount)
// 	assert.Equal(suite.T(), expectedMember.Status, response.Status)
// }

func TestMemberApiTestSuite(t *testing.T) {
	suite.Run(t, new(MemberApiTestSuite))
}
