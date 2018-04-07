# Taskmanager

This a simple RESTful API application that can be used to manage tasks and give updated and notes about each task. The application is written in go and uses MongoDB database for storage. This TaskManager RESTful API will be a JSON-based API in
which the JSON data format will be used for sending and receiving messages among client applications and
the RESTful API server. This is base one  the book Web Development with Go.

The RESTful API application has been divided into the following packages:
 - common: Implements some utility functions and provides initialization logic for the
application
 - controllers: Implements the application’s application handlers
 - data: Implements the persistence logic with the MongoDB database
 - models: Describes the data model of the application
 - routers: Implements the HTTP request routers for the RESTful API

| URI | HTTP Verb | Functionality |
| :--- |:---------| :------------- |
|/users/register| Post | Creates a new user.|
|/users/login| | Post | User logs in to the system, which returns a JWT if the login is successful. |
|/tasks | Post | Creates a new task. |
| /tasks/{id} | Put | Updates an existing task. |
| /tasks | Get | Gets all tasks.| 
| /tasks/{id} | Get | Gets a single task for a given ID. The value of the ID comes from the route parameter.| 
| /tasks/users/{id} | Get | Gets all tasks associated with a user. The value of the user ID comes from the route parameter.| 
| /tasks/{id} | Delete | Deletes an existing task for a given ID. The value of the ID comes from the route parameter.| 
| /notes | Post | Creates a new note against an existing task.| 
| /notes/{id} | Put | Updates an existing task note.| 
| /notes | Get | Gets all task notes.| 
| /notes/{id} | Get | Gets a single note for a given ID. The value of the ID comes from the route parameter.| 
| /notes/tasks/{id} | Get | Gets all task notes for a given task ID. The value of the ID comes from the route parameter.| 
| /notes/{id} | Delete | Deletes an existing note for a given ID. The value of the ID comes from the route parameter.| 
