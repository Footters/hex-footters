package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Footters/hex-footters/content"
	"github.com/Footters/hex-footters/db/mysqldb"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := mySQLConnection()
	defer db.Close()
	cRepo := mysqldb.NewMysqlContentRepository(db)
	cService := content.NewService(cRepo)
	cHandler := content.NewHandler(cService)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/contents", cHandler.Get).Methods("GET")
	router.HandleFunc("/contents/{id}", cHandler.GetByID).Methods("GET")
	router.HandleFunc("/contents", cHandler.Create).Methods("POST")

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
	log.Println("Connect to: ", os.Getenv("MYSQL_CONNECTION"))
	db, err = gorm.Open("mysql", os.Getenv("MYSQL_CONNECTION"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&content.Content{})

	return db
}
