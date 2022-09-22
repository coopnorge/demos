## GCP to Azure Auth Demo
This is a demo of authenticating towards Azure resources using Workload Identity Federation and exchanging a GCP
identity token for an Azure access token. 

The [Google credentials package](https://pkg.go.dev/cloud.google.com/go/iam/credentials/apiv1) used to create a IamCredentialsClient will create a client based 
on the current Application Default Credentials (ADC), usually the service account connected to an application.
When we run the demo locally we need to use impersonation to acquire the correct credentials for the `helloworld-to-azure` service account.

To add yourself as a principal that is allowed to impersonate the service account, add the following resource in
[helloworld-infrastructure repo](https://github.com/coopnorge/helloworld-infrastructure/blob/main/terraform/main.tf)
in `terraform/main.tf`.
**NB!** Remember to remove the resource after you are done testing.

```terraform
resource "google_service_account_iam_member" "impersonation-account-iam" {
  service_account_id = google_service_account.helloworld-to-azure.name
  role               = "roles/iam.serviceAccountTokenCreator"
  member             = "user:<your email>"
}
```

Then run
```shell
gcloud auth application-default login --impersonate-service-account=helloworld-to-azure@helloworld-production-5120.iam.gserviceaccount.com
``` 
in your terminal (with the gcloud CLI installed ) to enable impersonation with the correct service account credentials.


Furthermore, you must create a `.env` file in `/config` with the following content:

```
GCP_SERVICE_ACCOUNT = "helloworld-to-azure@helloworld-production-5120.iam.gserviceaccount.com"
AZURE_CLIENT_ID = ""
AZURE_AD_TENANT_ID = ""
CONTAINER_URL = "https://helloworldstoracc.blob.core.windows.net/helloworld-stor-account-content"
ACCOUNT_NAME = "helloworldstoracc"
CONTAINER_NAME = "helloworld-stor-account-content"
AZURE_SCOPE = "https://helloworldstoracc.blob.core.windows.net/.default"
BLOB_NAME = "text.txt"
```

The values above are pointing to a storage account and blob in Azure that has been created in
[helloworld-infrastructure](https://github.com/coopnorge/helloworld-infrastructure).
`AZURE_CLIENT_ID` and `AZURE_AD_TENANT_ID` can be
found by searching for "helloworld-application" in
[App registrations](https://portal.azure.com/#view/Microsoft_AAD_RegisteredApps/ApplicationsListBlade) in Azure Portal.
