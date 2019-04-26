#### database/tidb

##### Project Description
TiDB database driver wraps the mysql driver

##### Features
1. Support discovery service discovery Multi-node direct connection
2. Support connection through lvs single address
3. Support prepare to bind multiple nodes
4. Support dynamic increase and decrease node load balancing
5. Log distinguishes running nodes

##### Dependency package
1.[Go-MySQL-Driver](https://github.com/go-sql-driver/mysql)