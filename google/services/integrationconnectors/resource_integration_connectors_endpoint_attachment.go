// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package integrationconnectors

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceIntegrationConnectorsEndpointAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceIntegrationConnectorsEndpointAttachmentCreate,
		Read:   resourceIntegrationConnectorsEndpointAttachmentRead,
		Update: resourceIntegrationConnectorsEndpointAttachmentUpdate,
		Delete: resourceIntegrationConnectorsEndpointAttachmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIntegrationConnectorsEndpointAttachmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Location in which Endpoint Attachment needs to be created.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of Endpoint Attachment needs to be created.`,
			},
			"service_attachment": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The path of the service attachment.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Description of the resource.`,
			},
			"endpoint_global_access": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Enable global access for endpoint attachment.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Resource labels to represent user provided metadata.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Namespace was created in UTC.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"endpoint_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The Private Service Connect connection endpoint ip.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the Namespace was updated in UTC.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceIntegrationConnectorsEndpointAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandIntegrationConnectorsEndpointAttachmentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	serviceAttachmentProp, err := expandIntegrationConnectorsEndpointAttachmentServiceAttachment(d.Get("service_attachment"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service_attachment"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceAttachmentProp)) && (ok || !reflect.DeepEqual(v, serviceAttachmentProp)) {
		obj["serviceAttachment"] = serviceAttachmentProp
	}
	endpointGlobalAccessProp, err := expandIntegrationConnectorsEndpointAttachmentEndpointGlobalAccess(d.Get("endpoint_global_access"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("endpoint_global_access"); !tpgresource.IsEmptyValue(reflect.ValueOf(endpointGlobalAccessProp)) && (ok || !reflect.DeepEqual(v, endpointGlobalAccessProp)) {
		obj["endpointGlobalAccess"] = endpointGlobalAccessProp
	}
	labelsProp, err := expandIntegrationConnectorsEndpointAttachmentEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IntegrationConnectorsBasePath}}projects/{{project}}/locations/{{location}}/endpointAttachments?endpointAttachmentId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new EndpointAttachment: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EndpointAttachment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating EndpointAttachment: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/endpointAttachments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = IntegrationConnectorsOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating EndpointAttachment", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create EndpointAttachment: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/endpointAttachments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating EndpointAttachment %q: %#v", d.Id(), res)

	return resourceIntegrationConnectorsEndpointAttachmentRead(d, meta)
}

func resourceIntegrationConnectorsEndpointAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IntegrationConnectorsBasePath}}projects/{{project}}/locations/{{location}}/endpointAttachments/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EndpointAttachment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IntegrationConnectorsEndpointAttachment %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading EndpointAttachment: %s", err)
	}

	if err := d.Set("create_time", flattenIntegrationConnectorsEndpointAttachmentCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointAttachment: %s", err)
	}
	if err := d.Set("update_time", flattenIntegrationConnectorsEndpointAttachmentUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointAttachment: %s", err)
	}
	if err := d.Set("description", flattenIntegrationConnectorsEndpointAttachmentDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointAttachment: %s", err)
	}
	if err := d.Set("labels", flattenIntegrationConnectorsEndpointAttachmentLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointAttachment: %s", err)
	}
	if err := d.Set("service_attachment", flattenIntegrationConnectorsEndpointAttachmentServiceAttachment(res["serviceAttachment"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointAttachment: %s", err)
	}
	if err := d.Set("endpoint_ip", flattenIntegrationConnectorsEndpointAttachmentEndpointIp(res["endpointIp"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointAttachment: %s", err)
	}
	if err := d.Set("endpoint_global_access", flattenIntegrationConnectorsEndpointAttachmentEndpointGlobalAccess(res["endpointGlobalAccess"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointAttachment: %s", err)
	}
	if err := d.Set("terraform_labels", flattenIntegrationConnectorsEndpointAttachmentTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointAttachment: %s", err)
	}
	if err := d.Set("effective_labels", flattenIntegrationConnectorsEndpointAttachmentEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading EndpointAttachment: %s", err)
	}

	return nil
}

func resourceIntegrationConnectorsEndpointAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EndpointAttachment: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandIntegrationConnectorsEndpointAttachmentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	endpointGlobalAccessProp, err := expandIntegrationConnectorsEndpointAttachmentEndpointGlobalAccess(d.Get("endpoint_global_access"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("endpoint_global_access"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, endpointGlobalAccessProp)) {
		obj["endpointGlobalAccess"] = endpointGlobalAccessProp
	}
	labelsProp, err := expandIntegrationConnectorsEndpointAttachmentEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IntegrationConnectorsBasePath}}projects/{{project}}/locations/{{location}}/endpointAttachments/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating EndpointAttachment %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("endpoint_global_access") {
		updateMask = append(updateMask, "endpointGlobalAccess")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating EndpointAttachment %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating EndpointAttachment %q: %#v", d.Id(), res)
		}

		err = IntegrationConnectorsOperationWaitTime(
			config, res, project, "Updating EndpointAttachment", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceIntegrationConnectorsEndpointAttachmentRead(d, meta)
}

func resourceIntegrationConnectorsEndpointAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for EndpointAttachment: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{IntegrationConnectorsBasePath}}projects/{{project}}/locations/{{location}}/endpointAttachments/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting EndpointAttachment %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "EndpointAttachment")
	}

	err = IntegrationConnectorsOperationWaitTime(
		config, res, project, "Deleting EndpointAttachment", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting EndpointAttachment %q: %#v", d.Id(), res)
	return nil
}

func resourceIntegrationConnectorsEndpointAttachmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/endpointAttachments/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/endpointAttachments/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIntegrationConnectorsEndpointAttachmentCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIntegrationConnectorsEndpointAttachmentUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIntegrationConnectorsEndpointAttachmentDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIntegrationConnectorsEndpointAttachmentLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenIntegrationConnectorsEndpointAttachmentServiceAttachment(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIntegrationConnectorsEndpointAttachmentEndpointIp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIntegrationConnectorsEndpointAttachmentEndpointGlobalAccess(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIntegrationConnectorsEndpointAttachmentTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenIntegrationConnectorsEndpointAttachmentEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIntegrationConnectorsEndpointAttachmentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIntegrationConnectorsEndpointAttachmentServiceAttachment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIntegrationConnectorsEndpointAttachmentEndpointGlobalAccess(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIntegrationConnectorsEndpointAttachmentEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
