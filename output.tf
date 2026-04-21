output "vm_ip_addresses" {
  value      = [for vm in data.harvester_virtualmachine.vm : "${join(", ", vm.network_interface[*].ip_address)}"]
  depends_on = [harvester_virtualmachine.vm]
}

output "vm_lb_ip_address" {
  value      = var.create_lb ? [for lb in data.harvester_loadbalancer.vm_lb : "${lb.ip_address}"] : null
}