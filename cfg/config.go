package cfg

type Config struct {
	SQLite *SQLiteConfig
	IMAP   *IMAPConfig
	SMTP   *SMTPConfig
	TLS    *TLSConfig
	HTTP   *HTTPConfig
}

type IMAPConfig struct {
	ListenIP string
	Port     int
}

type SMTPConfig struct {
	ListenIP string
	Port     int
}

type TLSConfig struct {
	CertPath string
	KeyPath  string
}

type HTTPConfig struct {
	Listen string
}
