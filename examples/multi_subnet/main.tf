## ---------------------------------------------------------------------------------------------------------------------
## ALL MODULE RESOURCES
## Define all module resources in this file.
## ---------------------------------------------------------------------------------------------------------------------

resource "azurerm_resource_group" "network_rg" {
  name     = var.resource_group_name
  location = var.location
  tags     = var.tags
}

module "secure_network" {
  source              = "../../"
  resource_group_name = azurerm_resource_group.network_rg.name
  location            = azurerm_resource_group.network_rg.location
  network = {
    address_space = ["10.19.0.0/23"]
    subnets = {
      firewall-subnet = {
        address_prefixes  = ["10.19.1.0/24"]
        name              = "AzureFirewallSubnet"
        disable_nsg       = true
      }
      be-subnet = { address_prefixes = ["10.19.2.0/24"] }
      aci-subnet = {
        address_prefixes = ["10.19.3.0/23"]
        service_delegations = {
          aci_delegation = {
            name    = "Microsoft.ContainerInstance/containerGroups"
            actions = ["Microsoft.Network/virtualNetworks/subnets/action"]
          }
        }
      }      
    }
  }
  tags = var.tags
}