package helpers

import (
	"os"
	"strings"

	"github.com/pocketbase/pocketbase/models/settings"
)

type AppConfigEnv struct {
	Debug                bool
	SmtpUserName         string
	SmtpPassword         string
	GoogleClientId       string
	GoogleClientSecret   string
	CloudFlareBackupR2   settings.S3Config
	CloudFlareFileR2     settings.S3Config
	PbAdminAcceptDomains []string
	PbEnableWebUI        bool
}

var ConfigEnv *AppConfigEnv

func init() {
	ConfigEnv = &AppConfigEnv{
		Debug:              os.Getenv("DEBUG") != "1",
		SmtpUserName:       os.Getenv("SMTP_USER_NAME"),
		SmtpPassword:       os.Getenv("SMTP_PASSWORD"),
		GoogleClientId:     os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		CloudFlareBackupR2: settings.S3Config{
			Bucket:    os.Getenv("CF_BACKUP_R2_BUCKET"),
			Region:    os.Getenv("CF_BACKUP_R2_REGION"),
			Endpoint:  os.Getenv("CF_BACKUP_R2_BUCKET"),
			AccessKey: os.Getenv("CF_BACKUP_R2_ACCESS_KEY_ID"),
			Secret:    os.Getenv("CF_BACKUP_R2_SECRET_ACCESS_KEY"),
		},
		CloudFlareFileR2: settings.S3Config{
			Bucket:    os.Getenv("CF_FILE_R2_BUCKET"),
			Region:    os.Getenv("CF_FILE_R2_REGION"),
			Endpoint:  os.Getenv("CF_FILE_R2_ENDPOINT"),
			AccessKey: os.Getenv("CF_FILE_R2_ACCESS_KEY_ID"),
			Secret:    os.Getenv("CF_FILE_R2_SECRET_ACCESS_KEY"),
		},
		PbAdminAcceptDomains: strings.Split(os.Getenv("PB_ADMIN_ACCEPT_DOMAINS"), ","),
		PbEnableWebUI:        os.Getenv("PB_ENABLE_WEB_UI") == "1",
	}
}
