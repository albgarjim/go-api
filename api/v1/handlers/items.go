package handlers

import (
	"encoding/json"
	"net/http"

	inp "go-api/api/v1/input"
	out "go-api/api/v1/output"

	rdb "go-api/api/v1/models/rethinkdb"

	log "github.com/sirupsen/logrus"
)

func GetItem(_w http.ResponseWriter, _r *http.Request) {
	log.Info("-------- Call GetItem route --------")
	var err error
	var params *inp.URLParams

	if params, err = inp.ExtractParameters(_r); err != nil {
		log.Error("Incorrect url parameters: ", err)
		out.RespondWithError(_w, err)
		return
	}

	if (*params)[inp.QItemID] == nil {
		log.Error("Error missing item id")
		out.RespondWithError(_w, err)
		return
	}

	items, err := rdb.RetrieveItemDB(params)

	if err != nil {
		log.Error("Error retrieving items from database: ", err)
		out.RespondWithError(_w, err)
		return
	}

	log.Info("-------- Finish GetItem route --------")
	out.RespondWithJSON(_w, http.StatusOK, items)
}

func UpdateItem(_w http.ResponseWriter, _r *http.Request) {
	log.Info("-------- Call to UpdateItem route --------")
	var err error
	var params *inp.URLParams
	var u inp.ItemData

	if err := json.NewDecoder(_r.Body).Decode(&u); err != nil {
		log.Error("Error decoding body to ItemData struct: ", err)
		out.RespondWithError(_w, out.ErrBodyParams)
		return
	}

	if params, err = inp.ExtractParameters(_r); err != nil {
		log.Error("Incorrect url parameters: ", err)
		out.RespondWithError(_w, err)
		return
	}

	if err = rdb.ItemUpdateDB(params, u); err != nil {
		log.Error("Error updating items on database: ", err)
		out.RespondWithError(_w, out.ErrInternalServer)
		return
	}

	log.Info("-------- Finish UpdateItem route --------")
	out.RespondWithJSON(_w, http.StatusCreated, out.SuccessResource)
}

func DeleteItem(_w http.ResponseWriter, _r *http.Request) {
	log.Info("-------- Call to DeleteItem route --------")
	var err error
	var params *inp.URLParams

	if params, err = inp.ExtractParameters(_r); err != nil {
		log.Error("Incorrect url parameters: ", err)
		out.RespondWithError(_w, err)
		return
	}

	if err = rdb.ItemDeleteDB(params); err != nil {
		log.Error("Error deleting items on database: ", err)
		out.RespondWithError(_w, out.ErrInternalServer)
		return
	}

	log.Info("-------- Finish DeleteItem route --------")
	out.RespondWithJSON(_w, http.StatusCreated, out.SuccessResource)
}

func InsertItem(_w http.ResponseWriter, _r *http.Request) {
	log.Info("-------- Call to InsertItem route --------")
	var err error
	var u inp.ItemData

	if err := json.NewDecoder(_r.Body).Decode(&u); err != nil {
		log.Error("Error decoding body to ItemData struct: ", err)
		out.RespondWithError(_w, out.ErrBodyParams)
		return
	}

	if err = rdb.ItemInsertDB(u); err != nil {
		log.Error("Error inserting items on database: ", err)
		out.RespondWithError(_w, out.ErrInternalServer)
		return
	}

	log.Info("-------- Finish InsertItem route --------")
	out.RespondWithJSON(_w, http.StatusCreated, out.SuccessResource)
}
