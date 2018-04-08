package data

import (
	"log"
	"time"

	"github.com/ezechidc/taskmanager/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TaskRepository struct {
	C *mgo.Collection
}

func (r *TaskRepository) Create(task *models.Task) error {
	obj_id := bson.NewObjectId()
	task.ID = obj_id
	task.CreatedOn = time.Now()
	task.Status = "Created"
	log.Print(task)
	err := r.C.Insert(&task)
	return err
}
func (r *TaskRepository) Update(task *models.Task) error {
	// partial update on MogoDB
	err := r.C.Update(bson.M{"_id": task.ID},
		bson.M{"$set": bson.M{
			"name":        task.Name,
			"description": task.Description,
			"due":         task.Due,
			"status":      task.Status,
			"tags":        task.Tags,
		}})
	return err
}

func (r *TaskRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
func (r *TaskRepository) GetAll() []models.Task {
	var tasks []models.Task
	iter := r.C.Find(nil).Iter()
	result := models.Task{}
	for iter.Next(&result) {
		tasks = append(tasks, result)
	}
	return tasks
}
func (r *TaskRepository) GetById(id string) (task models.Task, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&task)
	return
}

func (r *TaskRepository) GetByUser(user string) []models.Task {
	var tasks []models.Task
	iter := r.C.Find(bson.M{"createdby": user}).Iter()
	result := models.Task{}
	for iter.Next(&result) {
		tasks = append(tasks, result)
	}
	return tasks
}
