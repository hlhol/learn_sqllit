package main

import (
    "database/sql"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os"

    _ "github.com/mattn/go-sqlite3"
)

type Errors struct {
    Err    int
    ErrStr string
}

type User struct {
    Username string
    Password string
}

type Data struct {
    Users []User
}

type PageData struct {
    Data   Data
    Error  Errors
    Logged bool
    SignUp bool
}

var db *sql.DB

func main() {
    var err error
    db, err = sql.Open("sqlite3", "./learn_sqllit.db")
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    defer db.Close()

    if len(os.Args) > 1 {
        log.Println("Use without any argument \nlike go run . or go run main.go")
        return
    }

    fmt.Println("Starting the web app...")

    http.HandleFunc("/", mainPage)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/signup", signupHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Serve the web app on localhost
    fmt.Println("Thanks, now the server is listening at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainPage(res http.ResponseWriter, req *http.Request) {
    if req.URL.Path != "/" {
        errorData := PageData{
            Error: Errors{
                Err:    http.StatusNotFound,
                ErrStr: "404 Not Found",
            },
        }
        log.Printf("404 Not Found: %s", req.URL.Path)
        temp, _ := template.ParseFiles("temp/error.html")
        temp.Execute(res, errorData)
        return
    }

    temp, err := template.ParseFiles("static/index.html")
    if err != nil {
        errorData := PageData{
            Error: Errors{
                Err:    http.StatusInternalServerError,
                ErrStr: fmt.Sprintf("Error parsing template: %v", err),
            },
        }
        log.Printf("Error parsing template: %v", err)
        temp, _ := template.ParseFiles("temp/error.html")
        temp.Execute(res, errorData)
        return
    }

    pageData := PageData{
        Data: Data{
            Users: []User{},
        },
    }

    // Example: Check if user is logged in (you need to implement this logic)
    // For now, assume no user is logged in
    pageData.Logged = false

    // Example: Check if it's a sign-up page request
    // For now, assume it's not a sign-up request
    pageData.SignUp = false

    if err := temp.Execute(res, pageData); err != nil {
        errorData := PageData{
            Error: Errors{
                Err:    http.StatusInternalServerError,
                ErrStr: fmt.Sprintf("Error executing template: %v", err),
            },
        }
        log.Printf("Error executing template: %v", err)
        temp, _ := template.ParseFiles("temp/error.html")
        temp.Execute(res, errorData)
    }
}

func loginHandler(res http.ResponseWriter, req *http.Request) {
    if req.Method != http.MethodPost {
        http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    username := req.FormValue("username")
    password := req.FormValue("password")

    user, err := authenticateUser(username, password)
    if err != nil {
        http.Error(res, err.Error(), http.StatusUnauthorized)
        return
    }

    fmt.Fprintf(res, "Welcome, %s!", user.Username)
}

func signupHandler(res http.ResponseWriter, req *http.Request) {
    if req.Method != http.MethodPost {
        http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse form data
    err := req.ParseForm()
    if err != nil {
        http.Error(res, "Failed to parse form data", http.StatusInternalServerError)
        return
    }

    // Extract username and password from form data
    username := req.Form.Get("username")
    password := req.Form.Get("password")

    // Validate username and password (add more validation as needed)
    if username == "" || password == "" {
        http.Error(res, "Username and password are required", http.StatusBadRequest)
        return
    }

    // Check if the username already exists in the database
    var existingUsername string
    err = db.QueryRow("SELECT username FROM users WHERE username = ?", username).Scan(&existingUsername)
    switch {
    case err == sql.ErrNoRows:
        // Username does not exist, proceed to create the new user
        _, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
        if err != nil {
            http.Error(res, "Failed to create user", http.StatusInternalServerError)
            log.Printf("Failed to create user: %v", err)
            return
        }
        // Redirect or respond with a success message
        http.Redirect(res, req, "/login", http.StatusSeeOther)
    case err != nil:
        // Database error
        http.Error(res, "Database error", http.StatusInternalServerError)
        log.Printf("Database error: %v", err)
        return
    default:
        // Username already exists
        http.Error(res, "Username already exists", http.StatusConflict)
        return
    }
}


func authenticateUser(username, password string) (User, error) {
    var user User
    err := db.QueryRow("SELECT username, password FROM users WHERE username = ?", username).Scan(&user.Username, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return User{}, fmt.Errorf("invalid username or password")
        }
        return User{}, err
    }

    if user.Password != password {
        return User{}, fmt.Errorf("invalid username or password")
    }

    return user, nil
}

func createUser(username, password string) error {
    _, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
    if err != nil {
        return fmt.Errorf("username already taken")
    }
    return nil
}
