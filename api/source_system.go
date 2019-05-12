package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//func PostDataSource(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
//	var jsonSourceSystem model.SourceSystem
//	jsonDecoder := json.NewDecoder(req.Body).Decode(&jsonSourceSystem)
//}

func GetDataSources(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {}
