To-Do List Web Application
Overview
This project is a simple web application for managing a to-do list, built using Go (Golang) for the backend and SQLite for the database. It provides a basic interface to add, view, and manage tasks.

Features
Add new tasks to the to-do list
Mark tasks as done
View a list of tasks categorized as To-Do, Done, and In Progress
Delete tasks
Prerequisites
Before you begin, ensure you have the following installed on your system:

Go (version 1.16 or higher)
SQLite
Project Structure
css
Copy code
.
├── go.mod
├── main.go
├── README.md
├── temp
│   ├── index.html
│   ├── index.js
│   └── main.css
└── database
    └── todo.db
main.go: The main Go application file.
temp/: Directory containing the HTML template, JavaScript, and CSS files.
database/: Directory containing the SQLite database file (todo.db).
Setup Instructions
Step 1: Clone the Repository
Clone the repository to your local machine:

bash
Copy code
git clone https://github.com/yourusername/todo-list-app.git
cd todo-list-app
Step 2: Install Dependencies
Ensure you have the necessary Go dependencies. Initialize the module if it's not already initialized:

bash
Copy code
go mod init todo-list-app
go mod tidy
Step 3: Set Up the Database
Create an SQLite database file in the database directory. You can use the following SQL script to create the required tables:

sql
Copy code
CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    status TEXT NOT NULL
);
You can create the database and run the script using the SQLite CLI:

bash
Copy code
sqlite3 database/todo.db < create_table.sql
Step 4: Running the Application
To run the application, use the following command:

bash
Copy code
go run .  # or `go run main.go`
The application will start a web server listening on http://localhost:8080.

Step 5: Accessing the Application
Open your web browser and navigate to http://localhost:8080 to access the To-Do List web application.

Usage
Adding a Task
Enter the task description in the input field.
Click the "Add Task" button to add the task to your to-do list.
Viewing Tasks
Tasks are categorized into three sections:

To-Do: Tasks that need to be done.
In Progress: Tasks that are currently being worked on.
Done: Completed tasks.
Marking Tasks as Done
Click the "Done" button next to a task to mark it as completed.

Deleting Tasks
Click the "Delete" button next to a task to remove it from the list.

Contributing
Contributions are welcome! Please fork this repository and submit a pull request for any features, bug fixes, or improvements.

License
This project is licensed under the MIT License. See the LICENSE file for details.