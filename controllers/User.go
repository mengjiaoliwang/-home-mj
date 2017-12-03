package controllers

import (
	"LvHm/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/cache"
	//	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	//	//	_ "github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	//	"time"
	//	"github.com/astaxie/beego/session"
)

type RegReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Sms_code string `json:"sms_code"`
}

type RegResp struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
}
type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp interface{}) {
	//给客户端返回json数据
	this.Data["json"] = resp
	//将json写回客户端
	this.ServeJSON()

}

func (this *UserController) RegPost() {
	beego.Debug("this is RegPost********")
	resp := RegResp{Errno: models.RECODE_OK, Errmsg: models.RecodeText(models.RECODE_OK)}
	defer this.RetData(&resp)
	//获取用户的POST的请求
	var RegReq_Data RegReq
	json.Unmarshal(this.Ctx.Input.RequestBody, &RegReq_Data)
	fmt.Printf("请求的数据为:%+v\n", RegReq_Data)
	//校验信息
	if RegReq_Data.Mobile == "" || RegReq_Data.Password == "" || RegReq_Data.Sms_code == "" {
		resp.Errno = models.RECODE_NODATA
		resp.Errmsg = models.RecodeText(resp.Errno)
		return
	}
	//对短信进行校验

	//将用户信息存储到mysql中
	user := models.User{}
	user.Mobile = RegReq_Data.Mobile
	user.Password_hash = RegReq_Data.Password
	user.Name = RegReq_Data.Mobile
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err != nil {
		beego.Debug("Reg Insert error,err=", err)
		resp.Errno = models.RECODE_NODATA
		resp.Errmsg = models.RecodeText(resp.Errno)
		return

	}
	beego.Info("Reg insert succ Id =", id)
	//将用户名存入session
	this.SetSession("user_id", user.Id)
	this.SetSession("name", user.Name)
	this.SetSession("mobile", user.Mobile)

	return
}
