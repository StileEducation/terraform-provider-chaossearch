package tests

import (
	"context"
	"cs-tf-provider/client"
	"cs-tf-provider/provider/models"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccView(t *testing.T) {
	bucketName := generateName("acc-test-tf-provider-view-og")
	viewName := generateName("acc-test-tf-provider-view")
	resourceName := "chaossearch_view.create-view"
	resource.Test(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccViewConfig(viewName, bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccViewExists(resourceName, bucketName),
				),
			},
		},
	})
}

func TestAccViewPreds(t *testing.T) {
	bucketName := generateName("acc-test-tf-provider-view-og")
	viewName := generateName("acc-test-tf-provider-view-preds")
	resourceName := "chaossearch_view.create-view"
	resource.Test(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccViewPredsConfig(viewName, bucketName),
				Check: resource.ComposeTestCheckFunc(
					testAccViewExists(resourceName, bucketName),
				),
			},
		},
	})
}

func testAccViewConfig(viewName, bucketName string) string {
	return fmt.Sprintf(`
		%s
	    resource "chaossearch_view" "create-view" {
			bucket           = "%s"
			case_insensitive = false
			index_pattern    = ".*"
			index_retention  = -1
			overwrite        = true
			sources          = ["%s"]
			time_field_name  = "timestamp"
			filter {
			  predicate {
				type = "chaossumo.query.NIRFrontend.Request.Predicate.Negate"
				pred {
				  type = "chaossumo.query.NIRFrontend.Request.Predicate.TextMatch"
				  field = "cs_partition_key_0"
				  query = "Test"
				  state {
					type = "chaossumo.query.QEP.Predicate.TextMatchState.Exact"
				  }
				}
			  }
			}
			depends_on = [
				chaossearch_index_model.model
			]
		  }
	`, testAccIndexConfig(bucketName), viewName, bucketName)
}

func testAccViewPredsConfig(viewName, bucketName string) string {
	return fmt.Sprintf(`
	%s
	resource "chaossearch_view" "create-view" {
		bucket           = "%s"
		case_insensitive = false
		index_pattern    = ".*"
		index_retention  = -1
		overwrite        = true
		sources          = ["%s"]
		time_field_name  = "@timestamp"
		filter {
		  predicate {
			type = "chaossumo.query.NIRFrontend.Request.Predicate.Or"
			preds = [
				jsonencode(
				  {
					"_type" = "chaossumo.query.NIRFrontend.Request.Predicate.TextMatch",
					"field" = "Subject",
					"query" = "Test"
					"state" = {
					  "_type" = "chaossumo.query.QEP.Predicate.TextMatchState.Exact"
					},
				  }
				),
				jsonencode(
				  {
					"_type" = "chaossumo.query.NIRFrontend.Request.Predicate.TextMatch",
					"field" = "Series_title_1",
					"query" = "Test2"
					"state" = {
					  "_type" = "chaossumo.query.QEP.Predicate.TextMatchState.Exact"
					},
				  }
				)
			  ]
		  }
		}
		depends_on = [
			chaossearch_index_model.model
		]
	  }
`, testAccIndexConfig(bucketName), viewName, bucketName)
}

func testAccViewExists(resourceName, bucketName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if res.Primary.ID == "" {
			return fmt.Errorf("View ID is not set")
		}

		providerMeta := testAccProvider.Meta().(*models.ProviderMeta)
		csClient := providerMeta.CSClient
		ctx := context.Background()
		response, err := csClient.ReadView(ctx, &client.BasicRequest{
			AuthToken: providerMeta.Token,
			Id:        bucketName,
		})

		if err != nil {
			return err
		}

		if response.ID == "" && response.Bucket != bucketName {
			return fmt.Errorf("View not found")
		}

		return nil
	}
}
