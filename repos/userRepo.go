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
	UserDelete(obj *models.User) (bool, string)

	UserGetall() ([]models.User, bool)
	UserGetById(obj *models.User) (models.User, bool, string)
	GetByUserMobileno(obj *models.User) (models.User, bool)

	Userverify(obj *models.Userverify) ( int64,bool, string)
	UserCheckOtp(obj *models.Userverify) (models.UserPassword, bool, string)
	UserUpdatePassword(obj *models.UserPassword) (string, bool)

}

type UserRepo struct {
}

func (user *UserRepo) UserCreate(obj *models.User) (string, bool) {
	Db, isConnected := utls.OpenDbConnection()

	if !isConnected {
		fmt.Println("DB Disconnented in CreateUser Repo ")
	}
	query, err := Db.Query(`SELECT email,mobileno FROM "user" WHERE isdeleted=0`)
	if err != nil {
		log.Println("Error in CreateUser Checking User verfiy QueryRow :", err)
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
		token,
		isdeleted,
		otppassword
	)values($1,$2,$3,$4,$5,$6,$7,$8,$9)RETURNING id`,
		obj.Name,
		obj.Email,
		obj.Mobileno,
		"3",
		obj.Password,
		utls.GetCurrentDateTime(),
		" NULL TOKEN",
		0,
		models.GenerateOtp(),
	).Scan(&obj.Id)

	if err != nil {
		fmt.Println("Error in CreateUser  QueryRow Scan :", err)
		return "User Create Failed", false
	}
	defer Db.Close()
	return "User created Successfully", true
}



func (user *UserRepo) UserLogin(obj *models.LoginUser) (models.User, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB DisConnected in UserLogin ")
	}
	userStruct := models.User{}
	query, err := Db.Prepare(`SELECT id,
	                                name,
									email,
									mobileno,
									role,
									createdon
									from "user" where mobileno=$1 and password=$2 and isdeleted=0`)
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
		log.Println("Error in User Login QueryRow Scan :", err)
		return userStruct, false, "User Login  Failed"
	}

	Token := utls.GenerateJwtToken(userStruct.Id, userStruct.Mobileno, userStruct.Email)
	// log.Println(userStruct.Id, Token)
	_,err= Db.Query( `UPDATE "user" SET token = $2 WHERE id=$1 and isdeleted=0`,&userStruct.Id,&Token)

	if err != nil {
		log.Println("Error in User Login Update TOKEN QueryRow :", err)
		return userStruct, false, "User Login  Failed"
	}
defer Db.Close()
	return userStruct, true, "User Login Successfully Completed"
}



func (user *UserRepo) UserUpdate(obj *models.UserUpdate) (models.UserUpdate, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB disconnceted in UserUpdate ")
	}
	err :=Db.QueryRow( `UPDATE "user" SET name=$2,email=$3,mobileno=$4 WHERE id=$1 and isdeleted=0`,
	&obj.Id,&obj.Name,&obj.Email,&obj.Mobileno)
	
	if err != nil {
		fmt.Println("Error in User Update QueryRow :", err)
		return *obj, false, "User Update Failed"
	}
	defer Db.Close()
	return *obj, true, "User Updated Successfully"
}



func (user *UserRepo) UserDelete(obj *models.User) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB disconnceted in User Delete ")
	}
	err :=Db.QueryRow( `UPDATE "user" SET isdeleted=1 WHERE id=$1 and isdeleted=0`,&obj.Id)
	
	if err != nil {
		fmt.Println("Error in User Delete QueryRow :", err)
		return false, "User Delete Failed"
	}
	defer Db.Close()
	return true, "User Deleted Successfully Completed"
}



func (user *UserRepo) UserGetall() ([]models.User, bool) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in User Getall")
	}
	userStruct := models.User{}
	result := []models.User{}

	query, err := Db.Query(`SELECT id,name,email,mobileno,role,createdon FROM "user" WHERE isdeleted=0`)
	if err != nil {
		log.Println("Error in User Get All QueryRow :",err)
	}

	for query.Next() {
		err := query.Scan(
			&userStruct.Id,
			&userStruct.Name,
			&userStruct.Email,
			&userStruct.Mobileno,
			&userStruct.Role,
			&userStruct.CreatedOn,
		)
		if err != nil {
			fmt.Println("Error in User GetAll QueryRow Scan :", err)
			return result, false
		}
		result = append(result, userStruct)
	}
defer Db.Close()
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
	return userStruct, true, "Successfully Compelted"
}



func (user *UserRepo) GetByUserMobileno(obj *models.User) (models.User, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in GetByUser Mobileno")
	}
	userStruct := models.User{}

	res, err := Db.Prepare(`SELECT id,name,email,mobileno,role,token from "user" where mobileno=$1 and isdeleted=0`)
	if err != nil {
		fmt.Println("Error in GetByUser Mobileno QueryRow :", err)
		return userStruct, false
	}
	err = res.QueryRow(obj.Mobileno).Scan(&userStruct.Id,
		&userStruct.Name,
		&userStruct.Email,
		&userStruct.Mobileno,
		&userStruct.Role,
		&userStruct.Token)

	if err != nil {
		fmt.Println("Error in GetByUser Mobileno QueryRow Scan :", err)
		return userStruct, false
	}
	return userStruct, true
}



func (user *UserRepo) Userverify(obj *models.Userverify) ( int64,bool, string) {
	Db, isConnected := utls.OpenDbConnection()

	if !isConnected {
		fmt.Println("DB Disconnented in  User verify Repo ")
	}
	query, err := Db.Query(`SELECT id,mobileno,email FROM "user" WHERE isdeleted=0`)
	if err != nil {
		log.Println("Error in  User verify QueryRow :", err)
	}
	fmt.Print(obj)
	userStruct := models.User{}
	otp := models.GenerateOtp()
	for query.Next() {
		query.Scan(&userStruct.Id,&userStruct.Mobileno,&userStruct.Email)

		if obj.VerifyUser == userStruct.Mobileno || obj.VerifyUser == userStruct.Email{
			obj.OTP = otp
			err :=Db.QueryRow( `UPDATE "user" SET otppassword=$2 WHERE id=$1 and isdeleted=0`,&userStruct.Id,obj.OTP)

			if err != nil {
				log.Println("Error in  User verify Update Otp QueryRow :", err)
				// missing  OTP SEND TO MOBILENUMBER
				return 0,false, "Invaild User"
			}
			fmt.Println(" userVerify DETAILS :", obj.VerifyUser)
			fmt.Println("OTP :",obj.OTP)
			return userStruct.Id,true, "Successfully Compeleted"
		}
	}
	defer Db.Close()
	return 0,false, "Invaild User"
}



func (user *UserRepo) UserCheckOtp(obj *models.Userverify) (models.UserPassword, bool, string) { 
	Db, conncet := utls.OpenDbConnection()
	if !conncet {
		fmt.Println("DB Disconnceted in User User Check Otp ")
	}
	userStruct := models.UserPassword{}

	query, err := Db.Prepare(`SELECT id,mobileno, password from "user" where id=$1 and otppassword=$2 and isdeleted=0`)
	if err != nil {
		fmt.Println("Error in User Check Otp  QueryRow :", err)
		return userStruct, false, "Failed"
	}

	err = query.QueryRow(obj.Id, obj.OTP).Scan(&userStruct.Id,&userStruct.Mobileno,&userStruct.Password)

	if err != nil {
		fmt.Println("Error in User Check Otp QueryRow Scan :", err)
		return userStruct, false, "Failed"
	}
defer query.Close()
	return userStruct, true, "Successfully Compelted"
}



func (user *UserRepo) UserUpdatePassword(obj *models.UserPassword) (string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in User Update Password ")
	}

	err :=Db.QueryRow( `UPDATE "user" SET password = $2 WHERE id=$1 and isdeleted=0`,&obj.Id,&obj.Password)
	
	if err != nil {
		fmt.Println("Error in User Update Password QueryRow :", err)
		return "User Password Update Failed", false
	}
	return "User Password Update Successfully Completed", true
}



