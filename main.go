package main

import (
	"GoTicket/general"
	"encoding/json"
	"fmt"
	"github.com/tebeka/selenium"
	"math"
	"os"
	"strings"
	"time"
)

var config map[string]string
var webDriver selenium.WebDriver

func init() {
	config = general.TicketConfig()
	webDriver = general.Driver()
}

func main() {
	get_cookie()
	set_cookie()
	//order_ticket()
}

func login() {

}

func get_cookie() {
	var web_title string

	if err := webDriver.Get(config["damai_url"]); err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}

	fmt.Println("### 请点击登录 ###")

	for {
		web_title, _ = webDriver.Title()
		if strings.Contains(web_title, "大麦网-全球演出赛事官方购票平台") == true {
			time.Sleep(1)
		} else {
			break
		}
	}

	fmt.Println("### 请选择扫码登陆 ###")

	for {
		web_title, _ = webDriver.Title()
		if web_title == "大麦登录" || strings.Contains(web_title, "大麦网-全球演出赛事官方购票平台") != true {
			time.Sleep(1)
		} else {
			break
		}

	}

	cookie, err := webDriver.GetCookies()

	if err != nil {
		fmt.Printf("Cookie获取失败\n%v", err)
	}

	general.Writefile(cookie)
	fmt.Println("### 登录成功 ###")
}

//TODO 自动读取cookie
func set_cookie() {
	var cookies []selenium.Cookie

	filePtr, err := os.Open("cookies.pkl")
	if err != nil {
		fmt.Println(err)
	}
	defer filePtr.Close()
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&cookies)
	if err != nil {
		fmt.Println(err)
	}
	co,_ := webDriver.GetCookies()
	fmt.Println(co)

	for _, v := range cookies {
		session := &selenium.Cookie{
			Name: v.Name,
			Value: v.Value,
			Expiry: math.MaxUint32,
		}
		session.Domain = v.Domain
		session.Path = "/"
		err = webDriver.AddCookie(session)

	}
	webDriver.Refresh()
	co,_ = webDriver.GetCookies()
	//fmt.Println(cookie)
	fmt.Println(co)


}

func order_ticket()  {
	webDriver.Get(config["target_url"])
	webDriver.Refresh()

}