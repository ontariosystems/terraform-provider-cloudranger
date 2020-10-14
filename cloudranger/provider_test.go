package cloudranger

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"cloudranger": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("CLOUDRANGER_ACCOUNT_ID"); v == "" {
		t.Fatal("CLOUDRANGER_ACCOUNT_ID must be set for acceptance tests")
	}
	if v := os.Getenv("CLOUDRANGER_ORGANIZATION_ID"); v == "" {
		t.Fatal("CLOUDRANGER_ORGANIZATION_ID must be set for acceptance tests")
	}
	if v := os.Getenv("CLOUDRANGER_API_KEY"); v == "" {
		t.Fatal("CLOUDRANGER_API_KEY must be set for acceptance tests")
	}
}
