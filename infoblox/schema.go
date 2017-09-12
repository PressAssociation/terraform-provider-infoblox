package infoblox

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

type ResourceAttr struct {
	Name  string
	Type  schema.ValueType
	Value interface{}
}

//GetAttrs - returns the list of attributes names and types
func GetAttrs(resource *schema.Resource) []ResourceAttr {
	attrs := make([]ResourceAttr, 0)

	s := resource.Schema

	str := spew.Sdump(s)
	log.Println("Schema:\n", str)

	for key, _ := range s {
		var attr ResourceAttr
		attr.Name = key
		attr.Type = s[key].Type
		attrs = append(attrs, attr)
	}
	return attrs
}

func GetValue(attr ResourceAttr) interface{} {
    var value interface{}

	switch attr.Type {
	case schema.TypeSet:
		v := attr.Value.(*schema.Set)
		value = v.List()
	}
	return attr.Value
}