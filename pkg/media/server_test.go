package media_test

import (
	"github.com/Footters/hex-footters/pkg/media"
	"github.com/Footters/hex-footters/pkg/media/mocks"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type MediaServerTestSuite struct {
	suite.Suite
	svc *mocks.MockService

	getContent     *httptransport.Server
	getAllContents *httptransport.Server
	createContent  *httptransport.Server
	toLiveContent  *httptransport.Server
}

func (suite *MediaServerTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	suite.svc = mocks.NewMockService(mockCtrl)

	//Endpoints
	getContentEndpoint := media.MakeGetContentEndpoint(suite.svc)
	getAllContentsEndpoint := media.MakeGetAllContentsEndpoint(suite.svc)
	createContentEndpoint := media.MakeCreateContentEndpoint(suite.svc)
	toLiveContentEndpoint := media.MakeSetContentLiveEndpoint(suite.svc)

	suite.getContent = httptransport.NewServer(
		getContentEndpoint,
		media.DecodeGetContentRequest,
		media.EncodeResponse,
	)

	suite.getAllContents = httptransport.NewServer(
		getAllContentsEndpoint,
		media.DecodeGetAllContentsRequest,
		media.EncodeResponse,
	)

	suite.createContent = httptransport.NewServer(
		createContentEndpoint,
		media.DecodeCreateContentRequest,
		media.EncodeResponse,
	)

	suite.toLiveContent = httptransport.NewServer(
		toLiveContentEndpoint,
		media.DecodeSetContentLiveRequest,
		media.EncodeResponse,
	)
}
