package media_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"

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

func (suite *MediaServerTestSuite) TestGetContent() {
	c := &media.Content{
		URLName:     "sevillafc-betis",
		Title:       "Sevilla FC vs Real Betis",
		Description: "Derbi sevillano",
		Status:      "pending",
		Free:        1,
		Visible:     1,
	}
	suite.svc.EXPECT().FindContentByID(c.ID).Return(c, nil)

	url := "/contents/" + strconv.Itoa(int(c.ID))
	r, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()

	suite.getContent.ServeHTTP(w, r)

	response := w.Result()
	suite.Equal("200 OK", response.Status)
	defer response.Body.Close()

	// res := new(media.Content)
	// json.NewDecoder(response.Body).Decode(res)

	// suite.Equal(c.URLName, result.URLName, "should be the same")
	// suite.Equal(c.Status, result.Status, "should be the same")
}
func (suite *MediaServerTestSuite) TestGetAllContents() {

}
func (suite *MediaServerTestSuite) TestCreateContent() {
	c := &media.Content{
		URLName:     "sevillafc-betis",
		Title:       "Sevilla FC vs Real Betis",
		Description: "Derbi sevillano",
		Status:      "pending",
		Free:        1,
		Visible:     1,
	}
	suite.svc.EXPECT().CreateContent(gomock.Eq(c)).Return(nil)

	body, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/contents", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	suite.createContent.ServeHTTP(w, r)

	response := w.Result()
	suite.Equal("200 OK", response.Status)

	defer response.Body.Close()
	// result := new(media.Content)
	// json.NewDecoder(response.Body).Decode(result)

	// suite.Equal("Register OK", result.Msg)
}
func (suite *MediaServerTestSuite) TestToLiveContent() {

}
