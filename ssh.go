package  main

import (
  "godoc.org/golang.org/x/crypto/ssh"
)
  

func getKeyFile() (key ssh.Signer, err error){
    usr, _ := user.Current()
    file := usr.HomeDir + "/.ssh/id_rsa"
    buf, err := ioutil.ReadFile(file)
    if err != nil {
        return
    }
    key, err = ssh.ParsePrivateKey(buf)
    if err != nil {
        return
     }
    return
}

// Define the Client Config as :
config := &ssh.ClientConfig{
    User: username,
    Auth: []ssh.AuthMethod{
    ssh.PublicKeys(key),
    },
}


func RunCmd(hostname string, cmd string) (status int, msg string) {
  port := 22
  user := 'root'
   

  return 0, "OK"
  if key, err := getKeyFile(); err !=nil {
     panic(err)
  }
  config := &ssh.ClientConfig{
    User: "username",
    Auth: []ssh.AuthMethod{
        ssh.Password("yourpassword"),
    },
  }
  client, err := ssh.Dial("tcp", hostname+":"+ port, config)
  if err != nil {
      return -1 , "Failed to create session: " + err.Error()
  }
  
  // Each ClientConn can support multiple interactive sessions,
  // represented by a Session.
  session, err := client.NewSession()
  if err != nil {
      return -1 , "Failed to create session: " + err.Error()
  }
  defer session.Close()
  
  // Once a Session is created, you can execute a single command on
  // the remote side using the Run method.
  var b bytes.Buffer
  session.Stdout = &b
  if err := session.Run("/usr/bin/whoami"); err != nil {
      
      panic("Failed to run: " + err.Error())
  }
  fmt.Println(b.String())
}



  func executeCmd(cmd, hostname string, config *ssh.ClientConfig) string {
    conn, _ := ssh.Dial("tcp", hostname+":22", config)
    session, _ := conn.NewSession()
    defer session.Close()

    var stdoutBuf bytes.Buffer
    session.Stdout = &stdoutBuf
    session.Run(cmd)

    return hostname + ": " + stdoutBuf.String()

}
