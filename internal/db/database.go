package database

import (
	"encoding/json"
	"os"
	"time"
)

// create a new Client struct that will hold all database interactions our API will user
type Client struct {
	DbPath string
}

func(c Client) NewClient(path string) Client {
	return Client{
		DbPath: path,
	}
}

func (c Client) createDB() error {

	byteData, marshalErr := json.Marshal("")

	if marshalErr != nil {
		return marshalErr
	}

	writeErr := os.WriteFile("./db.json", byteData, 0666)

	if writeErr != nil {
		return writeErr
	}

	return nil
}

func (c Client) EnsureDB() error {

	_, err:= os.ReadFile("./db.json")

	if err != nil {
		c.createDB()
		return err
	}

	return nil
}

// Below is out database schema where we're using structs with JSON tags
/*
When you encode and decode a struct to JSON, the key of the JSON object will be the name of 
the struct field unless you give the field an explicit JSON tag.
*/

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

// User -
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

// Post -
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}
