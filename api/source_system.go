package api

import (
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

func GeSourceSystems(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	sourceSystems, err := repository.GetSourceSystems()
	if nil != err {
		log.Println("Error while handling GET SOURCE: ", err)
	}
	err = json.NewEncoder(rw).Encode(sourceSystems)
	if nil != err {
		log.Println("Error while encoding response in GET SOURCE: ", err)
	}
}
