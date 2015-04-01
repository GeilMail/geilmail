package cfg

type Config struct {
	SQLite *SQLiteConfig
	IMAP   *IMAPConfig
	SMTP   *SMTPConfig
}

type IMAPConfig struct {
	ListenIP string
	Port     int
}

type SMTPConfig struct {
	ListenIP string
	Port     int
}
