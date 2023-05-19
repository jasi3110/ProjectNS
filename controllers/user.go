package controllers

import (
	"OnlineShop/models"
	"OnlineShop/repos"
	"OnlineShop/repos/masterRepo"
	"OnlineShop/utls"
	"fmt"

	// "context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	// "github.com/go-delve/delve/service"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type UserController struct {
}

func (user *UserController) UserCreate(w http.ResponseWriter, r *http.Request) {

	request := models.User{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Println("Error in Decoding UserCreate Request :", err)
	response := models.CommanRespones{
			Statuscode:  200,
			Status:      false,
			Descreption: "Check The Details You Entered",
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			fmt.Println("Error in Marshal CartUpdate Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}else{
	status, description := Validrequst(request)
	if !status {
		response := models.CommanRespones{
			Statuscode:  200,
			Status:      status,
			Descreption: description,
		}
		resp, err := json.Marshal(response)

		if err != nil {
			log.Println("Error in Marshal Validation Response :", err)
		}
		w.Write(resp)
	} else {
		repo := repos.UserInterface(&repos.UserRepo{})

		description, status := repo.UserCreate(&request)

		response := models.CommanRespones{
			Statuscode:  200,
			Status:      status,
			Descreption: description,
		}

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

	request := models.LoginUser{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Println("Error in Decoding UserLogin Request :", err)
	response := models.CommanRespones{
			Statuscode:  200,
			Status:      false,
			Descreption: "Failed",
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			fmt.Println("Error in Marshal CartUpdate Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}else{
	repo := repos.UserInterface(&repos.UserRepo{})

	value, status, descreption := repo.UserLogin(&request)

	response := models.UserResponseModel{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}

	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshal Login respone :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}
}

func (User *UserController) UserUpdateEmail(w http.ResponseWriter, r *http.Request) {
	request := models.UserUpdate{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error in Decoding UserUpdate Request :", err)
	}

	status := models.VerifyEmail(request.Email)
	if !status {
		response := models.CommanRespones{
			Statuscode:  200,
			Status:      status,
			Descreption: "Invalid Email",
		}
		resp, err := json.Marshal(response)

		if err != nil {
			log.Println("Error in marshal UserUpdate Validation Response :", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	} else {
		repo := repos.UserInterface(&repos.UserRepo{})
		value, status, descreption := repo.UserUpdateEmail(&request)
		respone := models.UserUpdateResponseModel{
			Statuscode:  200,
			Status:      status,
			Value:       value,
			Descreption: descreption,
		}
		resp, err := json.Marshal(&respone)
		if err != nil {
			log.Println("Error in marshal UserUpdate Response :", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}

func (User *UserController) UserChangePassword(w http.ResponseWriter, r *http.Request) {
	request := models.UserChangePassword{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error in Decoding UserChangePassword Request :", err)
		response := models.CommanRespones{
			Statuscode:  200,
			Status:      false,
			Descreption: "Failed",
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			fmt.Println("Error in Marshal CartUpdate Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}else{

	repo := repos.UserInterface(&repos.UserRepo{})
	descreption, status := repo.UserChangePassword(&request)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	resp, err := json.Marshal(&response)
	if err != nil {
		log.Println("Error in Marshal Change UserPassword Responsee:", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}
} 

func (User *UserController) UserUpdatePassword(w http.ResponseWriter, r *http.Request) {
	request := models.UserChangePassword{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error in Decoding UserUpdatePassword Request :", err)
	response := models.CommanRespones{
			Statuscode:  200,
			Status:      false,
			Descreption: "Failed",
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			fmt.Println("Error in Marshal CartUpdate Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}else{

	repo := repos.UserInterface(&repos.UserRepo{})
	descreption, status := repo.UserUpdatePassword(&request)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	resp, err := json.Marshal(&response)
	if err != nil {
		log.Println("Error in Marshal Update UserPassword Responsee:", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}
}

func (User *UserController) UserGetAll(w http.ResponseWriter, r *http.Request) {

	repo := repos.UserInterface(&repos.UserRepo{})
	value, status := repo.UserGetall()
	response := models.GetAllUserResponseModel{
		Status: status,
		Value:  value,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		log.Println("Error in Marshal GetAll User Response:", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (user *UserController) UserGetById(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	userid := models.User{
		Id: id,
	}
	if err != nil {
		log.Println("Error in Decoding UserGetById Request :", err)
	response := models.CommanRespones{
			Statuscode:  200,
			Status:      false,
			Descreption: "Failed",
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			fmt.Println("Error in Marshal CartUpdate Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}else{
	repo := repos.UserInterface(&repos.UserRepo{})

	value, status, descreption := repo.UserGetById(&userid)
	roleStruct := models.Role{
		Id: value.Role,
	}
	roleRepo := masterRepo.RoleInterface(&masterRepo.RoleStruct{})
	role, _, _ := roleRepo.RoleById(&roleStruct)
	value.Role = role.Type

	response := models.UserResponseModel{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshal UserGetById Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}
}

func (user *UserController) Userverfiy(w http.ResponseWriter, r *http.Request) {
	request := models.Userverify{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error in Decoding UserUpdatePassword Request :", err)
	response := models.CommanRespones{
			Statuscode:  200,
			Status:      false,
			Descreption: "Failed",
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			fmt.Println("Error in Marshal CartUpdate Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}else{
	log.Println(request)

	if !models.VerifyMobileno(request.VerifyUser) && !models.VerifyEmail(request.VerifyUser) {
		response := models.CommanRespones{
			Statuscode:  200,
			Status:      false,
			Descreption: "Check your Mobile Number Or Email",
		}
		resp, err := json.Marshal(response)

		if err != nil {
			log.Println("Error in Marshal UserUpdatePassword Validation Response :", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	} else {
		repo := repos.UserInterface(&repos.UserRepo{})
		value, status, descreption := repo.Userverify(&request)
		response := models.UserverfiyOtp{
			Statuscode:  200,
			Status:      status,
			Value:       value,
			Descreption: descreption,
		}

		resp, err := json.Marshal(&response)
		if err != nil {
			log.Println("Error in Marshal Update UserPassword Responsee:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}
}

func (user *UserController) UserCheckOtp(w http.ResponseWriter, r *http.Request) {
	request := models.Userverify{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error in Decoding UserUpdatePassword Request :", err)
	response := models.CommanRespones{
			Statuscode:  200,
			Status:      false,
			Descreption: "Failed",
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			fmt.Println("Error in Marshal CartUpdate Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}else{

	repo := repos.UserInterface(&repos.UserRepo{})
	value, status, descreption := repo.UserCheckOtp(&request)
	response := models.UserUpdatePassword{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(&response)
	if err != nil {
		log.Println("Error in Marshal Update UserPassword Responsee:", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}
}

func (user *UserController) UserDelete(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	userid := models.User{
		Id: id,
	}
	if err != nil {
		log.Println("Error in Decoding UserGetById Request :", err)
	response := models.CommanRespones{
			Statuscode:  200,
			Status:      false,
			Descreption: "Failed",
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			fmt.Println("Error in Marshal CartUpdate Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}else{
	repo := repos.UserInterface(&repos.UserRepo{})

	status, descreption := repo.UserDelete(&userid)

	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshal UserGetById Response :", err)
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
