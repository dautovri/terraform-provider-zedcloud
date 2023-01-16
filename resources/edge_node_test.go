package resources

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	apiclient "github.com/zededa/terraform-provider/client"
	"github.com/zededa/terraform-provider/client/edge_node_configuration"
	"github.com/zededa/terraform-provider/models"
)

func TestAccGetEdgeNode(t *testing.T) {
	var edgeNodeBefore models.DeviceConfig
	// var edgeNodeAfter models.DeviceConfig

	createConfigPath := "edge_node/required_only.tf"
	configCreate, err := getTestConfig(createConfigPath)
	if err != nil {
		t.Fatal(fmt.Sprintf("could not get testdata for %s", createConfigPath))
	}
	// removeCreateConfigPath := "edge_node/remove_required_only.tf"
	// removeConfigCreate, err := getTestConfig(removeCreateConfigPath)
	// if err != nil {
	// 	t.Fatal(fmt.Sprintf("could not get testdata for %s", removeCreateConfigPath))
	// }

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		// CheckDestroy: testEdgeNodeDestroy,
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configCreate,
				Check: resource.ComposeTestCheckFunc(
					testEdgeNodeExists("zedcloud_edgenode.required_only", &edgeNodeBefore),
				),
			},
			// {
			// 	Config: removeConfigCreate,
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testEdgeNodeExists("zedcloud_edgenode.remove_required_only", &edgeNodeAfter),
			// 	),
			// },
		},
	})
}

// testEdgeNodeDestroy verifies the EdgeNode has been destroyed.
func testEdgeNodeDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testAccProvider.Meta().(*apiclient.Zedcloudapi)

	// loop through the resources in state, verifying each EdgeNode is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_edgenode" {
			continue
		}

		// retrieve the EdgeNode by referencing it's state ID for API lookup
		params := edge_node_configuration.NewEdgeNodeConfigurationGetEdgeNodeParams()
		params.ID = rs.Primary.ID
		response, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationGetEdgeNode(params, nil)
		if err == nil {
			return fmt.Errorf("could not fetch EdgeNode (%s): %w", rs.Primary.ID, err)
		}

		// if the error is equivalent to 404 not found, the EdgeNode is destroyed.
		if response.IsCode(http.StatusNotFound) {
			return err
		}

		// 		deviceConfig := response.GetPayload()
		// 		if err == nil {
		// 			if deviceConfig != nil && deviceConfig.ID == rs.Primary.ID {
		// 				return fmt.Errorf("EdgeNode (%s) still exists.", rs.Primary.ID)
		// 			}
		// 			return nil
		// }
	}
	return nil
}

// testAccPreCheck validates the necessary test API keys exist
// in the testing environment
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("TF_CLI_CONFIG_FILE"); v == "" {
		t.Fatal("TF_CLI_CONFIG_FILE must be set for acceptance tests, it should contain the dev_overrides config that points to local instance of the provider")
	}
	if v := os.Getenv("TF_VAR_zedcloud_token"); v == "" {
		t.Fatal("TF_VAR_zedcloud_token must be set for acceptance tests to access the zedcloud API")
	}
}

// testEdgeNodeExists retrieves the EdgeNode and stores it in the provided *models.DeviceConfig.
func testEdgeNodeExists(resourceName string, device *models.DeviceConfig) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("EdgeNode not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("EdgeNode ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testAccProvider.Meta().(*apiclient.Zedcloudapi)

		// retrieve the EdgeNode by referencing it's state ID for API lookup
		params := edge_node_configuration.NewEdgeNodeConfigurationGetEdgeNodeParams()
		params.ID = rs.Primary.ID
		response, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationGetEdgeNode(params, nil)
		if err == nil {
			return fmt.Errorf("could not fetch EdgeNode (%s): %w", rs.Primary.ID, err)
		}

		deviceConfig := response.GetPayload()
		if err != nil {
			return fmt.Errorf("could not get response payload in EdgeNode existence test: %w", err)
		}
		if deviceConfig == nil {
			return errors.New("could not get response payload in EdgeNode existence test: deviceConfig is nil")
		}

		// store the resulting widget in the *example.WidgetDescription pointer
		*device = *deviceConfig
		return nil
	}
}
