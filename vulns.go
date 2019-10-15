package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"strings"
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
	log.Println("Here is exec")
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
