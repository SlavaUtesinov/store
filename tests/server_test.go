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

	if response, err := http.Post(fmt.Sprintf("http://localhost:%v/api/products", port), "application/json", bytes.NewBuffer(serialized)); err == nil {
		defer response.Body.Close()

		decoder := json.NewDecoder(response.Body)
		var productOut models.Product

		if err := decoder.Decode(&productOut); err == nil {
			if productIn != productOut {
				t.Errorf("Products are not equal: %v, %v", productIn, productOut)
			}
		} else {
			t.Errorf("Error has happened during deserialization: %v", err)
		}
	} else {
		t.Errorf("Error has happened during HTTP request: %v", err)
	}
}
