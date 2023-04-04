package setting

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var cfg *AppConfig

type AppConfig struct {
	ProdMode    bool   `mapstructure:"prod_mode"`
	LogLevel    string `mapstructure:"log_level"`
	LogFilePath string `mapstructure:"log_file_path"`
	ApiHost     string `mapstructure:"api_host"`
	AppHost     string `mapstructure:"app_host"`
	DBAdapter   string `mapstructure:"db_adapter"`
	DBHost      string `mapstructure:"db_host"`
	DBPort      string `mapstructure:"db_port"`
	DBName      string `mapstructure:"db_name"`
	DBUser      string `mapstructure:"db_user"`
	DBPassword  string `mapstructure:"db_password"`
	JwtSecret   string `mapstructure:"jwt_secret"`
	Domain      string `mapstructure:"DOMAIN"`
}

func InitConfig() (AppConfig, error) {
	cfg := new(AppConfig)

	err := godotenv.Load(os.Getenv("CONFIG_FILE"))
	if err != nil {
		return *cfg, err
	}
	v := viper.New()

	v.SetDefault("PROD_MODE", false)
	v.SetDefault("LOG_FILE_PATH", "~log.json")
	v.SetDefault("LOG_LEVEL", "info")
	v.SetDefault("API_HOST", "localhost:8411")
	v.SetDefault("APP_HOST", "localhost:8412")
	v.SetDefault("DB_ADAPTER", "postgres")

	v.SetDefault("DB_HOST", "localhost")
	v.SetDefault("DB_PORT", "5443")
	v.SetDefault("DB_NAME", "blocks")
	v.SetDefault("DB_USER", "blocks")
	v.SetDefault("DB_PASSWORD", "blocks")
	v.SetDefault("DOMAIN", "localhost")

	v.BindEnv("DOMAIN")
	v.BindEnv("PROD_MODE")
	v.BindEnv("LOG_LEVEL")
	v.BindEnv("LOG_FILE_PATH")
	v.BindEnv("API_HOST")
	v.BindEnv("APP_HOST")
	v.BindEnv("DB_ADAPTER")
	v.BindEnv("DB_HOST")
	v.BindEnv("DB_PORT")
	v.BindEnv("DB_NAME")
	v.BindEnv("DB_USER")
	v.BindEnv("DB_PASSWORD")

	if err := v.UnmarshalExact(cfg); err != nil {
		return *cfg, err
	}

	return *cfg, nil
}
func (c *AppConfig) DBConnectionString() string {
	return "user=" + c.DBUser + " dbname=" + c.DBName + " sslmode=disable port=" + c.DBPort + " host=" + c.DBHost + " password=" + c.DBPassword
}
func GetConfig() AppConfig {
	return *cfg
}
