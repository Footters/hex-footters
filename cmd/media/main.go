package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Footters/hex-footters/pkg/http/rest"
	"github.com/Footters/hex-footters/pkg/media"
	"github.com/Footters/hex-footters/pkg/provider/google"
	"github.com/Footters/hex-footters/pkg/storage/mysqldb"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := mySQLConnection()
	defer db.Close()

	mRepo := mysqldb.NewMysqlContentRepository(db)
	// mProv := ibm.NewIBMProvider()
	mProv := google.NewGoogleProvider()

	mService := media.NewService(mRepo, mProv)
	mHandler := rest.NewHandler(mService)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/contents", mHandler.Get).Methods("GET")
	router.HandleFunc("/contents/{id}", mHandler.GetByID).Methods("GET")
	router.HandleFunc("/contents", mHandler.Create).Methods("POST")
	router.HandleFunc("/contents/{id}/live", mHandler.SetToLive).Methods("GET")
	http.Handle("/", accessControl(router))

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port: 3000")
		errs <- http.ListenAndServe(":3000", nil)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("terminated %s", <-errs)
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
