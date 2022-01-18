package setting

import (
	"github.com/fsnotify/fsnotify"
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

	// 变更热更新的配置
	s := &Setting{vp: vp}
	s.WatchSettingChange()

	return s, nil
}

// WatchSettingChange 监听文件的热更新
func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig() // 对文件配置进行监听
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}
