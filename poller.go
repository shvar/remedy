package main


import( 
  "net/http"
  "io/ioutil"
  "log"
  "time"
  "encoding/json"
)

const PollURL = "https://raw.githubusercontent.com/shvar/remedy/master/payload_example.json"

type payload struct {
  Command string `json:"command"`
  Server string `json:"server"`
  Service string `json:"service"`
}

type operation struct {
  ActionName string `json:"action"`
  Payload payload `json:"payload"`
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
      WhatToDo.ActionName, 
      WhatToDo.Payload.Service,
      WhatToDo.Payload.Server,
      WhatToDo.Payload.Command)
 
    err = RunCmd(WhatToDo.Payload.Server, "service " + WhatToDo.Payload.Service + " " + WhatToDo.Payload.Command)
    if err != nil { panic(err) }
    log.Printf("Executed!\n")
    time.Sleep(time.Minute)
  }
}
