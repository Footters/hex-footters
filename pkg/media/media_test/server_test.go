package media_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Footters/hex-footters/pkg/media"
	"github.com/Footters/hex-footters/pkg/media/endpoint"
	"github.com/Footters/hex-footters/pkg/media/mocks"
	mediahttptransport "github.com/Footters/hex-footters/pkg/media/transport/http"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

func TestMediaServerSuite(t *testing.T) {
	suite.Run(t, new(MediaServerTestSuite))
}

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
	endpoints := endpoint.MakeServerEndpoints(suite.svc)

	suite.getContent = httptransport.NewServer(
		endpoints.GetContent,
		mediahttptransport.DecodeHTTPGetContentRequest,
		mediahttptransport.EncodeHTTPResponse,
	)

	suite.getAllContents = httptransport.NewServer(
		endpoints.GetAllContents,
		mediahttptransport.DecodeHTTPGetAllContentsRequest,
		mediahttptransport.EncodeHTTPResponse,
	)

	suite.createContent = httptransport.NewServer(
		endpoints.CreateContent,
		mediahttptransport.DecodeHTTPCreateContentRequest,
		mediahttptransport.EncodeHTTPResponse,
	)

	suite.toLiveContent = httptransport.NewServer(
		endpoints.SetContentToLive,
		mediahttptransport.DecodeHTTPSetContentLiveRequest,
		mediahttptransport.EncodeHTTPResponse,
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

	vars := map[string]string{
		"id": strconv.Itoa(int(c.ID)),
	}

	r, _ := http.NewRequest("GET", "/contents/", nil)
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	suite.getContent.ServeHTTP(w, r)

	response := w.Result()
	suite.Equal("200 OK", response.Status)
	defer response.Body.Close()

	res := new(endpoint.GetContentResponse)
	json.NewDecoder(response.Body).Decode(res)

	suite.Equal(c.URLName, res.Content.URLName, "should be the same")
	suite.Equal(c.Status, res.Content.Status, "should be the same")
}
func (suite *MediaServerTestSuite) TestGetAllContents() {
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
	suite.svc.EXPECT().FindAllContents().Return(cs, nil)

	r, _ := http.NewRequest("GET", "/contents", nil)
	w := httptest.NewRecorder()
	suite.getAllContents.ServeHTTP(w, r)

	response := w.Result()
	suite.Equal("200 OK", response.Status)

	defer response.Body.Close()
	res := new(endpoint.GetAllContentResponse)
	json.NewDecoder(response.Body).Decode(res)
	suite.Len(res.Contents, 2, "Should get two elements")
}
func (suite *MediaServerTestSuite) TestCreateContent() {
	cr := &endpoint.CreateContentRequest{
		Content: media.Content{
			URLName:     "sevillafc-betis",
			Title:       "Sevilla FC vs Real Betis",
			Description: "Derbi sevillano",
			Status:      "pending",
			Free:        1,
			Visible:     1,
		},
	}

	suite.svc.EXPECT().CreateContent(gomock.Eq(&cr.Content)).Return(nil)

	body, _ := json.Marshal(cr)
	r, _ := http.NewRequest("POST", "/contents", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	suite.createContent.ServeHTTP(w, r)

	response := w.Result()
	suite.Equal("200 OK", response.Status)

	defer response.Body.Close()
	res := new(endpoint.CreateContentResponse)
	json.NewDecoder(response.Body).Decode(res)
	suite.Equal("Created", res.Msg)
}
func (suite *MediaServerTestSuite) TestToLiveContent() {
	c := &media.Content{
		URLName:     "sevillafc-betis",
		Title:       "Sevilla FC vs Real Betis",
		Description: "Derbi sevillano",
		Status:      "pending",
		Free:        1,
		Visible:     1,
	}
	cu := c
	c.Status = "live"
	suite.svc.EXPECT().FindContentByID(c.ID).Return(c, nil)
	suite.svc.EXPECT().SetToLive(c).Return(nil)

	IDstr := strconv.Itoa(int(c.ID))
	url := "/contents/live"
	r, _ := http.NewRequest("PUT", url, nil)
	vars := map[string]string{
		"id": IDstr,
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	suite.toLiveContent.ServeHTTP(w, r)

	response := w.Result()
	suite.Equal("200 OK", response.Status)
	defer response.Body.Close()

	res := new(endpoint.GetContentResponse)
	json.NewDecoder(response.Body).Decode(res)

	suite.Equal(cu.Status, res.Content.Status, "Should be the same")
}
