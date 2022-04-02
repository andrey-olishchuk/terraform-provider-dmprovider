package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCompany() *schema.Resource {
	return &schema.Resource{
		Create: resourceCompanyCreate,
		Read:   resourceCompanyRead,
		Update: resourceCompanyUpdate,
		Delete: resourceCompanyDelete,

		Schema: map[string]*schema.Schema{
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCompanyCreate(d *schema.ResourceData, m interface{}) error {
	title := d.Get("title").(string)

	postBody, _ := json.Marshal(map[string]string{
		"title": title,
	})
	responseBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("https://api-nesebar.dorpm.sbs/api/rest/company", "application/json", responseBody)

	//Handle Error
	if err != nil {
		return err
	}

	b, err := io.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal(b, &result)
	company := result["insert_company_one"].(map[string]interface{})

	d.SetId(fmt.Sprint(company["id"]))

	defer resp.Body.Close()

	return resourceCompanyRead(d, m)
}

func resourceCompanyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCompanyUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceCompanyRead(d, m)
}

func resourceCompanyDelete(d *schema.ResourceData, m interface{}) error {
	resourceID := d.Id()
	url := "https://api-nesebar.dorpm.sbs/api/rest/company/delete"

	postBody, _ := json.Marshal(map[string]string{
		"id": resourceID,
	})

	responseBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(url, "application/json", responseBody)

	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	d.SetId("")
	return nil
}
