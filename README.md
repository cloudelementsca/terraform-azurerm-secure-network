# terraform-azurerm-secure-network
## Description
This module creates a NIST SP 800-53 and Canada Federal PBMM compliant Azure virtual network, and optional custom subnets, in an existing resource group (BYORG) and location that you specify. It's highly adjustable and takes the same input variables as [`azurerm_virtual_network`](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/virtual_network.html) and [`azurerm_subnet`](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/subnet) resource definitions have attributes. All attributes for the vnet and subnets are optional. The vnet output variables contain all the information for the vnet, subnets, and nsgs.

Meets MCSB control [NS-1: Establish Network Segmentation Boundaries](https://learn.microsoft.com/en-us/security/benchmark/azure/mcsb-network-security#ns-1-establish-network-segmentation-boundaries)

## Using the Module
See the `examples/` folders on how to use the module.

### Terraform Registry
```
module "network" {
  source = "cloudelementsca/terraform-azurerm-secure-network"

  resource_group_name = azurerm_resource_group.network_rg.name
  location            = azurerm_resource_group.network_rg.location
}
```

### GitHub
```
module "network" {
  source = "github.com/cloudelementsca/terraform-azurerm-secure-network"

  resource_group_name = azurerm_resource_group.network_rg.name
  location            = azurerm_resource_group.network_rg.location
}
```
