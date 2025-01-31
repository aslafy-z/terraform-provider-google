// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccComputeInterconnectAttachment_interconnectAttachmentBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeInterconnectAttachmentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInterconnectAttachment_interconnectAttachmentBasicExample(context),
			},
			{
				ResourceName:            "google_compute_interconnect_attachment.on_prem",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"router", "candidate_subnets", "region"},
			},
		},
	})
}

func testAccComputeInterconnectAttachment_interconnectAttachmentBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_interconnect_attachment" "on_prem" {
  name                     = "tf-test-on-prem-attachment%{random_suffix}"
  edge_availability_domain = "AVAILABILITY_DOMAIN_1"
  type                     = "PARTNER"
  router                   = google_compute_router.foobar.id
  mtu                      = 1500
}

resource "google_compute_router" "foobar" {
  name    = "router%{random_suffix}"
  network = google_compute_network.foobar.name
  bgp {
    asn = 16550
  }
}

resource "google_compute_network" "foobar" {
  name                    = "network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func testAccCheckComputeInterconnectAttachmentDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_interconnect_attachment" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeInterconnectAttachment still exists at %s", url)
			}
		}

		return nil
	}
}
