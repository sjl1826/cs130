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
	"github.com/lib/pq"
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

// CourseByID gets the course by ID
func CourseByID(db *gorm.DB, c *models.Course, w http.ResponseWriter) int {
	if err := c.GetCourse(db); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "Course not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return 0
	}
	return 1
}

// InvitationByID gets the invitation by ID
func InvitationByID(db *gorm.DB, i *models.Invitation, w http.ResponseWriter) int {
	if err := i.GetInvitation(db); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "Invitation not found")
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
	ID                int                 `json:"u_id"`
	CreatedAt         time.Time           `json:"CreatedAt"`
	UpdatedAt         time.Time           `json:"UpdatedAt"`
	FirstName         string              `json:"first_name"`
	LastName          string              `json:"last_name"`
	Email             string              `json:"u_email"`
	Biography         string              `json:"biography"`
	Discord			  string              `json:"discord"`
	Facebook          string              `json:"facebook"`
	Timezone          string              `json:"timezone"`
	SchoolName        string              `json:"school_name"`
	Courses           []models.Course     `json:"courses"`
	Groups            []models.Group      `json:"groups"`
	Listings          []models.Listing    `json:"listings"`
	Availability      []int64             `json:"availability"`
	Invitations       []models.Invitation `json:"invitations"`
}

func populateResponse(u *models.User, r *CreateResponse) {
	r.ID = u.ID
	r.CreatedAt = u.CreatedAt
	r.UpdatedAt = u.UpdatedAt
	r.FirstName = u.FirstName
	r.LastName = u.LastName
	r.Email = u.Email
	r.Biography = u.Biography
	r.Discord = u.Discord
	r.Facebook = u.Facebook
	r.Timezone = u.Timezone
	r.SchoolName = u.SchoolName
	r.Availability = u.Availability
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

	if err := p.GetCourses(db, &cr.Courses); err != nil {
		switch err {
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}
	if err := p.GetGroups(db, &cr.Groups); err != nil {
		switch err {
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}
	if err := p.GetListings(db, &cr.Listings); err != nil {
		switch err {
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}
	if err := p.GetInvitations(db, &cr.Invitations); err != nil {
		switch err {
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}

	respondWithJSON(w, http.StatusOK, cr)
}

// UpdateRequest for user requests parsing
type UpdateRequest struct {
	ID				int    	`json:"u_id"`
	FirstName		string 	`json:"first_name,omit_empty"`
	LastName		string 	`json:"last_name,omit_empty"`
	Email			string 	`json:"u_email"`
	Biography       string  `json:"biography"`
	Discord			string  `json:"discord"`
	Facebook        string  `json:"facebook"`
	Timezone        string  `json:"timezone"`
	SchoolName      string  `json:"school_name"`
	Availability    []int64 `json:"availability"`
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
	if p.Biography != "" {
		user.Biography = p.Biography
	}
	if p.Discord != "" {
		user.Discord = p.Discord
	}
	if p.Facebook != "" {
		user.Facebook = p.Facebook
	}
	if p.Timezone != "" {
		user.Timezone = p.Timezone
	}
	if p.SchoolName != "" {
		user.SchoolName = p.SchoolName
	}
	if p.Availability != nil {
		user.Availability = p.Availability
	}

	if err := user.UpdateUser(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var cr CreateResponse
	populateResponse(&user, &cr)

	respondWithJSON(w, http.StatusOK, cr)
}

// CourseRequest for addCourse requests parsing
type CourseRequest struct {
	ID				int      `json:"u_id"`
	CourseID		int 	 `json:"course_id,omit_empty"`
	CourseName		string   `json:"course_name"`
	Keywords		[]string `json:"keywords"`
	Categories		[]string `json:"categories"`
}

// AddCourse will add a course for the user
func AddCourse(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var p CourseRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	user := models.User{ID: p.ID}
	if GetUserByID(db, &user, w) == 0 {
		return
	}

	if p.CourseID == 0 {
		// Creates a new course if it doesn't exist
		var arr pq.Int64Array
		arr = append(arr, int64(p.ID))
		c := models.Course{Name: p.CourseName, Keywords: p.Keywords, Categories: p.Categories, StudyBuddies: arr}
		if err := c.CreateCourse(db); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		// Adds user to course if the course already exists
		course := models.Course{ID: p.CourseID}
		if CourseByID(db, &course, w) == 0 {
			return
		}

		if err := course.AddStudyBuddy(db, p.ID); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// RemoveCourse will remove a course for the user
func RemoveCourse(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var p CourseRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	user := models.User{ID: p.ID}
	if GetUserByID(db, &user, w) == 0 {
		return
	}

	course := models.Course{ID: p.CourseID}
	if CourseByID(db, &course, w) == 0 {
		return
	}

	if err := course.RemoveStudyBuddy(db, p.ID); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
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
