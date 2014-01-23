package main

import (
  "fmt"
  "os"
  "log"
  "io"
  "github.com/dchest/uniuri"
  "crypto/sha1"
  "bytes"
)

func main() {

  if len(os.Args) != 5 {
    log.Fatal("You must supply $difficulty, $tree, $parent, $timestamp")
  }

  difficulty := []byte(os.Args[1])
  tree := os.Args[2]
  parent := os.Args[3]
  timestamp := os.Args[4]

  success := false

  for success == false {
    nonce := uniuri.New()

    body := fmt.Sprintf(`tree %s
parent %s
author CTF user <me@example.com> %s +0000
committer CTF user <me@example.com> %s +0000

Give me a Gitcoin

%s
`, tree, parent, timestamp, timestamp, nonce)

    to_hash := fmt.Sprintf("commit %d\x00%s", len(body), body)

    h := sha1.New()
    io.WriteString(h, to_hash)
    result := []byte(fmt.Sprintf("%x", h.Sum(nil)))

    if bytes.Compare(result, difficulty) == -1 {
      // fmt.Printf("%s is < %s", result, difficulty)
      fmt.Printf(body)
      success = true
    } else {
      // fmt.Printf(".")
    //   fmt.Printf("%s is > %s", result, difficulty)
    }

    // success = true
  }

  // fmt.Println("done")
}