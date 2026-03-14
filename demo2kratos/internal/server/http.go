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
	pb "github.com/yylego/kratos-examples/demo2kratos/api/article"
	"github.com/yylego/kratos-examples/demo2kratos/internal/conf"
	"github.com/yylego/kratos-examples/demo2kratos/internal/data"
	"github.com/yylego/kratos-examples/demo2kratos/internal/pkg/dbauth"
	"github.com/yylego/kratos-examples/demo2kratos/internal/service"
	"github.com/yylego/kratos-static-auth/statickratosauth"
	"github.com/yylego/must"
)

func NewHTTPServer(c *conf.Server, dataData *data.Data, article *service.ArticleService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			NewRoleMiddleware(c, logger),
			NewUserMiddleware(dataData, logger),
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
	pb.RegisterArticleServiceHTTPServer(srv, article)
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
curl --location 'http://127.0.0.1:8002/v1/articles' --header 'Authorization: c98235f2b2f746408f212976bdfae467' --header 'AdminToken: 95d9fda7f675444d9acc3c8225dbf7de'
curl --location 'http://127.0.0.1:8002/v1/articles' --header 'Authorization: c98235f2b2f746408f212976bdfae467' --header 'AdminToken: 46421ed7de4a4fcc888ff84541defbc3'
*/

// NewRoleMiddleware creates auth middleware with token validation and route scope
//
// NewRoleMiddleware 创建认证中间件，进行令牌验证和路由范围控制
func NewRoleMiddleware(c *conf.Server, logger log.Logger) middleware.Middleware {
	routeScope := authkratos.NewInclude(
		pb.OperationArticleServiceCreateArticle,
		pb.OperationArticleServiceUpdateArticle,
		pb.OperationArticleServiceDeleteArticle,
		pb.OperationArticleServiceGetArticle,
		pb.OperationArticleServiceListArticles,
	)
	authTokens := map[string]string{
		"admin": must.Nice(c.Auth.AdminToken),
		"guest": must.Nice(c.Auth.GuestToken),
	}
	authConfig := statickratosauth.NewConfig(routeScope, authTokens).
		WithFieldName("Authorization").
		WithSimpleEnable().
		WithDebugMode(true)
	return statickratosauth.NewMiddleware(authConfig, logger)
}

// NewUserMiddleware creates user auth middleware with database token validation
//
// NewUserMiddleware 创建用户认证中间件，通过数据库验证令牌
func NewUserMiddleware(dataData *data.Data, logger log.Logger) middleware.Middleware {
	routeScope := authkratos.NewInclude(
		pb.OperationArticleServiceCreateArticle,
		pb.OperationArticleServiceUpdateArticle,
		pb.OperationArticleServiceDeleteArticle,
		pb.OperationArticleServiceGetArticle,
		pb.OperationArticleServiceListArticles,
	)

	checkAuthFunction := func(ctx context.Context, token string) (context.Context, *errors.Error) {
		ctx, erk := dbauth.CheckToken(ctx, dataData.DB(), token)
		if erk != nil {
			return nil, erk
		}
		return ctx, nil
	}
	authConfig := customkratosauth.NewConfig(routeScope, checkAuthFunction).
		WithFieldName("AdminToken").
		WithDebugMode(true)
	return customkratosauth.NewMiddleware(authConfig, logger)
}
