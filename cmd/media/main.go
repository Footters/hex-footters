package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Footters/hex-footters/pkg/media"
	mediaendpoint "github.com/Footters/hex-footters/pkg/media/endpoint"
	"github.com/Footters/hex-footters/pkg/media/provider/auth"
	"github.com/Footters/hex-footters/pkg/media/provider/google"
	"github.com/Footters/hex-footters/pkg/media/storage/mysqldb"
	mediahttptransport "github.com/Footters/hex-footters/pkg/media/transport/http"
	"github.com/go-kit/kit/log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Logger
	logger := log.NewLogfmtLogger(os.Stdout)
	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	// Repos
	db := mysqldb.NewMySQLConnection()
	defer db.Close()
	mRepo := mysqldb.NewMysqlContentRepository(db)
	// mProv := ibm.NewIBMProvider()
	mProv := google.NewGoogleProvider()

	//Service
	mSvc := media.NewService(mRepo, mProv)
	mSvc = media.NewLogginMiddleware(logger, mSvc)

	// Auth service provider
	grpcConn := auth.NewAuthServiceProviderConnection()
	defer func() {
		err := grpcConn.Close()
		if err != nil {
			fmt.Println("ConnectionError", err)
		}
	}()
	asp := auth.NewServiceProvider(context.Background(), grpcConn)

	// Endpoints
	endpoints := mediaendpoint.MakeServerEndpoints(mSvc, asp)

	// HTTPHandler
	httpHandler := mediahttptransport.NewHTTPHandler(endpoints)
	http.Handle("/", mediahttptransport.AccessControl(httpHandler))

	// Go!
	logger.Log("transport", "HTTP", "addr", ":3000")
	logger.Log("exit", http.ListenAndServe(":3000", nil))
}
