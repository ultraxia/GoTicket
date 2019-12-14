package main

import (
	"GoTicket/general"
	"encoding/json"
	"fmt"
	"github.com/tebeka/selenium"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var cookie = "cookies.pkl"
var config map[string]string
var webDriver selenium.WebDriver

func init() {
	config = general.TicketConfig()
	webDriver = general.Driver()
}

func main() {
	login()
	order_ticket()
}

func login() {
	if exist(cookie) == false {
		fmt.Println("### 未找到cookies，正在调用登录组件 ###")
		get_cookie()
	} else {
		set_cookie()
	}
}

func exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
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

func set_cookie() {
	webDriver.Get(config["target_url"])
	var cookies []selenium.Cookie

	filePtr, err := os.Open(cookie)
	if err != nil {
		fmt.Println(err)
	}
	defer filePtr.Close()
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&cookies)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range cookies {
		session := &selenium.Cookie{
			Name:   v.Name,
			Value:  v.Value,
			Expiry: math.MaxUint32,
		}
		session.Domain = v.Domain
		session.Path = "/"
		session.Secure = false

		err = webDriver.AddCookie(session)

	}

}

func order_ticket() {
	webDriver.Get(config["target_url"])
	//startTime := time.Now()
	//fmt.Println(startTime)

	fmt.Println("### 正在选择日期与票价 ###")
	num := 0
	//for {
	num += 1
	web_title, _ := webDriver.Title()
	if strings.Contains(web_title, "确认订单") == true {
		fmt.Printf("经过%v轮的努力，恭喜您成功抢到票", num)
	} else {
		if config["date"] != "0" || len(config["date"]) == 0 {
			select_date()
		} else {
			fmt.Println("无需选择场次")
		}
	}
	//}

}

//选择场次
func select_date() {
	datepickers, err := webDriver.FindElement(selenium.ByClassName, "perform__order__box")
	if err != nil {
		fmt.Println(err)
	}
	orderElements, err := datepickers.FindElements(selenium.ByClassName, "select_right_list")
	dateElements := orderElements[1]
	dateList, err := dateElements.FindElements(selenium.ByTagName, "span")
	if err != nil {
		fmt.Println(nil)
	}
	dateConfig, _ := strconv.Atoi(config["date"])
	err = dateList[dateConfig].Click()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("场次选择成功")
	}

}
