package controllers

import (
	"LvHm/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserIputInfo struct {
	Real_name string `json:"real_name"`
	Id_card   string `json:"id_card"`
}
type RespInfo struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}
type AuthController struct {
	beego.Controller
}

func (this *AuthController) RetData(Resp interface{}) {
	this.Data["json"] = Resp
	//beego.Info("Resp :", Resp)
	this.ServeJSON()
}

func (this *AuthController) AuthGet() {
	beego.Info("this is AuthGet*********")
	Resp := RespInfo{Errno: models.RECODE_OK, Errmsg: models.RecodeText(models.RECODE_OK)}
	defer this.RetData(&Resp)
	var user models.User
	user_id := this.GetSession("user_id")
	beego.Info("user_id:", user_id)
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("id", user_id).One(&user)
	if err != nil {
		beego.Info("数据库查询失败")
		Resp.Errno = models.RECODE_DBERR
		Resp.Errmsg = models.RecodeText(Resp.Errno)
		return
	}
	//beego.Info("数据库查询成功")
	Resp.Data = user
	//beego.Info("DATA:", Resp.Data)

	return
}
func (this *AuthController) AuthPost() {
	Resp := RespInfo{Errno: models.RECODE_OK, Errmsg: models.RecodeText(models.RECODE_OK)}
	defer this.RetData(&Resp)
	//从Session得到客户端的user_id
	user_id := this.GetSession("user_id")
	//beego.Info("user_id:", user_id)
	//得到客户的请求数据
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if user.Real_name == "" || user.Id_card == "" {
		Resp.Errno = models.RECODE_NODATA
		Resp.Errmsg = models.RecodeText(Resp.Errno)
		return
	}
	o := orm.NewOrm()
	_, err := o.QueryTable("user").Filter("id", user_id).Update(orm.Params{"real_name": user.Real_name, "id_card": user.Id_card})
	if err != nil {
		Resp.Errno = models.RECODE_DBERR
		Resp.Errmsg = models.RecodeText(Resp.Errno)
		return
	}
	//更新Session 中的real_name,id_card
	this.SetSession("real_name", user.Real_name)
	this.SetSession("id_card", user.Id_card)
	o.QueryTable("user").Filter("id", user_id).One(&user)
	//beego.Info("数据库更新成功")
	Resp.Data = user
	return
}
