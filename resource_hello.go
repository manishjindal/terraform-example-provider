package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceHello() *schema.Resource {
	return &schema.Resource{
		Create: resourceHelloCreate,
		Read:   resourceHelloRead,
		Update: resourceHelloUpdate,
		Delete: resourceHelloDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceHelloCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	d.SetId(address)
	return nil
}

func resourceHelloRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceHelloUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceHelloDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
