package main

import http

const PollURL = "https://raw.githubusercontent.com/shvar/remedy/master/payload_example.json"

res, err := http.Get(PollURL)
if err != nil { panic(err) }
 
RullerLoad, err := ioutil.ReadAll(res.Body)
res.Body.Close()
if err != nil { panic(err) }

log.Printf("JSON from ruller (%x)\n", RullerLoad)

