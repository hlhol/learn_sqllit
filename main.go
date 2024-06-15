package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Errors struct {
	Err    int
	ErrStr string
}

type Data struct {
	Todo    []string
	Done    []string
	Indoing []string
}

type PageData struct {
	Data  Data
	Error Errors
}

func main() {
	fmt.Println("Starting the web app...")

	http.HandleFunc("/", mainPage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("temp"))))

	// Serve the web app on localhost
	fmt.Println("Thanks, now the server is listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainPage(res http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseFiles("temp/index.html")
	if err != nil {
		errorData := PageData{
			Error: Errors{
				Err:    http.StatusInternalServerError,
				ErrStr: fmt.Sprintf("Error parsing template: %v", err),
			},
		}
		log.Printf("Error parsing template: %v", err)
		temp.Execute(res, errorData)
		return
	}

	if req.URL.Path != "/" {
		errorData := PageData{
			Error: Errors{
				Err:    http.StatusNotFound,
				ErrStr: "404 Not Found",
			},
		}
		log.Printf("404 Not Found: %s", req.URL.Path)
		temp.Execute(res, errorData)
		return
	}

	// Prepare empty data for the home page
	pageData := PageData{
		Data: Data{},
	}

	if err := temp.Execute(res, pageData); err != nil {
		errorData := PageData{
			Error: Errors{
				Err:    http.StatusInternalServerError,
				ErrStr: fmt.Sprintf("Error executing template: %v", err),
			},
		}
		log.Printf("Error executing template: %v", err)
		temp.Execute(res, errorData)
	}
}
