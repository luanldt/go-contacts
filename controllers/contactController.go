package controllers

import (
	"encoding/json"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Response(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Response(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {

	/*
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			//The passed path parameter is not an integer
			u.Response(w, u.Message(false, "There was an error in your request"))
			return
		}
	*/

	id := r.Context().Value("user").(uint)

	data := models.GetContacts(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Response(w, resp)
}
