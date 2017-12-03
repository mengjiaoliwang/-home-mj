package controllers

import (
	"LvHm/models"
	"encoding/json"
	//	"fmt"
	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/cache"
	//	//	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	//	//	//	_ "github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	//	//	"time"
	//	"github.com/astaxie/beego/session"
)

type LogReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LogResp struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
}
type LogController struct {
	beego.Controller
}

func (this *LogController) RetData(resp interface{}) {
	//给客户端返回json数据
	this.Data["json"] = resp
	//将json写回客户端
	this.ServeJSON()

}

func (this *LogController) LogPost() {
	beego.Debug("this is LogPost********")
	resp := LogResp{Errno: models.RECODE_OK, Errmsg: models.RecodeText(models.RECODE_OK)}
	defer this.RetData(&resp)
	//获取用户的POST请求，得到用户信息
	var LogIn_Data LogReq
	json.Unmarshal(this.Ctx.Input.RequestBody, &LogIn_Data)
	beego.Info("用户的登录信息为：", LogIn_Data)
	//校验信息的合法性
	if LogIn_Data.Mobile == "" || LogIn_Data.Password == "" {
		beego.Info("用户登录失败,输入内容为空")
		resp.Errno = models.RECODE_NODATA
		resp.Errmsg = models.RecodeText(models.RECODE_NODATA)
		return
	}
	//查询数据库mysql
	o := orm.NewOrm()
	user := models.User{}
	if err := o.QueryTable("user").Filter("mobile", LogIn_Data.Mobile).One(&user); err == orm.ErrNoRows {
		beego.Info("查询mysql user 表失败,用户不存在")
		beego.Info("登录失败")
		resp.Errno = models.RECODE_DBERR

		resp.Errmsg = models.RecodeText(resp.Errno)
		return

	}
	if user.Password_hash != LogIn_Data.Password {
		beego.Info("查询mysql user 表失败,密码错误")
		beego.Info("登录失败")
		resp.Errno = models.RECODE_DBERR

		resp.Errmsg = models.RecodeText(resp.Errno)
		return

	}
	beego.Info("查询mysql user成功")

	//将用户信息存入session
	this.SetSession("user_id", user.Id)
	beego.Info("Log user_id :", user.Id)
	this.SetSession("name", user.Name)
	this.SetSession("mobile", user.Mobile)

	beego.Info("登录成功")

	return

}
