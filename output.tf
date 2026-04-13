output "vm_ip_addresses" {
  value      = [for vm in data.harvester_virtualmachine.vm : "${vm.name}: ${join(", ", vm.network_interface[*].ip_address)}"]
  depends_on = [harvester_virtualmachine.vm]
}