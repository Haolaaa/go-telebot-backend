package config

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Timer  Timer  `mapstructure:"timer" json:"timer" yaml:"timer"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
}
