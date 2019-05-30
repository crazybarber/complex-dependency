package api

import (
	"docugraphy/model"
	"docugraphy/repository"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

//func PostDataSource(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
//	var jsonSourceSystem model.SourceSystem
//	jsonDecoder := json.NewDecoder(req.Body).Decode(&jsonSourceSystem)
//}

func GetSourceSystems(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	sourceSystems, err := repository.GetSourceSystems()
	if nil != err {
		log.Println("Error while handling GET SOURCE: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(rw).Encode(sourceSystems)
	if nil != err {
		log.Println("Error while encoding response in GET SOURCE: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func AddSourceSystem(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var sourceSystem model.SourceSystem
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&sourceSystem)
	if nil != err {
		log.Println("Error while decoding request in ADD SOURCE: ", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err = repository.AddSourceSystem(&sourceSystem)
	if nil != err {
		log.Println("Error while handling ADD SOURCE: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}
