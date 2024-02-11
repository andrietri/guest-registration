package httpresponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPResponder struct {
	Status       int
	ErrorMessage string
}

func (hr *HTTPResponder) ErrorMessageResonse() string {
	return hr.ErrorMessage
}

func (hr *HTTPResponder) StatusResponse() int {
	return hr.Status
}

// With Data and custome Response Status
// Reusable
func WithDataCustomResponseStatus(c *gin.Context, status, data, errmsg interface{}) {
	c.JSON(status.(int), gin.H{
		"status":        status,
		"data":          data,
		"error_message": errmsg,
	})
}

// WithData :nodoc
func WithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":        http.StatusOK,
		"data":          data,
		"error_message": nil,
	})
}

// DataOnly :nodoc
func DataOnly(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// NotFound :nodoc
func NotFound(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":        http.StatusNotFound,
		"data":          nil,
		"error_message": "No data found",
	})
}

// BadRequest :nodoc
func BadRequest(c *gin.Context, err interface{}) {
	if err == nil {
		err = "The request is not valid or not found"
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        http.StatusBadRequest,
		"data":          nil,
		"error_message": err,
	})
}

// Unauthorized :nodoc
func Unauthorized(c *gin.Context, err interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":        http.StatusUnauthorized,
		"data":          nil,
		"error_message": err,
	})
}

// InternalServerError :nodoc
func InternalServerError(c *gin.Context, err interface{}) {
	if err == nil {
		err = "An unexpected error has occured"
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        http.StatusInternalServerError,
		"data":          nil,
		"error_message": err,
	})
}

// Conflict :nodoc
func Conflict(c *gin.Context, err interface{}) {
	if err == nil {
		err = "Data already exist"
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        http.StatusConflict,
		"data":          nil,
		"error_message": err,
	})
}

// Forbidden :nodoc
func Forbidden(c *gin.Context, err interface{}) {
	if err == nil {
		err = "Forbidden"
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        http.StatusForbidden,
		"data":          nil,
		"error_message": err,
	})
}

// Error :nodoc:
func Error(c *gin.Context, hr *HTTPResponder) {
	c.JSON(http.StatusOK, gin.H{
		"status":        hr.StatusResponse(),
		"data":          nil,
		"error_message": hr.ErrorMessageResonse(),
	})
}
