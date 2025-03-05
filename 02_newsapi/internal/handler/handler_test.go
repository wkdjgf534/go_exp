package handler_test

import (
	"net/http"
	"net/http/httptest"
	"newsapi/internal/handler"
	"testing"
)

func Test_PostNews(t *testing.T){
	testCase := []struct{
		name string
		expectedStatus int
	}{
		{
			name: "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.PostNews()(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected : %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_GetAllNews(t *testing.T){
	testCase := []struct{
		name string
		expectedStatus int
	}{
		{
			name: "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.GetAllNews()(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected : %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_GetNewsByID(t *testing.T){
	testCase := []struct{
		name string
		expectedStatus int
	}{
		{
			name: "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.GetNewsByID()(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected : %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_UpdateNewByID(t *testing.T){
	testCase := []struct{
		name string
		expectedStatus int
	}{
		{
			name: "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.UpdateNewsByID()(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected : %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_DeleteNewsByID(t *testing.T){
	testCase := []struct{
		name string
		expectedStatus int
	}{
		{
			name: "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPost, "/", nil)

			// Act
			handler.DeleteNewsByID()(w, r)

			// Assert
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected : %d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}
