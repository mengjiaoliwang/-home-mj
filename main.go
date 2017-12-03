package main

import (
	"LvHm/models"
	_ "LvHm/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strings"
)

func init() {
	// set default database
	//绑定orm此时用的是哪个数据库的驱动
	//第四个参数表示 数据连接最大的空闲个数(可选)
	//第五个参数表示 数据库最大的链接个数(可选)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/LvHm?charset=utf8", 30)
	orm.RegisterModel(new(models.User), new(models.Area), new(models.Facility), new(models.House), new(models.HouseImage), new(models.OrderHouse))
	// create table
	//第二个参数表示是否强制替换
	//第三个表示 如果没有是否创建
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.SetStaticPath("/group1/M00", "fastdfs/storage_Data/data")
	ignoreStaticPath()
	beego.Run()
}
func ignoreStaticPath() {

	//透明static

	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	//如果请求uri还有api字段,说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)

	//将全部的静态资源重定向 加上/static/html路径
	//http://ip:port:8080/index.html----> http://ip:port:8080/static/html/index.html
	//如果restFUL api  那么就取消冲定向
	//http://ip:port:8080/api/v1.0/areas ---> http://ip:port:8080/static/html/api/v1.0/areas
}
