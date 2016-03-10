package main

import (
    "os"
    "net/http"
    l4g "github.com/shengkehua/xlog4go"
    "github.com/ant0ine/go-json-rest/rest"

    c "go-gorm/app/controllers" 
    env "go-gorm/app"
)

func main() {

    env.InitEnvs()

    if err := l4g.SetupLogWithConf(os.Getenv("file_log")); err != nil {
        panic(err)
    }
    defer l4g.Close()   

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
        l4g.Error(err.Error())
    }

    l4g.Trace("Init service go-gorm in port %s", os.Getenv("port"))    

    api.SetApp(router)

    l4g.Error(http.ListenAndServe(os.Getenv("port"), api.MakeHandler()).Error())

}

