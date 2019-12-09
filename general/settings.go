package general

import (
	"gopkg.in/ini.v1"
	"log"
	"path/filepath"
	"fmt"
)

func TicketConfig() map[string]string {
	configPath, err := filepath.Abs("config.ini")
	cfg, err := ini.Load(configPath)
	if err != nil {
		log.Fatalf("***[ERROR] 初始化失败，请检查配置文件是否正常***")
	}

	var result = map[string]string{}
	section := cfg.Section("ticket")
	for _, k := range section.KeyStrings() {
		key := section.Key(k)
		result[k] = key.Value()
	}
	fmt.Println("### 配置加载成功 ###")
	return result
}
