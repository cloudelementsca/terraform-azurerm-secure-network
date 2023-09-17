## ---------------------------------------------------------------------------------------------------------------------
## ONE LINE SUMMARY DESCRIBING WHAT IS BEING MANAGED IN THIS SECTION IN ALL CAPS
## The rest of the comments should be in standard casing. This section should contain an overall description of the
## component that is being managed, and highlight any unconventional workarounds or configurations that are in place.
## ---------------------------------------------------------------------------------------------------------------------


output "vnet" {
  description = "All vnet info."
  value       = azurerm_virtual_network.vnet
}

output "subnets" {
  description = "All subnets info."
  value       = azurerm_subnet.subnets
}

output "nsgs" {
  description = "All nsgs info."
  value       = azurerm_network_security_group.nsgs
}