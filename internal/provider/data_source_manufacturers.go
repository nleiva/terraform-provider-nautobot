package provider

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tidwall/gjson"

	nb "github.com/nautobot/go-nautobot"
)

func dataSourceManufacturers() *schema.Resource {
	return &schema.Resource{
		Description: "Manufacturer data source in the Terraform provider Nautobot.",

		ReadContext: dataSourceManufacturersRead,

		Schema: map[string]*schema.Schema{
			"manufacturers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created": {
							Description: "Manufacturer's creation date.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"custom_fields": {
							Description: "Manufacturer custom fields.",
							Type:        schema.TypeMap,
							Optional:    true,
						},
						"description": {
							Description: "Manufacturer's description.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"devicetype_count": {
							Description: "Manufacturer's device count.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"display": {
							Description: "Manufacturer's display name.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"id": {
							Description: "Manufacturer's UUID.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"inventoryitem_count": {
							Description: "Manufacturer's inventory item count.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"last_updated": {
							Description: "Manufacturer's last update.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Manufacturer's name.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"platform_count": {
							Description: "Manufacturer's platform count.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"slug": {
							Description: "Manufacturer's slug.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"url": {
							Description: "Manufacturer's URL.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

// Use this as reference: https://learn.hashicorp.com/tutorials/terraform/provider-setup?in=terraform/providers#implement-read
func dataSourceManufacturersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := meta.(*apiClient).Client
	s := meta.(*apiClient).Server

	rsp, err := c.DcimManufacturersListWithResponse(
		ctx,
		&nb.DcimManufacturersListParams{})

	if err != nil {
		return diag.Errorf("failed to get manufacturers list from %s: %s", s, err.Error())
	}

	results := gjson.Get(string(rsp.Body), "results")
	resultsReader := strings.NewReader(results.String())

	list := make([]map[string]interface{}, 0)

	err = json.NewDecoder(resultsReader).Decode(&list)
	if err != nil {
		return diag.Errorf("failed to decode manufacturers list from %s: %s", s, err.Error())
	}

	if err := d.Set("manufacturers", list); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
