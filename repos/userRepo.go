package repos

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
	"log"
)

type UserInterface interface {
	UserCreate(obj *models.User) (string, bool)
	UserLogin(obj *models.LoginUser) (models.User, bool, string)
	UserDelete(obj *models.User) (bool, string)
	UserChangePassword(obj *models.UserChangePassword) (string, bool)
	UserCheckingPassword(obj *models.User) (string, bool)

	UserGetall() ([]models.User, bool)
	UserGetById(obj *models.User) (models.User, bool, string)
	GetByUserMobileno(obj *models.User) (models.User, bool)

	UserUpdateMobileno(obj *models.UserverifyMobileno) (models.UserverifyMobileno, bool, string)
	UserUpdateName(obj *models.User) (bool, string)
	Userverify(obj *models.Userverify) (int64, bool, string)
	UserCheckOtp(obj *models.Userverify) ( bool, string)
	UserUpdatePassword(obj *models.UserChangePassword) (string, bool)
	UserUpdateEmail(obj *models.UserverifyEmail) (bool, string)
	UserverifyMobileno(obj *models.UserverifyMobileno) (bool, string)
	UserverifyById(obj *models.Userverify) (bool, string) 
	
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
			return "This Mobile Number or Email already Used by Other Customer ", false
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
		otppassword,
		otpemail,
		otpmobileno
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
		models.GenerateOtp(),
		models.GenerateOtp(),
	).Scan(&obj.Id)

	if err != nil {
		fmt.Println("Error in CreateUser  QueryRow Scan :", err)
		return " Creating User Failed", false
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return "WELCOME TO NATIONAL STORE ", true
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
									token,
									createdon
									from "user" where mobileno=$1 and password=$2 and isdeleted=0`)
	if err != nil {
		log.Println("Error in User Login QueryRow :", err)

		return userStruct, false, "Invaild Mobile Number Or Password"
	}

	err = query.QueryRow(obj.Mobileno, obj.Password).Scan(&userStruct.Id,
		&userStruct.Name,
		&userStruct.Email,
		&userStruct.Mobileno,
		&userStruct.Role,
		&userStruct.Token,
		&userStruct.CreatedOn,
	)
	if err != nil {
		log.Println("Error in User Login QueryRow Scan :", err)
		return userStruct, false, "Invaild Mobile Number Or Password"
	}

	Token := utls.GenerateJwtToken(userStruct.Id, userStruct.Mobileno, userStruct.Email)
	// log.Println(userStruct.Id, Token)
	_, err = Db.Query(`UPDATE "user" SET token = $2 WHERE id=$1 and isdeleted=0`, &userStruct.Id, &Token)
	userStruct.Token = Token
	if err != nil {
		log.Println("Error in User Login Update TOKEN QueryRow :", err)
		return userStruct, false, "Invaild Mobile Number Or Password"
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return userStruct, true, "WELCOME TO NATIONAL STORE "
}

func (user *UserRepo) UserUpdateEmail(obj *models.UserverifyEmail) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB disconnceted in UserUpdateEmail ")
	}
	userStruct := models.User{}
	query, err := Db.Query(`SELECT email FROM "user" WHERE isdeleted=0`)
	if err != nil {
		log.Println("Error in User User  Update TOKEN in Update Email QueryRow :", err)
		return false, "Update Email Failed"
	}
	for query.Next() {
		query.Scan(&userStruct.Email)
		if obj.Email == userStruct.Email {
			fmt.Println(" Email already Used By Other User")
			return false, "THis Email already Used By Other Customer"
		}
	}
	query1 := `UPDATE "user" SET email = $2 WHERE id = $1  RETURNING mobileno`
	err = Db.QueryRow(query1, &obj.Id, &obj.Email).Scan(&obj.Mobileno)
	fmt.Println("", obj)
	if err != nil {
		fmt.Println("Error in User Update Email QueryRow :", err)
		return false, "Update Email Failed"
	}
	Token := utls.GenerateJwtToken(obj.Id, obj.Mobileno, obj.Email)
	_, err = Db.Query(`UPDATE "user" SET token = $2 WHERE id=$1 and isdeleted=0`, &obj.Id, &Token)
	obj.Token = Token
	if err != nil {
		log.Println("Error in User User  Update TOKEN in Update Email QueryRow :", err)
		return false, "Update Email Failed"
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return true, "Email Updated  Successfully"
}

func (user *UserRepo) UserUpdateName(obj *models.User) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB disconnceted in UserUpdateName ")
	}

	query1 := `UPDATE "user" SET name = $2 WHERE id = $1 AND isdeleted=0`
	_, err := Db.Exec(query1, &obj.Id, &obj.Name)
	fmt.Println("", obj)
	if err != nil {
		fmt.Println("Error in User Update Name QueryRow :", err)
		return false, "Failed"
	}

	defer func() {
		Db.Close()
	}()
	return true, "Updated  Successfully"
}

func (user *UserRepo) UserUpdateMobileno(obj *models.UserverifyMobileno) (models.UserverifyMobileno, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB disconnceted in UserUpdate Mobileno ")
	}

	query := `UPDATE "user" SET mobileno= $2 WHERE id = $1 and otpmobileno=$3  RETURNING email`

	err := Db.QueryRow(query, &obj.Id, &obj.Mobileno, &obj.OTP).Scan(&obj.Email)
	fmt.Println("", obj)
	if err != nil {
		fmt.Println("Error in User Update mobileno QueryRow :", err)
		return *obj, false, "Failed"
	}
	Token := utls.GenerateJwtToken(obj.Id, obj.Mobileno, obj.Email)
	_, err = Db.Query(`UPDATE "user" SET token = $2 WHERE id=$1 and isdeleted=0`, &obj.Id, &Token)
	
	if err != nil {
		log.Println("Error in User User  Update TOKEN in Update Mobileno QueryRow :", err)
		return *obj, false, "Failed"
	}
	obj.Token = Token
	defer func() {
		Db.Close()
	}()
	return *obj, true, "Mobile Number Updated  Successfully"
}

func (user *UserRepo) UserDelete(obj *models.User) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB disconnceted in User Delete ")
	}
	query := `UPDATE "user" SET isdeleted=1 WHERE id=$1 and isdeleted=0`
	_, err := Db.Exec(query, obj.Id)

	if err != nil {
		fmt.Println("Error in User Delete QueryRow :")
		return false, "User Delete Failed"
	}
	defer func() {
		Db.Close()
	}()
	return true, "User Deleted Successfully Completed"
}

func (user *UserRepo) UserChangePassword(obj *models.UserChangePassword) (string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in User Update Password ")
	}

	query := `UPDATE "user" SET password = $2 WHERE id=$1  AND isdeleted=0`
	qq, err := Db.Exec(query, &obj.Id, &obj.Password)

	if err != nil {
		fmt.Println("Error in User Change Password QueryRow :", err, qq)
		return "incorrect Password", false
	}
	defer func() {
		Db.Close()
	}()

	return "Password Changed Successfully", true
}

func (user *UserRepo) UserCheckingPassword(obj *models.User) (string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in User Update Password ")
	}
	password := ""
	query := `SELECT password FROM  "user" WHERE id=$1  AND isdeleted=0`
	err := Db.QueryRow(query, &obj.Id).Scan(&password)

	if err != nil {
		fmt.Println("Error in User Change Password QueryRow :", err)
		return "incorrect Password", false
	}

	defer func() {
		Db.Close()
	}()
	if password == obj.Password {
		return "Password changed  Successfully", true
	}
	return "incorrect Password", false
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
		log.Println("Error in User Get All QueryRow :", err)
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
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true
}

func (user *UserRepo) UserGetById(obj *models.User) (models.User, bool, string) {
	Db, conncet := utls.OpenDbConnection()
	if !conncet {
		fmt.Println("DB Disconnceted in User GetById ")
	}
	userStruct := models.User{}
	fmt.Println("", obj)
	err := Db.QueryRow(`SELECT id,name,email,mobileno,role,token,createdon from "user" where id=$1`, obj.Id).Scan(&userStruct.Id,
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
	defer func() {
		Db.Close()
	}()
	return userStruct, true, "Successfully Compelted"
}

func (user *UserRepo) GetByUserMobileno(obj *models.User) (models.User, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in GetByUser Mobileno")
	}
	userStruct := models.User{}

	query, err := Db.Prepare(`SELECT id,name,email,mobileno,role,token from "user" where mobileno=$1 and isdeleted=0`)
	if err != nil {
		fmt.Println("Error in GetByUser Mobileno QueryRow :", err)
		return userStruct, false
	}
	err = query.QueryRow(obj.Mobileno).Scan(&userStruct.Id,
		&userStruct.Name,
		&userStruct.Email,
		&userStruct.Mobileno,
		&userStruct.Role,
		&userStruct.Token)

	if err != nil {
		fmt.Println("Error in GetByUser Mobileno QueryRow Scan :", err)
		return userStruct, false
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return userStruct, true
}

func (user *UserRepo) Userverify(obj *models.Userverify) (int64, bool, string) {
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
		query.Scan(&userStruct.Id, &userStruct.Mobileno, &userStruct.Email)

		if obj.VerifyUser == userStruct.Mobileno || obj.VerifyUser == userStruct.Email {
			obj.OTP = otp

			query := `UPDATE "user" SET otppassword=$2 WHERE id=$1 and isdeleted=0`
			_, err := Db.Exec(query, &userStruct.Id, obj.OTP)

			if err != nil {
				log.Println("Error in  User verify Update Otp QueryRow :", err)
				// missing  OTP SEND TO MOBILENUMBER
				return 0, false, "Invaild User"
			}
			fmt.Println(" userVerify DETAILS :", obj.VerifyUser)
			fmt.Println("OTP :", obj.OTP)
			return userStruct.Id, true, "Successfully Compeleted"
		}
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return 0, false, "Invaild User"
}

func (user *UserRepo) UserverifyById(obj *models.Userverify) (bool, string) {
	Db, isConnected := utls.OpenDbConnection()

	if !isConnected {
		fmt.Println("DB Disconnented in  User verify Repo ")
	}
			otp := models.GenerateOtp()
			query := `UPDATE "user" SET otppassword=$2 WHERE id=$1 and isdeleted=0`
			_, err := Db.Exec(query, &obj.Id, &otp)
			if err != nil{ 
				log.Println("Error in  User verify Update Otp QueryRow :", err)
				// missing  OTP SEND TO MOBILENUMBER
				return  false, "Invaild User"
			}
			fmt.Println(" userVerify DETAILS :", obj.VerifyUser)
			fmt.Println("OTP :", otp)
		
	defer func() {
		Db.Close()
	}()
	return  true, "Successfully Compeleted"
}

func (user *UserRepo) UserCheckOtp(obj *models.Userverify) (bool, string) {
	Db, conncet := utls.OpenDbConnection()
	if !conncet {
		fmt.Println("DB Disconnceted in User User Check Otp ")
	}
	otp := ""

	 err := Db.QueryRow(`SELECT otppassword from "user" where id=$1 and isdeleted=0`,obj.Id).Scan(&otp)
	if err != nil {
		fmt.Println("Error in User Check Otp  QueryRow :", err)
		return false, "Invaild Otp"
	}

	

	if otp == obj.OTP{
		return  true, "Validation Compelted"
	}
	defer func() {
		Db.Close()
	}()
	return false,"Invaild Otp"
	}
func (user *UserRepo) UserUpdatePassword(obj *models.UserChangePassword) (string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in User Update Password ")
	}

	query := `UPDATE "user" SET password = $2 WHERE id=$1 and isdeleted=0`
	_, err := Db.Exec(query, &obj.Id, &obj.NewPassword)

	if err != nil {
		fmt.Println("Error in User Update Password QueryRow :")
		return "Use Other Password", false
	}
	defer func() {
		Db.Close()
	}()
	return "User Password Changed Successfully Completed", true
}

func (user *UserRepo) UserverifyMobileno(obj *models.UserverifyMobileno) (bool, string) {
	Db, isConnected := utls.OpenDbConnection()

	if !isConnected {
		fmt.Println("DB Disconnented in  User verify Mobileno Repo ")
	}
	query, err := Db.Query(`SELECT mobileno FROM "user" WHERE isdeleted=0`)
	if err != nil {
		log.Println("Error in Update User mobileno Checking User verfiy mobileno QueryRow :", err)
		return false, "Invaild Email"
	}

	userStruct := models.User{}
	for query.Next() {
		query.Scan(&userStruct.Email)
		if obj.Mobileno == userStruct.Mobileno {
			fmt.Println(" mobileno already Used By Other User")
			return false, "THis Mobile Number already Used By Other Customer."
		}
	}
	otp := models.GenerateOtp()
	query1 := `UPDATE "user" SET otpmobileno=$2 WHERE id=$1 and isdeleted=0`
	_, err = Db.Exec(query1, &obj.Id, otp)

	if err != nil {
		log.Println("Error in  User verify mobileno Otp QueryRow :", err)
		// missing  OTP SEND TO MOBILENUMBER GATEWAY
		return false, "Invaild Mobile Number"
	}
	fmt.Println("Mobile Number :", obj.Mobileno)
	fmt.Println("OTP :",otp)

	defer func() {
		Db.Close()
		query.Close()
	}()
	return true, "OTP Sended Sucessfully"
}
