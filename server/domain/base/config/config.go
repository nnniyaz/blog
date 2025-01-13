package config

type Config struct {
	port      int
	isDevMode bool
	mongoUri  string
	smtpHost  string
	smtpPort  int
	smtpUser  string
	smtpPass  string
}

func NewConfig(port int, isDevMode bool, mongoUri, smtpHost string, smtpPort int, smtpUser, smtpPass string) *Config {
	return &Config{
		port:      port,
		isDevMode: isDevMode,
		mongoUri:  mongoUri,
		smtpHost:  smtpHost,
		smtpPort:  smtpPort,
		smtpUser:  smtpUser,
		smtpPass:  smtpPass,
	}
}

func (c *Config) GetPort() int {
	return c.port
}

func (c *Config) GetIsDevMode() bool {
	return c.isDevMode
}

func (c *Config) GetMongoUri() string {
	return c.mongoUri
}

func (c *Config) GetSmtpHost() string {
	return c.smtpHost
}

func (c *Config) GetSmtpPort() int {
	return c.smtpPort
}

func (c *Config) GetSmtpUser() string {
	return c.smtpUser
}

func (c *Config) GetSmtpPass() string {
	return c.smtpPass
}
