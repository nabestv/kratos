package main

const (
	_tplAppToml = `
# This is a TOML document. Boom~
`

	_tplMySQLToml = `
[demo]
	addr = "127.0.0.1:3306"
	dsn = "{user}:{password}@tcp(127.0.0.1:3306)/{database}?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8"
	readDSN = ["{user}:{password}@tcp(127.0.0.2:3306)/{database}?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8","{user}:{password}@tcp(127.0.0.3:3306)/{database}?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8,utf8mb4"]
	active = 20
	idle = 10
	idleTimeout ="4h"
	queryTimeout = "200ms"
	execTimeout = "300ms"
	tranTimeout = "400ms"
`
	_tplMCToml = `
demoExpire = "24h"

[demo]
	name = "{{.Name}}"
	proto = "tcp"
	addr = "127.0.0.1:11211"
	active = 50
	idle = 10
	dialTimeout = "100ms"
	readTimeout = "200ms"
	writeTimeout = "300ms"
	idleTimeout = "80s"
`
	_tplRedisToml = `
demoExpire = "24h"

[demo]
	name = "{{.Name}}"
	proto = "tcp"
	addr = "127.0.0.1:6389"
	idle = 10
	active = 10
	dialTimeout = "1s"
	readTimeout = "1s"
	writeTimeout = "1s"
	idleTimeout = "10s"
`

	_tplHTTPToml = `
[server]
    addr = "0.0.0.0:8000"
    timeout = "1s"
`
	_tplGRPCToml = `
[server]
    addr = "0.0.0.0:9000"
    timeout = "1s"
`

	_tplChangeLog = `## {{.Name}}

### v1.0.0
1. Online function xxx
`
	_tplMain = `package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"{{.Name}}/internal/server/grpc"
	"{{.Name}}/internal/server/http"
	"{{.Name}}/internal/service"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
)

func main() {
	flag.Parse()
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("{{.Name}} start")
	svc := service.New()
	grpcSrv := grpc.New(svc)
	httpSrv := http.New(svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			defer cancel()
			grpcSrv.Shutdown(ctx)
			httpSrv.Shutdown(ctx)
			log.Info("{{.Name}} exit")
			svc.Close()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
`

	_tplContributors = `# Owner
{{.Owner}}

# Author

# Reviewer
`

	_tplDao = `package dao

import (
	"context"
	"time"

	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	xtime "github.com/bilibili/kratos/pkg/time"
)

// Dao dao.
type Dao struct {
	db          *sql.DB
	redis       *redis.Pool
	redisExpire int32
	mc          *memcache.Memcache
	mcExpire    int32
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// New new a dao and return.
func New() (dao *Dao) {
	var (
		dc struct {
			Demo *sql.Config
		}
		rc struct {
			Demo       *redis.Config
			DemoExpire xtime.Duration
		}
		mc struct {
			Demo       *memcache.Config
			DemoExpire xtime.Duration
		}
	)
	checkErr(paladin.Get("mysql.toml").UnmarshalTOML(&dc))
	checkErr(paladin.Get("redis.toml").UnmarshalTOML(&rc))
	checkErr(paladin.Get("memcache.toml").UnmarshalTOML(&mc))
	dao = &Dao{
		// mysql
		db: sql.NewMySQL(dc.Demo),
		// redis
		redis:       redis.NewPool(rc.Demo),
		redisExpire: int32(time.Duration(rc.DemoExpire) / time.Second),
		// memcache
		mc:       memcache.New(mc.Demo),
		mcExpire: int32(time.Duration(mc.DemoExpire) / time.Second),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.mc.Close()
	d.redis.Close()
	d.db.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	if err = d.pingMC(ctx); err != nil {
		return
	}
	if err = d.pingRedis(ctx); err != nil {
		return
	}
	return d.db.Ping(ctx)
}

func (d *Dao) pingMC(ctx context.Context) (err error) {
	if err = d.mc.Set(ctx, &memcache.Item{Key: "ping", Value: []byte("pong"), Expiration: 0}); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}

func (d *Dao) pingRedis(ctx context.Context) (err error) {
	conn := d.redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}
`
	_tplReadme = `# {{.Name}}

## Project Description
1.
`
	_tplService = `package service

import (
	"context"

	"{{.Name}}/internal/dao"
	"github.com/bilibili/kratos/pkg/conf/paladin"
)

// Service service.
type Service struct {
	ac  *paladin.Map
	dao *dao.Dao
}

// New new a service and return.
func New() (s *Service) {
	var ac = new(paladin.TOML)
	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}
	s = &Service{
		ac:  ac,
		dao: dao.New(),
	}
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
`

	_tplGPRCService = `package service

import (
	"context"
	"fmt"

	pb "{{.Name}}/api"
	"{{.Name}}/internal/dao"
	"github.com/bilibili/kratos/pkg/conf/paladin"

	"github.com/golang/protobuf/ptypes/empty"
)

// Service service.
type Service struct {
	ac  *paladin.Map
	dao *dao.Dao
}

// New new a service and return.
func New() (s *Service) {
	var ac = new(paladin.TOML)
	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}
	s = &Service{
		ac:  ac,
		dao: dao.New(),
	}
	return s
}

// SayHello grpc demo func.
func (s *Service) SayHello(ctx context.Context, req *pb.HelloReq) (reply *empty.Empty, err error) {
	reply = new(empty.Empty)
	fmt.Printf("hello %s", req.Name)
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
`
	_tplHTTPServer = `package http

import (
	"net/http"

	"{{.Name}}/internal/model"
	"{{.Name}}/internal/service"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

var (
	svc *service.Service
)

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine) {
	var (
		hc struct {
			Server *bm.ServerConfig
		}
	)
	if err := paladin.Get("http.toml").UnmarshalTOML(&hc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	svc = s
	engine = bm.DefaultServer(hc.Server)
	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/{{.Name}}")
	{
		g.GET("/start", howToStart)
	}
}

func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang is good!",
	}
	c.JSON(k, nil)
}

`
	_tplAPIProto = `// The proto file that defines the project API can describe both gRPC and HTTP API
// protobuf file reference:
// - https://developers.google.com/protocol-buffers/
// - TODO: URL to be added
// protobuf generates an HTTP tool:
// - TODO: URL to be added
// gRPC Golang Model:
// - TODO: URL to be added
// gRPC Golang Warden Gen:
// - TODO: URL to be added
// gRPC http debugging tool (no pb file required):
// - TODO: URL to be added
// grpc command line debugging tool (no pb file required):
// - TODO: URL to be added
Syntax = "proto3";

import "github.com/gogo/protobuf/gogoproto/gogo.proto";
import "google/protobuf/empty.proto";

// package Named Use {appid}.{version} The way, version Shaped like v1, v2 ..
package demo.service.v1;

// NOTE: Finally, please delete these useless comments (゜-゜)つロ 

option go_package = "api";
// do not generate getXXX() method 
option (gogoproto.goproto_getters_all) = false;

service Demo {
	rpc SayHello (HelloReq) returns (.google.protobuf.Empty);
}

message HelloReq {
	string name = 1 [(gogoproto.moretags)='form:"name" validate:"required"'];
}
`
	_tplAPIGenerate = `package api

// Generate gRPC code
// go:generate kratos tool kprotoc
`
	_tplModel = `package model

// Kratos hello kratos.
type Kratos struct {
	Hello string
}`
	_tplGoMod = `module {{.Name}}

go 1.12

require (
	github.com/bilibili/kratos v0.0.0-20190417043059-a08e0f771592
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.3.1
	golang.org/x/net v0.0.0-20190420063019-afa5a82059c6
	google.golang.org/grpc v1.20.1
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.26.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190123085648-057139ce5d2b
	golang.org/x/lint => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/net => github.com/golang/net v0.0.0-20190311183353-d8887717615a
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sync => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20180905080454-ebe1bf3edb33
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190328211700-ab21143f2384
	google.golang.org/appengine => github.com/golang/appengine v1.1.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc => github.com/grpc/grpc-go v1.20.1
)
`
	_tplGRPCServer = `package grpc

import (
	pb "{{.Name}}/api"
	"{{.Name}}/internal/service"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
)

// New new a grpc server.
func New(svc *service.Service) *warden.Server {
	var rc struct {
		Server *warden.ServerConfig
	}
	if err := paladin.Get("grpc.toml").UnmarshalTOML(&rc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	ws := warden.NewServer(rc.Server)
	pb.RegisterDemoServer(ws.Server(), svc)
	ws, err := ws.Start()
	if err != nil {
		panic(err)
	}
	return ws
}
`
)
