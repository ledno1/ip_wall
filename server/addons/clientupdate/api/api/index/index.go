// Package index
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package index

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/clientupdate/model/input/sysin"
)

// TestReq 测试
type TestReq struct {
	g.Meta `path:"/index/test" method:"get" tags:"客户端更新" summary:"测试前台API"`
	sysin.IndexTestInp
}

type TestRes struct {
	*sysin.IndexTestModel
}

// TestReq 测试
type GetUpdateReq struct {
	g.Meta `path:"/index/getupdate" method:"post" tags:"客户端更新" summary:"获取是否需要更新和更新下载地址"`
	sysin.GetUpdateInp
}

type GetUpdateRes struct {
	*sysin.GetUpdateModel
}