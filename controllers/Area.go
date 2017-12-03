package controllers

import (
	"LvHm/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	//	_ "github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type AreaResp struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}
type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(resp interface{}) {
	//给客户端返回json数据
	this.Data["json"] = resp
	//将json写回客户端
	this.ServeJSON()

}

func (this *AreaController) AreaGet() {
	beego.Debug("this is Area********")
	resp := AreaResp{Errno: models.RECODE_OK, Errmsg: models.RecodeText(models.RECODE_OK)}
	defer this.RetData(&resp)

	//从redis里查询数据，看是否有缓存
	con_redis, err := cache.NewCache("redis", `{"key":"LvHm","conn":"127.0.0.1:6379","dbNum":"0"}`)
	if err != nil {
		beego.Debug("connect redis server fail")
		resp.Errno = models.RECODE_DATAERR
		resp.Errmsg = models.RecodeText(resp.Errno)
		return

	}
	areas_info_value := con_redis.Get("areas_info")
	if areas_info_value != nil {
		//代表缓存有数据， 直接将数据返回
		beego.Debug("==== get area_info from cache =======")

		//将areas_info_value字符串变成 go的结构体
		var areas_info interface{}
		json.Unmarshal(areas_info_value.([]byte), &areas_info)
		resp.Data = areas_info
		return

	}

	//从数据库里面查询area数据是否存在
	o := orm.NewOrm()
	var areas []models.Area
	qs := o.QueryTable("area")
	num, err := qs.All(&areas)
	if err != nil {
		resp.Errno = models.RECODE_DATAERR
		resp.Errmsg = models.RecodeText(resp.Errno)
		return
	}
	if num == 0 {
		//没有数据
		resp.Errno = models.RECODE_NODATA
		resp.Errmsg = models.RecodeText(resp.Errno)
		return
	}
	fmt.Printf("areas =%+v\n", areas)
	resp.Data = areas

	//将数据存储到缓存数据库
	areas_info_str, _ := json.Marshal(areas)
	if err := con_redis.Put("areas_info", areas_info_str, 3600*time.Second); err != nil {
		beego.Debug("set areas_info_str conn_cach error,err=", err)
		resp.Errno = models.RECODE_NODATA
		resp.Errmsg = models.RecodeText(resp.Errno)
		return
	}
	return

	/*	this.Data["json"] = Resp
		this.ServeJSON()*/
}
