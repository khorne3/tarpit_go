package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type resValue struct {
	OrderState string
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
			t.Execute(w, rv)
		}
	}
}

// /processOrder
func processHandler(w http.ResponseWriter, r *http.Request) {
	if checkAuth(w, r) {
		if r.Method == "POST" {
			orderName := r.FormValue("ordername")
			customerid := r.FormValue("customerid")
			if len(orderName) == 0 {
				http.Error(w, "missing ordername", http.StatusBadRequest)
				return
			}
			connection = getConnection()
			sql := fmt.Sprintf("SELECT orderState from %s where orderName=\"%s\";", tablename, orderName)
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
