package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_AllUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	r.GET("/allusers", AllUser)

	req, err := http.NewRequest(http.MethodGet, "/allusers", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

}

func Test_Get_UsersById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	r.GET("/user/:id", GetUserById)

	req, err := http.NewRequest(http.MethodGet, "/user/1", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

}

// func Test_Update_UsersById(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	r := gin.Default()

// 	r.PUT("/user/:id", UpdateUserById)

// 	req, err := http.NewRequest(http.MethodPut, "/user/1", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v\n", err)
// 	}

// 	w := httptest.NewRecorder()

// 	r.ServeHTTP(w, req)

// 	if w.Code != http.StatusOK {
// 		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
// 	}

// }

// func Test_Delete_UsersById(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	r := gin.Default()

// 	r.DELETE("/user/:id", DeleteUserById)

// 	req, err := http.NewRequest(http.MethodDelete, "/user/3", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v\n", err)
// 	}

// 	w := httptest.NewRecorder()

// 	r.ServeHTTP(w, req)

// 	if w.Code != http.StatusOK {
// 		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
// 	}

// }
