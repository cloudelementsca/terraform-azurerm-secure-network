output "vnet" {
  description = "Vnet info."
  value       = module.network.vnet
}

output "subnets" {
  description = "All subnets info."
  value       = module.network.subnets
}

output "nsgs" {
  description = "All subnets info."
  value       = module.network.nsgs
}