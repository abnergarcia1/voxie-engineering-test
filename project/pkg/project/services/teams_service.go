package services

import (
	"github.com/abnergarcia1/voxie-engineering-test/project/pkg/project/models"
	"time"
	"fmt"
)

type TeamService struct{
	contactService ContactService
}

func(s *TeamService) CreateTeam(team models.Team)(err error){
	db, err:=getDB()
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	strCommand, err:=db.Prepare("INSERT INTO teams(name, created_at, updated_at) VALUES(?,?,?)")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer strCommand.Close()

	res,err := strCommand.Exec(team.Name, time.Now(), time.Now())
	if err!=nil{
		fmt.Println(err)
		return
	}

	teamID,err:=res.LastInsertId()
	if err!=nil{
		fmt.Println(err)
		return
	}

	for _,contact:=range team.Contacts{

		contact.TeamID=int(teamID)
		err=s.contactService.CreateContact(contact,db)
		if err!=nil{
			fmt.Println(err)
			return
		}
	}

	return
}

func(s *TeamService) DeleteTeam(idTeam int)(err error){
	db, err:=getDB()
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	strCommand, err:=db.Prepare("DELETE FROM teams WHERE id=?")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer strCommand.Close()

	_,err = strCommand.Exec(idTeam)
	if err!=nil{
		fmt.Println(err)
		return
	}



	return
}

func(s *TeamService) GetTeams()(teams []models.Team, err error){
	teams=[]models.Team{}

	db, err:=getDB()
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	dataRows, err:=db.Query("SELECT id, `name`, created_at, updated_at FROM teams ORDER BY `name` ASC")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer dataRows.Close()

	for dataRows.Next(){
		tempTeam:=models.Team{}
		err=dataRows.Scan(&tempTeam.ID, &tempTeam.Name, &tempTeam.CreatedAt, &tempTeam.UpdatedAt)
		if err!=nil{
			fmt.Println(err)
			break
		}

		teams=append(teams, tempTeam)

	}

	for _,team:=range teams{
		team.Contacts,err=s.contactService.GetContacts(team.ID,db)
		if err!=nil{
			return
		}
	}
	return
}

func(s *TeamService) GetTeam(teamID int)(team models.Team, err error){
	team=models.Team{}

	db, err:=getDB()
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	dataRows, err:=db.Query("SELECT id, `name`, created_at, updated_at FROM teams WHERE id=? ", teamID)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer dataRows.Close()

	for dataRows.Next(){
		err=dataRows.Scan(&team.ID, &team.Name, &team.CreatedAt, &team.UpdatedAt)
		if err!=nil{
			fmt.Println(err)
			break
		}
	}

	team.Contacts,err=s.contactService.GetContacts(team.ID,db)

	return
}

func(s *TeamService) UpdateTeam(team models.Team)(err error){
	db, err:=getDB()
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	strCommand, err:=db.Prepare("UPDATE teams SET `name`=?, updated_at=? WHERE id=?")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer strCommand.Close()

	_,err = strCommand.Exec(team.Name, time.Now(), team.ID)
	if err!=nil{
		fmt.Println(err)
		return
	}

	return
}