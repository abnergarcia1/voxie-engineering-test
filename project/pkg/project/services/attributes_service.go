package services

import (
	"database/sql"
	"fmt"
	"github.com/abnergarcia1/voxie-engineering-test/project/pkg/project/models"
)

type CustomAttributesService struct{}

func(s *CustomAttributesService) CreateAttribute(attribute models.CustomAttribute, refDB *sql.DB)(err error){
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

	strCommand, err:=db.Prepare("INSERT INTO custom_attributes(contact_id, `key`, `value`) VALUES(?,?,?)")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer strCommand.Close()

	_,err = strCommand.Exec(attribute.ContactID, attribute.Key, attribute.Value)
	if err!=nil{
		fmt.Println(err)
		return
	}

	return
}

func(s *CustomAttributesService) DeleteAttribute(attributeID int,refDB *sql.DB)(err error){
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

	strCommand, err:=db.Prepare("DELETE FROM custom_attributes WHERE id=?")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer strCommand.Close()

	_,err = strCommand.Exec(attributeID)
	if err!=nil{
		fmt.Println(err)
		return
	}

	return
}

func(s *CustomAttributesService) GetAttributes(contactID int, refDB *sql.DB) (attributes []models.CustomAttribute, err error){
	attributes=[]models.CustomAttribute{}
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

	dataRows, err:=db.Query("SELECT id, contact_id, `key`, `value` FROM custom_attributes WHERE contact_id=? ORDER BY `key` ASC", contactID)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer dataRows.Close()

	for dataRows.Next(){
		tempAttribute:=models.CustomAttribute{}
		err=dataRows.Scan(&tempAttribute.ID, &tempAttribute.ContactID, &tempAttribute.Key, &tempAttribute.Value)
		if err!=nil{
			fmt.Println(err)
			break
		}

		attributes=append(attributes, tempAttribute)

	}

	return
}

func(s *CustomAttributesService) UpdateValueAttribute(attributeID int, attributeVal string, refDB *sql.DB)(err error){
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

	strCommand, err:=db.Prepare("UPDATE custom_attributes SET `value`=? WHERE id=?")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer strCommand.Close()

	_,err = strCommand.Exec(attributeVal, attributeID)
	if err!=nil{
		fmt.Println(err)
		return
	}

	return
}
