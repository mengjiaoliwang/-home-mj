package controllers

//beego.Router("api/v1.0/user/houses", &controllers.GetReleaseHousesController{}, "Get:GetReleaseHouses")
import (
	"LvHm/models"
	"github.com/astaxie/beego"
)

type GetReleaseHousesController struct {
	beego.Controller
}
type GetHousesInfoResp struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func (this *GetReleaseHousesController) RetData(Resp interface{}) {
	this.Data["json"] = Resp
	this.ServeJSON()
}
func (this *GetReleaseHousesController) GetReleaseHouses() {
	Resp := GetHousesInfoResp{Errno: models.RECODE_NODATA, Errmsg: models.RecodeText(models.RECODE_NODATA)}
	defer this.RetData(&Resp)
}
