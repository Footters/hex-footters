package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Footters/hex-footters/pkg/auth/endpoint"
	"github.com/Footters/hex-footters/pkg/auth/transport"

	"github.com/Footters/hex-footters/pkg/auth"
	"github.com/Footters/hex-footters/pkg/auth/mocks"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type UserServerTestSuite struct {
	suite.Suite
	svc      *mocks.MockService
	register *httptransport.Server
	login    *httptransport.Server
}

func TestUserServerSuite(t *testing.T) {
	suite.Run(t, new(UserServerTestSuite))
}

func (suite *UserServerTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	suite.svc = mocks.NewMockService(mockCtrl)
	endpoints := endpoint.MakeServerEndpoints(suite.svc)

	suite.register = httptransport.NewServer(
		endpoints.Register,
		transport.DecodeHTTPRegisterRequest,
		transport.EncodeHTTPResponse,
	)

	suite.login = httptransport.NewServer(
		endpoints.Login,
		transport.DecodeHTTPLoginRequest,
		transport.EncodeHTTPResponse,
	)
}

func (suite *UserServerTestSuite) TestRegister() {
	u := &auth.User{
		Email:    "david@lcarrascal.com",
		Password: "secret",
	}
	suite.svc.EXPECT().RegisterUser(gomock.Eq(u)).Return(nil)

	body, _ := json.Marshal(u)
	r, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))

	w := httptest.NewRecorder()
	suite.register.ServeHTTP(w, r)

	response := w.Result()
	suite.Equal("200 OK", response.Status)

	defer response.Body.Close()
	result := new(endpoint.RegisterResponse)
	json.NewDecoder(response.Body).Decode(result)

	suite.Equal("Register OK", result.Msg)
}

func (suite *UserServerTestSuite) TestLogin() {
	u := &auth.User{
		Email:    "david@lcarrascal.com",
		Password: "secret",
	}
	suite.svc.EXPECT().Login(u.Email, u.Password).Return(u, nil)

	body, _ := json.Marshal(u)
	r, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))

	w := httptest.NewRecorder()
	suite.login.ServeHTTP(w, r)

	response := w.Result()
	suite.Equal("200 OK", response.Status)

	defer response.Body.Close()
	result := new(endpoint.LoginResponse)
	json.NewDecoder(response.Body).Decode(result)

	suite.Equal("david@lcarrascal.com", result.User.Email)
	suite.Equal("secret", result.User.Password)
}
