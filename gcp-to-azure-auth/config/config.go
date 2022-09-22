package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	GcpServiceAccount string
	AzureClientID     string
	AzureAdTenantID   string
	ContainerUrl      string
	AccountName       string
	ContainerName     string
	AzureScope        string
	BlobName          string
}

func GetConfig() (*Config, error) {
	viper.SetConfigFile("config/.env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{
		GcpServiceAccount: viper.GetString("GCP_SERVICE_ACCOUNT"),
		AzureClientID:     viper.GetString("AZURE_CLIENT_ID"),
		AzureAdTenantID:   viper.GetString("AZURE_AD_TENANT_ID"),
		ContainerUrl:      viper.GetString("CONTAINER_URL"),
		AccountName:       viper.GetString("ACCOUNT_NAME"),
		ContainerName:     viper.GetString("CONTAINER_NAME"),
		AzureScope:        viper.GetString("AZURE_SCOPE"),
		BlobName:          viper.GetString("BLOB_NAME"),
	}
	return config, nil
}
