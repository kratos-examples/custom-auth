// Package dbauth provides database-based token validation and admin info extraction
//
// Package dbauth 提供基于数据库的令牌验证和管理员信息提取
package dbauth

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	pb "github.com/yylego/kratos-examples/demo2kratos/api/article"
	"github.com/yylego/kratos-examples/demo2kratos/internal/pkg/models"
	"github.com/yylego/must"
	"gorm.io/gorm"
)

var adminRepo = gormrepo.NewRepo(gormclass.Use(&models.Admin{})) // Admin table repo // Admin 表仓库
var tokenRepo = gormrepo.NewRepo(gormclass.Use(&models.Token{})) // Token table repo // Token 表仓库

// CheckToken validates token from database and stores admin info in context
// Returns updated context with auth info or error if validation fails
//
// CheckToken 从数据库验证令牌并将管理员信息存入上下文
// 返回包含认证信息的更新上下文，验证失败则返回错误
func CheckToken(ctx context.Context, db *gorm.DB, token string) (context.Context, *errors.Error) {
	// Step 1: Query token from database
	// 步骤1：从数据库查询令牌
	adminToken, erb := tokenRepo.With(ctx, db).FirstE(func(db *gorm.DB, cls *models.TokenColumns) *gorm.DB {
		return db.Where(cls.Token.Eq(token))
	})
	if erb != nil {
		if erb.NotExist {
			return nil, errors.Unauthorized("TOKEN_NOT_EXIST", "token not exist")
		}
		return nil, pb.ErrorUnknown("wrong db. cause=%v", erb.Cause)
	}
	must.Full(adminToken)
	must.Nice(adminToken.AdminID)

	// Step 2: Check token expiration
	// 步骤2：检查令牌是否过期
	if adminToken.ExpiresAt.Before(time.Now()) {
		return nil, errors.Unauthorized("TOKEN_EXPIRED", "token expired")
	}

	// Step 3: Query admin info by AdminID
	// 步骤3：根据 AdminID 查询管理员信息
	admin, erb := adminRepo.With(ctx, db).FirstE(func(db *gorm.DB, cls *models.AdminColumns) *gorm.DB {
		return db.Where(cls.ID.Eq(adminToken.AdminID))
	})
	if erb != nil {
		if erb.NotExist {
			return nil, errors.Unauthorized("ADMIN_NOT_EXIST", "admin not exist")
		}
		return nil, pb.ErrorUnknown("wrong db. cause=%v", erb.Cause)
	}
	must.Full(admin)
	must.Nice(admin.Username)
	must.Sane(admin.ID, adminToken.AdminID)

	// Step 4: Store auth info in context
	// 步骤4：将认证信息存入上下文
	ctx = context.WithValue(ctx, AuthInfo{}, &AuthInfo{
		Username: admin.Username,
		Mailbox:  admin.Mailbox,
		AdminID:  admin.ID,
	})
	return ctx, nil
}

// AuthInfo contains admin authentication information
//
// AuthInfo 包含管理员认证信息
type AuthInfo struct {
	Username string // Admin username // 管理员用户名
	Mailbox  string // Admin email address // 管理员邮箱地址
	AdminID  uint   // Admin unique ID // 管理员唯一ID
}

// GetAuthInfoFromContext extracts auth info from context
// Returns error if auth info not found in context
//
// GetAuthInfoFromContext 从上下文中提取认证信息
// 如果上下文中未找到认证信息则返回错误
func GetAuthInfoFromContext(ctx context.Context) (*AuthInfo, error) {
	res, ok := ctx.Value(AuthInfo{}).(*AuthInfo)
	if !ok {
		return nil, errors.Unauthorized("TOKEN_NOT_EXIST", "token not exist")
	}
	return res, nil
}

// GetAuthInfo is an alias of GetAuthInfoFromContext
//
// GetAuthInfo 是 GetAuthInfoFromContext 的别名
func GetAuthInfo(ctx context.Context) (*AuthInfo, error) {
	return GetAuthInfoFromContext(ctx)
}
