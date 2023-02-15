package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccOrganizationResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccOrganizationResourceConfig("one"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("supabase_organization.test", "name", "one"),
					resource.TestCheckResourceAttr("supabase_organization.test", "id", "example-id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "supabase_organization.test",
				ImportState:       true,
				ImportStateVerify: true,
				// This is not normally necessary, but is here because this
				// example code does not have an actual upstream service.
				// Once the Read method is able to refresh information from
				// the upstream service, this can be removed.
				ImportStateVerifyIgnore: []string{"name"},
			},
			// Update and Read testing
			{
				Config: testAccOrganizationResourceConfig("two"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("supabase_organization.test", "name", "two"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccOrganizationResourceConfig(configurableAttribute string) string {
	return fmt.Sprintf(`
resource "supabase_organization" "test" {
  name = %[1]q
}
`, configurableAttribute)
}
