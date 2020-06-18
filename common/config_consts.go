package common

// CONFIG 配置常量
type CONFIG string

const (
	// ConfigFileName 配置文件名
	ConfigFileName CONFIG = "conf"
	// ConfigFileType 配置文件类型
	ConfigFileType CONFIG = "json"
	// ConfigFilePath 配置文件路径
	ConfigFilePath CONFIG = "."

	// ServerConfigKey 服务器配置键
	ServerConfigKey CONFIG = "server"
	// StorageConfigKey 存储配置键
	StorageConfigKey CONFIG = "storage"
)

// String to string
func (config CONFIG) String() string {
	return string(config)
}
