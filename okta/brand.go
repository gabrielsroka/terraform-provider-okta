package okta

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

var brandResourceSchema = map[string]*schema.Schema{
	"id": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Brand ID",
	},
	"brand_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Brand ID - Note: Okta API for brands only reads and updates therefore the okta_brand resource needs to act as a quasi data source. Do this by setting brand_id.",
	},
	"agree_to_custom_privacy_policy": {
		Type:         schema.TypeBool,
		Optional:     true,
		Description:  "Consent for updating the custom privacy policy URL.",
		RequiredWith: []string{"custom_privacy_policy_url"},
	},
	"custom_privacy_policy_url": {
		Type:             schema.TypeString,
		Optional:         true,
		Description:      "Custom privacy policy URL",
		DiffSuppressFunc: suppressDuringCreate,
	},
	"links": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Link relations for this object - JSON HAL - Discoverable resources related to the brand",
	},
	"remove_powered_by_okta": {
		Type:             schema.TypeBool,
		Optional:         true,
		Default:          false,
		Description:      `Removes "Powered by Okta" from the Okta-hosted sign-in page and "© 2021 Okta, Inc." from the Okta End-User Dashboard`,
		DiffSuppressFunc: suppressDuringCreate,
	},
}

func suppressDuringCreate(k, old, new string, d *schema.ResourceData) bool {
	// If brand_id has changed assume this is create and treat the properties as readers not caring about what would otherwise apear to be drift.
	if d.HasChange("brand_id") {
		return true
	}
	return old == new
}

var brandDataSourceSchema = map[string]*schema.Schema{
	"brand_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Brand ID",
	},
	"custom_privacy_policy_url": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Custom privacy policy URL",
	},
	"links": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Link relations for this object - JSON HAL - Discoverable resources related to the brand",
	},
	"remove_powered_by_okta": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: `Removes "Powered by Okta" from the Okta-hosted sign-in page and "© 2021 Okta, Inc." from the Okta End-User Dashboard`,
	},
}

var brandsDataSourceSchema = map[string]*schema.Schema{
	"id": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Brand ID",
	},
	"custom_privacy_policy_url": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Custom privacy policy URL",
	},
	"links": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Link relations for this object - JSON HAL - Discoverable resources related to the brand",
	},
	"remove_powered_by_okta": {
		Type:        schema.TypeBool,
		Computed:    true,
		Description: `Removes "Powered by Okta" from the Okta-hosted sign-in page and "© 2021 Okta, Inc." from the Okta End-User Dashboard`,
	},
}

func flattenBrand(brand *okta.Brand) map[string]interface{} {
	attrs := map[string]interface{}{}
	// NOTE: explicitly set defaults on brand
	attrs["custom_privacy_policy_url"] = ""
	if brand.CustomPrivacyPolicyUrl != "" {
		attrs["custom_privacy_policy_url"] = brand.CustomPrivacyPolicyUrl
	}
	links, _ := json.Marshal(brand.Links)
	attrs["links"] = string(links)
	attrs["remove_powered_by_okta"] = false
	if brand.RemovePoweredByOkta != nil {
		attrs["remove_powered_by_okta"] = *brand.RemovePoweredByOkta
	}

	return attrs
}
