package auth_test

import (
	"testing"

	"github.com/Footters/hex-footters/pkg/auth"
	"github.com/Footters/hex-footters/pkg/auth/mocks"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	userRepo  *mocks.MockUserRepository
	underTest auth.Service
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

func (suite *UserServiceTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	suite.userRepo = mocks.NewMockUserRepository(mockCtrl)
	suite.underTest = auth.NewService(suite.userRepo)
}

func (suite *UserServiceTestSuite) TestRegisterUser() {
	// Arrange
	u := &auth.User{
		Email:    "david@lcarrascal.com",
		Password: "secret",
	}
	suite.userRepo.EXPECT().Create(gomock.AssignableToTypeOf(&auth.User{})).Return(nil)

	// Act
	err := suite.underTest.RegisterUser(u)

	// Assert
	suite.NoError(err, "Shouldn't error")
	suite.NotNil(u.Email, "should not be null")
	suite.NotNil(u.Password, "should not be null")
}

func (suite *UserServiceTestSuite) TestLogin() {
	// Arrange
	// When my service call to FindByEmail, it will return u auth.User
	u := &auth.User{
		Email:    "david@lcarrascal.com",
		Password: "secret",
	}
	suite.userRepo.EXPECT().FindByEmail(u.Email).Return(u, nil)

	// Act
	res, err := suite.underTest.Login(u.Email, u.Password)

	// Assert
	suite.NoError(err, "Shoulnd't error")
	suite.Equal(u, res, "should be pushing value returned from repo")
}
