package gin_json_cr

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndentedJSON(c *gin.Context, code int, obj interface{}) {
	c.Render(code, IndentedJsonRender{Data: obj})
}

// IndentedJSON contains the given interface object.
type IndentedJsonRender struct {
	Data interface{}
}

// Render writes data with custom ContentType.
func (r IndentedJsonRender) Render(w http.ResponseWriter) (err error) {
	r.WriteContentType(w)
	jsonBytes, err := json.MarshalIndent(r.Data, "", "    ")
	if err != nil {
		return err
	}
	jsonBytes = append(jsonBytes, '\n')
	_, err = w.Write(jsonBytes)
	return err
}

// WriteContentType writes custom ContentType.
func (r IndentedJsonRender) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}
