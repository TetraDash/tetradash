package helpers

import (
	"github.com/pocketbase/pocketbase"
)

func ApplySettings(app *pocketbase.PocketBase) {

	settings, err := app.Dao().FindSettings()
	if err == nil {
		if ConfigEnv.Debug {
			settings.Meta.HideControls = true
		}

		//basic settings
		settings.Meta.AppName = "Tetra Dash"
		settings.Meta.AppUrl = "https://tetradash.com"

		// cloudflare backup bucket
		if len(ConfigEnv.CloudFlareBackupR2.Bucket) > 0 &&
			len(ConfigEnv.CloudFlareBackupR2.Region) > 0 &&
			len(ConfigEnv.CloudFlareBackupR2.Endpoint) > 0 &&
			len(ConfigEnv.CloudFlareBackupR2.AccessKey) > 0 &&
			len(ConfigEnv.CloudFlareBackupR2.Secret) > 0 {

			settings.Backups.S3.Enabled = true
			settings.Backups.S3.Bucket = ConfigEnv.CloudFlareBackupR2.Bucket
			settings.Backups.S3.Region = ConfigEnv.CloudFlareBackupR2.Region
			settings.Backups.S3.Endpoint = ConfigEnv.CloudFlareBackupR2.Endpoint
			settings.Backups.S3.AccessKey = ConfigEnv.CloudFlareBackupR2.AccessKey
			settings.Backups.S3.Secret = ConfigEnv.CloudFlareBackupR2.Secret

			settings.Backups.Cron = "0 0 * * *"
			settings.Backups.CronMaxKeep = 10
		} else {
			settings.Backups.S3.Enabled = false
		}

		// cloudflare file storage bucket

		if len(ConfigEnv.CloudFlareFileR2.Bucket) > 0 &&
			len(ConfigEnv.CloudFlareFileR2.Region) > 0 &&
			len(ConfigEnv.CloudFlareFileR2.Endpoint) > 0 &&
			len(ConfigEnv.CloudFlareFileR2.AccessKey) > 0 &&
			len(ConfigEnv.CloudFlareFileR2.Secret) > 0 {

			settings.S3.Enabled = true
			settings.S3.Bucket = ConfigEnv.CloudFlareFileR2.Bucket
			settings.S3.Region = ConfigEnv.CloudFlareFileR2.Region
			settings.S3.Endpoint = ConfigEnv.CloudFlareFileR2.Endpoint
			settings.S3.AccessKey = ConfigEnv.CloudFlareFileR2.AccessKey
			settings.S3.Secret = ConfigEnv.CloudFlareFileR2.Secret

		} else {
			settings.S3.Enabled = false
		}

		settings.Logs.MaxDays = 2

		app.Dao().SaveSettings(settings)
	}
}
