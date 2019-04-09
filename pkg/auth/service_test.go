package auth_test

import (
	"testing"

	"github.com/Footters/hex-footters/pkg/auth"
	"github.com/Footters/hex-footters/pkg/auth/mocks"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/suite"
)

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

type UserServiceTestSuite struct {
	suite.Suite
	userRepo  *mocks.MockUserRepository
	underTest auth.Service
}

func (suite *UserServiceTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	suite.userRepo = mocks.NewMockUserRepository(mockCtrl)
	suite.underTest = auth.NewService(suite.userRepo)
}
