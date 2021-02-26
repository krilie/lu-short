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

	return cfg
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
