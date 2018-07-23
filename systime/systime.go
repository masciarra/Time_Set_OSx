package main

import (
    "os/exec"
    "time"
    "fmt"
    "ntp"
)

func main() {
    currtime, err1 := ntp.Time("216.239.35.0")
    if err1 != nil {
        fmt.Printf("Error1: %s", err1.Error())
    }

    err2 := SetSystemDate(currtime)
    if err2 != nil {
        fmt.Printf("Error2: %s", err2.Error())
    }

}

func SetSystemDate(newTime time.Time) error {
    _, lookErr := exec.LookPath("date")
    if lookErr != nil {
        fmt.Printf("Date binary not found, cannot set system date: %s\n", lookErr.Error())
        return lookErr
    } else {
        dateString := newTime.Format("0102150406")
        fmt.Printf("Setting system date to: %s\n", dateString)
        args := []string{"date", dateString}
        return exec.Command("sudo", args...).Run()
    }
}