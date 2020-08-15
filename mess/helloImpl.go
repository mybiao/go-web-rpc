package mess

import (
	"context"
	"strings"
)

// ImplHelloUpperServer sd
type ImplHelloUpperServer struct{}

// Hello impl
func (*ImplHelloUpperServer) Hello(c context.Context, u *User) (*User, error) {
	return &User{
		Str: strings.ToUpper(u.Str),
	}, nil
}
