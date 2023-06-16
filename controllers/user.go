package controllers

import (
	"OnlineShop/models"
	"OnlineShop/repos"
	"OnlineShop/utls"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type UserController struct {
}

func (user *UserController) UserCreate(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic(r)
		}
	}()

	request := models.User{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Println("Error in Decoding UserCreate Request :", err)
		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = false
		response.Descreption = "Check The Details You Entered"

		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal CartUpdate Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		status, description := Validrequst(request)
		if !status {
			response := models.CommanRespones{}
			response.Statuscode = 200
			response.Status = status
			response.Descreption = description
			resp, err := json.Marshal(response)

			if err != nil {
				log.Panic("Error in Marshal Validation Response :", err)
			}
			w.Write(resp)
		} else {
			repo := repos.UserInterface(&repos.UserRepo{})
			description, status := repo.UserCreate(&request)

			response := models.CommanRespones{}
			response.Statuscode = 200
			response.Status = status
			response.Descreption = description

			resp, err := json.Marshal(response)
			if err != nil {
				log.Println("Error in Marshal UserCreate Response : ", err)
			}
			w.Header().Set("Content-Type", "Application/json")
			w.Write(resp)
		}
	}
}

func (user *UserController) UserLogin(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic(r)
		}
	}()

	request := models.LoginUser{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Panic("Error in Decoding UserLogin Request :", err)
		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = false
		response.Descreption = "Login Failed"

		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal User Login Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	} else {
		repo := repos.UserInterface(&repos.UserRepo{})
		value, status, descreption := repo.UserLogin(&request)

		response := models.UserResponseModel{}
		response.Statuscode = 200
		response.Status = status
		response.Value = value
		response.Descreption = descreption

		resp, err := json.Marshal(response)
		if err != nil {
			log.Panic("Error in Marshal Login respone :", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}

func (User *UserController) UserUpdateEmail(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic(r)
		}
	}()

	request := models.UserverifyUpdate{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Panic("Error in Decoding UserUpdate Email Request :", err)
	}

	status := models.VerifyEmail(request.Email)
	if !status {
		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = status
		response.Descreption = "Invalid Email"

		resp, err := json.Marshal(response)

		if err != nil {
			log.Panic("Error in marshal UserUpdate Email Validation Response :", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	} else {
		repo := repos.UserInterface(&repos.UserRepo{})
		value,status, descreption := repo.UserUpdateEmail(&request)
		respone := models.UserVerifyUpdateResponseModel{}
		respone.Statuscode = 200
		respone.Status = status
		respone.Value=value
		respone.Descreption = descreption

		resp, err := json.Marshal(&respone)
		if err != nil {
			log.Panic("Error in marshal UserUpdate Email Response :", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}

func (User *UserController) UserUpdateMobileno(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := models.UserverifyUpdate{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Panic("Error in Decoding User Update  Mobile Number Request :", err)
	}

	status := models.VerifyMobileno(request.Mobileno)
	if !status {
		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = status
		response.Descreption = "Invailed Mobile Number"

		resp, err := json.Marshal(response)
		if err != nil {
			log.Panic("Error in marshal User Update Mobile Number Validation Response :", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		repo := repos.UserInterface(&repos.UserRepo{})
		value, status, descreption := repo.UserUpdateMobileno(&request)

		respone := models.UserVerifyUpdateResponseModel{}
		respone.Statuscode = 200
		respone.Status = status
		respone.Value = value
		respone.Descreption = descreption

		resp, err := json.Marshal(&respone)
		if err != nil {
			log.Panic("Error in marshal UserUpdate Response :", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}

func (User *UserController) UserChangePassword(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := models.UserChangePassword{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Panic("Error in Decoding UserChangePassword Request :", err)

		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = false
		response.Descreption = "Invaild Operation"

		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal UserChangePassword Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
	status := models.VerifyPassword(request.Password)
	if !status {
		response := models.CommanRespones{
			Statuscode:  200,
			Status:      status,
			Descreption: "Make Password Strong",
		}
		resp, err := json.Marshal(response)

		if err != nil {
			log.Panic("Error in marshal UserChangePassword Validation Response :", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	} else {

		repo := repos.UserInterface(&repos.UserRepo{})
		descreption, status := repo.UserChangePassword(&request)
		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = status
		response.Descreption = descreption

		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal Change UserPassword Responsee:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}

func (User *UserController) UserCheckingPassword(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := models.User{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Panic("Error in Decoding UserCheckingPassword Request :", err)

		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = false
		response.Descreption = "Failed"

		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal UserCheckingpassword Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		repo := repos.UserInterface(&repos.UserRepo{})
		descreption, status := repo.UserCheckingPassword(&request)
		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = status
		response.Descreption = descreption

		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal UserCheckingPassword Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}

func (User *UserController) UserUpdatePassword(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := models.UserChangePassword{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Panic("Error in Decoding UserUpdatePassword Request :", err)

		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = false
		response.Descreption = "Failed"

		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal UserUpdatePassword Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	} else {

		repo := repos.UserInterface(&repos.UserRepo{})
		descreption, status := repo.UserUpdatePassword(&request)
		response := models.CommanRespones{
			Statuscode:  200,
			Status:      status,
			Descreption: descreption,
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal UpdateUserPassword Responsee:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}

func (User *UserController) UserVerifyById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := models.Userverify{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Panic("Error in Decoding UserVerifyById Request :", err)

		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=false
		response.Descreption="Failed"
		
		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal UserVerifyById Response:", err)
		}
		
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		if !models.VerifyMobileno(request.VerifyUser) && !models.VerifyEmail(request.VerifyUser) {

			response := models.CommanRespones{}
			response.Statuscode=200
			response.Status=false
			response.Descreption="Check Your Mobile Number or Email"
			
			resp, err := json.Marshal(response)
			if err != nil {
				log.Panic("Error in Marshal UserVerfiyById Validation Response :", err)
			}
			w.Header().Set("Content-Type", "Application/json")
			w.Write(resp)
		} else {
			repo := repos.UserInterface(&repos.UserRepo{})
			status, descreption := repo.UserverifyById(&request)

			response := models.CommanRespones{}
			response.Statuscode=200
			response.Status=status
			response.Descreption=descreption

			resp, err := json.Marshal(&response)
			if err != nil {
				log.Panic("Error in Marshal UserVerfiyById Response:", err)
			}

			w.Header().Set("Content-Type", "Application/json")
			w.Write(resp)
		}
	}
}

func (User *UserController) UserUpdateName(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := models.User{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Panic("Error in Decoding UserUpdateName Request :", err)

		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=false
		response.Descreption="Failed"

		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal UserUpdateName Response:", err)
		}
		
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	} else {

		repo := repos.UserInterface(&repos.UserRepo{})
		status, descreption := repo.UserUpdateName(&request)
		response := models.CommanRespones{
			Statuscode:  200,
			Status:      status,
			Descreption: descreption,
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal  UserUpdateName Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}

func (User *UserController) UserGetAll(w http.ResponseWriter, r *http.Request) {

	repo := repos.UserInterface(&repos.UserRepo{})
	value, status := repo.UserGetall()
	response := models.GetAllUserResponseModel{}
	response.Statuscode=200
	response.Status=status
	response.Value=value
		
	resp, err := json.Marshal(response)
	if err != nil {
		log.Panic("Error in Marshal User GetAll Response:", err)
	}

	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (user *UserController) UserGetById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	userid := models.User{Id: id,}
	if err != nil {
		log.Panic("Error in Decoding UserGetById Request :", err)

		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=false
		response.Descreption="Failed"
		
		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal UserGetId Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		repo := repos.UserInterface(&repos.UserRepo{})
		value, status, descreption := repo.UserGetById(&userid)

		response := models.UserResponseModel{}
		response.Statuscode= 200
		response.Status=status
		response.Value=value
		response.Descreption=descreption
			
		
		resp, err := json.Marshal(response)
		if err != nil {
			log.Panic("Error in Marshal UserGetById Response :", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}

func (user *UserController) Userverify(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := models.Userverify{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Panic("Error in Decoding User Verify Request :", err)

		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=false
		response.Descreption="Failed"
		
		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal User Verify Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		if !models.VerifyMobileno(request.VerifyUser) && !models.VerifyEmail(request.VerifyUser) {

			response := models.CommanRespones{}
			response.Statuscode=200
			response.Status=false
			response.Descreption="Check Your Mobile Number or Email"
				
			resp, err := json.Marshal(response)
			if err != nil {
				log.Panic("Error in Marshal User verify Validation Response :", err)
			}
			
			w.Header().Set("Content-Type", "Application/json")
			w.Write(resp)
		} else {

			repo := repos.UserInterface(&repos.UserRepo{})
			value, status, descreption := repo.Userverify(&request)

			response := models.UserverfiyOtp{}
			response.Statuscode=200
			response.Status=status
			response.Value=value
			response.Descreption=descreption

			resp, err := json.Marshal(&response)
			if err != nil {
				log.Panic("Error in Marshal User verify  Response:", err)
			}

			w.Header().Set("Content-Type", "Application/json")
			w.Write(resp)
		}
	}
}

func (user *UserController) UserCheckOtp(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := models.Userverify{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Panic("Error in Decoding User Check OTP Request :", err)

		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=false
		response.Descreption="Failed"
			
		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal User Check OTP Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		repo := repos.UserInterface(&repos.UserRepo{})
		status, descreption := repo.UserCheckOtp(&request)

		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=status
		response.Descreption=descreption

		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal User Check OTP Responsee:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	}
}

func (user *UserController) UserDelete(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	userid := models.User{Id: id,}

	if err != nil {
		log.Panic("Error in Decoding User Delete Request :", err)

		response := models.CommanRespones{}
		response.Statuscode=  200
		response.Status=    false
		response.Descreption= "Failed"
		
		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal User Delete Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		repo := repos.UserInterface(&repos.UserRepo{})
		status, descreption := repo.UserDelete(&userid)

		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=status
		response.Descreption=descreption
		
		
		resp, err := json.Marshal(response)
		if err != nil {
			log.Panic("Error in Marshal User Delete Response :", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}

func (user *UserController) UserVerifyMobileno(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := models.UserverifyUpdate{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Panic("Error in Decoding User Verfiy Mobileno Request :", err)

		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=false
		response.Descreption= "Check The Details You Entered"
			
		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal CartUpdate Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else if !models.VerifyMobileno(request.Mobileno) {

		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=false
		response.Descreption="Please Use Vaild Mobile Number "
		
		resp, err := json.Marshal(response)
		if err != nil {
			log.Panic("Error in Marshal User Verfiy Mobileno Validation Response :", err)
		}
		w.Write(resp)

	} else {

		repo := repos.UserInterface(&repos.UserRepo{})
		status, description := repo.UserverifyMobileno(&request)

		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=status
		response.Descreption=description
		
		resp, err := json.Marshal(response)
		if err != nil {
			log.Panic("Error in Marshal UserVerfiy Mobileno Response : ", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}

// VALIDATION METHODS

func Validrequst(obj models.User) (bool, string) {

	switch {
	case !models.VerifyMobileno(obj.Mobileno):
		return false, "Invailed Mobile Number"
	case !models.VerifyEmail(obj.Email):
		return false, "Invailed Email Address"
	case !models.VerifyPassword(obj.Password):
		return false, "Your Password Should be Minimum 8 Characters"
	default:
		return true, " Sucessfully Completed"
	}
}

func ValidrequstUpdate(obj models.UserUpdate) (bool, string) {
	if models.VerifyMobileno(obj.Mobileno) {
		if models.VerifyEmail(obj.Email) {
			return true, "Validation Sucessfully Completed"
		}
	}
	return false, "Invalid Mobile Number Or Email "
}

func ValidrequstPassword(obj models.UserPassword) (bool, string) {
	// var err string
	if models.VerifyMobileno(obj.Mobileno) {
		if models.VerifyPassword(obj.Password) {
			return true, "Validation Sucessfully Completed"
		}
	}
	return false, "Invalid Mobile Number Or Password "
}

// VALIDATION FOR TOKEN

func GetUserFromRequest(r *http.Request) models.User {
	requestToken := r.Header.Get("Authorization")

	splitToken := strings.Split(requestToken, "Bearer")

	requestToken = strings.TrimSpace(splitToken[1])

	claims := jwt.MapClaims{}

	// validation token
	token, _ := jwt.ParseWithClaims(requestToken, claims, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use it public counter part to verfiykey

		return utls.VerificationKey, nil
	})

	myClaims := token.Claims.(jwt.MapClaims)

	user := models.User{
		Mobileno: myClaims["Mobileno"].(string),
	}

	Repo := repos.UserInterface(&repos.UserRepo{})

	userRow, _ := Repo.GetByUserMobileno(&user)

	return userRow
}

func CheckAuthenticLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Recovered from panic condition : ", r)
			}
		}()

		repo := repos.UserInterface(&repos.UserRepo{})
		me := GetUserFromRequest(r)
		userId := models.User{
			Mobileno: me.Mobileno,
		}
		myResult, _ := repo.GetByUserMobileno(&userId)

		if myResult.Id > 0 && me.Token == GetUserFromRequest(r).Token {
			log.Println("Token Authorization Sucess...")
			next.ServeHTTP(w, r)
			// myContext := context.WithValue(r.Context(), me)
			// next.ServeHTTP(w, r.WithContext(myContext))
		} else {
			response := models.UserResponseModel{
				Statuscode:  200,
				Status:      false,
				Descreption: "Token UnAthorization ",
			}
			log.Println("Token Authorization Failed...")
			respone, err := json.Marshal(response)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return

			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(respone)
		}

	})
}
