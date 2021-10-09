package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Mithilesh Ghadge 19BCE7309 mithilesh.19bce7309@vitap.ac.in

var p int
var e error
var id string
var name string
var email string
var pass string
var pid string
var caption string
var imgurl string
var posturl string
var geturl string
var posttime time.Time = time.Now()
var userC mongo.Collection
var postC mongo.Collection

type User struct {
	Id       string
	Name     string
	Email    string
	Password string
}

type Post struct {
	Id                string
	Caption           string
	ImageURL          string
	Postted_Timestamp time.Time
}

func menu() {
	fmt.Println()
	fmt.Println("================")
	fmt.Println("Instagram API!")
	fmt.Println("time : ", posttime)
	fmt.Println("1. Create an User")
	fmt.Println("2. Get a User")
	fmt.Println("3. Create a Post")
	fmt.Println("4. Get a Post")
	fmt.Println("5. List all Posts")
	fmt.Println("6. Exit")
	fmt.Println("================")
	fmt.Println("Enter the Serial Number of Task you want to perform: ")
	fmt.Scanf("%d", &p)
	fmt.Println()
}

func connectdb() {
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://cherry:<cherry123>@cluster0.q2gpl.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	client, e := mongo.Connect(context.TODO(), clientOptions)
	if e != nil {
		log.Fatal(e)
	}
	e = client.Ping(context.TODO(), nil)
	if e != nil {
		log.Fatal(e)
	}
	userc := client.Database("myFirstDatabase").Collection("Users")
	postc := client.Database("myFirstDatabase").Collection("Posts")
	userC = *userc
	postC = *postc
}

func createUser() {
	fmt.Scanf("%d", &posturl)

	fmt.Println("Give User Id")
	fmt.Scanf("%d", &id)
	fmt.Println("Give User Name")
	fmt.Scanf("%d", &name)
	fmt.Println("Give User Email")
	fmt.Scanf("%d", &email)
	fmt.Println("Give Password")
	fmt.Scanf("%d", &pass)

	var jsonData = []byte(`{
		"id": id,
		"name": name,
		"email": email,
		"password" : pass,
	}`)
	request, error := http.NewRequest("POST", posturl, bytes.NewBuffer(jsonData))
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	userf := User{id, name, email, pass}
	_, e = userC.InsertOne(context.TODO(), body)
	if e != nil {
		_, e = userC.InsertOne(context.TODO(), userf)
		if e != nil {
			log.Fatal(e)
		}
	}
}

func createPost() {
	fmt.Scanf("%d", posturl)

	fmt.Println("Enter Post Id")
	fmt.Scanf("%d", &pid)
	fmt.Println("Enter Caption")
	fmt.Scanf("%d", &caption)
	fmt.Println("Enter Image URL")
	fmt.Scanf("%d", &imgurl)
	posttime = time.Now()

	var jsonData = []byte(`{
		"id": pid,
		"caption": caption,
		"image": imgurl,
		"posted timestamp" : posttime,
	}`)
	request, error := http.NewRequest("POST", posturl, bytes.NewBuffer(jsonData))
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	postf := Post{pid, caption, imgurl, posttime}
	_, e = postC.InsertOne(context.TODO(), body)
	if e != nil {
		_, e = postC.InsertOne(context.TODO(), postf)
		if e != nil {
			log.Fatal(e)
		}
	}

	fmt.Println("Post created!")
}

func getUser() {
	fmt.Println("Enter User Id")
	fmt.Scanf("%d", &id)

	filter := bson.D{{"id", id}}
	var res User
	e = userC.FindOne(context.TODO(), filter).Decode(&res)
	fmt.Println(res)

	fmt.Println("User found!")
}

func getPost() {
	fmt.Println("Enter Post Id")
	fmt.Scanf("%d", &pid)

	filter := bson.D{{"pid", pid}}
	var res Post
	e = postC.FindOne(context.TODO(), filter).Decode(&res)
	fmt.Println(res)

	fmt.Println("Post found!")
}

func getAllPosts() {
	fmt.Println("Enter User/Post Id")
	fmt.Scanf("%d", &pid)

	filter := bson.D{{"pid", pid}}
	var res User
	e = postC.FindOne(context.TODO(), filter).Decode(&res)
	fmt.Println(res)

	fmt.Println("Posts found!")
}

func main() {

	connectdb()

	var APIrunning bool = true
	for APIrunning == true {

		menu()

		switch {
		case p == 1:
			createUser()

		case p == 2:
			getUser()

		case p == 3:
			createPost()

		case p == 4:
			getPost()

		case p == 5:
			getAllPosts()

		case p == 6:
			fmt.Println("Exited!")
			APIrunning = false

		default:
			fmt.Println("Invalid Input, Exited!")
			APIrunning = false
		}
	}

}

//Mithilesh Ghadge 19BCE7309 mithilesh.19bce7309@vitap.ac.in
