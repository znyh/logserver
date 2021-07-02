// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/znyh/middle-end/logserver/internal/dao"
	"github.com/znyh/middle-end/logserver/internal/server/grpc"
	"github.com/znyh/middle-end/logserver/internal/server/http"
	"github.com/znyh/middle-end/logserver/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
