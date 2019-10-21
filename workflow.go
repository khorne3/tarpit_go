package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type resValue struct {
	OrderState string
	Profile    string
}
type outputList struct {
	Name    []string
	Command string
}

var rv resValue

func checkAuth(w http.ResponseWriter, r *http.Request) bool {
	if authenticated == false {
		http.Redirect(w, r, "/login", http.StatusFound)
		return false
	}
	return true
}

// appHandler ... /app
func appHandler(w http.ResponseWriter, r *http.Request) {
	if checkAuth(w, r) {
		if t, err := template.ParseFiles("view/app.html"); err != nil {
			log.Println(err.Error())
		} else {
			rv.OrderState = ""
			rv.Profile = ""
			t.Execute(w, rv)
		}
	}
}

// /processOrder
func processHandler(w http.ResponseWriter, r *http.Request) {
	if checkAuth(w, r) {
		if r.Method == "GET" {
			q := r.URL.Query()
			orderName := q.Get("ordername")
			customerid := q.Get("customerid")
			if len(orderName) == 0 {
				http.Error(w, "missing ordername", http.StatusBadRequest)
				return
			}

			order := Order{OrderName: orderName, CustomerID: customerid}

			oen, err := order.gobEncode()
			if err != nil {
				log.Println(err)
			} else {
				log.Println(oen)
			}

			connection = getConnection()
			sql := fmt.Sprintf("SELECT orderState FROM %s WHERE orderName=\"%s\";", tablename, orderName)
			log.Println(sql)
			if connection == nil {
				log.Println("db connection is null")
			}
			row, err := connection.Query(sql)
			defer row.Close()
			if err != nil {
				log.Println("here for query")
				log.Fatal(err.Error())
			}

			if row.Next() {
				// if order is aleady exist
				var stateNum int64
				if err := row.Scan(&stateNum); err != nil {
					log.Fatal(err)
				}
				stateNum++
				sql = fmt.Sprintf("UPDATE %s SET orderState=%d WHERE orderName=\"%s\"", tablename, stateNum, orderName)

			} else {
				// if order is not exist
				sql = fmt.Sprintf("INSERT INTO %s (custId, ordername, orderState) VALUES (\"%s\", \"%s\", 0)", tablename, customerid, orderName)
			}
			stmt, err := connection.Prepare(sql)
			if err != nil {
				log.Println(err.Error())
			}

			_, err = stmt.Exec()
			if err != nil {
				log.Println(err.Error())
			}
		}
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Printf("Error reading body: %v", err)
				http.Error(w, "can't read body", http.StatusBadRequest)
				return
			}
			var order Order
			err = order.gobDecode(body)
			if err == nil {
				log.Println(err)
			}
			log.Println(order)
		}
		http.Redirect(w, r, "/app", http.StatusFound)
	}
}

// /getOrderStatus
func statusHandler(w http.ResponseWriter, r *http.Request) {
	if checkAuth(w, r) && (r.Method == "GET") {
		q := r.URL.Query()
		orderName := q.Get("ordername")
		if len(orderName) == 0 {
			http.Error(w, "missing ordername", http.StatusBadRequest)
			return
		}
		connection = getConnection()

		sql := fmt.Sprintf("SELECT orderState from %s where orderName=\"%s\"", tablename, orderName)
		row, err := connection.Query(sql)
		defer row.Close()
		if err != nil {
			log.Fatal(err.Error())
		}

		if row.Next() {
			var stateNum int64
			if err := row.Scan(&stateNum); err != nil {
				log.Fatal(err)
			}
			var orderState string
			switch stateNum {
			case 0:
				orderState = "Preparing"
			case 1:
				orderState = "Delivering"
			case 2:
				orderState = "Out for delivery in the post office"
			case 3:
				orderState = "Delivered"
			default:
				orderState = "Case error"

			}
			log.Printf("order status is : %s \n", orderState)
			rv.OrderState = orderState
			t, _ := template.ParseFiles("view/app.html")
			t.Execute(w, rv)
		}
	}
}

// /init
func initHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/init.html")
	t.Execute(w, nil)
}

// /profile
func proHandler(w http.ResponseWriter, r *http.Request) {
	if checkAuth(w, r) && (r.Method == "GET") {
		t, _ := template.ParseFiles("view/profile.html")
		t.Execute(w, rv)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

// /setprofile
func setproHandler(w http.ResponseWriter, r *http.Request) {
	if checkAuth(w, r) && (r.Method == "GET") {
		q := r.URL.Query()
		imagename := q.Get("image")
		rv.Profile = "image/" + imagename
		t, _ := template.ParseFiles("view/profile.html")
		t.Execute(w, rv)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

// /listdemo
func listdemoHandler(w http.ResponseWriter, r *http.Request) {
	// log.Println("listdemo")
	// files, err := ioutil.ReadDir("./demo")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	var ol outputList
	ol.Name = make([]string, 0, 1)
	ol.Command = "ls -a"
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
