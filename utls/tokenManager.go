package utls

import (
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type AppClaims struct {
	UserId   int64
	Mobileno string
	Email     string
	jwt.StandardClaims
}

type GeneralResponseResource struct {
	TokenStatus  string      `json:"tokenstatus"`
	ResponseData interface{} `json:"responsedata"`
}

const (
	privateKeyPath = "/keys/aim.rsa"
	publicKeyPath  = "/keys/aim.rsa.pub"
)

var (
	VerificationKey *rsa.PublicKey
	SigningKey      *rsa.PrivateKey
)

func initKeys() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("workingDir", workingDir)
	signedBytes, err := ioutil.ReadFile(workingDir + privateKeyPath)
	if err != nil {
		log.Fatal("RSA read file error", err)
		panic(err)
	}

	SigningKey, err = jwt.ParseRSAPrivateKeyFromPEM(signedBytes)
	if err != nil {
		log.Fatal("RSA Parse sign error", err)
	}

	verificationBytes, err := ioutil.ReadFile(workingDir + publicKeyPath)
	if err != nil {
		log.Fatal("RSA read file error ", err)
		panic(err)
	}

	VerificationKey, err = jwt.ParseRSAPublicKeyFromPEM(verificationBytes)
	if err != nil {
		log.Fatal("RSA Parse verif file error", err)
	}

}

func GenerateJwtToken(id int64, mobileno, email string) string {
	//create Claims
	claims := AppClaims{
		UserId:   id,
		Mobileno: mobileno,
		Email:     email,
		StandardClaims: jwt.StandardClaims{
			Issuer:"OnlineShopAdmin",
		},
	}
	//create a signer using RSA256 key
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	//using the key, Sign  and get a token
	tokenString, err := token.SignedString(SigningKey)
	if err != nil {
		log.Fatal(err)
	}
	// return token
	return tokenString
}

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware for validition jwt token
		objrespone := GeneralResponseResource{}

		reqToken := r.Header.Get("Authorization")

		splitToken := strings.Split(reqToken, "Bearer")

		reqToken = strings.TrimSpace(splitToken[1])

		claims := jwt.MapClaims{}

		// validation token
		token, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use it public counter part to verfiykey

			return VerificationKey, nil
		})

		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				log.Println("reason : ", "ValidaionErrorMalformed")
				w.WriteHeader(http.StatusUnauthorized)

				objrespone.TokenStatus = "false"
				objrespone.ResponseData = "Token Malformed !"
				jsonResponse(objrespone, w)
				return

			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				log.Println("Reason : ", "Token Expired")
				w.WriteHeader(http.StatusUnauthorized)
				objrespone.TokenStatus = "false"
				objrespone.ResponseData = "Token Expired !"
				jsonResponse(objrespone, w)
				return
			} else {
				log.Println("Reason : ", "StatusUnauthorized : ", ve.Errors)
				w.WriteHeader(http.StatusUnauthorized)
				objrespone.TokenStatus = "false"
				objrespone.ResponseData = "Token Invaild !"
				jsonResponse(objrespone, w)

			}
		}
		if token != nil {
			if token.Valid {
				log.Println("Request Authorized...")
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				objrespone.TokenStatus = "false"
				objrespone.ResponseData = "Invalid Token"
				jsonResponse(objrespone, w)
			}
		}

	})
}

func jsonResponse(respone interface{}, w http.ResponseWriter) {
	jsonData, err := json.Marshal(respone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
