package app

import (
    "os"
    "fmt"
    "encoding/json"
    "io/ioutil"
)

type env struct{
    Key string
    Value string
}

type envs []env

func InitEnvs(){

    file, e := ioutil.ReadFile("./config.json")

    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

    var jsonEnvs envs

    json.Unmarshal(file, &jsonEnvs)

    for _, e := range jsonEnvs {

        os.Setenv(e.Key, e.Value)

    }

}