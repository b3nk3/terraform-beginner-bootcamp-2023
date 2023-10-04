// In Go, the main package is a special package because it defines an executable program.
// It's the package where the main function resides, which serves as the entry point for your program when it is run.
package main

// Import the "fmt" package, which provides input/output functions.
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

type Config struct {
	Endpoint string
	Token    string
	UserUuid string
}

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
	// Format.PrintLine
	// Prints to standard output
	fmt.Println("Hello, world!")
}

// Provider - title case functions get exported
func Provider() *schema.Provider {
	var p *schema.Provider
	p = &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"terratowns_home": HomeResource(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Endpoint for the external service",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true, // make token sensitive - hide from logs
				Description: "Bearer token for auth",
			},
			"user_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "UUID for configuration",
				ValidateFunc: validateUUID,
			},
		},
	}
	p.ConfigureContextFunc = providerConfig(p)
	return p
}

func HomeResource() *schema.Resource {
	log.Print("Resource:start")
	resource := &schema.Resource{
		CreateContext: resourceHomeCreate,
		ReadContext:   resourceHomeRead,
		UpdateContext: resourceHomeUpdate,
		DeleteContext: resourceHomeDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of home",
			},
			"content_version": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Content version",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Description of home",
			},
			"domain_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain name of home e.g. *.cloudfront.net",
			},
			"town": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The town the home will belong to",
			},
		},
	}
	log.Print("Resource:end")
	return resource
}

func resourceHomeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("resourceHomeCreate:start")
	var diags diag.Diagnostics
	config := m.(*Config)

	// construct payload
	payload := map[string]interface{}{
		"name":            d.Get("name").(string),
		"content_version": d.Get("content_version").(int),
		"description":     d.Get("description").(string),
		"domain_name":     d.Get("domain_name").(string),
		"town":            d.Get("town").(string),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	// construct the HTTP request
	req, err := http.NewRequest("POST", config.Endpoint+"/u/"+config.UserUuid+"/homes", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return diag.FromErr(err)
	}

	// set Headers
	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// make client call
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	// parse response JSON
	var responseData map[string]interface {
	}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return diag.FromErr(err)
	}

	// StatusOK = 200
	if resp.StatusCode != http.StatusOK {
		return diag.FromErr(fmt.Errorf("failed to create home resource,status code: %d, status: %s, body: %s", resp.StatusCode, resp.Status, responseData))
	}

	homeUUID := responseData["uuid"].(string)
	d.SetId(homeUUID)

	log.Print("resourceHomeCreate:end")
	return diags
}

func resourceHomeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("resourceHomeRead:start")
	var diags diag.Diagnostics
	config := m.(*Config)

	// assign the homeUUID
	homeUUID := d.Id()

	// construct the HTTP request
	req, err := http.NewRequest("GET", config.Endpoint+"/u/"+config.UserUuid+"/homes/"+homeUUID, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	// set Headers
	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// make client call
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	// parse response JSON
	var responseData map[string]interface {
	}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return diag.FromErr(err)
	}
	if resp.StatusCode == http.StatusOK {
		d.Set("name", responseData["name"].(string))
		d.Set("content_version", responseData["content_version"].(float64))
		d.Set("description", responseData["description"].(string))
		d.Set("domain_name", responseData["domain_name"].(string))
	} else if resp.StatusCode == http.StatusNotFound {
		d.SetId("")
	} else if resp.StatusCode != http.StatusOK {
		return diag.FromErr(fmt.Errorf("failed to read home resource, status code: %d, status: %s, body: %s", resp.StatusCode, resp.Status, responseData))
	}

	log.Print("resourceHomeRead:end")
	return diags
}

func resourceHomeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("resourceHomeUpdate:start")
	var diags diag.Diagnostics
	config := m.(*Config)

	// construct payload
	payload := map[string]interface{}{
		"name":            d.Get("name").(string),
		"content_version": d.Get("content_version").(int),
		"description":     d.Get("description").(string),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	// assign the homeUUID
	homeUUID := d.Id()

	// construct the HTTP request
	req, err := http.NewRequest("PUT", config.Endpoint+"/u/"+config.UserUuid+"/homes/"+homeUUID, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return diag.FromErr(err)
	}
	// set Headers
	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// make client call
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	// StatusOK = 200
	if resp.StatusCode != http.StatusOK {
		return diag.FromErr(fmt.Errorf("failed to update home resource, status code: %d, status: %s", resp.StatusCode, resp.Status))
	}

	d.Set("name", payload["name"])
	d.Set("content_version", payload["content_version"])
	d.Set("description", payload["description"])

	log.Print("resourceHomeUpdate:end")
	return diags
}

func resourceHomeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("resourceHomeDelete:start")
	var diags diag.Diagnostics
	config := m.(*Config)

	// assign the homeUUID
	homeUUID := d.Id()

	// construct the HTTP request
	req, err := http.NewRequest("DELETE", config.Endpoint+"/u/"+config.UserUuid+"/homes/"+homeUUID, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	// set Headers
	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// make client call
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	// StatusOK = 200
	if resp.StatusCode != http.StatusOK {
		return diag.FromErr(fmt.Errorf("failed to delete home resource, status code: %d, status: %s", resp.StatusCode, resp.Status))
	}

	d.SetId("")
	log.Print("resourceHomeDelete:end")
	return diags
}

func providerConfig(p *schema.Provider) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		log.Print("providerConfig:start")
		config := Config{
			Endpoint: d.Get("endpoint").(string),
			Token:    d.Get("token").(string),
			UserUuid: d.Get("user_uuid").(string),
		}
		log.Print("providerConfig:end")
		return &config, nil
	}
}

func validateUUID(v interface{}, k string) (ws []string, errors []error) {
	log.Print("validateUUID:start")
	value := v.(string)

	// Attempt to parse the value as a UUID
	_, err := uuid.Parse(value)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q must be a valid UUID: %s", k, err))
	}
	log.Print("validateUUID:end")

	return
}
