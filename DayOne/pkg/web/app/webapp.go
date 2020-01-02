package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	models "github.com/jpramirez/HundredDaysOfCode/DayOne/pkg/models"

	storage "github.com/jpramirez/HundredDaysOfCode/DayOne/pkg/storage"

	"encoding/gob"

	"github.com/gorilla/mux"
)

type jresponse struct {
	ResponseCode string
	Message      string
	ResponseData string
}

// MainWebApp PHASE
type MainWebApp struct {
	Mux    *mux.Router
	Log    *log.Logger
	Config models.Config
}

//NewApp creates a new instances
func NewApp(config models.Config) (MainWebApp, error) {

	var err error
	var wapp MainWebApp
	mux := mux.NewRouter().StrictSlash(true)
	log := log.New(os.Stdout, "web ", log.LstdFlags)
	wapp.Mux = mux
	wapp.Config = config
	wapp.Log = log

	return wapp, err
}

func (a *MainWebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Mux.ServeHTTP(w, r)
}

//Liveness just keeps the connection alive
func (a *MainWebApp) Liveness(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var response jresponse
	response.ResponseCode = "200"
	response.Message = "Success"
	response.ResponseData = ""
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}

//ImportFullDataSet import the entire dataset
func (a *MainWebApp) ImportFullDataSet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var storageDB storage.DB
	storageDB, err := storage.NewBadgerDB("data/")

	defer storageDB.Close()

	/// READ FILE
	file := "news-feed-list-of-countries.json"
	listInit, err := os.Open(file)
	defer listInit.Close()

	jsonParser := json.NewDecoder(listInit)
	var feeds models.NewsFeed
	jsonParser.Decode(&feeds)

	for _, x := range feeds {
		fmt.Println(x.Name)

		var b bytes.Buffer
		e := gob.NewEncoder(&b)
		if err := e.Encode(x); err != nil {
			panic(err)
		}
		namespace := []byte("Sources")
		key := []byte(x.Name)
		data := b.Bytes()
		storageDB.Set(namespace, key, data)

	}

	var response jresponse
	response.ResponseCode = "200"
	response.Message = "Success"
	response.ResponseData = ""

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}

func (a *MainWebApp) GetFullDataSet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var storageDB storage.DB
	storageDB, err := storage.NewBadgerDB("data/")

	defer storageDB.Close()

	namespace := []byte("Sources")
	key := []byte("Mali")

	data, err := storageDB.Get(namespace, key)
	if err != nil {
		fmt.Println("Error ", err)
	}

	var feeds models.CountryFeed
	d := gob.NewDecoder(bytes.NewReader(data))

	d.Decode(&feeds)
	var response jresponse

	response.ResponseCode = "200"
	response.Message = "Success"

	b, err := json.Marshal(feeds)
	response.ResponseData = string(b)

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}

//CheckCountry will return the data for that country
func (a *MainWebApp) CheckCountry(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var storageDB storage.DB
	storageDB, err := storage.NewBadgerDB("data/")

	defer storageDB.Close()
	vars := mux.Vars(r)
	country := vars["country"]

	namespace := []byte("Sources")

	key := []byte(country)

	data, err := storageDB.Get(namespace, key)
	if err != nil {
		fmt.Println("Error ", err)
	}

	var countryFeeds models.CountryFeed
	d := gob.NewDecoder(bytes.NewReader(data))

	d.Decode(&countryFeeds)

	for _, i := range countryFeeds.Sources {
		fmt.Println(i.Feedlink)
	}

	var response jresponse
	response.ResponseCode = "200"
	response.Message = "Success"

	b, err := json.Marshal(countryFeeds)
	response.ResponseData = string(b)

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}
