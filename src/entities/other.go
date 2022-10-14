package entities

type (
	AppStruct struct {
		Addr          string `yaml:"Addr"`
		Name          string `yaml:"Name"`
		Timeout       int    `yaml:"Timeout"`
		Qps           int    `yaml:"Qps"`
		MaxPostMemory int64  `yaml:"MaxPostMemory"`
	}
	LogStruct struct {
		Level    string `yaml:"Level"`
		Output   string `yaml:"Output"`
		FilePath string `yaml:"FilePath"`
	}
	PostgreSqlStruct struct {
		Host     string `yaml:"Host"`
		Port     int    `yaml:"Port"`
		User     string `yaml:"User"`
		Pass     string `yaml:"Pass"`
		DB       string `yaml:"DB"`
		TimeZone string `yaml:"TimeZone"`
	}
	RedisStruct struct {
		Host string `yaml:"Host"`
		Port int    `yaml:"Port"`
		Pass string `yaml:"Pass"`
	}
	DatabaseStruct struct {
		PostgreSql PostgreSqlStruct `yaml:"PostgreSql"`
		Redis      RedisStruct      `yaml:"Redis"`
	}
	ConfigStruct struct {
		App      AppStruct      `yaml:"App"`
		Log      LogStruct      `yaml:"Log"`
		Database DatabaseStruct `yaml:"Database"`
	}
)
