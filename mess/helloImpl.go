package mess

import (
	"context"
	"fmt"
	"strings"
	"time"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// ImplHelloUpperServer sd
type ImplHelloUpperServer struct{}

// Hello impl
func (*ImplHelloUpperServer) Hello(c context.Context, u *User) (*User, error) {
	return &User{
		Str: strings.ToUpper(u.Str),
	}, nil
}

//GetCompany get company
func (*ImplHelloUpperServer) GetCompany(c context.Context, r *Request) (*Response, error) {
	return &Response{
		CompanyId:   r.CompanyId,
		CompanyName: "google",
		Address:     "New York",
		Domain:      "www.google.com",
		Enable:      true,
	}, nil
}

//GetStream get stream
func (*ImplHelloUpperServer) GetStream(req *ReqData, my MyStream_GetStreamServer) error {
	for i := 0; i < 10; i++ {
		my.Send(&ResData{Data: format()})
		time.Sleep(time.Millisecond * 1000)
	}
	return nil
}

//PutStream put stream
func (*ImplHelloUpperServer) PutStream(MyStream_PutStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PutStream not implemented")
}

//AllStream all stream
func (*ImplHelloUpperServer) AllStream(MyStream_AllStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AllStream not implemented")
}

func format() string {
	now := time.Now()
	return fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}
