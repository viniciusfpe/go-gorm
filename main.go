package main

import (
    "log"
    "net/http"
    "github.com/ant0ine/go-json-rest/rest"

    c "go-gorm/app/controllers" 
)

func main() {

    i := c.ImplGorm{}
    i.InitDB()
    i.InitSchema()

    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    
    router, err := rest.MakeRouter(
        rest.Get("/reminders", i.GetAllReminders),
        rest.Post("/reminders", i.PostReminder),
        rest.Get("/reminders/:id", i.GetReminder),
        rest.Put("/reminders/:id", i.PutReminder),
        rest.Delete("/reminders/:id", i.DeleteReminder),
    )
    
    if err != nil {
        log.Fatal(err)
    }
    
    api.SetApp(router)
    
    log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

