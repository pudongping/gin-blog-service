package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config") // 设置配置文件的名称为 config

	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
			// vp.AddConfigPath("configs/") // 设置其配置路径为相对路径 configs/
		}
	}

	vp.SetConfigType("yaml") // 配置类型为 yaml
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp: vp}, nil
}
