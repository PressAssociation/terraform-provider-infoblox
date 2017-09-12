package infoblox

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

type ResourceAttr struct {
    Name    string
    Type    schema.ValueType
}

func GetAttrs(resource *schema.Resource) []string {
	attrs := make([]string, 0)

	s := resource.Schema

	str := spew.Sdump(s)
	log.Println("Schema:\n", str)

	for key, _ := range s {
		attrs = append(attrs, key)
	}
	return attrs
}
