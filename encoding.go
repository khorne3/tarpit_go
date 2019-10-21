package main

import (
	"bytes"
	"encoding/gob"
)

// Order ... ordername and customer ID
type Order struct {
	OrderName  string
	CustomerID string
}

func (d *Order) gobEncode() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(d.OrderName)
	if err != nil {
		return nil, err
	}
	err = encoder.Encode(d.CustomerID)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (d *Order) gobDecode(buf []byte) error {
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)
	err := decoder.Decode(&d.OrderName)
	if err != nil {
		return err
	}
	return decoder.Decode(&d.CustomerID)
}
