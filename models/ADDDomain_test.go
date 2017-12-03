package models

//import "LvHm/models"
import "testing"
import "github.com/astaxie/beego"

func Test_ADd(t *testing.T) {
	ret := AddDomain2Url("http")
	beego.Info("返回值为：", ret)

}
