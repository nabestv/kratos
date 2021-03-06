# Quick Start

To quickly use the kratos project, you can use the `kratos tool` as follows:

```shell
go get -u github.com/bilibili/kratos/tool/kratos
kratos init
```
Follow the prompts to quickly create a project, such as [kratos-demo](https://github.com/bilibili/kratos-demo) which is created by the tool. The directory structure is as follows:

```
├── CHANGELOG.md # CHANGELOG
├── CONTRIBUTORS.md # CONTRIBUTORS
├── README.md # README
├── api # api directory is a foreign reserved proto file, and the generated pb.go file
│ ├── api.proto
│ ├── api.pb.go # pb.go file generated by go generate
│ └── generate.go
├── cmd # cmd directory is main
│ └── main.go # main.go
├── configs # configs is the configuration file directory
│ ├── application.toml # Application's custom configuration file, which may be some business switches such as: useABtest = true
│ ├── grpc.toml # grpc related configuration
│ ├── http.toml # httprelated configuration
│ ├── log.toml #log related configuration
│ ├── memcache.toml # memcache related configuration
│ ├── mysql.toml # mysql related configuration
│ └── redis.toml # redis related configuration
├── go.mod # go.mod
└── internal # internal is the project internal package, including the following directories:
    ├── dao # dao layer, used for database, cache, MQ, dependent on a business grpc|http and other resource access
    │ └── dao.go
    ├── model #model layer, used to declare the business structure
    │ └── model.go
    ├── server # server layer, used to initialize grpc and http server
    │ └── http # http layer, used to initialize http server and declare handler
    │ └── http.go
    │ └── grpc # grpc layer, used to initialize grpc server and define method
    │ └── grpc.go
    └── service # service layer, used for business logic processing, and for the convenience of http and grpc sharing methods, it is recommended to enter the parameters and the parameters to maintain grpc style, and use pb file to generate code
        └── service.go
```

Once generated, it can be run directly as follows:

```shell
cd kratos-demo/cmd
go build
./cmd -conf ../configs
```

Open your browser and visit: [http://localhost:8000/kratos-demo/start](http://localhost:8000/kratos-demo/start), you will see the output of `Golang Dafa! ! ! `

# Map

![kratos init](/doc/img/kratosinit.gif)

-------------

[document directory tree](summary.md)
