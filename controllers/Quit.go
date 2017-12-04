package controllers

//beego.Router("api/v1.0/session", &controllers.QuitController{}, "delete:QuitDelete")
import (
	"LvHm/models"
	"github.com/astaxie/beego"
)

type QuitRespInfo struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
}
type QuitController struct {
	beego.Controller
}

func (this *QuitController) RetData(Resp interface{}) {
	this.Data["json"] = Resp
	this.ServeJSON()
}

func (this *QuitController) QuitDelete() {
	Resp := QuitRespInfo{Errno: models.RECODE_OK, Errmsg: models.RecodeText(models.RECODE_OK)}

	defer this.RetData(&Resp)
	this.DelSession("user_id")
	this.DelSession("name")
	//	beego.Info("用户退出")
	return
}
