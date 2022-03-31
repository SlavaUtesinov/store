package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"sync"

	"github.com/SlavaUtesinov/store/models"
)

var products = make([]models.Product, 0)

var mutex sync.Mutex

type routeHandler func(writer http.ResponseWriter, request *http.Request)

type httpHandler struct {
	CreateProduct routeHandler `url:"/api/products" method:"POST"`
	GetProducts   routeHandler `url:"/api/products" method:"GET"`
}

func (h httpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	t := reflect.TypeOf(h)
	v := reflect.ValueOf(h)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		url := f.Tag.Get("url")
		method := f.Tag.Get("method")

		if url == request.URL.Path && method == request.Method {
			handler := v.FieldByName(f.Name)
			args := []reflect.Value{reflect.ValueOf(writer), reflect.ValueOf(request)}
			handler.Call(args)
			return
		}
	}

	writer.WriteHeader(404)
}

func CreateHandler() httpHandler {
	handler := httpHandler{
		CreateProduct: func(writer http.ResponseWriter, request *http.Request) {
			decoder := json.NewDecoder(request.Body)
			defer request.Body.Close()

			var product models.Product
			if err := decoder.Decode(&product); err == nil {
				mutex.Lock()
				product.Id = len(products) + 1
				products = append(products, product)
				mutex.Unlock()

				if serialized, err := json.Marshal(product); err == nil {
					writer.Header().Add("Content-Type", "application/json")
					writer.Write(serialized)
				}
			} else {
				writer.WriteHeader(400)
				writer.Write([]byte(fmt.Sprintf("Error has hapened during request's body deserialization: %v", err)))
			}
		},
		GetProducts: func(writer http.ResponseWriter, request *http.Request) {
			if serialized, err := json.Marshal(products); err == nil {
				writer.Header().Add("Content-Type", "application/json")
				writer.Write(serialized)
			}
		},
	}

	return handler
}
