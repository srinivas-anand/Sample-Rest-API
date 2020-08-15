package api

import (
	"encoding/json"
	"net/http"

	"github.com/Sample-Rest-API/dao"
	"github.com/Sample-Rest-API/utils"
)

func getAll(w http.ResponseWriter, r *http.Request) {

	tableName := utils.GetConfig().Table
	response := dao.DDBGetAll(tableName)
	json.NewEncoder(w).Encode(response)

}
