package client

type Client interface {
	FetchData() string
}

type RealClient struct{}

func (RealClient) FetchData() string {
	return "real data"
}

func Sample(x int) int {
	return x + 3
}
