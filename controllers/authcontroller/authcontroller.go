package authcontroller

import (
	"encoding/json"
	"jwt/config"
	"jwt/helper"
	"jwt/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	// Get JSON BODY
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": "Gagal Decode"}
		helper.ResponeJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// Get Data User Berdasarkan Username
	var user models.User
	if err := models.DB.Where("username =?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Username Tidak Ditemukan"}
			helper.ResponeJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponeJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// Cek Password Valid
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		response := map[string]string{"message": "Password Salah"}
		helper.ResponeJSON(w, http.StatusUnauthorized, response)
		return
	}

	// Generate Token
	expTime := time.Now().Add(time.Minute * 5)
	claims := &config.JWTClaim{
		Id:       user.Id,
		Username: user.Username,
		Nama:     user.Nama,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// Algoritma Sign
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign Token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponeJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Set Token Cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "Login Berhasil"}
	helper.ResponeJSON(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {

	var userInput models.User
	// Get JSON BODY
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": "Gagal Decode"}
		helper.ResponeJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// HASHING PASS
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPass)

	// Exec DB
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": "Internal Server Error"}
		helper.ResponeJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Response
	response := map[string]string{"message": "Sukses"}
	helper.ResponeJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "Logout Berhasil"}
	helper.ResponeJSON(w, http.StatusOK, response)
}
