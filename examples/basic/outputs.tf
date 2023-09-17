output "vnet" {
  description = "Vnet info."
  value       = module.secure_network.vnet
}

output "subnets" {
  description = "All subnets info."
  value       = module.secure_network.subnets
}

output "nsgs" {
  description = "All subnets info."
  value       = module.secure_network.nsgs
}