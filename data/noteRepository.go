package data

import (
	"log"
	"time"

	"github.com/ezechidc/taskmanager/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//NoteRepository struct to
type NoteRepository struct {
	C *mgo.Collection
}

//Create a note
func (r *NoteRepository) Create(note *models.TaskNote) error {
	objID := bson.NewObjectId()
	note.ID = objID
	note.CreatedOn = time.Now()
	err := r.C.Insert(&note)
	log.Print(err)
	return err
}

//Update note with id provided
func (r *NoteRepository) Update(note *models.TaskNote) error {
	// partial update on MogoDB
	err := r.C.Update(bson.M{"_id": note.ID},
		bson.M{"$set": bson.M{
			"description": note.Description,
		}})
	return err
}

//Delete note using the id
func (r *NoteRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//GetByTask note using the associated taskid
func (r *NoteRepository) GetByTask(id string) []models.TaskNote {
	var notes []models.TaskNote
	taskid := bson.ObjectIdHex(id)
	iter := r.C.Find(bson.M{"taskid": taskid}).Iter()
	result := models.TaskNote{}
	for iter.Next(&result) {
		notes = append(notes, result)
	}
	return notes
}

//GetAll gets all notes
func (r *NoteRepository) GetAll() []models.TaskNote {
	var notes []models.TaskNote
	iter := r.C.Find(nil).Iter()
	result := models.TaskNote{}
	for iter.Next(&result) {
		notes = append(notes, result)
	}
	return notes
}

//GetByID get a note with a given id
func (r *NoteRepository) GetByID(id string) (note models.TaskNote, err error) {
	log.Print(id)
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&note)
	log.Print(bson.ObjectIdHex(id))
	return
}
