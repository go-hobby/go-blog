package component

import (
	"github.com/google/wire"
	"go-blog/internal/app/component/casbin"
	"go-blog/internal/app/component/client/grpc"
	"go-blog/internal/app/component/discovery"
	"go-blog/internal/app/component/orm"
	"go-blog/internal/app/component/redis"
	"go-blog/internal/app/component/trace"
	"go-blog/internal/app/component/uid"
)

var ProviderSet = wire.NewSet(
	wire.NewSet(orm.New), //gorm 初始化
	//wire.NewSet(data.New),//ent 初始化
	wire.NewSet(redis.New),     // redis 初始化
	wire.NewSet(trace.New),     //链路追踪初始化
	wire.NewSet(discovery.New), //注册到etcd 或  Consul
	wire.NewSet(casbin.New),    //casbin 权限验证
	wire.NewSet(wire.Bind(new(uid.Generator), new(*uid.Uid)), uid.New), //uuid
	grpc.New, //grpc
)
