package controllers

import (
	"DataCertPaltPhone/blockchain"
	"DataCertPaltPhone/models"
	"DataCertPaltPhone/utils"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"time"
)

/*
用于处理文件上传的功能
*/
type HomeController struct {
	beego.Controller
}

/*
该post方法用于处理用户在客户端提交的文件
*/
func (h *HomeController) Post() {
	//1、解析客户端提交的数据和文件
	phone := h.Ctx.Request.PostFormValue("phone")      //获取用户phone的信息
	title := h.Ctx.Request.PostFormValue("home_title") //用户输入标题
	fmt.Println(title)
	file, header, err := h.GetFile("hejiancheng") //封装好，到下面使用
	if err != nil {                               //解析客户端提交的文件出现的错误
		h.Ctx.WriteString("抱歉，文件解析失败，请重试")
		return
	}
	defer file.Close() //延迟执行   空指针错误：invalid memorey or nil pointer dereferenece
	//2、调用工具函数保存文件到本地
	saveFilePath := "static/home/" + header.Filename
	_, err = utils.SavaFile(saveFilePath, file)
	if err != nil {
		h.Ctx.WriteString("抱歉，文件认证失败，请重试")
		return
	}
	//3、计算文件的SHA256值
	fileHash, err := utils.Sha256Reader(file)
	fmt.Println(fileHash)
	//先查询用户id
	user, err := models.User{Phone: phone}.QueryUserByPhone()
	if err != nil {
		fmt.Println("查询用户", err.Error())
		h.Ctx.WriteString("抱歉，电子数据认证失败，请稍后再试")
		return
	}
	//把上传的文件作为记录保存到数据库当中
	//①计算文件的md5值
	saveFile,err:= os.Open(saveFilePath)
	Md5String, err := utils.Md5HashReader(saveFile)
	if err != nil {
		h.Ctx.WriteString("抱歉。数据认证失败，请重试")
	}
	record := models.HomeRecord{
		UserId:    user.Id,
		FileName:  header.Filename,
		FileSize:  header.Size,
		FileCert:  Md5String,
		FileTitle: title,
		CertTime:  time.Now().Unix(),
	}
	//②保存认证数据到数据库中
	_, err = record.SavaRecord()
	if err != nil {
		fmt.Println(err.Error())
		h.Ctx.WriteString("抱歉，文件认证保存失败，请稍后再试")
		return
	}
	//③经用户上传的文件的md5值和sha256值保存到区块链上，即数据上链
	blockchain.CHAIN.SaveData([]byte(fileHash))
	//上传文件保存到数据库成功
	records, err := models.QueryRecordsByUserId(user.Id)
	if err != nil {
		h.Ctx.WriteString("抱歉，获取电子数据列表失败，请重新尝试")
		return
	}
	h.Data["Records"] = records
	h.TplName = "list_record.html"
}

/*
该post方法用于文件上传
*/

//func (h *HomeController) Post() {
//	h.TplName = "home.html"
//	/*
//		1、解析用户上传的数据
//	*/
//	//用户上传的自定义的标题
//	title := h.Ctx.Request.PostFormValue("home_title") //用户输入标题
//	//用于用户上传的文件
//	file, header, err := h.GetFile("hejiancheng") //封装好，到下面使用
//	if err != nil {                               //解析客户端提交的文件出现的错误
//		h.Ctx.WriteString("抱歉，文件解析失败，请重试")
//		fmt.Println(err.Error())
//		return
//	}
//	defer file.Close()
//	fmt.Println("自定义的标题：", title)
//	//获得到上传的文件
//	fmt.Println("上传的文件名称：", header.Filename)
//	//eg:支持jpg,png类型，不支持jpeg，gif 类型
//	//文件名
//	fileNameSlice := strings.Split(header.Filename, ".")
//	fileType := fileNameSlice[1]
//	fmt.Println(".", )
//	if fileType != "jpg" && fileType != "png" {
//		//文件类型不支持
//		h.Ctx.WriteString("抱歉，文件类型不符合，请上传符合格式的文件")
//		return
//	}
//	//文件大小，200kb
//	config := beego.AppConfig
//	fileSize, err := config.Int64("file_size")
//	if header.Size/1024 > fileSize {
//		h.Ctx.WriteString("抱歉，文件大小超出范围，请上传符合要求的文件")
//		return
//	}
//	fmt.Println("文件的大小：", header.Size) //字节大小
//	//perm : permission  权限
//	//权限的组成：a+b+c
//	//a:文件所有者对文件的操作权限    读4、写2、执行2
//	//b:文件所有者所在组的用户的操作权限   读4、写2、执行2
//	//c:其他用户的操作权限  读4、写2、执行2
//	//eg :m文件，权限：451
//	//判断题  文件所有者对该m文件有写权限（对）
//	//文件的所有组所在组用户对该文件有写权限（错误）
//	savaDir := "static/home"
//	//①打开文件
//	_, err = os.Open(savaDir)
//	os.OpenFile("文件名", os.O_CREATE|os.O_RDONLY, 777)
//	if err != nil {
//		//②创建文件夹
//		err = os.Mkdir("static/home", 777)
//		if err != nil {
//			h.Ctx.WriteString("抱歉，文件认证遇到错误，请重试")
//			return
//		}
//	}
//	//fmt.Println(f)
//	//文件名：文件路径 + 文件名 +"."+文件扩展名
//	savaName := savaDir + "/" + header.Filename // static/home/xxx.jpg
//	fmt.Println("saveName", savaName)
//	fmt.Println("要保存的文件名：", savaName)
//	//fromFile：文件，
//	//toFile:要保存的文件路径
//	err = h.SaveToFile("hejiancheng", savaName)
//	if err != nil {
//		h.Ctx.WriteString("抱歉，文件认证失败，请重试")
//		return
//	}
//	fmt.Println("上传的文件", file)
//	h.Ctx.WriteString("已经获取到上传文件。")
//
//}
