package controllers

import (
    "os"
    "log"
    "net/http"
    "github.com/jinzhu/gorm"
    _"github.com/go-sql-driver/mysql"
    "github.com/ant0ine/go-json-rest/rest"

    m "go-gorm/app/models"
)


type ImplGorm struct {
    DB *gorm.DB
}

func (i *ImplGorm) InitDB() {
    
    var err error
    
    i.DB, err = gorm.Open(os.Getenv("db"), os.Getenv("connection"))
    
    if err != nil {
        log.Fatalf("Got error when connect database, the error is '%v'", err)
    }
    
    i.DB.LogMode(true)
}

func (i *ImplGorm) InitSchema() {
    
    i.DB.AutoMigrate(&m.Reminder{})

}

func (i *ImplGorm) GetAllReminders(w rest.ResponseWriter, r *rest.Request) {
    
    reminders := []m.Reminder{}
    
    i.DB.Find(&reminders)
    
    w.WriteJson(&reminders)
}

func (i *ImplGorm) GetReminder(w rest.ResponseWriter, r *rest.Request) {
    
    id := r.PathParam("id")
    
    reminder := m.Reminder{}
    
    if i.DB.First(&reminder, id).Error != nil {
        rest.NotFound(w, r)
        return
    }
    
    w.WriteJson(&reminder)
}

func (i *ImplGorm) PostReminder(w rest.ResponseWriter, r *rest.Request) {
    
    reminder := m.Reminder{}
    
    if err := r.DecodeJsonPayload(&reminder); err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    
    }
    
    if err := i.DB.Save(&reminder).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteJson(&reminder)
}

func (i *ImplGorm) PutReminder(w rest.ResponseWriter, r *rest.Request) {

    id := r.PathParam("id")
    
    reminder := m.Reminder{}
    
    if i.DB.First(&reminder, id).Error != nil {
        rest.NotFound(w, r)
        return
    }

    updated := m.Reminder{}
    
    if err := r.DecodeJsonPayload(&updated); err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    reminder.Message = updated.Message

    if err := i.DB.Save(&reminder).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteJson(&reminder)
}

func (i *ImplGorm) DeleteReminder(w rest.ResponseWriter, r *rest.Request) {
    
    id := r.PathParam("id")
    reminder := m.Reminder{}
    
    if i.DB.First(&reminder, id).Error != nil {
        rest.NotFound(w, r)
        return
    }
    
    if err := i.DB.Delete(&reminder).Error; err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusOK)
}

