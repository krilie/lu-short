package ncfg

import "github.com/mitchellh/mapstructure"

func (cfg *NConfig) GetConfigItem(path string, cfgItem interface{}) {
	eCfg := cfg.V.GetStringMap(path)
	err := mapstructure.Decode(eCfg, cfgItem)
	if err != nil {
		panic(err.Error())
	}
}

func (cfg *NConfig) GetLogCfg() *Log {
	var logCfg = &Log{}
	cfg.GetConfigItem("log", logCfg)
	return logCfg
}

func (cfg *NConfig) GetAliSmsCfg() *AliSms {
	var cfgItem = &AliSms{}
	cfg.GetConfigItem("ali_sms", cfgItem)
	return cfgItem
}

func (cfg *NConfig) GetFileSaveCfg() *FileSave {
	var cfgItem = &FileSave{}
	cfg.GetConfigItem("file_save", cfgItem)
	return cfgItem
}

func (cfg *NConfig) GetDbCfg() *DB {
	var cfgItem = &DB{}
	cfg.GetConfigItem("db", cfgItem)
	return cfgItem
}

func (cfg *NConfig) GetEmailCfg() *Email {
	var cfgItem = &Email{}
	cfg.GetConfigItem("email", cfgItem)
	return cfgItem
}

func (cfg *NConfig) GetHttpCfg() *Http {
	var cfgItem = &Http{}
	cfg.GetConfigItem("http", cfgItem)
	return cfgItem
}

func (cfg *NConfig) GetJwtCfg() *JWT {
	var cfgItem = &JWT{}
	cfg.GetConfigItem("jwt", cfgItem)
	return cfgItem
}
