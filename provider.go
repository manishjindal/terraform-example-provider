package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// func Provider() *schema.Provider {
// 	return &schema.Provider{
// 		ResourcesMap: map[string]*schema.Resource{},
// 	}
// }

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"example_hello":  resourceHello(),
			"example_server": resourceServer(),
		},
	}
}
