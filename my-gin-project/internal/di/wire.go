// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"my-gin-project/internal/dao"
	"my-gin-project/internal/server/grpc"
	"my-gin-project/internal/server/http"
	"my-gin-project/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
