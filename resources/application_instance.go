// // Code generated by go-swagger; DO NOT EDIT.
package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"errors"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider/client"
	config "github.com/zededa/terraform-provider/client/edge_application_instance_configuration"
	"github.com/zededa/terraform-provider/models"
	zschema "github.com/zededa/terraform-provider/schemas"
)

func ApplicationInstanceResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateApplicationInstance,
		ReadContext:   GetApplicationInstanceByName,
		UpdateContext: UpdateApplicationInstance,
		DeleteContext: DeleteApplicationInstance,
		Schema:        zschema.ApplicationInstance(),
	}
}

func ApplicationInstanceDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: GetApplicationInstanceByName,
		Schema:      zschema.ApplicationInstance(),
	}
}

func CreateApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.ApplicationInstanceModel(d)
	params := config.CreateParams()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ApplicationInstance.Create(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}

	// note, we need to set the ID in any case, even GetByName endpoint seems to require items
	// but doesn't return any error if it's not set.
	d.SetId(responseData.ObjectID)

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetApplicationInstanceByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func GetApplicationInstanceByName(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.GetByNameParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	nameVal, nameIsSet := d.GetOk("name")
	if nameIsSet {
		params.Name = nameVal.(string)
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: name")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ApplicationInstance.GetByName(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	appInstance := resp.GetPayload()
	zschema.SetApplicationInstanceResourceData(d, appInstance)
	d.SetId(appInstance.ID)

	return diags
}

func UpdateApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.UpdateParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	params.SetBody(zschema.ApplicationInstanceModel(d))

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ApplicationInstance.Update(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return append(diags, diag.Errorf("unexpected: %s", err)...)
	}

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}

	// note, we need to set the ID in any case, even GetByName endpoint seems to require items
	// but doesn't return any error if it's not set.
	d.SetId(responseData.ObjectID)

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	if errs := GetApplicationInstanceByName(ctx, d, m); err != nil {
		return append(diags, errs...)
	}

	return diags
}

func DeleteApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := config.DeleteParams()

	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
	if xRequestIdIsSet {
		params.XRequestID = xRequestIdVal.(*string)
	}

	idVal, idIsSet := d.GetOk("id")
	if idIsSet {
		id, _ := idVal.(string)
		params.ID = id
	} else {
		diags = append(diags, diag.Errorf("missing client parameter: id")...)
		return diags
	}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.ApplicationInstance.Delete(params, nil)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diags
	}

	d.SetId("")
	return diags
}

// func EdgeApplicationInstanceConfiguration_ActivateEdgeApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := edge_application_instance_configuration.NewEdgeApplicationInstanceConfigurationActivateEdgeApplicationInstanceParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.EdgeApplicationInstanceConfiguration.EdgeApplicationInstanceConfigurationActivateEdgeApplicationInstance(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func EdgeApplicationInstanceConfiguration_ConnectToEdgeApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := edge_application_instance_configuration.NewEdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstanceParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.EdgeApplicationInstanceConfiguration.EdgeApplicationInstanceConfigurationConnectToEdgeApplicationInstance(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func EdgeApplicationInstanceConfiguration_DeActivateEdgeApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := edge_application_instance_configuration.NewEdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstanceParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.EdgeApplicationInstanceConfiguration.EdgeApplicationInstanceConfigurationDeActivateEdgeApplicationInstance(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func EdgeApplicationInstanceConfiguration_RefreshEdgeApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := edge_application_instance_configuration.NewEdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstanceParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.EdgeApplicationInstanceConfiguration.EdgeApplicationInstanceConfigurationRefreshEdgeApplicationInstance(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func EdgeApplicationInstanceConfiguration_RefreshPurgeEdgeApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := edge_application_instance_configuration.NewEdgeApplicationInstanceConfigurationRefreshPurgeEdgeApplicationInstanceParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.EdgeApplicationInstanceConfiguration.EdgeApplicationInstanceConfigurationRefreshPurgeEdgeApplicationInstance(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }

// func EdgeApplicationInstanceConfiguration_RestartEdgeApplicationInstance(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	d.Partial(true)

// 	params := edge_application_instance_configuration.NewEdgeApplicationInstanceConfigurationRestartEdgeApplicationInstanceParams()

// 	xRequestIdVal, xRequestIdIsSet := d.GetOk("x_request_id")
// 	if xRequestIdIsSet {
// 		params.XRequestID = xRequestIdVal.(*string)
// 	}

// 	idVal, idIsSet := d.GetOk("id")
// 	if idIsSet {
// 		id, _ := idVal.(string)
// 		params.ID = id
// 	} else {
// 		diags = append(diags, diag.Errorf("missing client parameter: id")...)
// 		return diags
// 	}

// 	// makes a bulk update for all properties that were changed
// 	client := m.(*apiclient.Zedcloudapi)
// 	resp, err := client.EdgeApplicationInstanceConfiguration.EdgeApplicationInstanceConfigurationRestartEdgeApplicationInstance(params, nil)
// 	log.Printf("[TRACE] response: %v", resp)
// 	if err != nil {
// 		return append(diags, diag.Errorf("unexpected: %s", err)...)
// 	}

// 	responseData := resp.GetPayload()
// 	if responseData != nil && len(responseData.Error) > 0 {
// 		for _, err := range responseData.Error {
// 			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
// 		}
// 		return diags
// 	}

// 	// the zedcloud API does not return the partially updated object but a custom response.
// 	// thus, we need to fetch the object and populate the state.
// 	if errs := GetDevice(ctx, d, m); err != nil {
// 		return append(diags, errs...)
// 	}

// 	d.Partial(false)

// 	return diags
// }