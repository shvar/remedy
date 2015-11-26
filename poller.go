package main


import( 
  "net/http"
  "io/ioutil"
  "log"
)

const PollURL = "https://raw.githubusercontent.com/shvar/remedy/master/payload_example.json"


func main() {
  res, err := http.Get(PollURL)
  if err != nil { panic(err) }
   
  RullerLoad, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil { panic(err) }

  log.Printf("JSON from ruller (%x)\n", RullerLoad)
}