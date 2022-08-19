package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User structure
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Count    interface{}        `bson:"count"`
}

// Initalize mongodb connection
func initmongo(uri string) (context.Context, *mongo.Client, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return ctx, client, err
}

func launchAndServe() {
	// Get the heroku port assigned by heroku itself
	heroku_port := os.Getenv("PORT")
	if heroku_port == "" {
		fmt.Println("No port found from heroku to listen")
		return
	}

	if USE_HEROKU == true {
		fmt.Println("Launching server at port " + heroku_port)
		if err := http.ListenAndServe(":"+heroku_port, nil); err != nil {
			fmt.Println("Failed to launch the server, error: " + err.Error())
			return
		}
	} else {
		fmt.Println("Lauching server at " + DOMAIN + PORT)
		if err := http.ListenAndServe(DOMAIN+PORT, nil); err != nil {
			fmt.Println("Failed to launch the server, error: " + err.Error())
			return
		}
	}
}

func main() {
	// Issue initmongo with the URI
	ctx, client, err := initmongo(MONGODB_URL)
	if err != nil {
		fmt.Println("MongoDB connection error: " + err.Error())
		return
	} else {
		fmt.Println("Connected with MongoDB")
	}

	// Root page, basically the base URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if len(r.UserAgent()) != 0 {
			if !strings.Contains(r.UserAgent(), "github-camo") {
				http.Error(w, "URL access only allowed on GitHub readme", http.StatusForbidden)
				return
			}
		}
		usr_name := r.URL.Query().Get("username")
		first_color := r.URL.Query().Get("first_color")
		second_color := r.URL.Query().Get("second_color")

		var result bson.M

		// Check username length
		if len(usr_name) == 0 {
			http.Error(w, "Please provide a username", http.StatusNotFound)
			return
		}

		// Check if username length is more than 39 charecters
		if usr_name != "" && len(usr_name) > 39 {
			http.Error(w, "User name length must be under 39 charecters", http.StatusNotFound)
			return
		}

		// Index database and get the collection
		collection := client.Database("users").Collection("list")
		e := collection.FindOne(ctx, bson.M{"username": usr_name}).Decode(&result)
		if e != nil {
			if e == mongo.ErrNoDocuments {
				// If username doesn't exist in the database, create one
				fmt.Println("Requested for new account, username: " + usr_name)
				user := &User{
					ID:       primitive.NewObjectID(),
					Username: usr_name,
					Count:    1,
				}
				_, err := collection.InsertOne(ctx, *user)

				// Check if valid params exist
				if len(first_color) != 0 {
					if len(second_color) != 0 {
						svg_image(w, r, "#"+first_color, "#"+second_color, 0)
					} else {
						svg_image(w, r, "#"+first_color, DEFAULT_SECOND_COLOR, 0)
					}
				} else {
					svg_image(w, r, DEFAULT_FIRST_COLOR, DEFAULT_SECOND_COLOR, 0)
				}

				if err != nil {
					fmt.Println("InsertOne returned an error: " + err.Error())
					return
				}
				return
			} else {
				fmt.Println("FindOne returned an error: " + e.Error())
				return
			}
		}

		// Get key, value from result
		for _, value := range result {
			fmtval := fmt.Sprint(value)
			var subresult struct {
				Username string
				Count    uint64
			}
			ff := bson.M{"username": usr_name}
			ex := collection.FindOne(ctx, ff).Decode(&subresult)

			if ex != nil {
				fmt.Println("FineOne returned an error: " + ex.Error())
				return
			}

			// If fmtval variable's value is same as usr_name, render the SVG image
			if fmtval == usr_name {
				filter := bson.M{"username": usr_name}
				collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": bson.M{"count": subresult.Count + 1}})

				// Check if valid params exist
				if len(first_color) != 0 {
					if len(second_color) != 0 {
						svg_image(w, r, "#"+first_color, "#"+second_color, subresult.Count)
					} else {
						svg_image(w, r, "#"+first_color, DEFAULT_SECOND_COLOR, subresult.Count)
					}
				} else {
					svg_image(w, r, DEFAULT_FIRST_COLOR, DEFAULT_SECOND_COLOR, subresult.Count)
				}
				return
			}
		}
	})

	// Launch the server
	launchAndServe()
}
