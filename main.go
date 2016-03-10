package main

import (
    "os"
    "log"
    "net/http"
    "github.com/ant0ine/go-json-rest/rest"

    c "go-gorm/app/controllers" 
    env "go-gorm/app"
)

func main() {

    env.InitEnvs()

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
    
    log.Fatal(http.ListenAndServe(os.Getenv("port"), api.MakeHandler()))
}

