package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eranamarante/go-react-expense-tracker/server/database"
	"github.com/eranamarante/go-react-expense-tracker/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var usersCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	var result bson.M
	filter := bson.M{"email": user.Email}
	if err := usersCollection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		log.Fatal(err)
	}

	insertResult, err := usersCollection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
	json.NewEncoder(w).Encode(user)

	// err := json.NewDecoder(r.Body).Decode(&user)
	// if err != nil {
	// 	var err Error
	// 	err = SetError(err, "Error in reading payload.")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(err)
	// 	return
	// }

	// var dbuser User
	// connection.Where("email = ?", user.Email).First(&dbuser)

	// //check email is alredy registered or not
	// if dbuser.Email != "" {
	// 	var err Error
	// 	err = SetError(err, "Email already in use")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(err)
	// 	return
	// }

	// user.Password, err = GeneratehashPassword(user.Password)
	// if err != nil {
	// 	log.Fatalln("Error in password hashing.")
	// }

	// //insert user details in database
	// connection.Create(&user)
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(user)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	// connection := GetDatabase()
	// defer CloseDatabase(connection)

	// var authDetails Authentication

	// err := json.NewDecoder(r.Body).Decode(&authDetails)
	// if err != nil {
	// 	var err Error
	// 	err = SetError(err, "Error in reading payload.")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(err)
	// 	return
	// }

	// var authUser User
	// connection.Where("email = 	?", authDetails.Email).First(&authUser)

	// if authUser.Email == "" {
	// 	var err Error
	// 	err = SetError(err, "Username or Password is incorrect")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(err)
	// 	return
	// }

	// check := CheckPasswordHash(authDetails.Password, authUser.Password)

	// if !check {
	// 	var err Error
	// 	err = SetError(err, "Username or Password is incorrect")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(err)
	// 	return
	// }

	// validToken, err := GenerateJWT(authUser.Email, authUser.Role)
	// if err != nil {
	// 	var err Error
	// 	err = SetError(err, "Failed to generate token")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(err)
	// 	return
	// }

	// var token Token
	// token.Email = authUser.Email
	// token.Role = authUser.Role
	// token.TokenString = validToken
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(token)
}
