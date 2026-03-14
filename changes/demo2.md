# Changes

Code differences compared to source project.

## cmd/demo2kratos/wire_gen.go (+3 -3)

```diff
@@ -24,9 +24,9 @@
 		return nil, nil, err
 	}
 	articleUsecase := biz.NewArticleUsecase(dataData, logger)
-	articleService := service.NewArticleService(articleUsecase)
-	grpcServer := server.NewGRPCServer(confServer, articleService, logger)
-	httpServer := server.NewHTTPServer(confServer, articleService, logger)
+	articleService := service.NewArticleService(articleUsecase, logger)
+	grpcServer := server.NewGRPCServer(confServer, dataData, articleService, logger)
+	httpServer := server.NewHTTPServer(confServer, dataData, articleService, logger)
 	app := newApp(logger, grpcServer, httpServer)
 	return app, func() {
 		cleanup()
```

## configs/config.yaml (+3 -0)

```diff
@@ -5,6 +5,9 @@
   grpc:
     address: 0.0.0.0:9002
     timeout: 1s
+  auth:
+    adminToken: "c98235f2b2f746408f212976bdfae467"
+    guestToken: "61d1a2e27b734d959670112389f7b2c2"
 data:
   database:
     driver: sqlite3
```

## internal/conf/conf.pb.go (+87 -17)

```diff
@@ -78,6 +78,7 @@
 	state         protoimpl.MessageState `protogen:"open.v1"`
 	Http          *Server_HTTP           `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
 	Grpc          *Server_GRPC           `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
+	Auth          *Server_Auth           `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
 	unknownFields protoimpl.UnknownFields
 	sizeCache     protoimpl.SizeCache
 }
@@ -126,6 +127,13 @@
 	return nil
 }
 
