package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"testing"
)

func TestAccInfobloxPermissionBasic(t *testing.T) {

	randomInt := acctest.RandInt()
	adminRoleName := fmt.Sprintf("acctest-infoblox-permission-role-%d", randomInt)
	permissionResource := "infoblox_permission.permission_acctest"

	testPermission := model.Permission{
		Role:         adminRoleName,
		Permission:   "READ",
		ResourceType: "AAAA",
	}

	updatedTestPermission := model.Permission{
		Role:         adminRoleName,
		Permission:   "WRITE",
		ResourceType: "AAAA",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccInfobloxPermissionCheckDestroy(state, testPermission)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccInfobloxPermissionCreateTemplate(adminRoleName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxPermissionCheckExists(testPermission),
					resource.TestCheckResourceAttr(permissionResource, "role", adminRoleName),
					resource.TestCheckResourceAttr(permissionResource, "permission", "READ"),
					resource.TestCheckResourceAttr(permissionResource, "resource_type", "AAAA"),
				),
			},
			{
				Config: testAccInfobloxPermissionUpdateTemplate(adminRoleName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxPermissionCheckExists(updatedTestPermission),
					resource.TestCheckResourceAttr(permissionResource, "role", adminRoleName),
					resource.TestCheckResourceAttr(permissionResource, "permission", "WRITE"),
					resource.TestCheckResourceAttr(permissionResource, "resource_type", "AAAA"),
				),
			},
		},
	})
}

func testAccInfobloxPermissionCheckExists(testPermision model.Permission) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		client := GetClient()
		recs, err := client.ReadAll(model.PermissionObj)
		if err != nil {
			return fmt.Errorf("Error retrieving the list of permission objects: ", err)
		}
		for _, permission := range recs {
			if permission["role"] == testPermision.Role {
				if permission["permission"] == testPermision.Permission {
					if permission["resource_type"] == testPermision.ResourceType {
						return nil
					}
				}
			}
		}
		return fmt.Errorf("Permission wasn't found on remote Infoblox server")
	}
}

func testAccInfobloxPermissionCheckDestroy(state *terraform.State, testPermision model.Permission) error {
	client := GetClient()
	recs, err := client.ReadAll(model.PermissionObj)
	if err != nil {
		return fmt.Errorf("Infoblox - error occurred whilst retrieving a list of permissions")
	}
	for _, permission := range recs {
		if permission["role"] == testPermision.Role {
			if permission["permission"] == testPermision.Permission {
				if permission["resource_type"] == testPermision.ResourceType {
					return fmt.Errorf("Infoblox Permission still exists")
				}
			}
		}
	}

	return nil
}

func testAccInfobloxPermissionCreateTemplate(roleName string) string {
	return fmt.Sprintf(`
resource "infoblox_admin_role" "role_acctest" {
name = "%s"
comment = "Infoblox Terraform Role for Permission Acceptance test"
disable = true
}

resource "infoblox_permission" "permission_acctest" {
role = "${infoblox_admin_role.role_acctest.name}"
permission = "READ"
resource_type = "AAAA"
}
`, roleName)

}

func testAccInfobloxPermissionUpdateTemplate(roleName string) string {
	return fmt.Sprintf(`
resource "infoblox_admin_role" "role_acctest" {
name = "%s"
comment = "Infoblox Terraform Role for Permission Acceptance test"
disable = true
}

resource "infoblox_permission" "permission_acctest" {
role = "${infoblox_admin_role.role_acctest.name}"
permission = "WRITE"
resource_type = "AAAA"
}
`, roleName)
}
