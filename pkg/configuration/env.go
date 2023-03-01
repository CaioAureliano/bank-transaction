package configuration

type EnvironmentVariable struct {
	PORT                    string
	JWTSECRET               string
	DBURI                   string
	DBNAME                  string
	DBUSER                  string
	DBPASSWORD              string
	QUEUEURI                string
	QUEUENAME               string
	REDISURI                string
	AUTHENTICATORSERVICEURL string
}

var Env = EnvironmentVariable{
	PORT:                    getenv("PORT"),
	JWTSECRET:               getenv("JWT_SECRET"),
	DBURI:                   getenv("DB_URI"),
	DBNAME:                  getenv("DB_NAME"),
	DBUSER:                  getenv("DB_USER"),
	DBPASSWORD:              getenv("DB_PASSWORD"),
	QUEUEURI:                getenv("RABBITMQ_URI"),
	QUEUENAME:               getenv("RABBITMQ_QUEUE_NAME"),
	REDISURI:                getenv("REDIS_URI"),
	AUTHENTICATORSERVICEURL: getenv("AUTHENTICATOR_SERVICE_URL"),
}
