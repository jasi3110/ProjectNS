package repos

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
	"log"
)

type UserInterface interface {
	UserCreate(obj *models.User) (string, bool)
	UserUpdate(obj *models.UserUpdate) (models.UserUpdate, bool, string)
	UserLogin(obj *models.LoginUser) (models.User, bool, string)
	UserGetall() ([]models.User, bool)
	UserGetById(obj *models.User) (models.User, bool, string)
	UserUpdatePassword(obj *models.UserPassword) (string, bool)
	GetByUserMobileno(obj *models.User) (models.User, bool)
}

type UserRepo struct {
}

func (user *UserRepo) UserCreate(obj *models.User) (string, bool) {
	Db, isConnected := utls.OpenDbConnection()

	if !isConnected {
		fmt.Println("DB Disconnented in CreateUser Repo ")
	}
	query, err := Db.Query(`SELECT email,mobileno FROM "user"`)
	if err != nil {
		log.Println("Error in CreateUser  QueryRow :", err)
	}
	userStruct := models.User{}
	for query.Next() {
		query.Scan(&userStruct.Email, &userStruct.Mobileno)
		if obj.Mobileno == userStruct.Mobileno || obj.Email == userStruct.Email {
			fmt.Println("This Mobile number or Email already Used ")
			return "this Mobile number or Email already Used by Other User ", false
		}
	}
	err = Db.QueryRow(`INSERT into "user"(
		name,
		email,
		mobileno,
		role,
		password,
		createdon,
		token
	)values($1,$2,$3,$4,$5,$6,$7)RETURNING id`,
		obj.Name,
		obj.Email,
		obj.Mobileno,
		obj.Role,
		obj.Password,
		utls.GetCurrentDateTime(),
		" NULL TOKEN",
	).Scan(&obj.Id)

	if err != nil {
		fmt.Println("Error in CreateUser  QueryRow :", err)
		return "User Create Failed", false
	}
	return "User created Sucessfully", true
}

func (user *UserRepo) UserLogin(obj *models.LoginUser) (models.User, bool, string) {
	myDb, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB DisConnected in UserLogin ")
	}
	userStruct := models.User{}
	query, err := myDb.Prepare(`SELECT id,
	                                name,
									email,
									mobileno,
									role,
									createdon
									from "user" where mobileno=$1 AND password=$2`)
	if err != nil {
		log.Println("Error in User Login QueryRow :", err)
		return userStruct, false, "User Login  Failed"
	}

	err = query.QueryRow(obj.Mobileno, obj.Password).Scan(&userStruct.Id,
		&userStruct.Name,
		&userStruct.Email,
		&userStruct.Mobileno,
		&userStruct.Role,
		&userStruct.CreatedOn,
	)
	if err != nil {
		log.Println("Error in User Login QueryRow :", err)
		return userStruct, false, "User Login  Failed"
	}

	Token := utls.GenerateJwtToken(userStruct.Id, userStruct.Mobileno, userStruct.Name)
	log.Println(userStruct.Id, Token)
	updatequery := `UPDATE  "user" SET  token=$2 WHERE id=$1;`
	_, err = myDb.Exec(updatequery, userStruct.Id, Token)

	userStruct.Token = Token

	if err != nil {
		log.Println("Error in User Login QueryRow :", err)
		return userStruct, false, "User Login  Failed"
	}

	return userStruct, true, "User Login Sucessfully Completed"
}

func (user *UserRepo) UserUpdate(obj *models.UserUpdate) (models.UserUpdate, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB disconnceted in UserUpdate ")
	}
	query := `UPDATE "user"
	SET name      =$2,
	    email     =$3,
	    mobileno  =$4,
	    role      =$5
	    WHERE id=$1 `
	_, err := Db.Exec(query,
		&obj.Id,
		&obj.Name,
		&obj.Email,
		&obj.Mobileno,
		&obj.Role,
	)
	if err != nil {
		fmt.Println("Error in User Update QueryRow :", err)
		return *obj, false, "User Update Failed"
	}
	return *obj, true, "User Updated Sucessfully Completed"
}

func (user *UserRepo) UserUpdatePassword(obj *models.UserPassword) (string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in User Update Password ")
	}

	query := `UPDATE "user" SET password = $3 WHERE id=$1 and mobileno=$2`
	_, err := Db.Exec(query, &obj.Id, &obj.Mobileno, &obj.Password)
	if err != nil {
		fmt.Println("Error in User Update Password QueryRow :", err)
		return "User Password Update Failed", false
	}
	return "User Password Update Sucessfully Completed", true
}

func (user *UserRepo) UserGetall() ([]models.User, bool) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in User Getall")
	}
	userStruct := models.User{}
	result := []models.User{}

	query, err := Db.Query(`SELECT id,name,email,mobileno,role,password,createdon FROM "user"`)
	if err != nil {
		log.Println(err)
	}

	for query.Next() {
		err := query.Scan(
			&userStruct.Id,
			&userStruct.Name,
			&userStruct.Email,
			&userStruct.Mobileno,
			&userStruct.Role,
			&userStruct.Password,
			&userStruct.CreatedOn,
		)
		if err != nil {
			fmt.Println("Error in User GetAll QueryRow :", err)
			return result, false
		}
		result = append(result, userStruct)
	}

	return result, true
}

func (user *UserRepo) UserGetById(obj *models.User) (models.User, bool, string) {
	Db, conncet := utls.OpenDbConnection()
	if !conncet {
		fmt.Println("DB Disconnceted in User GetById ")
	}
	userStruct := models.User{}

	query, _ := Db.Prepare(`SELECT id,name,email,mobileno,role,token,createdon from "user" where id=$1`)

	err := query.QueryRow(obj.Id).Scan(&userStruct.Id,
		&userStruct.Name,
		&userStruct.Email,
		&userStruct.Mobileno,
		&userStruct.Role,
		&userStruct.Token,
		&userStruct.CreatedOn)

	if err != nil {
		fmt.Println("Error in User GetById QueryRow :", err)
		return userStruct, false, "Failed"
	}
	return userStruct, true, "Sucessfully Compelted"
}

func (user *UserRepo) GetByUserMobileno(obj *models.User) (models.User, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in GetByUser Mobileno")
	}
	userStruct := models.User{}

	res, _ := Db.Prepare(`SELECT id,name,email,mobileno,role,token from "user" where mobileno=$1`)

	err := res.QueryRow(obj.Mobileno).Scan(&userStruct.Id,
		&userStruct.Name,
		&userStruct.Email,
		&userStruct.Mobileno,
		&userStruct.Role,
		&userStruct.Token)

	if err != nil {
		fmt.Println("Error in GetByUser Mobileno QueryRow :", err)
		return userStruct, false
	}
	return userStruct, true
}
