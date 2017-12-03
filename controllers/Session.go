package controllers

import (
	"LvHm/models"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session"
)

type SessionResp struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}
type Name struct {
	Name string `json:"name"`
}
type SessionController struct {
	beego.Controller
}

func (this *SessionController) RetData(Resp interface{}) {
	this.Data["json"] = Resp
	this.ServeJSON()

}

func (this *SessionController) SessionGet() {
	beego.Debug("this is Session********")
	//name := Name{Name: "张3"}
	Resp := SessionResp{Errno: models.RECODE_OK, Errmsg: models.RecodeText(models.RECODE_OK)}
	defer this.RetData(&Resp)
	name := this.GetSession("name")
	if name == nil {
		Resp.Errno = models.RECODE_SESSIONERR
		Resp.Errmsg = models.RecodeText(Resp.Errno)
		return
	}
	NameData := Name{Name: name.(string)}
	Resp.Data = NameData

}
