// This file was generated by protogen. DO NOT EDIT.

package sdm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	sdm "github.com/strongdm/terraform-provider-sdm/sdm/internal/sdk"
)

func dataSourceControlPanelSSHCAPublicKey() *schema.Resource {
	return &schema.Resource{
		ReadContext: wrapCrudOperation(dataSourceControlPanelSSHCAPublicKeyGet),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier of the SSH CA Public Key.",
			},
			"public_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The SSH Certificate Authority Public Key.",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(60 * time.Second),
		},
	}
}

func dataSourceControlPanelSSHCAPublicKeyGet(ctx context.Context, d *schema.ResourceData, cc *sdm.Client) error {
	resp, err := cc.ControlPanel().GetSSHCAPublicKey(ctx)
	if err != nil {
		return fmt.Errorf("cannot get SSH CA Public Key %s: %w", d.Id(), err)
	}

	d.Set("public_key", resp.PublicKey)
	d.SetId("ControlPanelSSHCAPublicKey")

	return nil
}
