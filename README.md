# dbx

数据库客户端管理工具，基于 GORM 封装，支持 MySQL/PostgreSQL 多数据源。

## 安装

```bash
go get github.com/go-xuan/dbx
```

## 快速开始

在 `conf/database.yaml` 中配置：

```yaml
source: "default"
driver: "gorm"
enable: true
dialect: "mysql"
host: "127.0.0.1"
port: 3306
username: "root"
password: "root"
database: "demo"
maxOpenConns: 100
maxIdleConns: 10
logLevel: "warn"
slowThreshold: 200
```

```go
import "github.com/go-xuan/dbx"

func main() {
    dbx.Initialize()                    // 加载配置并初始化连接
    db := dbx.GetGormDB("default")      // 获取 *gorm.DB
    db.First(&user)
}
```

## 主要功能

- **多数据源** — 支持同时连接多个数据库，通过 source 名称区分
- **MySQL/PostgreSQL** — 内置两种方言，可扩展自定义 Builder
- **连接池管理** — MaxOpenConns / MaxIdleConns / MaxLifetime / MaxIdleTime
- **慢查询日志** — 可配置阈值，通过 logrus 输出
- **表结构迁移** — InitGormTable 自动建表、AutoMigrate、添加注释
- **配置驱动** — 配合 configx 自动从 nacos / 本地文件加载
