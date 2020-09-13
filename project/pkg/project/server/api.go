package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abnergarcia1/voxie-engineering-test/project/pkg/project/models"
	"github.com/abnergarcia1/voxie-engineering-test/project/pkg/project/services"
	"github.com/gorilla/mux"
)

type api struct {
	router          http.Handler
	teamsService    *services.TeamService
	contactsService *services.ContactService
	attributesService *services.CustomAttributesService
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()

	fmt.Println("Starting REST endpoints...")
	r.HandleFunc("/import", a.ImportData).Methods(http.MethodPost)
	r.HandleFunc("/api/teams", a.GetTeams).Methods(http.MethodGet)
	r.HandleFunc("/api/team/{ID}", a.GetTeam).Methods(http.MethodGet)
	r.HandleFunc("/api/team/", a.UpdateTeam).Methods(http.MethodPut)
	r.HandleFunc("/api/team/{ID}", a.DeleteTeam).Methods(http.MethodDelete)

	r.HandleFunc("/api/contacts/", a.CreateContact).Methods(http.MethodPost)
	r.HandleFunc("/api/contacts/{TeamID}", a.GetContacts).Methods(http.MethodGet)
	r.HandleFunc("/api/contacts/{ID}", a.GetContact).Methods(http.MethodGet)
	r.HandleFunc("/api/contacts/", a.UpdateContact).Methods(http.MethodPut)
	r.HandleFunc("/api/contacts/{ID}", a.DeleteContact).Methods(http.MethodDelete)

	r.HandleFunc("/api/attributes/", a.CreateAttribute).Methods(http.MethodPost)
	r.HandleFunc("/api/attributes/{ContactID}", a.GetAttributes).Methods(http.MethodGet)
	r.HandleFunc("/api/attributes/{ID}", a.DeleteAttribute).Methods(http.MethodDelete)
	r.HandleFunc("/api/attributes/", a.UpdateAttributes).Methods(http.MethodPut)

	fmt.Println("Running REST endpoints!")
	r.PathPrefix("/webclient/").Handler(http.StripPrefix("/webclient/",
		http.FileServer(http.Dir("../../pkg/project/server/static"))))
	a.router = r
	a.attributesService=&services.CustomAttributesService{}
	a.contactsService=&services.ContactService{}
	a.teamsService = &services.TeamService{}



	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

/* Teams */
func (a *api) ImportData(w http.ResponseWriter, r *http.Request) {
	var team models.Team

	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.teamsService.CreateTeam(team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (a *api) GetTeams(w http.ResponseWriter, r *http.Request) {

	teams, err := a.teamsService.GetTeams()
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(teams)

}

func (a *api) GetTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID, _ := strconv.Atoi(vars["ID"])

	team, err := a.teamsService.GetTeam(teamID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(team)
}

func (a *api) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	var team models.Team

	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.teamsService.UpdateTeam(team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (a *api) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID, _ := strconv.Atoi(vars["ID"])

	err := a.teamsService.DeleteTeam(teamID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode(err)
		return
	}
}

/*-***********Contacts*****************-*/
func (a *api) CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact

	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.contactsService.CreateContact(contact, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (a *api) DeleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactID, _ := strconv.Atoi(vars["ID"])

	err := a.contactsService.DeleteContact(contactID, nil)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode(err)
		return
	}
}

func (a *api) GetContacts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID, _ := strconv.Atoi(vars["TeamID"])

	contacts, err := a.contactsService.GetContacts(teamID, nil)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(contacts)

}

func (a *api) GetContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactID, _ := strconv.Atoi(vars["ID"])

	contact, err := a.contactsService.GetContact(contactID, nil)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(contact)
}

func (a *api) UpdateContact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact

	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.contactsService.UpdateContact(contact, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

/*--------------Attributes----------------------*/
func (a *api) CreateAttribute(w http.ResponseWriter, r *http.Request) {
	var attribute models.CustomAttribute

	err := json.NewDecoder(r.Body).Decode(&attribute)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.attributesService.CreateAttribute(attribute, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (a *api) DeleteAttribute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	attributeID, _ := strconv.Atoi(vars["ID"])

	err := a.attributesService.DeleteAttribute(attributeID, nil)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode(err)
		return
	}
}

func (a *api) GetAttributes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactID, _ := strconv.Atoi(vars["ContactID"])

	attributes, err := a.attributesService.GetAttributes(contactID, nil)
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(attributes)
}

func (a *api) UpdateAttributes(w http.ResponseWriter, r *http.Request) {
	var attribute models.CustomAttribute

	err := json.NewDecoder(r.Body).Decode(&attribute)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.attributesService.UpdateValueAttribute(attribute.ID, attribute.Value, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
