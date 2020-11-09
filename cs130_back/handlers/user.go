// user.go

package handlers

import (
	"cs130_back/hash"
	"cs130_back/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// GetUserByID gets the user by ID
func GetUserByID(db *gorm.DB, u *models.User, w http.ResponseWriter) int {
	if err := u.GetUser(db); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return 0
	}
	return 1
}

// CreateRequest required fields to create a user
type CreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"u_email"`
	Password  string `json:"password"`
}

// CreateResponse fields to send back
// HTTP status code 201 and user model in data
type CreateResponse struct {
	ID                int       `json:"u_id"`
	CreatedAt         time.Time `json:"CreatedAt"`
	UpdatedAt         time.Time `json:"UpdatedAt"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"u_email"`
}

func populateResponse(u *models.User, r *CreateResponse) {
	r.ID = u.ID
	r.CreatedAt = u.CreatedAt
	r.UpdatedAt = u.UpdatedAt
	r.FirstName = u.FirstName
	r.LastName = u.LastName
	r.Email = u.Email
}

// CreateUser initializes a new user in the database
func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var p CreateRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	encryptedPassword := hash.Salt([]byte(p.Password))

	user := models.User{FirstName: p.FirstName, LastName: p.LastName, Email: p.Email, Password: encryptedPassword}
	if err := user.CreateUser(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := user.GetByEmail(db); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	var cr CreateResponse
	populateResponse(&user, &cr)

	respondWithJSON(w, http.StatusCreated, cr)
}

// LoginResponse is the response with information to access other methods
// HTTP status code 200 and user login model in data
// swagger:response loginResp
type LoginResponse struct {
	AccessToken      string        `json:"access_token"`
	ExpiresIn        time.Duration `json:"expires_in"`
	RefreshToken     string        `json:"refresh_token"`
	RefreshExpiresIn time.Duration `json:"refresh_expires_in"`
}

// LoginUser checks for a valid email and password to generate an access token
func LoginUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := models.User{Email: email}
	if err := user.GetByEmail(db); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	user.GetPassword(db)
	if hash.ComparePasswords(user.Password, []byte(password)) {
		newToken := models.Token{UserID: user.ID}
		if err := newToken.New(db, &user); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		var response LoginResponse
		response.AccessToken = newToken.AccessToken
		response.ExpiresIn = newToken.AccessExpiresIn
		response.RefreshToken = newToken.RefreshToken
		response.RefreshExpiresIn = newToken.RefreshExpiresIn
		respondWithJSON(w, http.StatusOK, response)
	} else {
		respondWithError(w, http.StatusBadRequest, "Invalid user credentials")
	}
}

// RefreshToken generates a new token
func RefreshToken(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	refresh := r.FormValue("refresh_token")

	token := models.Token{RefreshToken: refresh}
	if err := token.GetTokenByRefresh(db); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "Token not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	expirationTime := token.AccessCreateAt.Add(token.AccessExpiresIn)
	diff := expirationTime.Sub(time.Now())
	if diff <= 0 {
		respondWithError(w, http.StatusNotFound, "Token expired")
		return
	}

	user := models.User{ID: token.UserID}
	if GetUserByID(db, &user, w) == 0 {
		return
	}
	newToken := models.Token{UserID: user.ID}
	if err := newToken.New(db, &user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	var response LoginResponse
	response.AccessToken = newToken.AccessToken
	response.ExpiresIn = newToken.AccessExpiresIn
	response.RefreshToken = newToken.RefreshToken
	response.RefreshExpiresIn = newToken.RefreshExpiresIn
	respondWithJSON(w, http.StatusOK, response)
}

// GetUser retrieves and returns the user
func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id, ok := strconv.Atoi(vars["u_id"][0])
	if ok != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user id")
		return
	}
	p := models.User{ID: id}
	if GetUserByID(db, &p, w) == 0 {
		return
	}

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		respondWithError(w, http.StatusBadRequest, "Invalid token format")
		return
	}

	reqToken = strings.TrimSpace(splitToken[1])
	fullToken := models.Token{AccessToken: reqToken}
	if err := fullToken.GetTokenByAccess(db); err != nil {
		respondWithError(w, http.StatusNotFound, "Invalid Token")
		return
	}
	if fullToken.UserID != id {
		respondWithError(w, http.StatusNotFound, "Invalid Token")
		return
	}

	var cr CreateResponse
	populateResponse(&p, &cr)

	respondWithJSON(w, http.StatusOK, cr)
}

// CoursesResponse is the structure in which the courses are returned
type CoursesResponse struct {
	Courses             []models.Course    `json:"courses"`
}

// GetCourses retrieves and returns the user's course objects
func GetCourses(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id, ok := strconv.Atoi(vars["u_id"][0])
	if ok != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user id")
		return
	}
	p := models.User{ID: id}
	if GetUserByID(db, &p, w) == 0 {
		return
	}
	var courses []models.Course
	if err := p.GetCourses(db, &courses); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sr := CoursesResponse{Courses: courses}

	respondWithJSON(w, http.StatusOK, sr)
}

// UpdateRequest for user requests parsing
type UpdateRequest struct {
	ID				int    	`json:"u_id"`
	FirstName		string 	`json:"first_name,omit_empty"`
	LastName		string 	`json:"last_name,omit_empty"`
	Email			string 	`json:"u_email"`
}

// UpdateUser will update the values of the specified user
func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var p UpdateRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	user := models.User{ID: p.ID, Email: p.Email}
	if GetUserByID(db, &user, w) == 0 {
		return
	}

	if p.FirstName != "" {
		user.FirstName = p.FirstName
	}
	if p.LastName != "" {
		user.LastName = p.LastName
	}

	if err := user.UpdateUser(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var cr CreateResponse
	populateResponse(&user, &cr)

	respondWithJSON(w, http.StatusOK, cr)
}

// DeleteUser deletes the user permanently
func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id, ok := strconv.Atoi(vars["u_id"][0])
	if ok != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	p := models.User{ID: id}
	if GetUserByID(db, &p, w) == 0 {
		return
	}

	if err := p.DeleteUser(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
