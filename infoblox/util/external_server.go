package util

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common"
)

// ExternalServerListSchema - returns the schema for a list of external servers
func ExternalServerListSchema(optional, required bool) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "The primary preference list with Grid member names and/or External Server structs for this member.",
		Optional:    optional,
		Required:    required,
		Elem:        externalServerSchema(),
	}
}

// ExternalServerSetSchema - returns the schema for a set of external servers
func ExternalServerSetSchema(optional, required bool) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Description: "The primary preference set with Grid member names and/or External Server structs for this member.",
		Optional:    optional,
		Required:    required,
		Elem:        externalServerSchema(),
	}
}

// externalServerSchema - returns an external server resource
func externalServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Description: "The IPv4 Address or IPv6 Address of the server.",
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "A resolvable domain name for the external DNS server.",
				Required:    true,
			},
			"shared_with_ms_parent_delegation": {
				Type:        schema.TypeBool,
				Description: "This flag represents whether the name server is shared with the parent Microsoft primary zone’s delegation server.",
				Optional:    true,
				Computed:    true,
			},
			"stealth": {
				Type:        schema.TypeBool,
				Description: "Set this flag to hide the NS record for the primary name server from DNS queries.",
				Optional:    true,
			},
			"tsig_key": {
				Type:        schema.TypeString,
				Description: "A generated TSIG key. Values with leading or trailing whitespace are not valid for this field.",
				Optional:    true,
			},
			"tsig_key_alg": {
				Type:        schema.TypeString,
				Description: "The TSIG key algorithm. Valid values: HMAC-MD5 or HMAC-SHA256. The default value is HMAC-MD5.",
				Optional:    true,
				Computed:    true,
			},
			"tsig_key_name": {
				Type:        schema.TypeString,
				Description: "The TSIG key name.",
				Optional:    true,
			},
			"use_tsig_key_name": {
				Type:        schema.TypeBool,
				Description: "Use flag for: tsig_key_name",
				Optional:    true,
			},
		},
	}
}

// BuildExternalServerListFromT - Builds a list of external servers given the corresponding list of
// items from state
func BuildExternalServerListFromT(extServerListFromT []map[string]interface{}) []common.ExternalServer {

	es := []common.ExternalServer{}
	for _, item := range extServerListFromT {
		var extServer common.ExternalServer

		if v, ok := item["address"]; ok {
			extServer.Address = v.(string)
		}

		if v, ok := item["name"]; ok {
			extServer.Name = v.(string)
		}

		if v, ok := item["stealth"]; ok {
			b := v.(bool)
			extServer.Stealth = &b
		}

		if v, ok := item["tsig_key"]; ok {
			extServer.TsigKey = v.(string)
		}

		if v, ok := item["tsig_key_alg"]; ok {
			extServer.TsigKeyAlg = v.(string)
		}

		if v, ok := item["tsig_key_name"]; ok {
			extServer.TsigKeyName = v.(string)
		}

		if v, ok := item["use_tsig_key_name"]; ok {
			b := v.(bool)
			extServer.UseTsigKeyName = &b
		}

		es = append(es, extServer)
	}
	return es
}

// BuildExternalServerSetFromT : Builds a list of external servers from a template set
func BuildExternalServerSetFromT(extServerSetFromT *schema.Set) []common.ExternalServer {

	var externalServer common.ExternalServer
	externalServers := make([]common.ExternalServer, 0)
	for _, item := range extServerSetFromT.List() {
		castItem := item.(map[string]interface{})
		if v, ok := castItem["address"]; ok && v != "" {
			externalServer.Address = castItem["address"].(string)
		}
		if v, ok := castItem["name"]; ok && v != "" {
			externalServer.Name = castItem["name"].(string)
		}
		if v, ok := castItem["stealth"]; ok {
			b := v.(bool)
			externalServer.Stealth = &b
		}
		if v, ok := castItem["tsig_key"]; ok && v != "" {
			externalServer.TsigKey = castItem["tsig_key"].(string)
		}
		if v, ok := castItem["tsig_key_alg"]; ok && v != "" {
			externalServer.TsigKeyAlg = castItem["tsig_key_alg"].(string)
		}
		if v, ok := castItem["tsig_key_name"]; ok && v != "" {
			externalServer.TsigKeyName = castItem["tsig_key_name"].(string)
		}
		if v, ok := castItem["use_tsig_key_name"]; ok {
			b := v.(bool)
			externalServer.UseTsigKeyName = &b
		}
		externalServers = append(externalServers, externalServer)
	}
	return externalServers
}

// BuildExternalServersListFromIBX - builds a list of external servers for terraform
// given the corresponding struct from IBX
func BuildExternalServersListFromIBX(IBXExtServersList []common.ExternalServer) []map[string]interface{} {
	es := make([]map[string]interface{}, 0)
	for _, IBXExtServer := range IBXExtServersList {
		server := make(map[string]interface{})

		if IBXExtServer.Address != "" {
			server["address"] = IBXExtServer.Address
		}

		if IBXExtServer.Name != "" {
			server["name"] = IBXExtServer.Name
		}

		if IBXExtServer.Stealth != nil {
			server["stealth"] = *IBXExtServer.Stealth
		}

		if IBXExtServer.TsigKey != "" {
			server["tsig_key"] = IBXExtServer.TsigKey
		}

		if IBXExtServer.TsigKeyAlg != "" {
			server["tsig_key_alg"] = IBXExtServer.TsigKeyAlg
		}

		if IBXExtServer.TsigKeyName != "" {
			server["tsig_key_name"] = IBXExtServer.TsigKeyName
		}

		if IBXExtServer.UseTsigKeyName != nil {
			server["use_tsig_key_name"] = *IBXExtServer.UseTsigKeyName
		}

		es = append(es, server)
	}

	return es
}