+func (x *Server) GetAuth() *Server_Auth {
+	if x != nil {
+		return x.Auth
+	}
+	return nil
+}
+
 type Data struct {
 	state         protoimpl.MessageState `protogen:"open.v1"`
 	Database      *Data_Database         `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
@@ -290,6 +298,58 @@
 	return nil
 }
 
+type Server_Auth struct {
+	state         protoimpl.MessageState `protogen:"open.v1"`
+	AdminToken    string                 `protobuf:"bytes,1,opt,name=adminToken,proto3" json:"adminToken,omitempty"`
+	GuestToken    string                 `protobuf:"bytes,2,opt,name=guestToken,proto3" json:"guestToken,omitempty"`
+	unknownFields protoimpl.UnknownFields
+	sizeCache     protoimpl.SizeCache
+}
+
+func (x *Server_Auth) Reset() {
+	*x = Server_Auth{}
+	mi := &file_conf_conf_proto_msgTypes[5]
+	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
+	ms.StoreMessageInfo(mi)
+}
+
+func (x *Server_Auth) String() string {
+	return protoimpl.X.MessageStringOf(x)
+}
+
+func (*Server_Auth) ProtoMessage() {}
+
+func (x *Server_Auth) ProtoReflect() protoreflect.Message {
+	mi := &file_conf_conf_proto_msgTypes[5]
+	if x != nil {
+		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
+		if ms.LoadMessageInfo() == nil {
+			ms.StoreMessageInfo(mi)
+		}
+		return ms
+	}
+	return mi.MessageOf(x)
+}
+
+// Deprecated: Use Server_Auth.ProtoReflect.Descriptor instead.
+func (*Server_Auth) Descriptor() ([]byte, []int) {
+	return file_conf_conf_proto_rawDescGZIP(), []int{1, 2}
+}
+
+func (x *Server_Auth) GetAdminToken() string {
+	if x != nil {
+		return x.AdminToken
+	}
+	return ""
+}
+
+func (x *Server_Auth) GetGuestToken() string {
+	if x != nil {
+		return x.GuestToken
+	}
+	return ""
+}
+
 type Data_Database struct {
 	state         protoimpl.MessageState `protogen:"open.v1"`
 	Driver        string                 `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
@@ -300,7 +360,7 @@
 
 func (x *Data_Database) Reset() {
 	*x = Data_Database{}
-	mi := &file_conf_conf_proto_msgTypes[5]
+	mi := &file_conf_conf_proto_msgTypes[6]
 	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
 	ms.StoreMessageInfo(mi)
 }
@@ -312,7 +372,7 @@
 func (*Data_Database) ProtoMessage() {}
 
 func (x *Data_Database) ProtoReflect() protoreflect.Message {
-	mi := &file_conf_conf_proto_msgTypes[5]
+	mi := &file_conf_conf_proto_msgTypes[6]
 	if x != nil {
 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
 		if ms.LoadMessageInfo() == nil {
@@ -350,10 +410,11 @@
 	"kratos.api\x1a\x1egoogle/protobuf/duration.proto\"]\n" +
 	"\tBootstrap\x12*\n" +
 	"\x06server\x18\x01 \x01(\v2\x12.kratos.api.ServerR\x06server\x12$\n" +
-	"\x04data\x18\x02 \x01(\v2\x10.kratos.api.DataR\x04data\"\xc4\x02\n" +
+	"\x04data\x18\x02 \x01(\v2\x10.kratos.api.DataR\x04data\"\xb9\x03\n" +
 	"\x06Server\x12+\n" +
 	"\x04http\x18\x01 \x01(\v2\x17.kratos.api.Server.HTTPR\x04http\x12+\n" +
-	"\x04grpc\x18\x02 \x01(\v2\x17.kratos.api.Server.GRPCR\x04grpc\x1ao\n" +
+	"\x04grpc\x18\x02 \x01(\v2\x17.kratos.api.Server.GRPCR\x04grpc\x12+\n" +
+	"\x04auth\x18\x03 \x01(\v2\x17.kratos.api.Server.AuthR\x04auth\x1ao\n" +
 	"\x04HTTP\x12\x18\n" +
 	"\anetwork\x18\x01 \x01(\tR\anetwork\x12\x18\n" +
 	"\aaddress\x18\x02 \x01(\tR\aaddress\x123\n" +
@@ -361,7 +422,14 @@
 	"\x04GRPC\x12\x18\n" +
 	"\anetwork\x18\x01 \x01(\tR\anetwork\x12\x18\n" +
 	"\aaddress\x18\x02 \x01(\tR\aaddress\x123\n" +
-	"\atimeout\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\atimeout\"y\n" +
+	"\atimeout\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\atimeout\x1aF\n" +
+	"\x04Auth\x12\x1e\n" +
+	"\n" +
+	"adminToken\x18\x01 \x01(\tR\n" +
+	"adminToken\x12\x1e\n" +
+	"\n" +
+	"guestToken\x18\x02 \x01(\tR\n" +
+	"guestToken\"y\n" +
 	"\x04Data\x125\n" +
 	"\bdatabase\x18\x01 \x01(\v2\x19.kratos.api.Data.DatabaseR\bdatabase\x1a:\n" +
 	"\bDatabase\x12\x16\n" +
@@ -380,29 +448,31 @@
 	return file_conf_conf_proto_rawDescData
 }
 
-var file_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
+var file_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
 var file_conf_conf_proto_goTypes = []any{
 	(*Bootstrap)(nil),           // 0: kratos.api.Bootstrap
 	(*Server)(nil),              // 1: kratos.api.Server
 	(*Data)(nil),                // 2: kratos.api.Data
 	(*Server_HTTP)(nil),         // 3: kratos.api.Server.HTTP
 	(*Server_GRPC)(nil),         // 4: kratos.api.Server.GRPC
-	(*Data_Database)(nil),       // 5: kratos.api.Data.Database
-	(*durationpb.Duration)(nil), // 6: google.protobuf.Duration
+	(*Server_Auth)(nil),         // 5: kratos.api.Server.Auth
+	(*Data_Database)(nil),       // 6: kratos.api.Data.Database
+	(*durationpb.Duration)(nil), // 7: google.protobuf.Duration
 }
 var file_conf_conf_proto_depIdxs = []int32{
 	1, // 0: kratos.api.Bootstrap.server:type_name -> kratos.api.Server
 	2, // 1: kratos.api.Bootstrap.data:type_name -> kratos.api.Data
 	3, // 2: kratos.api.Server.http:type_name -> kratos.api.Server.HTTP
 	4, // 3: kratos.api.Server.grpc:type_name -> kratos.api.Server.GRPC
-	5, // 4: kratos.api.Data.database:type_name -> kratos.api.Data.Database
-	6, // 5: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
-	6, // 6: kratos.api.Server.GRPC.timeout:type_name -> google.protobuf.Duration
-	7, // [7:7] is the sub-list for method output_type
-	7, // [7:7] is the sub-list for method input_type
-	7, // [7:7] is the sub-list for extension type_name
-	7, // [7:7] is the sub-list for extension extendee
-	0, // [0:7] is the sub-list for field type_name
+	5, // 4: kratos.api.Server.auth:type_name -> kratos.api.Server.Auth
+	6, // 5: kratos.api.Data.database:type_name -> kratos.api.Data.Database
+	7, // 6: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
+	7, // 7: kratos.api.Server.GRPC.timeout:type_name -> google.protobuf.Duration
+	8, // [8:8] is the sub-list for method output_type
+	8, // [8:8] is the sub-list for method input_type
+	8, // [8:8] is the sub-list for extension type_name
+	8, // [8:8] is the sub-list for extension extendee
+	0, // [0:8] is the sub-list for field type_name
 }
 
 func init() { file_conf_conf_proto_init() }
@@ -416,7 +486,7 @@
 			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
 			RawDescriptor: unsafe.Slice(unsafe.StringData(file_conf_conf_proto_rawDesc), len(file_conf_conf_proto_rawDesc)),
 			NumEnums:      0,
-			NumMessages:   6,
+			NumMessages:   7,
 			NumExtensions: 0,
 			NumServices:   0,
 		},
```

## internal/conf/conf.proto (+5 -0)

```diff
@@ -21,8 +21,13 @@
     string address = 2;
     google.protobuf.Duration timeout = 3;
   }
+  message Auth {
+    string adminToken = 1;
+    string guestToken = 2;
+  }
   HTTP http = 1;
   GRPC grpc = 2;
+  Auth auth = 3;
 }
 
 message Data {
```

## internal/data/data.go (+62 -5)

```diff
@@ -1,27 +1,84 @@
 package data
 
 import (
+	"time"
+
 	"github.com/go-kratos/kratos/v2/log"
 	"github.com/google/wire"
 	"github.com/yylego/kratos-examples/demo2kratos/internal/conf"
+	"github.com/yylego/kratos-examples/demo2kratos/internal/pkg/models"
 	"github.com/yylego/must"
 	"github.com/yylego/rese"
 	"gorm.io/driver/sqlite"
 	"gorm.io/gorm"
+	loggergorm "gorm.io/gorm/logger"
 )
 
 var ProviderSet = wire.NewSet(NewData)
 
 type Data struct {
-	db *gorm.DB
+	db *gorm.DB // GORM database connection instance // GORM 数据库连接实例
 }
 
 func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
-	must.Same(c.Database.Driver, "sqlite3")
-	db := rese.P1(gorm.Open(sqlite.Open(c.Database.Source), &gorm.Config{}))
+	dsn := must.Nice(c.Database.Source)
+	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
+		Logger: loggergorm.Default.LogMode(loggergorm.Info),
+	}))
+
+	must.Done(db.AutoMigrate(&models.Admin{}, &models.Token{}))
+
+	mockAdmin1(db)
+	mockAdmin2(db)
+
 	cleanup := func() {
 		log.NewHelper(logger).Info("closing the data resources")
-		_ = rese.P1(db.DB()).Close()
+		must.Done(rese.P1(db.DB()).Close())
 	}
