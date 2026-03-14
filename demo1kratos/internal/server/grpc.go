// Package server provides HTTP and gRPC with two-step auth middleware
//
// Package server 提供带双层认证中间件的 HTTP 和 gRPC 服务
package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	pb "github.com/yylego/kratos-examples/demo1kratos/api/student"
	"github.com/yylego/kratos-examples/demo1kratos/internal/conf"
	"github.com/yylego/kratos-examples/demo1kratos/internal/data"
	"github.com/yylego/kratos-examples/demo1kratos/internal/service"
)

// NewGRPCServer creates gRPC with two-step auth middleware
// Step 1: Role-based auth from config (NewRoleMiddleware)
// Step 2: User-based auth from database (NewUserMiddleware)
//
// NewGRPCServer 创建带双层认证中间件的 gRPC 服务
// 第一层：基于配置的角色认证（NewRoleMiddleware）
// 第二层：基于数据库的用户认证（NewUserMiddleware）
func NewGRPCServer(c *conf.Server, dataData *data.Data, student *service.StudentService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			NewRoleMiddleware(c, logger),        // Role-based auth from config // 基于配置的角色认证
			NewUserMiddleware(dataData, logger), // User-based auth from database // 基于数据库的用户认证
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Address != "" {
		opts = append(opts, grpc.Address(c.Grpc.Address))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	pb.RegisterStudentServiceServer(srv, student)
	return srv
}
