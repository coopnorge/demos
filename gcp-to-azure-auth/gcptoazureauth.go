package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"

	"gcp-to-azure-auth/config"

	credentials "cloud.google.com/go/iam/credentials/apiv1"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
	credentialspb "google.golang.org/genproto/googleapis/iam/credentials/v1"
)

func main() {
	ctx := context.Background()

	gcpToAzureAuthConfig, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to read config: %s", err)
	}

	accessToken, err := exchangeIdTokenForAzureADAccessToken(ctx, gcpToAzureAuthConfig)
	if err != nil {
		log.Fatalf("%s", err)
	}

	content, err := readBlobData(ctx, accessToken, gcpToAzureAuthConfig)
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println("Blob content:", content)
}

func readBlobData(ctx context.Context, token string, config *config.Config) (string, error) {
	credential := azblob.NewTokenCredential(token, nil)

	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	URL, err := url.Parse(config.ContainerUrl)
	if err != nil {
		return "", fmt.Errorf("failed to parse url %w", err)
	}
	containerURL := azblob.NewContainerURL(*URL, p)

	blobURL := containerURL.NewBlockBlobURL(config.BlobName)

	get, err := blobURL.Download(ctx, 0, 0, azblob.BlobAccessConditions{}, false, azblob.ClientProvidedKeyOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to read blob data %w", err)
	}

	downloadedData := &bytes.Buffer{}
	reader := get.Body(azblob.RetryReaderOptions{})
	downloadedData.ReadFrom(reader)
	reader.Close()
	return downloadedData.String(), nil
}

func exchangeIdTokenForAzureADAccessToken(ctx context.Context, config *config.Config) (string, error) {
	assertion := confidential.NewCredFromAssertionCallback(func(c context.Context, o confidential.AssertionRequestOptions) (string, error) {
		return getGcpIdToken(ctx, config)
	})

	authority := fmt.Sprintf("https://login.microsoftonline.com/%s", config.AzureAdTenantID)
	client, err := confidential.New(config.AzureClientID, assertion, confidential.WithAuthority(authority))
	if err != nil {
		return "", fmt.Errorf("failed to create client %w", err)
	}

	token, err := client.AcquireTokenByCredential(ctx, []string{config.AzureScope})
	if err != nil {
		return "", fmt.Errorf("failed to exchange id token for access token %w", err)
	}

	return token.AccessToken, nil
}

func getGcpIdToken(ctx context.Context, config *config.Config) (string, error) {
	c, err := credentials.NewIamCredentialsClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create credentials client %w", err)
	}
	defer c.Close()

	req := &credentialspb.GenerateIdTokenRequest{
		Name:     config.GcpServiceAccount,
		Audience: "api://AzureADTokenExchange",
	}
	resp, err := c.GenerateIdToken(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to generate id token %w", err)
	}

	return resp.GetToken(), nil
}
