package server

import (
	"fmt"
	"io"
	"net/http"
)

type MockServer struct {
}

func (m *MockServer) GetFromMock() []byte {
	mockServerURL := "http://localhost:3000/v1/api"

	resp, err := http.Get(mockServerURL + "/todolists/getall")
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil
	}
	return body
}
