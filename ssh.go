package  main

import (
    "code.google.com/p/go.crypto/ssh"
)


func RunCmd(host string, cmd string) (status int, msg string) {


  func executeCmd(cmd, hostname string, config *ssh.ClientConfig) string {
    conn, _ := ssh.Dial("tcp", hostname+":22", config)
    session, _ := conn.NewSession()
    defer session.Close()

    var stdoutBuf bytes.Buffer
    session.Stdout = &stdoutBuf
    session.Run(cmd)

    return hostname + ": " + stdoutBuf.String()

}
