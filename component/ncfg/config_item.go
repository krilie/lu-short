package ncfg

type Config struct {
	Http     Http     `mapstructure:"http" json:"http" toml:"http"`
	Log      Log      `mapstructure:"log" json:"log" toml:"log"`
	DB       DB       `mapstructure:"db" json:"db" toml:"db"`
	FileSave FileSave `mapstructure:"file_save" json:"file_save" toml:"file_save"`
	JWT      JWT      `mapstructure:"jwt" json:"jwt" toml:"jwt"`
	Email    Email    `mapstructure:"email" json:"email" toml:"email"`
	AliSms   AliSms   `mapstructure:"ali_sms" json:"ali_sms" toml:"ali_sms"`
}

// Http http相关配置
type Http struct {
	EnableSwagger bool   `mapstructure:"enable_swagger" json:"enable_swagger" toml:"enable_swagger"`
	GinMode       string `mapstructure:"gin_mode" json:"gin_mode" toml:"gin_mode"`
	Port          int    `mapstructure:"port" json:"port" toml:"port"`
	SslPri        string `mapstructure:"ssl_pri" json:"ssl_pri" toml:"ssl_pri"`
	SslPub        string `mapstructure:"ssl_pub" json:"ssl_pub" toml:"ssl_pub"`
	Url           string `mapstructure:"url" json:"url" toml:"url"`
}

// Log 日志相关配置
type Log struct {
	LogFile  string `mapstructure:"log_file" json:"log_file" toml:"log_file"`    // 配置文件 空为控制台
	LogLevel uint32 `mapstructure:"log_level" json:"log_level" toml:"log_level"` // 0-6 fatal ... trace
}

// DB 数据库相关配置
type DB struct {
	ConnStr         string `mapstructure:"conn_str" json:"conn_str" toml:"conn_str"`
	MaxOpenConn     int    `mapstructure:"max_open_conn" json:"max_open_conn" toml:"max_open_conn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn" json:"max_idle_conn" toml:"max_idle_conn"`
	ConnMaxLeftTime int    `mapstructure:"conn_max_left_time" json:"conn_max_left_time" toml:"conn_max_left_time"` // 秒数
}

type FileSave struct {
	OssKey      string `mapstructure:"oss_key" json:"oss_key" toml:"oss_key"`
	OssSecret   string `mapstructure:"oss_secret" json:"oss_secret" toml:"oss_secret"`
	OssEndPoint string `mapstructure:"oss_end_point" json:"oss_end_point" toml:"oss_end_point"`
	OssBucket   string `mapstructure:"oss_bucket" json:"oss_bucket" toml:"oss_bucket"`
	Channel     string `mapstructure:"channel" json:"channel" toml:"channel"` // local qiniu aliyun minio aws
}

type JWT struct {
	NormalExpDuration int    `mapstructure:"normal_exp_duration" json:"normal_exp_duration" toml:"normal_exp_duration"` // 秒
	HS256key          string `mapstructure:"hs_256_key" json:"hs_256_key" toml:"hs256_key"`
}

type Email struct {
	Address  string `mapstructure:"address" json:"address" toml:"address"`
	Host     string `mapstructure:"host" json:"host" toml:"host"`
	Port     int    `mapstructure:"port" json:"port" toml:"port"`
	UserName string `mapstructure:"user_name" json:"user_name" toml:"user_name"`
	Password string `mapstructure:"password" json:"password" toml:"password"`
}

type AliSms struct {
	Key    string `mapstructure:"key" json:"key" toml:"key"`
	Secret string `mapstructure:"secret" json:"secret" toml:"secret"`
}
