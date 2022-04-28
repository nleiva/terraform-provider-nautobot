package provider

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tidwall/gjson"

	nb "github.com/nautobot/go-nautobot"
)

func resourceManufacturer() *schema.Resource {
	return &schema.Resource{
		Description: "This object manages a manufacturer in Nautobot",

		CreateContext: resourceManufacturerCreate,
		ReadContext:   resourceManufacturerRead,
		UpdateContext: resourceManufacturerUpdate,
		DeleteContext: resourceManufacturerDelete,

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
	}
}

func resourceManufacturerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := meta.(*apiClient).Client
	s := meta.(*apiClient).Server

	var m nb.Manufacturer

	name, ok := d.GetOk("name")
	n := name.(string)
	if ok {
		m.Name = n
	}

	m.Description = &n
	description, ok := d.GetOk("description")
	if ok {
		t := description.(string)
		m.Description = &t
	}

	sl := strings.ReplaceAll(strings.ToLower(n), " ", "-")
	m.Slug = &sl
	slug, ok := d.GetOk("slug")
	if ok {
		t := slug.(string)
		m.Slug = &t
	}

	rsp, err := c.DcimManufacturersCreateWithResponse(
		ctx,
		nb.DcimManufacturersCreateJSONRequestBody(m))
	if err != nil {
		return diag.Errorf("failed to create manufacturer %s on %s: %s", name.(string), s, err.Error())
	}
	data := string(rsp.Body)

	dataName := gjson.Get(data, "name.0")

	if dataName.String() == "manufacturer with this name already exists." {
		rsp, err := c.DcimManufacturersListWithResponse(
			ctx,
			&nb.DcimManufacturersListParams{
				NameIe: &[]string{n},
			})

		if err != nil {
			return diag.Errorf("failed to get manufacturer %s from %s: %s", n, s, err.Error())
		}
		id := gjson.Get(string(rsp.Body), "results.0.id")

		d.Set("id", id.String())
		d.SetId(id.String())
		resourceManufacturerRead(ctx, d, meta)

		return diags
	}

	tflog.Trace(ctx, "manufacturer created", map[string]interface{}{
		"name": name.(string),
		"data": []interface{}{description, slug},
	})

	id := gjson.Get(data, "id")

	//d.Set("id", id.String())
	d.SetId(id.String())
	resourceManufacturerRead(ctx, d, meta)

	return diags
}

func resourceManufacturerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := meta.(*apiClient).Client
	s := meta.(*apiClient).Server

	name := d.Get("name").(string)
	id := d.Get("id").(string)

	rsp, err := c.DcimManufacturersListWithResponse(
		ctx,
		&nb.DcimManufacturersListParams{
			IdIe: &[]types.UUID{types.UUID(id)},
		})

	if err != nil {
		return diag.Errorf("failed to get manufacturer %s from %s: %s", name, s, err.Error())
	}

	results := gjson.Get(string(rsp.Body), "results.0")
	resultsReader := strings.NewReader(results.String())

	item := make(map[string]interface{})

	err = json.NewDecoder(resultsReader).Decode(&item)
	if err != nil {
		return diag.Errorf("failed to decode manufacturer %s from %s: %s", name, s, err.Error())
	}

	d.Set("name", item["name"].(string))
	d.Set("created", item["created"].(string))
	d.Set("description", item["description"].(string))
	// TODO: Fix issue with display going away
	d.Set("display", item["display"].(string))
	d.Set("id", item["id"].(string))
	// TODO: Fix issue with slug going away
	d.Set("slug", item["slug"].(string))
	// TODO: Fix issue with url going away
	d.Set("url", item["url"].(string))
	d.Set("last_updated", item["last_updated"].(string))

	switch v := item["devicetype_count"].(type) {
	case int:
		d.Set("devicetype_count", v)
	case float64:
		d.Set("devicetype_count", int(v))
	default:
	}
	switch v := item["inventoryitem_count"].(type) {
	case int:
		d.Set("inventoryitem_count", v)
	case float64:
		d.Set("inventoryitem_count", int(v))
	default:
	}
	switch v := item["platform_count"].(type) {
	case int:
		d.Set("platform_count", v)
	case float64:
		d.Set("platform_count", int(v))
	default:
	}

	return diags
}

func resourceManufacturerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

// DcimManufacturersPartialUpdateWithBodyWithResponse request with arbitrary body returning *DcimManufacturersPartialUpdateResponse
// func (c *ClientWithResponses) DcimManufacturersPartialUpdateWithBodyWithResponse(ctx context.Context, id openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DcimManufacturersPartialUpdateResponse, error) {
// 	rsp, err := c.DcimManufacturersPartialUpdateWithBody(ctx, id, contentType, body, reqEditors...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ParseDcimManufacturersPartialUpdateResponse(rsp)
// }

// func (c *ClientWithResponses) DcimManufacturersPartialUpdateWithResponse(ctx context.Context, id openapi_types.UUID, body DcimManufacturersPartialUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*DcimManufacturersPartialUpdateResponse, error) {
// 	rsp, err := c.DcimManufacturersPartialUpdate(ctx, id, body, reqEditors...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ParseDcimManufacturersPartialUpdateResponse(rsp)
// }



	return diags
}

func resourceManufacturerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}
