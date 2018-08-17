package main

type App struct {
	ServerAddr     string `toml:"server_addr"`
	ApiKey         string `toml:"api_key"`
	MailServerAddr string `toml:"mail_server_addr"`
	ServerAccount  string `toml:"server_account"`
	MailNotify     string `toml:"mail_notice"`
}

type Users struct {
	MailList []string `toml:"mail_list"`
}

type Config struct {
	App   App   `toml:"app"`
	Users Users `toml:"users"`
}
