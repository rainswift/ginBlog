package main

import (
	"ginBlog/router"
)

func main() {
	//user := models.BlogUser{
	//	Username: "qq",
	//	Password: "123456",
	//}
	//dao.Mgr.AddUser(&user)
	router.Start()

	//name,_ := dao.Mgr.GetByName("qq")
	//fmt.Printf("% + v\n",name)

	//dao.Mgr.Cs_text()
}
