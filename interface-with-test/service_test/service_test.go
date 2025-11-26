package service_test

import (
	"testing"

	"example.com/project/client"
	"example.com/project/service"
)

// MockClient implements client.Client for testing
type MockClient struct {
	Data string
}

var _ client.Client = (*MockClient)(nil)

func (m MockClient) FetchData() string {
	return m.Data
}

func TestServiceProcess(t *testing.T) {
	mock := MockClient{Data: "fake data"}
	s := service.Service{C: mock}

	got := s.Process()
	want := "processed: fake data"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
