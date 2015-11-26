package main

import (
    "encoding/json"
    "os"
    "fmt"
    "log"
)

type Configuration struct {
    Operations    []operation
}

func ReadConfig()(configuration Configuration) {
  file, _ := os.Open("conf.json")
  decoder := json.NewDecoder(file)
  configuration = Configuration{}
  err := decoder.Decode(&configuration)
  if err != nil {
    fmt.Println("error:", err)
  }
  //fmt.Println(configuration.Operations)
  log.Println(configuration.Operations)
  return
}
