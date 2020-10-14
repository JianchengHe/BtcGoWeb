package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)
/*
用于处理文件上传的功能
 */
type HomeController struct {
	beego.Controller
}
/*
该post方法用于文件上传
 */
func (h *HomeController) Post(){
	h.TplName = "home.html"
	/*
	1、解析用户上传的数据
	 */
	//用户上传的自定义的标题
	title := h.Ctx.Request.PostFormValue("home_title")//用户输入标题
	//用于用户上传的文件
	file,header,err := h.GetFile("hejiancheng")
	if err != nil {//解析客户端提交的文件出现的错误
		h.Ctx.WriteString("抱歉，文件解析失败，请重试")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("自定义的标题：",title)
	//获得到上传的文件
	fmt.Println("上传的文件名称：",header.Filename)
	//eg:支持jpg,png类型，不支持jpeg，gif 类型
	//文件名
	fileNameSlice := strings.Split(header.Filename,".")
	fileType := fileNameSlice[1]
	if fileType != "jpg" || fileType != "png"{
		//文件类型不支持
		h.Ctx.WriteString("抱歉，文件类型不符合，请上传符合格式的文件")
		return
	}
	//文件大小，200kb
	config := beego.AppConfig
	fileSize,err := config.Int64("file_size")
	if header.Size / 1024 > fileSize {
		h.Ctx.WriteString("抱歉，文件大小超出范围，请上传符合要求的文件")
		return
	}
	fmt.Println("文件的大小：",header.Size)//字节大小
	fmt.Println("上传的文件",file)
	h.Ctx.WriteString("已经获取到上传文件。")
}