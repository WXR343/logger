package config

type Log struct {
	Level      string   `mapstructure:"level" json:"level" yaml:"level"`                   //日志等级
	RootDir    string   `mapstructure:"root_dir" json:"root_dir" yaml:"root_dir"`          //日志根目录
	Filename   []string `mapstructure:"filename" json:"filename" yaml:"filename"`          //日志文件名称
	Format     string   `mapstructure:"format" json:"format" yaml:"format"`                //写入格式 可选json
	ShowLine   bool     `mapstructure:"show_line" json:"show_line" yaml:"show_line"`       //是否显示调用行
	MaxBackups int      `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"` //旧文件的最大个数
	MaxSize    int      `mapstructure:"max_size" json:"max_size" yaml:"max_size"`          // 日志文件最大大小（MB）
	MaxAge     int      `mapstructure:"max_age" json:"max_age" yaml:"max_age"`             // 旧文件的最大保留天数
	Compress   bool     `mapstructure:"compress" json:"compress" yaml:"compress"`          //是否压缩
}

type Configuration struct {
	Log Log `mapstructure:"log" json:"log" yaml:"log"`
}
