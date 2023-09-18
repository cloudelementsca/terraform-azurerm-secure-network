package network_tests

import (
	"encoding/json"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

type VnetStruct struct {
	Address_space           []string
	Bgp_community           string
	Ddos_protection_plan    []string
	Dns_servers             []string
	Edge_zone               string
	Flow_timeout_in_minutes int
	Location                string
	Resource_group_name     string
	Tags                    map[string]interface{}
}

type ServiceDelegation struct {
	Name    string
	Actions []string
}

type Delegation struct {
	Name               string
	Service_delegation []ServiceDelegation
}

type Subnet struct {
	Address_prefixes                              []string
	Delegation                                    []Delegation
	Private_endpoint_network_policies_enabled     bool
	Private_link_service_network_policies_enabled bool
}

func TestMultiSubnetNetworkModule(t *testing.T) {

	expectedVnetOutput := VnetStruct{
		Address_space:           []string{"10.19.0.0/16"},
		Bgp_community:           "",
		Ddos_protection_plan:    []string{},
		Dns_servers:             []string{},
		Edge_zone:               "",
		Flow_timeout_in_minutes: 0,
		Location:                "canadacentral",
		Resource_group_name:     "vnet-rg",
		Tags:                    map[string]interface{}{"environment": "dev"},
	}

	expectedPeSubnetOutupt := Subnet{
		Address_prefixes: []string{"10.19.1.0/24"},
		Delegation:       []Delegation{},
		Private_endpoint_network_policies_enabled:     false,
		Private_link_service_network_policies_enabled: false,
	}

	expectedAciSubnetOutupt := Subnet{
		Address_prefixes: []string{"10.19.2.0/24"},
		Delegation: []Delegation{
			{
				Name: "aci_delegation",
				Service_delegation: []ServiceDelegation{
					{
						Name:    "Microsoft.ContainerInstance/containerGroups",
						Actions: []string{"Microsoft.Network/virtualNetworks/subnets/action"},
					},
				},
			},
		},
		Private_endpoint_network_policies_enabled:     true,
		Private_link_service_network_policies_enabled: true,
	}

	expectedFeSubnetOutupt := Subnet{
		Address_prefixes: []string{"10.19.3.0/24"},
		Delegation:       []Delegation{},
		Private_endpoint_network_policies_enabled:     true,
		Private_link_service_network_policies_enabled: true,
	}

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/multi_subnet",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualVnetObject := VnetStruct{}
	strVnet := terraform.OutputJson(t, terraformOptions, "vnet")
	json.Unmarshal([]byte(strVnet), &actualVnetObject)
	assert.Equal(t, expectedVnetOutput, actualVnetObject, &actualVnetObject)

	strSubnets := terraform.OutputJson(t, terraformOptions, "subnets")

	subnets := map[string]Subnet{}
	json.Unmarshal([]byte(strSubnets), &subnets)
	assert.Equal(t, expectedPeSubnetOutupt, subnets["pe-subnet"], &subnets)
	assert.Equal(t, expectedFeSubnetOutupt, subnets["fe-subnet"], &subnets)
	assert.Equal(t, expectedAciSubnetOutupt, subnets["aci-subnet"], &subnets)
}
