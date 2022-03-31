package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/SlavaUtesinov/store/models"
	"github.com/SlavaUtesinov/store/server"
)

func Test_Product_Should_Be_Created(t *testing.T) {
	port := 3000
	go server.Run(port)
	time.Sleep(3 * time.Second)

	productIn := models.Product{Id: 1, Price: 123.123, Name: "DVD"}
	serialized, _ := json.Marshal(productIn)
	url := fmt.Sprintf("http://localhost:%v/api/products", port)

	if _, err := http.Post(url, "application/json", bytes.NewBuffer(serialized)); err != nil {
		t.Errorf("Error has happened during POST request: %v", err)
	}

	if response, err := http.Get(url); err == nil {
		defer response.Body.Close()

		decoder := json.NewDecoder(response.Body)
		productsOut := make([]models.Product, 0)

		if err := decoder.Decode(&productsOut); err == nil {
			if productIn != productsOut[0] {
				t.Errorf("Products are not equal: %v, %v", productIn, productsOut[0])
			}
		} else {
			t.Errorf("Error has happened during deserialization: %v", err)
		}
	} else {
		t.Errorf("Error has happened during GET request: %v", err)
	}
}
