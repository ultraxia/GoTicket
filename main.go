package main

import (
	"GoTicket/general"
	"fmt"
	"github.com/tebeka/selenium"
	"log"
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
		log.Printf("Cookie获取失败\n%v", err)
	}

	general.Writefile(cookie[0])
	fmt.Println("### 登录成功 ###")
}
