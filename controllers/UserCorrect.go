package controllers

import (
	"LvHm/models"
	"encoding/json"
	//	"fmt"
	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/cache"
	//	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	//	//	_ "github.com/garyburd/redigo/redis"
	//	_ "github.com/go-sql-driver/mysql"
	//	"time"
	"path"
)

type UserCorrectName_Data struct {
	Name string `json:"name"`
}
type Avatar struct {
	Url  string `json:"avatar_url"`
	Name string `json:"name"`
}
type UserCorrectPostResp struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
	Data   Avatar `json:"data"`
}
type UserCorrectPostController struct {
	beego.Controller
}

func (this *UserCorrectPostController) RetData(resp interface{}) {
	//给客户端返回json数据
	this.Data["json"] = resp
	//将json写回客户端
	this.ServeJSON()

}

func (this *UserCorrectPostController) UserCorrectPost() {
	beego.Debug("this is UserCorrectPost********")
	Resp := UserCorrectPostResp{Errno: models.RECODE_OK, Errmsg: models.RecodeText(models.RECODE_OK)}
	defer this.RetData(&Resp)

	file, header, err := this.GetFile("avatar")

	if err != nil {
		Resp.Errno = models.RECODE_SERVERERR
		Resp.Errmsg = models.RecodeText(Resp.Errno)
		beego.Info("get file error")
		return
	}
	defer file.Close()
	//创建一个文件的缓冲
	fileBuffer := make([]byte, header.Size)

	_, err = file.Read(fileBuffer)
	if err != nil {
		Resp.Errno = models.RECODE_IOERR
		Resp.Errmsg = models.RecodeText(Resp.Errno)
		beego.Info("read file error")
		return
	}

	//home1.jpg
	suffix := path.Ext(header.Filename) // suffix = ".jpg"
	groupName, fileId, err1 := models.FDFSUploadByBuffer(fileBuffer, suffix[1:])
	if err1 != nil {
		Resp.Errno = models.RECODE_IOERR
		Resp.Errmsg = models.RecodeText(Resp.Errno)
		beego.Info("fdfs upload  file error")
		return
	}

	beego.Info("groupname,", groupName, " file id ", fileId)

	//通过session得到当前用户
	user_id := this.GetSession("user_id")
	beego.Info(user_id)

	//添加Avatar_url字段到数据库中
	o := orm.NewOrm()
	user := models.User{Id: user_id.(int), Avatar_url: fileId}

	if _, err := o.Update(&user, "avatar_url"); err != nil {
		Resp.Errno = models.RECODE_DBERR
		Resp.Errmsg = models.RecodeText(Resp.Errno)
		return
	}

	//拼接一个完整的路径
	avatar_url := "http://192.168.237.132:9000/" + fileId
	beego.Info(avatar_url)
	Resp.Data.Url = avatar_url
	beego.Info(Resp.Data.Url)
	return

}
func (this *UserCorrectPostController) UserCorrectName() {
	beego.Debug("this is UserCorrectName********")
	Resp := UserCorrectPostResp{Errno: models.RECODE_OK, Errmsg: models.RecodeText(models.RECODE_OK)}
	defer this.RetData(&Resp)

	//得到用户的请求数据name
	var cor_name_data UserCorrectName_Data
	json.Unmarshal(this.Ctx.Input.RequestBody, &cor_name_data)
	beego.Info("cor_name_data: ", cor_name_data)
	if cor_name_data.Name == "" {
		beego.Info("修改用户名为空，请填入信息")
		Resp.Errno = models.RECODE_NODATA
		Resp.Errmsg = models.RecodeText(models.RECODE_NODATA)
		return

	}
	//从session得到user_id
	user_id := this.GetSession("user_id")
	//去数据库查找信息
	o := orm.NewOrm()
	user := models.User{Id: user_id.(int), Name: cor_name_data.Name}
	//更新数据库
	if _, err := o.Update(&user, "name"); err != nil {
		Resp.Errno = models.RECODE_DBERR
		Resp.Errmsg = models.RecodeText(Resp.Errno)
		return
	}
	//更新session中的user_id,name
	this.SetSession("user_id", user.Id)
	this.SetSession("name", user.Name)

	//拼接一个完整的路径

	Resp.Data.Name = user.Name
	beego.Info(Resp.Data.Name)

	return
}
