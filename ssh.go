package  main

import (
  "golang.org/x/crypto/ssh"
  "bytes"
)
  

func RunCmd( hostname string, cmd string) (err error) {

  user := "root"
  password := "m0vefast"
  port := "22"

  // An SSH client is represented with a ClientConn. Currently only
  // the "password" authentication method is supported.
  //
  // To authenticate with the remote server you must pass at least one
  // implementation of AuthMethod via the Auth field in ClientConfig.
  config := &ssh.ClientConfig{
      User: user,
      Auth: []ssh.AuthMethod{
          ssh.Password(password),
      },
  }

  client, err := ssh.Dial("tcp", hostname + ":" + port , config)
  if err != nil { return err }
  
  // Each ClientConn can support multiple interactive sessions,
  // represented by a Session.
  session, err := client.NewSession()
  if err != nil { return err }
  defer session.Close()
  
  // Once a Session is created, you can execute a single command on
  // the remote side using the Run method.
  var b bytes.Buffer
  session.Stdout = &b
  err = session.Run(cmd)

  return
}
