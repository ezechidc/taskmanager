package controllers

import "github.com/ezechidc/taskmanager/models"

//Models for JSON resources
type (
	//UserResource For Post - /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}
	//LoginResource For Post - /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	//AuthUserResource Response for authorized user Post - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	//TaskResource  For Post/Put - /tasks
	//TaskResource  For Get - /tasks/id
	TaskResource struct {
		Data models.Task `json:"data"`
	}
	//TasksResource For Get - /tasks
	TasksResource struct {
		Data []models.Task `json:"data"`
	}
	//NoteResource For Post/Put - /notes
	NoteResource struct {
		Data NoteModel `json:"data"`
	}
	//NotesResource For Get - /notes
	//NotesResource For /notes/tasks/id
	NotesResource struct {
		Data []models.TaskNote `json:"data"`
	}
	//LoginModel Model for authentication
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	//AuthUserModel Model for authorized user with access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
	//Model for a TaskNote
	NoteModel struct {
		TaskId      string `json:"taskid"`
		Description string `json:"description"`
	}
)
