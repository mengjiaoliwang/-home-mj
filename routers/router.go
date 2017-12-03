package routers

import (
	"LvHm/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//请求地域信息
	beego.Router("/api/v1.0/areas", &controllers.AreaController{}, "get:AreaGet")
	//请求用户Session信息
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:SessionGet")
	//注册用户
	beego.Router("/api/v1.0/users", &controllers.UserController{}, "post:RegPost")
	//用户登陆
	beego.Router("/api/v1.0/sessions", &controllers.LogController{}, "post:LogPost")
	//获取用户信息
	beego.Router("/api/v1.0/user", &controllers.UserInfoController{}, "get:UserInfoGet")
	//上传头像
	beego.Router("api/v1.0/user/avatar", &controllers.UserCorrectPostController{}, "post:UserCorrectPost")
	//更新用户名
	beego.Router("api/v1.0/user/name", &controllers.UserCorrectPostController{}, "put:UserCorrectName")
	//请求查看房东/租客订单信息
	beego.Router("api/v1.0/user/orders", &controllers.GetOrderController{}, "get:GetOrder")
	//发布订单
	//beego.Router("api/v1.0/orders", &controllers.PostOrderReleaseController{}, "post:PostOrderRelease")
	//发布房源信息
	beego.Router("api/v1.0/houses", &controllers.HousesController{}, "post:ReleaseHouses")

}
