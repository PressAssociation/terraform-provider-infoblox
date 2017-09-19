package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"testing"
)

func TestAccAdminUserResource(t *testing.T) {
	randomInt := acctest.RandIntRange(1, 10000)
	recordUserName := fmt.Sprintf("testadminuser%d", randomInt)
	resourceName := "infoblox_admin_user.testadmin"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.AdminuserObj, "name", recordUserName)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAdminUserNameCreateTemplate(recordUserName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceAdminUserExists("name", recordUserName),
					resource.TestCheckResourceAttr(resourceName, "name", recordUserName),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is a comment"),
					resource.TestCheckResourceAttr(resourceName, "email", "exampleuser@domain.internal.com"),
					//resource.TestCheckResourceAttr(resourceName, "admin_groups", []string{"APP-OVP-INFOBLOX-READONLY"}),
				),
			}, {
				Config: testAccResourceAdminUserNameUpdateTemplate(recordUserName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceAdminUserExists("name", recordUserName),
					resource.TestCheckResourceAttr(resourceName, "name", recordUserName),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is a comment updated"),
					resource.TestCheckResourceAttr(resourceName, "email", "user@domain.internal.com"),
					//resource.TestCheckResourceAttr(resourceName, "admin_groups", []string{"APP-OVP-INFOBLOX-READONLY"}),
				),
			},
		},
	})

}

func testAccResourceAdminUserExists(key, value string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.AdminuserObj, key, value)
	}
}

func testAccResourceAdminUserNameCreateTemplate(username string) string {
	return fmt.Sprintf(`
	resource "infoblox_admin_user" "testadmin" {
	name = "%s"
	comment = "this is a comment"
	email = "exampleuser@domain.internal.com"
	admin_groups = ["APP-OVP-INFOBLOX-READONLY"]
	password = "c0a6264f0f128d94cd8ef26652e7d9fd"}`, username)
}

func testAccResourceAdminUserNameUpdateTemplate(username string) string {
	return fmt.Sprintf(`
	resource "infoblox_admin_user" "testadmin" {
  		name = "%s"
		comment = "this is a comment updated"
		email = "user@domain.internal.com"
		admin_groups = ["APP-OVP-INFOBLOX-READONLY"]
		password = "c0a6264f0f128d94cd8ef26652e7d9fd"
	}
	`, username)
}
