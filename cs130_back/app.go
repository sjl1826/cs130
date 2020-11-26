// app.go

package cs130_back

import (
	"cs130_back/handlers"
	"cs130_back/models"
	"cs130_back/seeds"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {	
	Router *mux.Router
	DB *gorm.DB
}

const APP_DB_USERNAME = "postgres"
const APP_DB_PASSWORD = "admin"
const APP_DB_NAME = "cs130"

// Start starts the server
func Start() {
	a := App{}
	a.Initialize(
		APP_DB_USERNAME,
		APP_DB_PASSWORD,
		APP_DB_NAME)
	a.Run(":8080")
}

// Initialize creates the database and sets up the server
func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)

	var err error
	a.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	//a.DB = models.DBMigrate(a.DB)  <-- remove in a following commit if this doesn't break anything
	a.DB = seeds.Seed(a.DB)
	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	routes := a.Router.PathPrefix("/api/v1").Subrouter()
	routes.Use(loggingMiddleware)

	// User Routes
	userRoutes := routes.PathPrefix("/user").Subrouter()
	userRoutes.HandleFunc("/register", a.handleRequest(handlers.CreateUser)).Methods("POST")

	userRoutes.HandleFunc("/login", a.handleRequest(handlers.LoginUser)).Methods("POST")
	userRoutes.HandleFunc("/refresh", a.handleRequest(handlers.RefreshToken)).Methods("POST")
	userRoutes.HandleFunc("/classes-info", a.handleRequest(handlers.GetClassesInfo)).Methods("GET")

	authUserRoutes := userRoutes.PathPrefix("/o").Subrouter()
	authUserRoutes.Use(a.loginMiddleware)
	authUserRoutes.HandleFunc("", a.handleRequest(handlers.GetUser)).Methods("GET")
	authUserRoutes.HandleFunc("/update", a.handleRequest(handlers.UpdateUser)).Methods("PUT")
	authUserRoutes.HandleFunc("/delete", a.handleRequest(handlers.DeleteUser)).Methods("DELETE")
	authUserRoutes.HandleFunc("/addCourse", a.handleRequest(handlers.AddCourse)).Methods("PUT")
	authUserRoutes.HandleFunc("/removeCourse", a.handleRequest(handlers.RemoveCourse)).Methods("PUT")
	authUserRoutes.HandleFunc("/updateListing", a.handleRequest(handlers.UpdateListing)).Methods("PUT")
	authUserRoutes.HandleFunc("/getUserGroups", a.handleRequest(handlers.GetUserGroups)).Methods("GET")
	authUserRoutes.HandleFunc("/updateInvitation", a.handleRequest(handlers.UpdateInvitation)).Methods("PUT")
	authUserRoutes.HandleFunc("/getBuddiesListings", a.handleRequest(handlers.GetBuddiesAndListings)).Methods("GET")
	authUserRoutes.HandleFunc("/deleteListing", a.handleRequest(handlers.DeleteListing)).Methods("DELETE")


	//Group Routes
	groupRoutes := routes.PathPrefix("/group").Subrouter()
	groupRoutes.HandleFunc("/create", a.handleRequest(handlers.CreateGroup)).Methods("POST")

	groupRoutes.HandleFunc("", a.handleRequest(handlers.GetGroup)).Methods("GET")
	groupRoutes.HandleFunc("/update", a.handleRequest(handlers.UpdateGroup)).Methods("PUT")
	groupRoutes.HandleFunc("/delete", a.handleRequest(handlers.DeleteGroup)).Methods("DELETE")

	//Course Routes
	courseRoutes := routes.PathPrefix("/course").Subrouter()
	courseRoutes.HandleFunc("/addListing", a.handleRequest(handlers.CreateListing)).Methods("POST")

	//Miscellaneous
	routes.HandleFunc("/getAllUsers", a.handleRequest(handlers.GetAllUsers)).Methods("GET")
	
}

func (a *App) loginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid token format"})
			return
		}

		reqToken = strings.TrimSpace(splitToken[1])
		fullToken := models.Token{AccessToken: reqToken}

		if err := fullToken.GetTokenByAccess(a.DB); err != nil {
			respondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
			return
		}

		expirationTime := fullToken.AccessCreateAt.Add(fullToken.AccessExpiresIn)
		diff := expirationTime.Sub(time.Now())
		if diff <= 0 {
			respondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "Token Expired"})
			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

// RequestHandlerFunction for functions to handle
type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
