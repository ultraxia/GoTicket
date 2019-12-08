package general

import (
	"github.com/Unknwon/goconfig"
	"log"
	"path/filepath"
)

func TicketConfig() map[string]string {
	configPath, err := filepath.Abs("config.ini")
	if err != nil {
		log.Println(err)
	}

	cfg, err := goconfig.LoadConfigFile(configPath)
	if err != nil {
		log.Println(err)
		log.Println("***[ERROR] 初始化失败，请检查配置文件是否正常***")
	}

	sess, _ := cfg.GetValue("ticket", "sess")
	price, _ := cfg.GetValue("ticket", "price")
	date, _ := cfg.GetValue("ticket", "date")
	real_name, _ := cfg.GetValue("ticket", "real_name")
	nick_name, _ := cfg.GetValue("ticket", "nick_name")
	ticket_mum, _ := cfg.GetValue("ticket", "ticket_num")
	damai_url, _ := cfg.GetValue("ticket", "damai_url")
	target_url, _ := cfg.GetValue("ticket", "target_url")
	total_wait_time, _ := cfg.GetValue("ticket", "total_wait_time")
	refresh_wait_time, _ := cfg.GetValue("ticket", "refresh_wait_time")
	intersect_wait_time, _ := cfg.GetValue("ticket", "intersect_wait_time")

	config := make(map[string]string, 0)
	config["secc"] = sess
	config["price"] = price
	config["date"] = date
	config["real_name"] = real_name
	config["nick_name"] = nick_name
	config["ticket_num"] = ticket_mum
	config["damai_url"] = damai_url
	config["target_url"] = target_url
	config["total_wait_time"] = total_wait_time
	config["refresh_wait_time"] = refresh_wait_time
	config["intersect_wait_time"] = intersect_wait_time

	return config
}
