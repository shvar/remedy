package main


import( 
  "net/http"
  "io/ioutil"
  "log"
  "time"

  "encoding/json"
)

const PollURL = "https://raw.githubusercontent.com/shvar/remedy/master/payload_example.json"

type action struct {
  Command string `json:"command"`
  Server string `json:"server"`
  Service string `json:"service"`
}

type operation struct {
  Command string `json:"command"`
  Action action `json:"payload"`
}

func main() {
  for {
    res, err := http.Get(PollURL)
    if err != nil { panic(err) }
     
    RullerLoad, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil { panic(err) }

    log.Printf("JSON from ruller (%s)\n", RullerLoad)
    
    WhatToDo := operation{}
    err = json.Unmarshal(RullerLoad, &WhatToDo);
    if err != nil { panic(err) }

    log.Printf("I have to execute %s for service '%s' on server '%s' with service command '%s'\n", 
      WhatToDo.Command, 
      WhatToDo.Action.Service,
      WhatToDo.Action.Server,
      WhatToDo.Action.Command)

    time.Sleep(time.Minute)
  }
}