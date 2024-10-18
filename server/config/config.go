package config

var C = config{
	RDBConnectionString: "host=localhost user=postgres password=postgres dbname=realworld sslmode=disable timezone=UTC",
	PasswordAuthHmacKey: []byte("sample-key"),
}

type config struct {
	RDBConnectionString string
	PasswordAuthHmacKey []byte
}
