// This file was generated by protogen. DO NOT EDIT.

package sdm

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

const userAgent = "terraform-provider-sdm/2.6.3"

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SDM_API_ACCESS_KEY", nil),
				Description: "A GUID identifying the API key used to authenticate with the StrongDM API.",
			},

			"api_secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SDM_API_SECRET_KEY", nil),
				Description: "A base64 encoded secret key used to authenticate with the StrongDM API.",
			},

			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SDM_API_HOST", "api.strongdm.com:443"),
				Description: "The host and port of the StrongDM API endpoint.",
			},
			"retry_rate_limit_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SDM_RETRY_RATE_LIMITS", "true"),
				Description: "Whether experienced rate limits should cause the client to sleep instead of erroring out",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sdm_account_attachment": resourceAccountAttachment(),
			"sdm_account_grant":      resourceAccountGrant(),
			"sdm_account":            resourceAccount(),
			"sdm_node":               resourceNode(),
			"sdm_remote_identity":    resourceRemoteIdentity(),
			"sdm_resource":           resourceResource(),
			"sdm_role_attachment":    resourceRoleAttachment(),
			"sdm_role_grant":         resourceRoleGrant(),
			"sdm_role":               resourceRole(),
			"sdm_secret_store":       resourceSecretStore(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"sdm_ssh_ca_pubkey":         dataSourceControlPanelSSHCAPublicKey(),
			"sdm_account_attachment":    dataSourceAccountAttachment(),
			"sdm_account_grant":         dataSourceAccountGrant(),
			"sdm_account":               dataSourceAccount(),
			"sdm_node":                  dataSourceNode(),
			"sdm_remote_identity":       dataSourceRemoteIdentity(),
			"sdm_remote_identity_group": dataSourceRemoteIdentityGroup(),
			"sdm_resource":              dataSourceResource(),
			"sdm_role_attachment":       dataSourceRoleAttachment(),
			"sdm_role_grant":            dataSourceRoleGrant(),
			"sdm_role":                  dataSourceRole(),
			"sdm_secret_store":          dataSourceSecretStore(),
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			host := d.Get("host").(string)
			opts := []sdm.ClientOption{
				sdm.WithUserAgentExtra(userAgent),
				sdm.WithHost(host),
			}
			if strings.HasPrefix(host, "localhost:") ||
				strings.HasPrefix(host, "127.0.0.1:") ||
				strings.HasPrefix(host, "host.docker.internal:") {
				opts = append(opts, sdm.WithInsecure())
			}
			if retry, ok := d.Get("retry_rate_limit_errors").(bool); ok {
				opts = append(opts, sdm.WithRateLimitRetries(retry))
			}
			client, err := sdm.New(d.Get("api_access_key").(string), d.Get("api_secret_key").(string), opts...)
			if err != nil {
				return nil, fmt.Errorf("cannot dial API server: %w", err)
			}
			return client, nil
		},
	}
}
