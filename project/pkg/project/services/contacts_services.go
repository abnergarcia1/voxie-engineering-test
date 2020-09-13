package services

import (
	"database/sql"
	"github.com/abnergarcia1/voxie-engineering-test/project/pkg/project/models"
	"time"
	"fmt"
)

type ContactService struct{
	attributesService CustomAttributesService
}

func(s *ContactService) CreateContact(contact models.Contact, refDB *sql.DB)(err error){
	db:=&sql.DB{}

	if refDB==nil {
		db, err = getDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()
	}else{
		db=refDB
	}

	strCommand, err:=db.Prepare("INSERT INTO contacts(team_id, `name`, phone, email, created_at, updated_at) VALUES(?,?,?,?,?,?)")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer strCommand.Close()

	res,err := strCommand.Exec(contact.TeamID, contact.Name, contact.Phone, contact.Email, time.Now(), time.Now())
	if err!=nil{
		fmt.Println(err)
		return
	}

	contactID,err:=res.LastInsertId()
	if err!=nil{
		fmt.Println(err)
		return
	}

	for _,attribute:=range contact.CustomAttributes{

		attribute.ContactID=int(contactID)
		err=s.attributesService.CreateAttribute(attribute,db)
		if err!=nil{
			fmt.Println(err)
			return
		}
	}

	return
}

func(s *ContactService) DeleteContact(contactID int,refDB *sql.DB)(err error){
	db:=&sql.DB{}

	if refDB==nil {
		db, err = getDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()
	}else{
		db=refDB
	}

	strCommand, err:=db.Prepare("DELETE FROM contacts WHERE id=?")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer strCommand.Close()

	_,err = strCommand.Exec(contactID)
	if err!=nil{
		fmt.Println(err)
		return
	}

	return
}

func(s *ContactService) GetContacts(teamID int, refDB *sql.DB)(contacts []models.Contact, err error){
	contacts=[]models.Contact{}
	db:=&sql.DB{}

	if refDB==nil {
		db, err = getDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()
	}else{
		db=refDB
	}

	dataRows, err:=db.Query("SELECT id, `name`, phone, email, created_at, updated_at FROM contacts WHERE team_id=? ORDER BY `name` ASC", teamID)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer dataRows.Close()

	for dataRows.Next(){
		tempContact:=models.Contact{}
		err=dataRows.Scan(&tempContact.ID, &tempContact.Name, &tempContact.Phone, &tempContact.Email, &tempContact.CreatedAt, &tempContact.UpdatedAt)
		if err!=nil{
			fmt.Println(err)
			break
		}

		contacts=append(contacts, tempContact)

	}

	for _,contact:=range contacts{
		contact.CustomAttributes,err=s.attributesService.GetAttributes(contact.ID,db)
		if err!=nil{
			return
		}
	}

	return
}

func(s *ContactService) GetContact(contactID int, refDB *sql.DB)(contact models.Contact, err error){
	contact=models.Contact{}
	db:=&sql.DB{}

	if refDB==nil {
		db, err = getDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()
	}else{
		db=refDB
	}

	dataRows, err:=db.Query("SELECT id, `name`, phone, email, created_at, updated_at FROM contacts WHERE id=?", contactID)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer dataRows.Close()

	for dataRows.Next(){
		err=dataRows.Scan(&contact.ID, &contact.Name, &contact.Phone, &contact.Email, &contact.CreatedAt, &contact.UpdatedAt)
		if err!=nil{
			fmt.Println(err)
			break
		}
	}

	contact.CustomAttributes,err=s.attributesService.GetAttributes(contact.ID,db)

	return
}

func(s *ContactService) UpdateContact(contact models.Contact, refDB *sql.DB)(err error){
	db:=&sql.DB{}

	if refDB==nil {
		db, err = getDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()
	}else{
		db=refDB
	}

	strCommand, err:=db.Prepare("UPDATE contacts SET name=?, phone=?, email=?, updated_at=?")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer strCommand.Close()

	_,err = strCommand.Exec(contact.Name, contact.Phone, contact.Email, time.Now())
	if err!=nil{
		fmt.Println(err)
		return
	}





	return
}
