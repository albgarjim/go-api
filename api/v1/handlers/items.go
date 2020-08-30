package handlers

import (
	"encoding/json"
	"net/http"

	inp "goggers/api/v1/input"
	out "goggers/api/v1/output"

	rdb "goggers/api/v1/models/rethinkdb"

	log "github.com/sirupsen/logrus"
)

func GetItem(_w http.ResponseWriter, _r *http.Request) {
	log.Info("-------- Call GetSingleItem route --------")
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

	log.Info("-------- Finish GetSingleItem route --------")
	out.RespondWithJSON(_w, http.StatusOK, items)
}

func UpdateItem(_w http.ResponseWriter, _r *http.Request) {
	log.Info("-------- Call to RevertDislikedItem route --------")
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
		log.Error("Error reverting update items on database: ", err)
		out.RespondWithError(_w, out.ErrInternalServer)
		return
	}

	log.Info("-------- Finish RevertDislikedItem route --------")
	out.RespondWithJSON(_w, http.StatusCreated, out.SuccessResource)
}

func DeleteItem(_w http.ResponseWriter, _r *http.Request) {
	log.Info("-------- Call to DislikeItem route --------")
	var err error
	var params *inp.URLParams

	if params, err = inp.ExtractParameters(_r); err != nil {
		log.Error("Incorrect url parameters: ", err)
		out.RespondWithError(_w, err)
		return
	}

	if err = rdb.ItemDeleteDB(params); err != nil {
		log.Error("Error disliking items on database: ", err)
		out.RespondWithError(_w, out.ErrInternalServer)
		return
	}

	log.Info("-------- Finish DislikeItem route --------")
	out.RespondWithJSON(_w, http.StatusCreated, out.SuccessResource)
}

func InsertItem(_w http.ResponseWriter, _r *http.Request) {
	log.Info("-------- Call to RevertDislikedItem route --------")
	var err error
	var u inp.ItemData

	if err := json.NewDecoder(_r.Body).Decode(&u); err != nil {
		log.Error("Error decoding body to ItemData struct: ", err)
		out.RespondWithError(_w, out.ErrBodyParams)
		return
	}

	if err = rdb.ItemInsertDB(u); err != nil {
		log.Error("Error reverting disliked items on database: ", err)
		out.RespondWithError(_w, out.ErrInternalServer)
		return
	}

	log.Info("-------- Finish RevertDislikedItem route --------")
	out.RespondWithJSON(_w, http.StatusCreated, out.SuccessResource)
}
