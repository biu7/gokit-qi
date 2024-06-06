package logto

import (
	"fmt"
	"testing"
)

func TestLogto_Parse(t *testing.T) {
	logto, err := NewLogto("login.randome.chat")
	if err != nil {
		panic(err)
	}
	token, err := logto.Parse("eyJhbGciOiJFUzM4NCIsInR5cCI6IkpXVCIsImtpZCI6IjZibmVFcjIyX0F4MWh5MURMZkZpR1k5SEpGdENNOS1lZzJHRzhIa253ZE0ifQ.eyJzdWIiOiJwNjM1YWtrZ211azAiLCJuYW1lIjoi5p-Q5LiDIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FDZzhvY0pqelhiczYxMjdfaHdTQllObE9uNG9pNXk3eTB6Z25zRFF6VmJJM2lSV0lJNV9RZz1zOTYtYyIsInVwZGF0ZWRfYXQiOjE3MTc2Nzc1ODkzOTksInVzZXJuYW1lIjpudWxsLCJjcmVhdGVkX2F0IjoxNjkxOTI0MTE3Njg0LCJvcmdhbml6YXRpb25zIjpbXSwib3JnYW5pemF0aW9uX3JvbGVzIjpbXSwiYXV0aF90aW1lIjoxNzE3Njc3NTkwLCJhdF9oYXNoIjoiYndKQlo4SDBjVlFTa1lMNG1ydmFhbURzRHlFd1I3M1EiLCJhdWQiOiJkZW1vLWFwcCIsImV4cCI6MTcxNzY4MTE5MywiaWF0IjoxNzE3Njc3NTkzLCJpc3MiOiJodHRwczovL2JkcnJxay5sb2d0by5hcHAvb2lkYyJ9.P6JjGSaIYn0yKi04OT-1CjOAZ7BtzSsq9eThkOmHtjm77O49ZnMIfbJDXEa_dqbE9Ft_VSwj_NSHzwtJqX433t9OgcZ07dL8jY93VljxzaAf3xymrVqvZRrqmW3naBWg")
	if err != nil {
		panic(err)
	}
	fmt.Println(token.Valid)
	fmt.Println(token.Claims)
}