-	return &Data{db: db}, cleanup, nil
+	return &Data{
+		db: db,
+	}, cleanup, nil
+}
+
+// DB returns the database instance
+//
+// DB 返回数据库实例
+func (d *Data) DB() *gorm.DB {
+	return d.db
+}
+
+// mockAdmin1 creates test admin with username "abc" and access token
+//
+// mockAdmin1 创建测试管理员（用户名 "abc"）及其访问令牌
+func mockAdmin1(db *gorm.DB) {
+	must.Done(db.Create(&models.Admin{
+		Username: "abc",
+		Password: "123",
+		Mailbox:  "",
+		Status:   0,
+	}).Error)
+	must.Done(db.Create(&models.Token{
+		AdminID:   1,
+		Token:     "95d9fda7f675444d9acc3c8225dbf7de",
+		Type:      "access",
+		ExpiresAt: time.Now().UTC().Add(30 * time.Minute),
+	}).Error)
+}
+
+// mockAdmin2 creates test admin with username "xyz" and access token
+//
+// mockAdmin2 创建测试管理员（用户名 "xyz"）及其访问令牌
+func mockAdmin2(db *gorm.DB) {
+	must.Done(db.Create(&models.Admin{
+		Username: "xyz",
+		Password: "456",
+		Mailbox:  "",
+		Status:   0,
+	}).Error)
+	must.Done(db.Create(&models.Token{
+		AdminID:   2,
+		Token:     "46421ed7de4a4fcc888ff84541defbc3",
+		Type:      "access",
+		ExpiresAt: time.Now().UTC().Add(30 * time.Minute),
+	}).Error)
 }
