// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
	"log"
)

var EdgeAppDataSourceSchema = &schema.Resource{
	ReadContext: readEdgeApp,
	Schema:      zschemas.EdgeAppSchema,
	Description: "Schema for data source zedcloud_edgeapp. Must specify id or name",
}

// The schema for a resource group data source
func getEdgeAppDataSourceSchema() *schema.Resource {
	return EdgeAppDataSourceSchema
}

func getEdgeApp(client *zedcloudapi.Client,
	name, id string) (*swagger_models.App, error) {
	rspData := &swagger_models.App{}
	err := client.GetObj(edgeAppUrlExtension, name, id, false, rspData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] FAILED to get EdgeApp %s ( id: %s). Err: %s",
			name, id, err.Error())
	}
	return rspData, nil
}

func flattenEdgeAppConfig(cfg *swagger_models.App) map[string]interface{} {
	return map[string]interface{}{
		"name":                 ptrValStr(cfg.Name),
		"id":                   cfg.ID,
		"title":                ptrValStr(cfg.Title),
		"description":          cfg.Description,
		"cpus":                 int(cfg.Cpus),
		"drives":               int(cfg.Drives),
		"memory":               int(cfg.Memory),
		"networks":             int(cfg.Networks),
		"origin_type":          ptrValStr(cfg.OriginType),
		"parent_detail":        flattenObjectParentDetail(cfg.ParentDetail),
		"revision":             flattenObjectRevision(cfg.Revision),
		"storage":              int(cfg.Storage),
		"user_defined_version": cfg.UserDefinedVersion,
	}
}

// Read the Resource Group
func readEdgeApp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	id := rdEntryStr(d, "id")
	name := rdEntryStr(d, "name")

	if client == nil {
		return diag.Errorf("nil Client.")
	}
	log.Printf("PROVIDER:  readEdgeAppResource - id: %s, name: %s", id, name)
	if (id == "") && (name == "") {
		return diag.Errorf("The arguments \"id\" or \"name\" are required, but no definition was found.")
	}
	cfg, err := getEdgeApp(client, name, id)
	if err != nil {
		return diag.Errorf("[ERROR] Edge App %s (id: %s) not found. Err: %s",
			name, id, err.Error())
	}

	// Take the Config and convert it to terraform object
	marshalData(d, flattenEdgeAppConfig(cfg))
	return diags
}