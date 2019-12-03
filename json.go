package gin_json_cr

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

var jsonContentType = []string{"application/json; charset=utf-8"}

func JSON(c *gin.Context, code int, obj interface{}) {
	c.Render(code, JsonRender{Data: obj})
}

type JsonRender struct {
	Data interface{}
}

// Render writes data with custom ContentType.
func (r JsonRender) Render(w http.ResponseWriter) (err error) {
	if err = WriteJSON(w, r.Data); err != nil {
		panic(err)
	}
	return
}

// WriteContentType writes custom ContentType.
func (r JsonRender) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

// WriteJSON marshals the given interface object and writes it with custom ContentType.
func WriteJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	jsonBytes = append(jsonBytes, '\n')
	_, err = w.Write(jsonBytes)
	return err
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
