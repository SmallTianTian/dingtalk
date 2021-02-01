package test

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

var (
	corpID       string
	suiteKey     string
	suiteSecrect string
	suiteTicket  string
)

func init() {
	p1 := "env.txt"
	p2 := "test/env.txt"
	var configPath string
	if isExist(p1) {
		configPath = p1
	} else if isExist(p2) {
		configPath = p2
	}
	if configPath != "" {
		readFileConfig(configPath)
	}
	envConfig()

	pretty()
}

func pretty() {
	println("corpID", corpID)
	println("suiteKey", suiteKey)
	println("suiteSecrect", suiteSecrect)
	println("suiteTicket", suiteTicket)
}

func envConfig() {
	if c := os.Getenv("CROP_ID"); c != "" {
		corpID = c
	}
	if c := os.Getenv("SUITE_KEY"); c != "" {
		corpID = c
	}
	if c := os.Getenv("SUITE_SECRECT"); c != "" {
		corpID = c
	}
	if c := os.Getenv("SUITE_TICKET"); c != "" {
		corpID = c
	}
}

func readFileConfig(p string) {
	f, err := os.Open(p)
	if err != nil {
		return
	}
	reader := bufio.NewReader(f)
	m := make(map[string]string)
	var line []byte
	for {
		if line, _, err = reader.ReadLine(); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return
		}
		if id := bytes.IndexByte(line, '='); id > 0 {
			m[string(line[:id])] = string(line[id+1:])
		}
	}
	corpID = m["corpID"]
	suiteKey = m["suiteKey"]
	suiteSecrect = m["suiteSecrect"]
	suiteTicket = m["suiteTicket"]
}

func isExist(p string) bool {
	_, err := os.Stat(p)
	return err == nil || os.IsExist(err)
}
