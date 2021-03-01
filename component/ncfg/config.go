package ncfg

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"lu-short/common/run_env"
	"os"
	"strings"
)

type NConfig struct {
	V      *viper.Viper
	RunEnv *run_env.RunEnv
}

func NewNConfig() *NConfig {
	var cfg = &NConfig{V: viper.New(), RunEnv: run_env.RunEnvLocal}

	//读取环境变量值
	cfg.V.SetEnvPrefix("LUSHORT")
	cfg.V.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	cfg.V.SetEnvKeyReplacer(replacer)

	// 命令行位置 > 环境变量位置 > 环境变量配置 > 默认位置

	// 默认位置的配置文件
	err := cfg.LoadConfigByFile("config.toml")
	if err != nil {
		println("warn:默认位置 未能加载 " + err.Error())
	}
	// 默认环境变量位置
	err = cfg.LoadConfigByFile(os.Getenv("LUSHORT_CFG_PATH"))
	if err != nil {
		println("warn:环境变量位置 未能加载" + err.Error())
	}
	// 默认环境变量内容
	err = cfg.LoadFromConfigJsonStr(os.Getenv("LUSHORT_CFG_CONTENT"))
	if err != nil {
		println("warn:环境变量配置内容 未能加载" + err.Error())
	}
	// 配置位置 --config_path
	configPath := func() string {
		for _, value := range os.Args[1:] {
			if strings.HasPrefix(value, "config_path") ||
				strings.HasPrefix(value, "-config_path") ||
				strings.HasPrefix(value, "--config_path") {
				split := strings.Split(value, "=")
				return split[1]
			}
		}
		return ""
	}()
	if configPath != "" {
		err = cfg.LoadConfigByFile(configPath)
		if err != nil {
			println("warn:命令行配置文件位置 未能加载" + err.Error())
		}
	} else {
		println("warn:命令行配置文件位置 未能加载")
	}

	return cfg
}

func NewNConfigFromTomlFilePathEnv(cfgPathEnvName string) func() *NConfig {
	return func() *NConfig {
		var cfg = NewNConfig()
		err := cfg.LoadConfigByFile(os.Getenv(cfgPathEnvName))
		if err != nil {
			panic(err)
		}
		return cfg
	}
}

func NewNConfigFromFilePathJsonContent(cfgContentEnvName string) func() *NConfig {
	return func() *NConfig {
		var cfg = NewNConfig()
		err := cfg.LoadFromConfigJsonStr(os.Getenv(cfgContentEnvName))
		if err != nil {
			panic(err)
		}
		return cfg
	}
}

func (cfg *NConfig) LoadConfigByFile(name string) error {
	open, err := os.Open(name)
	if err != nil {
		return err
	}
	defer open.Close()
	cfgStr, err := ioutil.ReadAll(open)
	if err != nil {
		return err
	}
	return cfg.LoadFromConfigTomlStr(string(cfgStr))
}

func (cfg *NConfig) LoadFromConfigTomlStr(cfgStr string) error {
	cfg.V.SetConfigType("toml")
	if err := cfg.V.MergeConfig(strings.NewReader(cfgStr)); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			log.Println("no config find on cfg str gen and use default:", err)
		default:
			log.Println(err)
		}
		return err
	} else {
		return nil
	}
}

func (cfg *NConfig) LoadFromConfigJsonStr(cfgStr string) error {
	cfg.V.SetConfigType("json")
	if err := cfg.V.MergeConfig(strings.NewReader(cfgStr)); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			log.Println("no config find on cfg str gen and use default:", err)
		default:
			log.Println(err)
		}
		return err
	} else {
		return nil
	}
}
