package config

type Config struct {
	port          int
	isDevMode     bool
	mongoUri      string
	smtpHost      string
	smtpPort      int
	smtpUser      string
	smtpPass      string
	spaceBucket   string
	spaceKey      string
	spaceSecret   string
	spaceEndPoint string
	spaceRegion   string
	spaceName     string
	spaceHost     string
}

func NewConfig(port int, isDevMode bool, mongoUri, smtpHost string, smtpPort int, smtpUser, smtpPass, spaceBucket, spaceKey, spaceSecret, spaceEndPoint, spaceRegion, spaceName, spaceHost string) *Config {
	return &Config{
		port:          port,
		isDevMode:     isDevMode,
		mongoUri:      mongoUri,
		smtpHost:      smtpHost,
		smtpPort:      smtpPort,
		smtpUser:      smtpUser,
		smtpPass:      smtpPass,
		spaceBucket:   spaceBucket,
		spaceKey:      spaceKey,
		spaceSecret:   spaceSecret,
		spaceEndPoint: spaceEndPoint,
		spaceRegion:   spaceRegion,
		spaceName:     spaceName,
		spaceHost:     spaceHost,
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

func (c *Config) GetSpaceBucket() string {
	return c.spaceBucket
}

func (c *Config) GetSpaceKey() string {
	return c.spaceKey
}

func (c *Config) GetSpaceSecret() string {
	return c.spaceSecret
}

func (c *Config) GetSpaceEndPoint() string {
	return c.spaceEndPoint
}

func (c *Config) GetSpaceRegion() string {
	return c.spaceRegion
}

func (c *Config) GetSpaceName() string {
	return c.spaceName
}

func (c *Config) GetSpaceHost() string {
	return c.spaceHost
}