```

## internal/pkg/dbauth/db_auth.go (+100 -0)

```diff
@@ -0,0 +1,100 @@
+// Package dbauth provides database-based token validation and admin info extraction
+//
+// Package dbauth 提供基于数据库的令牌验证和管理员信息提取
+package dbauth
+
+import (
+	"context"
+	"time"
+
+	"github.com/go-kratos/kratos/v2/errors"
+	"github.com/yylego/gormrepo"
+	"github.com/yylego/gormrepo/gormclass"
+	pb "github.com/yylego/kratos-examples/demo2kratos/api/article"
+	"github.com/yylego/kratos-examples/demo2kratos/internal/pkg/models"
+	"github.com/yylego/must"
+	"gorm.io/gorm"
+)
+
+var adminRepo = gormrepo.NewRepo(gormclass.Use(&models.Admin{})) // Admin table repo // Admin 表仓库
+var tokenRepo = gormrepo.NewRepo(gormclass.Use(&models.Token{})) // Token table repo // Token 表仓库
+
+// CheckToken validates token from database and stores admin info in context
+// Returns updated context with auth info or error if validation fails
+//
+// CheckToken 从数据库验证令牌并将管理员信息存入上下文
+// 返回包含认证信息的更新上下文，验证失败则返回错误
+func CheckToken(ctx context.Context, db *gorm.DB, token string) (context.Context, *errors.Error) {
+	// Step 1: Query token from database
+	// 步骤1：从数据库查询令牌
+	adminToken, erb := tokenRepo.With(ctx, db).FirstE(func(db *gorm.DB, cls *models.TokenColumns) *gorm.DB {
+		return db.Where(cls.Token.Eq(token))
+	})
+	if erb != nil {
+		if erb.NotExist {
+			return nil, errors.Unauthorized("TOKEN_NOT_EXIST", "token not exist")
+		}
+		return nil, pb.ErrorUnknown("wrong db. cause=%v", erb.Cause)
+	}
+	must.Full(adminToken)
+	must.Nice(adminToken.AdminID)
+
+	// Step 2: Check token expiration
+	// 步骤2：检查令牌是否过期
+	if adminToken.ExpiresAt.Before(time.Now()) {
+		return nil, errors.Unauthorized("TOKEN_EXPIRED", "token expired")
+	}
+
+	// Step 3: Query admin info by AdminID
+	// 步骤3：根据 AdminID 查询管理员信息
+	admin, erb := adminRepo.With(ctx, db).FirstE(func(db *gorm.DB, cls *models.AdminColumns) *gorm.DB {
+		return db.Where(cls.ID.Eq(adminToken.AdminID))
+	})
+	if erb != nil {
+		if erb.NotExist {
+			return nil, errors.Unauthorized("ADMIN_NOT_EXIST", "admin not exist")
+		}
+		return nil, pb.ErrorUnknown("wrong db. cause=%v", erb.Cause)
+	}
+	must.Full(admin)
+	must.Nice(admin.Username)
+	must.Sane(admin.ID, adminToken.AdminID)
+
+	// Step 4: Store auth info in context
+	// 步骤4：将认证信息存入上下文
+	ctx = context.WithValue(ctx, AuthInfo{}, &AuthInfo{
+		Username: admin.Username,
+		Mailbox:  admin.Mailbox,
+		AdminID:  admin.ID,
+	})
+	return ctx, nil
+}
+
+// AuthInfo contains admin authentication information
+//
+// AuthInfo 包含管理员认证信息
+type AuthInfo struct {
+	Username string // Admin username // 管理员用户名
+	Mailbox  string // Admin email address // 管理员邮箱地址
+	AdminID  uint   // Admin unique ID // 管理员唯一ID
+}
+
+// GetAuthInfoFromContext extracts auth info from context
+// Returns error if auth info not found in context
+//
+// GetAuthInfoFromContext 从上下文中提取认证信息
+// 如果上下文中未找到认证信息则返回错误
+func GetAuthInfoFromContext(ctx context.Context) (*AuthInfo, error) {
+	res, ok := ctx.Value(AuthInfo{}).(*AuthInfo)
+	if !ok {
+		return nil, errors.Unauthorized("TOKEN_NOT_EXIST", "token not exist")
+	}
+	return res, nil
+}
+
+// GetAuthInfo is an alias of GetAuthInfoFromContext
+//
+// GetAuthInfo 是 GetAuthInfoFromContext 的别名
+func GetAuthInfo(ctx context.Context) (*AuthInfo, error) {
+	return GetAuthInfoFromContext(ctx)
+}
```

## internal/pkg/models/admin.go (+19 -0)

```diff
@@ -0,0 +1,19 @@
+package models
+
+import "gorm.io/gorm"
+
+// Admin model stores admin user info
+// Admin 模型存储管理员用户信息
+type Admin struct {
+	gorm.Model
+	Username string `gorm:"uniqueIndex;size:64" json:"username"` // Unique username // 唯一用户名
+	Password string `gorm:"size:128" json:"-"`                   // Encrypted password // 加密密码
+	Mailbox  string `gorm:"size:128" json:"mailbox"`             // Email address // 邮箱地址
+	Status   int    `gorm:"default:1" json:"status"`             // Status: 1=active, 0=disabled // 状态：1=启用，0=禁用
+}
+
+// TableName sets custom table name
+// TableName 设置自定义表名
+func (*Admin) TableName() string {
+	return "tb_admin"
+}
```

## internal/pkg/models/gormcnm.gen.go (+75 -0)

```diff
@@ -0,0 +1,75 @@
+// Code generated using gormcngen. DO NOT EDIT.
+// This file was auto generated via github.com/yylego/gormcngen
+
+//go:build !gormcngen_generate
+
+// Generated from: gormcnm.gen_test.go:35 -> models_test.TestGenerateColumns
+// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========
+
+// Code generated using gormcngen. DO NOT EDIT.
+// This file was auto generated via github.com/yylego/gormcngen
+
+package models
+
+import (
+	"time"
+
+	"github.com/yylego/gormcnm"
+	"gorm.io/gorm"
+)
+
+func (c *Admin) Columns() *AdminColumns {
+	return &AdminColumns{
+		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
+		ID:        gormcnm.Cnm(c.ID, "id"),
+		CreatedAt: gormcnm.Cnm(c.CreatedAt, "created_at"),
+		UpdatedAt: gormcnm.Cnm(c.UpdatedAt, "updated_at"),
+		DeletedAt: gormcnm.Cnm(c.DeletedAt, "deleted_at"),
+		Username:  gormcnm.Cnm(c.Username, "username"),
+		Password:  gormcnm.Cnm(c.Password, "password"),
+		Mailbox:   gormcnm.Cnm(c.Mailbox, "mailbox"),
+		Status:    gormcnm.Cnm(c.Status, "status"),
+	}
+}
+
+type AdminColumns struct {
+	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
+	gormcnm.ColumnOperationClass
+	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
+	ID        gormcnm.ColumnName[uint]
+	CreatedAt gormcnm.ColumnName[time.Time]
+	UpdatedAt gormcnm.ColumnName[time.Time]
+	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
+	Username  gormcnm.ColumnName[string]
+	Password  gormcnm.ColumnName[string]
+	Mailbox   gormcnm.ColumnName[string]
+	Status    gormcnm.ColumnName[int]
+}
+
+func (c *Token) Columns() *TokenColumns {
+	return &TokenColumns{
+		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
+		ID:        gormcnm.Cnm(c.ID, "id"),
+		CreatedAt: gormcnm.Cnm(c.CreatedAt, "created_at"),
+		UpdatedAt: gormcnm.Cnm(c.UpdatedAt, "updated_at"),
+		DeletedAt: gormcnm.Cnm(c.DeletedAt, "deleted_at"),
+		AdminID:   gormcnm.Cnm(c.AdminID, "admin_id"),
+		Token:     gormcnm.Cnm(c.Token, "token"),
+		Type:      gormcnm.Cnm(c.Type, "type"),
+		ExpiresAt: gormcnm.Cnm(c.ExpiresAt, "expires_at"),
+	}
+}
+
+type TokenColumns struct {
+	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
+	gormcnm.ColumnOperationClass
+	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
+	ID        gormcnm.ColumnName[uint]
+	CreatedAt gormcnm.ColumnName[time.Time]
+	UpdatedAt gormcnm.ColumnName[time.Time]
+	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
+	AdminID   gormcnm.ColumnName[uint]
+	Token     gormcnm.ColumnName[string]
+	Type      gormcnm.ColumnName[string]
+	ExpiresAt gormcnm.ColumnName[time.Time]
+}
```

## internal/pkg/models/gormcnm.gen_test.go (+37 -0)

```diff
@@ -0,0 +1,37 @@
+package models_test
+
+import (
+	"testing"
+
+	"github.com/yylego/gormcngen"
+	"github.com/yylego/kratos-examples/demo2kratos/internal/pkg/models"
+	"github.com/yylego/osexistpath/osmustexist"
+	"github.com/yylego/runpath/runtestpath"
+)
+
+// Auto generate columns with go generate command
+// Support execution via: go generate ./...
+// Delete this comment block if auto generation is not needed
+//
+// 使用 go generate 命令自动生成列定义
+// 支持通过以下命令执行：go generate ./...
+// 如果不需要自动生成，可以删除此注释块
+//
+//go:generate go test -v -run TestGenerateColumns
+func TestGenerateColumns(t *testing.T) {
+	absPath := osmustexist.FILE(runtestpath.SrcPath(t))
+	t.Log(absPath)
+
+	objects := []any{
+		&models.Admin{},
+		&models.Token{},
+	}
+
+	options := gormcngen.NewOptions().
+		WithColumnClassExportable(true).
+		WithColumnsMethodRecvName("c").
+		WithColumnsCheckFieldType(true)
+
+	cfg := gormcngen.NewConfigs(objects, options, absPath)
+	cfg.Gen()
+}
```

## internal/pkg/models/token.go (+23 -0)

```diff
@@ -0,0 +1,23 @@
+package models
+
+import (
+	"time"
+
+	"gorm.io/gorm"
+)
+
+// Token model stores auth tokens
+// Token 模型存储认证令牌
+type Token struct {
+	gorm.Model
+	AdminID   uint      `gorm:"index" json:"admin_id"`         // Related admin ID // 关联的管理员ID
+	Token     string    `gorm:"uniqueIndex;size:128" json:"-"` // Token value // 令牌值
+	Type      string    `gorm:"size:32" json:"type"`           // Token type: access, refresh // 令牌类型：access、refresh
+	ExpiresAt time.Time `json:"expires_at"`                    // Expire time // 过期时间
+}
+
+// TableName sets custom table name
+// TableName 设置自定义表名
+func (*Token) TableName() string {
+	return "tb_token"
+}
```

## internal/server/grpc.go (+14 -1)

```diff
@@ -1,3 +1,6 @@
+// Package server provides HTTP and gRPC with two-step auth middleware
+//
+// Package server 提供带双层认证中间件的 HTTP 和 gRPC 服务
 package server
 
 import (
@@ -6,13 +9,23 @@
 	"github.com/go-kratos/kratos/v2/transport/grpc"
 	pb "github.com/yylego/kratos-examples/demo2kratos/api/article"
 	"github.com/yylego/kratos-examples/demo2kratos/internal/conf"
+	"github.com/yylego/kratos-examples/demo2kratos/internal/data"
 	"github.com/yylego/kratos-examples/demo2kratos/internal/service"
 )
 
