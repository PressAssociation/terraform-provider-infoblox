package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"testing"
)

func TestAccResourceARecord(t *testing.T) {

	randInt := acctest.RandInt()
	recordName := fmt.Sprintf("a-record-test-%d.slupaas.bskyb.com", randInt)
	resourceName := "infoblox_arecord.acctest"

	fmt.Printf("\n\nAcc Test record name is %s\n\n", recordName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.RecordAObj, "name", recordName)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceARecordCreateTemplate(recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceARecordExists(recordName, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", recordName),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.0.0.10"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "9000"),
				),
			},
			{
				Config: testAccResourceARecordUpdateTemplate(recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceARecordExists(recordName, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", recordName),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.0.0.10"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "900"),
				),
			},
		},
	})
}

/*
func testAccResourceARecordDestroy(state *terraform.State, recordName string) error {

	client := GetClient()
	recs, err := client.ReadAll(model.RecordAObj)
	if err != nil {
		return err
	}
	for _, rec := range recs {
		if rec["name"] == recordName {
			return fmt.Errorf("A record %s still exists!!", recordName)
		}
	}
	return nil
}
*/
func testAccResourceARecordExists(recordName, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("\nInfoblox A record resource %s not found in resources: ", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("\nInfoblox A record resource %s ID not set", resourceName)
		}
		client := GetClient()
		recs, err := client.ReadAll(model.RecordAObj)
		if err != nil {
			return fmt.Errorf("Error getting the A records: %+v", err)
		}
		for _, x := range recs {
			if x["name"] == recordName {
				return nil
			}
		}
		return fmt.Errorf("Could not find %s", recordName)

	}
}

func testAccResourceARecordCreateTemplate(arecordName string) string {
	return fmt.Sprintf(`
	resource "infoblox_arecord" "acctest"{
	name = "%s"
	ipv4addr = "10.0.0.10"
	ttl = 9000
	}`, arecordName)
}

func testAccResourceARecordUpdateTemplate(arecordName string) string {
	return fmt.Sprintf(`
	resource "infoblox_arecord" "acctest"{
	name = "%s"
	ipv4addr = "10.0.0.10"
	ttl = 900
    use_ttl = false
	}`, arecordName)
}
