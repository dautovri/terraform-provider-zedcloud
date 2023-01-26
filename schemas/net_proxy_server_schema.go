package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func NetworkProxyServerModel(d *schema.ResourceData) *models.NetProxyServer {
	portInt, _ := d.Get("port").(int)
	port := int64(portInt)
	var proto *models.NetworkProxyProto // NetworkProxyProto
	protoInterface, protoIsSet := d.GetOk("proto")
	if protoIsSet {
		protoModel := protoInterface.(string)
		proto = models.NewNetworkProxyProto(models.NetworkProxyProto(protoModel))
	}
	server, _ := d.Get("server").(string)
	return &models.NetProxyServer{
		Port:   port,
		Proto:  proto,
		Server: server,
	}
}

func NetworkProxyServerModelFromMap(m map[string]interface{}) *models.NetProxyServer {
	port := int64(m["port"].(int))      // int64
	var proto *models.NetworkProxyProto // NetworkProxyProto
	protoInterface, protoIsSet := m["proto"]
	if protoIsSet {
		protoModel := protoInterface.(string)
		proto = models.NewNetworkProxyProto(models.NetworkProxyProto(protoModel))
	}
	server := m["server"].(string)
	return &models.NetProxyServer{
		Port:   port,
		Proto:  proto,
		Server: server,
	}
}

func SetworkNetProxyServerResourceData(d *schema.ResourceData, m *models.NetProxyServer) {
	d.Set("port", m.Port)
	d.Set("proto", m.Proto)
	d.Set("server", m.Server)
}

func SetNetworkProxyServerSubResourceData(m []*models.NetProxyServer) (d []*map[string]interface{}) {
	for _, NetProxyServerModel := range m {
		if NetProxyServerModel != nil {
			properties := make(map[string]interface{})
			properties["port"] = NetProxyServerModel.Port
			properties["proto"] = NetProxyServerModel.Proto
			properties["server"] = NetProxyServerModel.Server
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetProxyServer resource defined in the Terraform configuration
func NetworkProxyServer() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"port": {
			Description: `Net Proxy Port`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"proto": {
			Description: `Net Proxy protocol:
NETWORK_PROXY_PROTO_HTTP
NETWORK_PROXY_PROTO_HTTPS
NETWORK_PROXY_PROTO_SOCKS
NETWORK_PROXY_PROTO_FTP
NETWORK_PROXY_PROTO_OTHER`,
			Type:     schema.TypeString,
			Optional: true,
		},

		"server": {
			Description: `Net Proxy Server`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the NetProxyServer resource
func NetworkProxyServerPropertyFields() (t []string) {
	return []string{
		"port",
		"proto",
		"server",
	}
}
