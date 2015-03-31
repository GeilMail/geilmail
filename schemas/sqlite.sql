CREATE TABLE domains (
	id INTEGER PRIMARY KEY,
	domain VARCHAR(255)
);

CREATE TABLE users (
	user_id INTEGER PRIMARY KEY,
	domain_id INTEGER,
	mail VARCHAR(200),
	salt VARCHAR(30),
	password VARCHAR(32),
	FOREIGN KEY(domain_id) REFERENCES domains(id)
);

CREATE TABLE mails (
	mail_id INTEGER PRIMARY KEY,
	user INTEGER,
	folder VARCHAR(100),
	FOREIGN KEY(user) REFERENCES users(user_id)
);
