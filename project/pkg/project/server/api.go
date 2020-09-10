package server

import(
	"fmt"
	"github.com/abnergarcia1/voxie-engineering-test/project/pkg/project/models"
	"github.com/abnergarcia1/voxie-engineering-test/project/pkg/project/services"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"strconv"
)

type api struct{
	router http.Handler
	teamsService *services.TeamService
	//contactsService services.ContactService
	//attributesService services.AttributeService
}

type Server interface{
	Router() http.Handler
}

func New() Server {
	a:= &api{}

	r:=mux.NewRouter()
	fmt.Println("Starting REST endpoints...")
	r.HandleFunc("/import", a.ImportData).Methods(http.MethodPost)
	r.HandleFunc("/api/teams", a.GetTeams).Methods(http.MethodGet)
	r.HandleFunc("/api/team/{ID:[a-zA-Z0-9_]+}", a.GetTeam).Methods(http.MethodGet)
	r.HandleFunc("/api/team/{ID:[a-zA-Z0-9_]+}", a.UpdateTeam).Methods(http.MethodPut)
	r.HandleFunc("/api/team/{ID:[a-zA-Z0-9_]+}", a.DeleteTeam).Methods(http.MethodDelete)
	fmt.Println("Running REST endpoints!")
	r.PathPrefix("/webclient/").Handler(http.StripPrefix("/webclient/",
		http.FileServer(http.Dir("../../pkg/project/server/static"))))
	a.router = r
	a.teamsService=&services.TeamService{}

	return a
}

func (a *api) Router() http.Handler{
	return a.router
}

/* Teams */
func(a *api) ImportData(w http.ResponseWriter, r *http.Request){
	var team models.Team

	err:=json.NewDecoder(r.Body).Decode(&team)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err=a.teamsService.CreateTeam(team)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func(a *api) GetTeams(w http.ResponseWriter, r *http.Request){

	teams,err:=a.teamsService.GetTeams()
	if err!=nil{
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(teams)

}

func (a *api) GetTeam(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	teamID,_ := strconv.Atoi(vars["ID"])

	team, err:=a.teamsService.GetTeam(teamID)
	if err!=nil{
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(team)
}

func (a *api) UpdateTeam(w http.ResponseWriter, r *http.Request){}

func (a *api) DeleteTeam(w http.ResponseWriter, r *http.Request){}