package routers

import (
	"DataCertPaltPhone/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //用户注册
    beego.Router("/register",&controllers.ResgiterController{})
    //用户登录的接口
    beego.Router("/login",&controllers.LoginController{})
	//用户上传文件的功能
    beego.Router("/home",&controllers.HomeController{})
    //查看认证数据证书
    beego.Router("/cert_detail.html",&controllers.CertDetailController{})
    //用户实名认证
    beego.Router("/user_kyc",&controllers.UserKycController{})
}
