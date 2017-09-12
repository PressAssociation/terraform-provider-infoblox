package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

// CreateResource - Creates a new resource provided its resource schema
func CreateResource(name string, resource *schema.Resource, d *schema.ResourceData, m interface{}) error {

	obj := make(map[string]interface{})
	attrs := GetAttrs(resource)
	for _, attr := range attrs {
		key := attr.Name
		log.Println("Found attribute: ", key)
		if v, ok := d.GetOk(key); ok {
			attr.Value = v
			obj[key] = GetValue(attr)
		}
	}

	client := GetClient()

	log.Printf("Going to create an %s object: %+v", name, obj)
	ref, err := client.Create(name, obj)
	if err != nil {
		d.SetId("")
		return err
	}

	d.SetId(ref)
	return ReadResource(resource, d, m)
}

// ReadResource - Reads a resource provided its resource schema
func ReadResource(resource *schema.Resource, d *schema.ResourceData, m interface{}) error {

	client := GetClient()

	ref := d.Id()
	obj := make(map[string]interface{})

	attrs := GetAttrs(resource)
	keys := []string{}
	for _, attr := range attrs {
		keys = append(keys, attr.Name)
	}
	err := client.Read(ref, keys, &obj)
	if err != nil {
		d.SetId("")
		return err
	}

	delete(obj, "_ref") // TODO  should we move this to the binding side ?
	for key := range obj {
		d.Set(key, obj[key])
	}

	return nil
}

// DeleteResource - Deletes a resource
func DeleteResource(d *schema.ResourceData, m interface{}) error {

	client := GetClient()

	ref := d.Id()
	ref, err := client.Delete(ref)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

// UpdateResource - Updates a resource provided its schema
func UpdateResource(resource *schema.Resource, d *schema.ResourceData, m interface{}) error {

	needsUpdate := false

	client := GetClient()

	ref := d.Id()
	obj := make(map[string]interface{})

	attrs := GetAttrs(resource)
	for _, attr := range attrs {
		key := attr.Name
		if d.HasChange(key) {
			attr.Value = d.Get(key)
			obj[key] = GetValue(attr)
			log.Printf("Updating field %s, value: %+v\n", key, obj[key])
			needsUpdate = true
		}
	}

	log.Printf("UPDATE: going to update reference %s with obj: \n%+v\n", obj)

	if needsUpdate {
		ref, err := client.Update(ref, obj)
		if err != nil {
			return err
		}
		d.SetId(ref)
	}

	return nil
}
