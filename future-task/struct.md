```go
// good example is
func (r *Request) Cookie(name string) (*Cookie, error) {
	if name == "" {
		return nil, ErrNoCookie
	}
	for _, c := range readCookies(r.Header, name) {
		return c, nil
	}
	return nil, ErrNoCookie
}

/usr/local/go/src/net/http/request.go
```
