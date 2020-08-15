package api

import (
	"encoding/json"
	"net/http"

	"github.com/Sample-Rest-API/dao"
	"github.com/Sample-Rest-API/utils"
	"github.com/gorilla/mux"
)

func getOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	tableName := utils.GetConfig().Table
	response := dao.DDBGetOne(ID, tableName)
	json.NewEncoder(w).Encode(response)

}
