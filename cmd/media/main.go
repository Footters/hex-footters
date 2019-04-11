package main

import (
	"net/http"
	"os"

	"github.com/Footters/hex-footters/pkg/media"
	"github.com/Footters/hex-footters/pkg/media/provider/google"
	"github.com/Footters/hex-footters/pkg/media/storage/mysqldb"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {

	// Logging
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
	}

	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	db := mySQLConnection()
	defer db.Close()

	mRepo := mysqldb.NewMysqlContentRepository(db)
	// mProv := ibm.NewIBMProvider()
	mProv := google.NewGoogleProvider()

	mService := media.NewService(mRepo, mProv)
	mService = media.NewLogginMiddleware(logger, mService)

	//Endpoint
	getContentEndpoint := media.MakeGetContentEndpoint(mService)
	getAllContentsEndpoint := media.MakeGetAllContentsEndpoint(mService)
	createContentEndpoint := media.MakeCreateContentEndpoint(mService)
	toLiveContentEndpoint := media.MakeSetContentLiveEndpoint(mService)

	// Transport
	getContentHandler := httptransport.NewServer(
		getContentEndpoint,
		media.DecodeGetContentRequest,
		media.EncodeResponse,
	)

	getAllContentsHandler := httptransport.NewServer(
		getAllContentsEndpoint,
		media.DecodeGetAllContentsRequest,
		media.EncodeResponse,
	)

	createContentHandler := httptransport.NewServer(
		createContentEndpoint,
		media.DecodeCreateContentRequest,
		media.EncodeResponse,
	)

	toLiveContentHandler := httptransport.NewServer(
		toLiveContentEndpoint,
		media.DecodeSetContentLiveRequest,
		media.EncodeResponse,
	)

	r := mux.NewRouter()
	r.Handle("/contents", createContentHandler).Methods("POST")
	r.Handle("/contents", getAllContentsHandler).Methods("GET")
	r.Handle("/contents/{id}", getContentHandler).Methods("GET")
	r.Handle("/contents/{id}/live", toLiveContentHandler).Methods("PUT")
	http.Handle("/", accessControl(r))

	// Go!
	logger.Log("transport", "HTTP", "addr", ":3000")
	logger.Log("exit", http.ListenAndServe(":3000", nil))
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func mySQLConnection() *gorm.DB {

	var db *gorm.DB
	var err error

	db, err = gorm.Open("mysql", os.Getenv("MYSQL_CONNECTION"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&media.Content{})

	return db
}
