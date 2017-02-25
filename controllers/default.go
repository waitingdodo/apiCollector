package controllers

import (
	"github.com/astaxie/beego"
	"time"
	//"strings"
	//"encoding/json"
	//"bytes"
	//"math"
	"strconv"
	"os"
)

type MainController struct {
	beego.Controller
	string
}
var lastTime int64=0
var timeLimit=int64(time.Millisecond*1000)
const minTimeLimitMilliSecond=100
const myStopPwd=953503192007

func (c *MainController) Get() {//默认get方法
	if time.Now().UnixNano()-lastTime<timeLimit{//允许一秒访问一次
		//c.Ctx.Abort(500,"service pause!")
		c.Ctx.Output.Body([]byte(""))
		return
	}
	lastTime=time.Now().UnixNano()
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.Ctx.Output.Body([]byte("visit successful!"))
}




func (c *MainController) StopService() {//关闭应用

	if time.Now().UnixNano()-lastTime<timeLimit{//允许一秒访问一次
		//c.Ctx.Abort(500,"service pause!")
		c.Ctx.Output.Body([]byte(""))
		return
	}
	pwd,err:=c.GetInt("pwd",1000)
	if err!=nil || pwd!=myStopPwd{
		c.Ctx.Output.Body([]byte(""))
		return
	}


	c.Ctx.Output.Body([]byte("stop service successful!"))

	go stopApp()

	//bytes.
	//json.Marshal()
}

func stopApp(){
	time.Sleep(time.Second*5)
	os.Exit(-1)
}

func (c *MainController) UpdateTimeLimit() {//修改间隔时间

	if time.Now().UnixNano()-lastTime<timeLimit{//允许一秒访问一次
		//c.Ctx.Abort(500,"service pause!")
		c.Ctx.Output.Body([]byte(""))
		//c.StopRun()
		beego.Error("time limit  "+strconv.FormatInt(timeLimit,10))
		return
	}

	newTimeLimt,err:=c.GetInt("timeLimit",1000)
	if err!=nil{
		c.Ctx.Output.Body([]byte(""))
		return
	}

	lastTime=time.Now().UnixNano()

	if newTimeLimt<minTimeLimitMilliSecond{//防止时间间隔过小
		newTimeLimt=minTimeLimitMilliSecond
	}
	timeLimit=int64(time.Millisecond)*int64(newTimeLimt)

	c.Ctx.Output.Body([]byte("UpdateTimeLimit Successful! \n now time limit："+strconv.FormatInt(timeLimit,10)  +" millsSecond "))

	//bytes.
	//json.Marshal()
}
