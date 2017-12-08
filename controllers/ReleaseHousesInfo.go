package controllers

//beego.Router("api/v1.0/houses", &controllers.ReleaseHousesController{}, "post:ReleaseHouses")

import (
	"LvHm/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type HouseReleaseInfo struct {
	Title      string   `json:"title"`
	Price      string   `json:"price"`
	Area_id    string   `json:"area_id"`
	Address    string   `json:'address"`
	Room_count string   `json:"room_count"`
	Acreage    string   `json:"acreage"`
	Unit       string   `json:"unit"`
	Capacity   string   `json:"capacity"`
	Beds       string   `json:'beds"`
	Deposit    string   `json:"deposit"`
	Min_days   string   `json:"min_days"`
	Max_days   string   `json:"max_days"`
	Facility   []string `json:"facility"`
}
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
	var houseReleaseInfo HouseReleaseInfo
	json.Unmarshal(this.Ctx.Input.RequestBody, &houseReleaseInfo)
	beego.Info("输入的信息为：", houseReleaseInfo)
	var house models.House
	house.Acreage, _ = strconv.Atoi(houseReleaseInfo.Acreage)
	house.Address = houseReleaseInfo.Address
	house.Beds = houseReleaseInfo.Beds
	house.Capacity, _ = strconv.Atoi(houseReleaseInfo.Capacity)
	house.Deposit, _ = strconv.Atoi(houseReleaseInfo.Deposit)
	house.Max_days, _ = strconv.Atoi(houseReleaseInfo.Max_days)
	house.Min_days, _ = strconv.Atoi(houseReleaseInfo.Min_days)
	house.Price, _ = strconv.Atoi(houseReleaseInfo.Price)
	house.Room_count, _ = strconv.Atoi(houseReleaseInfo.Room_count)
	house.Title = houseReleaseInfo.Title
	house.Unit = houseReleaseInfo.Unit
	area_id, _ := strconv.Atoi(houseReleaseInfo.Area_id)
	house.Area = &models.Area{Id: area_id}
	house.User = &models.User{Id: this.GetSession("user_id").(int)}
	//onetoone
	o := orm.NewOrm()
	houseId, err := o.Insert(&house)
	if err != nil {
		Resp.Errno = models.RECODE_DBERR
		Resp.Errmsg = models.RecodeText(Resp.Errno)
		beego.Info("插入数据错误")
		return
	}
	//moretomore
	m2m := o.QueryM2M(&house, "Facilities")

	//讲用户输入的数据变成数据可以识别的格式，讲[]string变成model.Facility
	var facilities []*models.Facility

	for _, fid := range houseReleaseInfo.Facility {
		fcl_id, _ := strconv.Atoi(fid)
		facityone := &models.Facility{Id: fcl_id}
		facilities = append(facilities, facityone)
	}

	num, err := m2m.Add(facilities)
	if err != nil {
		Resp.Errno = models.RECODE_DBERR
		Resp.Errmsg = models.RecodeText(Resp.Errno)
		beego.Info("facility数据插入错误,num=", num)
		return
	}
	Resp.Data = House_Id{House_id: strconv.FormatInt(houseId, 10)}
	return
}
