## 客户端更新

### 简介

配套的客户端更新接口更新包地址


### 使用说明




### 迁移或安装

1、安装 HotGo (2.9.3及以上)

项目介绍：https://github.com/bufanyun/hotgo

2、将当前插件项目拷贝进 HotGo 根目录的 server/addons 目录下

3、在 HotGo 根目录的 server/addons/modules 目录下创建go文件:clientupdate.go，内容如下：
```go
package modules

import _ "hotgo/addons/clientupdate"
```

4、HotGo 后台进入 开发工具->插件管理->找到 客户端更新 (clientupdate) 进行安装

5、重启服务即可生效


### 常用命令行

```shell
# 接口维护-gen service
gf gen service -s=addons/clientupdate/logic -d=addons/clientupdate/service

```
