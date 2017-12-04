package controllers

//beego.Router("api/v1.0/houses", &controllers.ReleaseHousesController{}, "post:ReleaseHouses")

import (
	"LvHm/models"
	"github.com/astaxie/beego"
)

type ReleaseHousesController struct {
	beego.Controller
}
type House_Id struct {
	House_id string `json:"house_id'`
}
type HousesInfoResp struct {
	Errno  string   `json:"errno"`
	Errmsg string   `json:"errmsg"`
	Data   House_Id `json:"data"`
}

func (this *ReleaseHousesController) RetData(Resp interface{}) {
	this.Data["json"] = Resp
	this.ServeJSON()
}
func (this *ReleaseHousesController) ReleaseHouses() {
	Resp := HousesInfoResp{Errno: models.RECODE_OK, Errmsg: models.RecodeText(models.RECODE_OK)}
	defer this.RetData(&Resp)
}
