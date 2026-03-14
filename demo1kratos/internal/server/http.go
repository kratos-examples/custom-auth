// Package server provides HTTP and gRPC with two-step auth middleware
//
// Package server 提供带双层认证中间件的 HTTP 和 gRPC 服务
package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/yylego/kratos-auth/authkratos"
	"github.com/yylego/kratos-custom-auth/customkratosauth"
	pb "github.com/yylego/kratos-examples/demo1kratos/api/student"
	"github.com/yylego/kratos-examples/demo1kratos/internal/conf"
	"github.com/yylego/kratos-examples/demo1kratos/internal/data"
	"github.com/yylego/kratos-examples/demo1kratos/internal/pkg/dbauth"
	"github.com/yylego/kratos-examples/demo1kratos/internal/service"
	"github.com/yylego/kratos-static-auth/statickratosauth"
	"github.com/yylego/must"
)

func NewHTTPServer(c *conf.Server, dataData *data.Data, student *service.StudentService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			NewRoleMiddleware(c, logger),        // Role-based auth from config // 基于配置的角色认证
			NewUserMiddleware(dataData, logger), // User-based auth from database // 基于数据库的用户认证
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Address != "" {
		opts = append(opts, http.Address(c.Http.Address))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	pb.RegisterStudentServiceHTTPServer(srv, student)
	return srv
}

// Requires both Authorization (role token) and AdminToken (user token) headers
// Authorization: Role-based token from config file (admin/guest)
// AdminToken: User-specific token from database (which admin)
//
// 需要同时提供 Authorization（角色令牌）和 AdminToken（用户令牌）两个请求头
// Authorization: 来自配置文件的角色令牌（admin/guest）
// AdminToken: 来自数据库的用户令牌（具体哪个管理员）
/*
curl --location 'http://127.0.0.1:8001/v1/students' --header 'Authorization: 63a16b29e5bc4a28a880de1b2e707cc6' --header 'AdminToken: 95d9fda7f675444d9acc3c8225dbf7de'
curl --location 'http://127.0.0.1:8001/v1/students' --header 'Authorization: 63a16b29e5bc4a28a880de1b2e707cc6' --header 'AdminToken: 46421ed7de4a4fcc888ff84541defbc3'
*/

// NewRoleMiddleware creates auth middleware with token validation and route scope
// Configure which routes need auth and setup valid tokens
//
// NewRoleMiddleware 创建认证中间件，进行令牌验证和路由范围控制
// 配置需要认证的路由并设置有效令牌
func NewRoleMiddleware(c *conf.Server, logger log.Logger) middleware.Middleware {
	routeScope := authkratos.NewInclude( // Create INCLUDE mode route scope // 创建 INCLUDE 模式的路由范围
		pb.OperationStudentServiceCreateStudent,
		pb.OperationStudentServiceUpdateStudent,
		pb.OperationStudentServiceDeleteStudent,
		pb.OperationStudentServiceGetStudent,
		pb.OperationStudentServiceListStudents,
	)
	authTokens := map[string]string{ // Setup valid tokens map // 设置有效令牌映射表
		"admin": must.Nice(c.Auth.AdminToken),
		"guest": must.Nice(c.Auth.GuestToken),
	}
	authConfig := statickratosauth.NewConfig(routeScope, authTokens).
		WithFieldName("Authorization").
		WithSimpleEnable(). // Enable simple token type // 启用简单令牌类型
		WithDebugMode(true) // Enable debug mode to log auth process // 启用调试模式记录认证过程
	return statickratosauth.NewMiddleware(authConfig, logger)
}

// NewUserMiddleware creates user auth middleware with database token validation
// Validate admin tokens from database and check expiration
//
// NewUserMiddleware 创建用户认证中间件，通过数据库验证令牌
// 从数据库验证管理员令牌并检查是否过期
func NewUserMiddleware(dataData *data.Data, logger log.Logger) middleware.Middleware {
	routeScope := authkratos.NewInclude( // Create INCLUDE mode route scope // 创建 INCLUDE 模式的路由范围
		pb.OperationStudentServiceCreateStudent,
		pb.OperationStudentServiceUpdateStudent,
		pb.OperationStudentServiceDeleteStudent,
		pb.OperationStudentServiceGetStudent,
		pb.OperationStudentServiceListStudents,
	)

	checkAuthFunction := func(ctx context.Context, token string) (context.Context, *errors.Error) {
		ctx, erk := dbauth.CheckToken(ctx, dataData.DB(), token) // Check token from database // 从数据库检查令牌
		if erk != nil {
			return nil, erk
		}
		return ctx, nil
	}
	authConfig := customkratosauth.NewConfig(routeScope, checkAuthFunction).
		WithFieldName("AdminToken"). // Set token field name in request header // 设置请求头中的令牌字段名
		WithDebugMode(true)          // Enable debug mode to log auth process // 启用调试模式记录认证过程
	return customkratosauth.NewMiddleware(authConfig, logger)
}
