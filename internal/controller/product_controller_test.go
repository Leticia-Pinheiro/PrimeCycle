package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	controller := &ProductController{}

	t.Run("should return that the product was created", func(t *testing.T) {
		requestBody := ProductRequest{Name: "Camiseta"}
		jsonBody, _ := json.Marshal(requestBody)

		_, rr := createTestContext(http.MethodPost, "/product", jsonBody, controller.CreateProduct)

		assert.Equal(t, http.StatusOK, rr.Code, "Wait for OK Status")
		assert.Equal(t, "application/json; charset=utf-8", rr.Header().Get("Content-Type"), "Wait for a JSON Content-Type")

		var responseBody ProductResponse
		err := json.NewDecoder(rr.Body).Decode(&responseBody)
		assert.NoError(t, err, "Don't have to exist error decodificating the response")
		assert.Equal(t, "product created", responseBody.Message, "Wait for the product created message")
	})
}

func createTestContext(method, path string, bodyBytes []byte, handler gin.HandlerFunc) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	rr := httptest.NewRecorder()

	var req *http.Request
	if bodyBytes != nil {
		req = httptest.NewRequest(method, path, bytes.NewBuffer(bodyBytes))
	} else {
		req = httptest.NewRequest(method, path, nil)
	}

	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		req.Header.Set("Content-Type", "application/json")
	}

	router.Handle(method, path, handler)
	router.ServeHTTP(rr, req)
	c, _ := gin.CreateTestContext(rr)
	return c, rr
}
