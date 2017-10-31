package main

import(
	"encoding/json"
    "io/ioutil"
	"time"
    "net"
    "fmt"
    "./common"
    "os"
	"path"
	"encoding/gob"
)

type Post struct {
    Content string
	Timestr string
    Time time.Time
}

type User struct {
    Username string
    Hash []byte
    SessionID string
    Created time.Time
    Follows []string
    Posts []*Post
}

func handleConnection(conn net.Conn) {
    request := common.Request{}
    conn.SetReadDeadline(time.Now().Add(30 * time.Second))
    defer conn.Close()
    dec := gob.NewDecoder(conn)
    dec.Decode(&request)
    fmt.Println(request)
    switch request.Action {
        case common.LOGIN:
            fmt.Println("Handling login action")
        case common.LOGOUT:
            fmt.Println("Handling logout action")
        case common.REGISTER:
			fmt.Println("Handling register action")
            // db_register(db_JSON_to_user())
        case common.DELETE:
            fmt.Println("Handling delete action")
			// db_delete_user()
        case common.FOLLOW:
            fmt.Println("Handling follow action")
        case common.POST:
            fmt.Println("Handling post action")
        case common.FEED:
            fmt.Println("Handling feed action")
        case common.PROFILE:
            fmt.Println("Handling profile action")
        // case default:
            // fmt.Println("Unrecognized action")
    }
}

func db_update_user(username string, sessionid string, follows []string, post Post){
	
}

func db_register(user User) {
	fmt.Println("JSON DATA:")
	newUserBytes := db_user_to_JSON(user)
	fmt.Println(string(newUserBytes)[:])
	writeerr := ioutil.WriteFile(path.Join("db/users", user.Username+".json"), newUserBytes, 0644)
	if writeerr != nil {
		panic(writeerr)
	}
}

func db_delete_user(username string) {
	err := os.Remove(path.Join("db/users", username+".json"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	fmt.Println("User Removed: ", username)
}

func db_check_user_exists(username string) bool {
	if _, err := os.Stat(path.Join("db/users", username+".json")); !os.IsNotExist(err) {
		return true
	}
	return false
}

// converting user struct to JSON string
func db_user_to_JSON(user User) []byte {
	JSON_string, _ := json.MarshalIndent(user, "", "    ")
	return JSON_string
}

// converting JSON string to user struct
func db_JSON_to_user(username string) User {
	dat, err := ioutil.ReadFile(path.Join("db/users", username+".json"))
	if err != nil {
		panic(err.Error())
	}
	
	var user User
	if err := json.Unmarshal(dat, &user); err != nil {
		panic(err)
	}
	return user
}

func mainLoop() {
    ln, err := net.Listen("tcp", ":1338")
    if err != nil {
        fmt.Println("Error listening on port 1338")
        return
    }
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting connection")
        }
        handleConnection(conn)
    }
}

func main() {
    fmt.Println("Hello world!")
    mainLoop()
}
