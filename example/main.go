package main

import (
  log "github.com/Sirupsen/logrus"
  scribelog "github.com/sagar8192/logrus-scribe-hook"
)

func main() {
  // Disable stdout and stderr
  hook, err := scribelog.NewScribeHook("tmp_sagarp_logrus_testing", true, "localhost:1463")

  if err == nil {
    log.AddHook(hook)
  }

  log.Print("It works!!!")
}
