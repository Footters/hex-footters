package media_test

import (
	"testing"

	"github.com/Footters/hex-footters/pkg/media"
	"github.com/Footters/hex-footters/pkg/media/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type MediaServiceTestSuite struct {
	suite.Suite
	contentRepo  *mocks.MockContentRepository
	providerRepo *mocks.MockProviderRepository
	underTest    media.Service
}

func TestMediaServiceSuite(t *testing.T) {
	suite.Run(t, new(MediaServiceTestSuite))
}

func (suite *MediaServiceTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	suite.contentRepo = mocks.NewMockContentRepository(mockCtrl)
	suite.providerRepo = mocks.NewMockProviderRepository(mockCtrl)

	suite.underTest = media.NewService(suite.contentRepo, suite.providerRepo)
}

func (suite *MediaServiceTestSuite) TestCreateContent() {
	c := &media.Content{
		URLName:     "sevillafc-betis",
		Title:       "Sevilla FC vs Real Betis",
		Description: "Derbi sevillano",
		Status:      "pending",
		Free:        1,
		Visible:     1,
	}
	suite.contentRepo.EXPECT().Create(gomock.AssignableToTypeOf(&media.Content{})).Return(nil)

	err := suite.underTest.CreateContent(c)

	suite.NoError(err, "Shouldn't error")
	suite.NotNil(c.ID, "should not be null")
	suite.NotNil(c.URLName, "should not be null")
	suite.NotNil(c.Title, "should not be null")
	suite.NotNil(c.Description, "should not be null")
	suite.NotNil(c.Status, "should not be null")
	suite.NotNil(c.Free, "should not be null")
	suite.NotNil(c.Visible, "should not be null")
}

func (suite *MediaServiceTestSuite) TestFindContentByID() {
	c := &media.Content{
		URLName:     "sevillafc-betis",
		Title:       "Sevilla FC vs Real Betis",
		Description: "Derbi sevillano",
		Status:      "pending",
		Free:        1,
		Visible:     1,
	}
	suite.contentRepo.EXPECT().FindByID(c.ID).Return(c, nil)

	res, err := suite.underTest.FindContentByID(c.ID)
	suite.NoError(err, "Shouldn't error")
	suite.Equal(c, res, "should be the same")
}

func (suite *MediaServiceTestSuite) TestFindAllContents() {
	cs := []media.Content{
		media.Content{
			URLName:     "sevillafc-betis",
			Title:       "Sevilla FC vs Real Betis",
			Description: "Derbi sevillano",
			Status:      "pending",
			Free:        1,
			Visible:     1,
		},
		media.Content{
			URLName:     "barca-madrid",
			Title:       "Barcelona vs Real Madrid",
			Description: "El clásico de España",
			Status:      "pending",
			Free:        1,
			Visible:     1,
		},
	}

	suite.contentRepo.EXPECT().FindAll().Return(cs, nil)
	res, err := suite.underTest.FindAllContents()

	suite.NoError(err, "Shouldn't error")
	suite.Len(res, 2, "Should get two results")
}

func (suite *MediaServiceTestSuite) TestSetToLive() {
	c := &media.Content{
		URLName:     "sevillafc-betis",
		Title:       "Sevilla FC vs Real Betis",
		Description: "Derbi sevillano",
		Status:      "pending",
		Free:        1,
		Visible:     1,
	}

	cu := c
	cu.Status = "live"

	suite.providerRepo.EXPECT().CreateEvent()
	suite.contentRepo.EXPECT().Update(c).Return(nil)

	err := suite.underTest.SetToLive(c)
	suite.NoError(err, "Shouldn't error")
	suite.Equal(c, cu, "Should be the same")
}
