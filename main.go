package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"

	"github.com/EnesToraman/Go-KPI-Dashboard/model"
	"github.com/EnesToraman/Go-KPI-Dashboard/utils"
)

var db *sql.DB

var jwtKey = []byte("secret_key")

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "airline",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	r := httprouter.New()
	r.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			header := w.Header()
			header.Set("Content-Type", "application/json")
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
			header.Set("Access-Control-Allow-Credentials", "true")
			header.Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		}
		w.WriteHeader(http.StatusNoContent)
	})

	r.POST("/sign-up", SignUp)
	r.POST("/login", Login)
	r.GET("/authUser", AuthUser)
	r.GET("/ticketData", GetTicketData)
	r.GET("/ticketData/:date", GetTicketDataPerDate)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

	// if err := http.ListenAndServe(":"+os.Getenv("PORT"), r); err != nil {
	// 	log.Fatal(err)
	// }

}

func SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	CORS(w, r)
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pwd := user.Password
	email := user.Email
	hashedPassword, err := utils.HashPassword(pwd)
	if err != nil {
		fmt.Printf("CreateUser: %v", err)
	}

	_, err2 := db.Exec("INSERT INTO user (email, hashed_password, role) VALUES (?, ?, 'STAFF')", email, hashedPassword)
	if err2 != nil {
		fmt.Printf("CreateUser: %v", err2)
	}
}
func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	CORS(w, r)
	var credentials model.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	email := credentials.Email
	password := credentials.Password

	row := db.QueryRow("SELECT hashed_password, role FROM user WHERE email = ?", email)
	var hashedPassword string
	var role string
	err = row.Scan(&hashedPassword, &role)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println(role)
	err = utils.CheckPassword(password, hashedPassword)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(time.Minute * 60)

	claims := &model.Claims{
		Email: credentials.Email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	fmt.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	})
}

func AuthUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	CORS(w, r)
	cookie, err := r.Cookie("token")
	if err != nil {
		fmt.Printf("Home: %v", err)
	}
	tokenStr := cookie.Value
	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		fmt.Printf("Home: %v", err)
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Println(claims)
	json.NewEncoder(w).Encode(&claims)
}

func CORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
	if r.Method == "OPTIONS" {
		http.Error(w, "No Content", http.StatusNoContent)
		return
	}
}

func GetTicketData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	CORS(w, r)
	var passenger model.PassengersPerDate
	var passengers []model.PassengersPerDate
	query := `
	SELECT COUNT(t.passenger_id), DATE(f.dep_time) AS date FROM ticket AS t
	INNER JOIN flight AS f ON t.flight_id = f.flight_id
	GROUP BY date
	ORDER BY date
	`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("PassengersPerDate: %v", err)
	}
	for rows.Next() {
		if err := rows.Scan(&passenger.NumberOfPassengers, &passenger.Date); err != nil {
			fmt.Printf("PassengersPerDate: %v", err)
		}
		passengers = append(passengers, passenger)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("PassengersPerDate: %v", err)
	}
	json.NewEncoder(w).Encode(passengers)
}

func GetTicketDataPerDate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	CORS(w, r)
	date := ps.ByName("date")
	fmt.Println(date)
	var passenger model.PassengersPerDate
	query := `
	SELECT COUNT(t.passenger_id), DATE(f.dep_time) FROM ticket AS t
	INNER JOIN flight AS f ON t.flight_id = f.flight_id
	WHERE DATE(f.dep_time) = ?
	GROUP BY DATE(f.dep_time)
	`
	rows, err := db.Query(query, date)
	if err != nil {
		fmt.Printf("PassengersPerDate: %v", err)
	}
	for rows.Next() {
		if err := rows.Scan(&passenger.NumberOfPassengers, &passenger.Date); err != nil {
			fmt.Printf("PassengersPerDate: %v", err)
		}
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("PassengersPerDate: %v", err)
	}
	json.NewEncoder(w).Encode(passenger)
}