-func NewGRPCServer(c *conf.Server, article *service.ArticleService, logger log.Logger) *grpc.Server {
+// NewGRPCServer creates gRPC with two-step auth middleware
+// Step 1: Role-based auth from config (NewRoleMiddleware)
+// Step 2: User-based auth from database (NewUserMiddleware)
+//
+// NewGRPCServer 创建带双层认证中间件的 gRPC 服务
+// 第一层：基于配置的角色认证（NewRoleMiddleware）
+// 第二层：基于数据库的用户认证（NewUserMiddleware）
+func NewGRPCServer(c *conf.Server, dataData *data.Data, article *service.ArticleService, logger log.Logger) *grpc.Server {
 	var opts = []grpc.ServerOption{
 		grpc.Middleware(
 			recovery.Recovery(),
+			NewRoleMiddleware(c, logger),
+			NewUserMiddleware(dataData, logger),
 		),
 	}
 	if c.Grpc.Network != "" {
```

## internal/server/http.go (+75 -1)

```diff
@@ -1,18 +1,33 @@
+// Package server provides HTTP and gRPC with two-step auth middleware
+//
+// Package server 提供带双层认证中间件的 HTTP 和 gRPC 服务
 package server
 
 import (
+	"context"
+
+	"github.com/go-kratos/kratos/v2/errors"
 	"github.com/go-kratos/kratos/v2/log"
+	"github.com/go-kratos/kratos/v2/middleware"
 	"github.com/go-kratos/kratos/v2/middleware/recovery"
 	"github.com/go-kratos/kratos/v2/transport/http"
+	"github.com/yylego/kratos-auth/authkratos"
+	"github.com/yylego/kratos-custom-auth/customkratosauth"
 	pb "github.com/yylego/kratos-examples/demo2kratos/api/article"
 	"github.com/yylego/kratos-examples/demo2kratos/internal/conf"
+	"github.com/yylego/kratos-examples/demo2kratos/internal/data"
+	"github.com/yylego/kratos-examples/demo2kratos/internal/pkg/dbauth"
 	"github.com/yylego/kratos-examples/demo2kratos/internal/service"
+	"github.com/yylego/kratos-static-auth/statickratosauth"
+	"github.com/yylego/must"
 )
 
-func NewHTTPServer(c *conf.Server, article *service.ArticleService, logger log.Logger) *http.Server {
+func NewHTTPServer(c *conf.Server, dataData *data.Data, article *service.ArticleService, logger log.Logger) *http.Server {
 	var opts = []http.ServerOption{
 		http.Middleware(
 			recovery.Recovery(),
+			NewRoleMiddleware(c, logger),
+			NewUserMiddleware(dataData, logger),
 		),
 	}
 	if c.Http.Network != "" {
@@ -27,4 +42,63 @@
 	srv := http.NewServer(opts...)
 	pb.RegisterArticleServiceHTTPServer(srv, article)
 	return srv
+}
+
+// Requires both Authorization (role token) and AdminToken (user token) headers
+// Authorization: Role-based token from config file (admin/guest)
+// AdminToken: User-specific token from database (which admin)
+//
+// 需要同时提供 Authorization（角色令牌）和 AdminToken（用户令牌）两个请求头
+// Authorization: 来自配置文件的角色令牌（admin/guest）
+// AdminToken: 来自数据库的用户令牌（具体哪个管理员）
+/*
+curl --location 'http://127.0.0.1:8002/v1/articles' --header 'Authorization: c98235f2b2f746408f212976bdfae467' --header 'AdminToken: 95d9fda7f675444d9acc3c8225dbf7de'
+curl --location 'http://127.0.0.1:8002/v1/articles' --header 'Authorization: c98235f2b2f746408f212976bdfae467' --header 'AdminToken: 46421ed7de4a4fcc888ff84541defbc3'
+*/
+
+// NewRoleMiddleware creates auth middleware with token validation and route scope
+//
+// NewRoleMiddleware 创建认证中间件，进行令牌验证和路由范围控制
+func NewRoleMiddleware(c *conf.Server, logger log.Logger) middleware.Middleware {
+	routeScope := authkratos.NewInclude(
+		pb.OperationArticleServiceCreateArticle,
+		pb.OperationArticleServiceUpdateArticle,
+		pb.OperationArticleServiceDeleteArticle,
+		pb.OperationArticleServiceGetArticle,
+		pb.OperationArticleServiceListArticles,
+	)
+	authTokens := map[string]string{
+		"admin": must.Nice(c.Auth.AdminToken),
+		"guest": must.Nice(c.Auth.GuestToken),
+	}
+	authConfig := statickratosauth.NewConfig(routeScope, authTokens).
+		WithFieldName("Authorization").
+		WithSimpleEnable().
+		WithDebugMode(true)
+	return statickratosauth.NewMiddleware(authConfig, logger)
+}
+
+// NewUserMiddleware creates user auth middleware with database token validation
+//
+// NewUserMiddleware 创建用户认证中间件，通过数据库验证令牌
+func NewUserMiddleware(dataData *data.Data, logger log.Logger) middleware.Middleware {
+	routeScope := authkratos.NewInclude(
+		pb.OperationArticleServiceCreateArticle,
+		pb.OperationArticleServiceUpdateArticle,
+		pb.OperationArticleServiceDeleteArticle,
+		pb.OperationArticleServiceGetArticle,
+		pb.OperationArticleServiceListArticles,
+	)
+
+	checkAuthFunction := func(ctx context.Context, token string) (context.Context, *errors.Error) {
+		ctx, erk := dbauth.CheckToken(ctx, dataData.DB(), token)
+		if erk != nil {
+			return nil, erk
+		}
+		return ctx, nil
+	}
+	authConfig := customkratosauth.NewConfig(routeScope, checkAuthFunction).
+		WithFieldName("AdminToken").
+		WithDebugMode(true)
+	return customkratosauth.NewMiddleware(authConfig, logger)
 }
```

## internal/service/article.go (+27 -4)

```diff
@@ -2,27 +2,50 @@
 
 import (
 	"context"
+	"fmt"
 
+	"github.com/go-kratos/kratos/v2/log"
 	pb "github.com/yylego/kratos-examples/demo2kratos/api/article"
 	"github.com/yylego/kratos-examples/demo2kratos/internal/biz"
+	"github.com/yylego/kratos-examples/demo2kratos/internal/pkg/dbauth"
+	"github.com/yylego/kratos-static-auth/statickratosauth"
+	"github.com/yylego/must"
 )
 
 type ArticleService struct {
 	pb.UnimplementedArticleServiceServer
 
-	uc *biz.ArticleUsecase
+	uc  *biz.ArticleUsecase
+	log *log.Helper
 }
 
-func NewArticleService(uc *biz.ArticleUsecase) *ArticleService {
-	return &ArticleService{uc: uc}
+func NewArticleService(uc *biz.ArticleUsecase, logger log.Logger) *ArticleService {
+	return &ArticleService{uc: uc, log: log.NewHelper(logger)}
 }
 
 func (s *ArticleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
+	// Extract role name from config-based auth
+	//
+	// 从基于配置的认证中提取角色名
+	roleName, ok := statickratosauth.GetUsername(ctx)
+	must.True(ok)
+	must.Nice(roleName)
+	s.log.WithContext(ctx).Infof("CreateArticle roleName=%s", roleName)
+
+	// Extract user info from database-based auth
+	//
+	// 从基于数据库的认证中提取用户信息
+	authInfo, erk := dbauth.GetAuthInfo(ctx)
+	if erk != nil {
+		return nil, erk
+	}
+	s.log.WithContext(ctx).Infof("CreateArticle userName=%s", authInfo.Username)
+
 	v, ebz := s.uc.CreateArticle(ctx, nil)
 	if ebz != nil {
 		return nil, ebz.Erk
 	}
-	return &pb.CreateArticleReply{Article: &pb.ArticleInfo{Id: v.ID, Title: v.Title, Content: v.Content, StudentId: v.StudentID}}, nil
+	return &pb.CreateArticleReply{Article: &pb.ArticleInfo{Id: v.ID, Title: fmt.Sprintf("%s (role=%s, user=%s)", v.Title, roleName, authInfo.Username), Content: v.Content, StudentId: v.StudentID}}, nil
 }
 
 func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
```

