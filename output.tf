output "cp_node_ip_addresses" {
  value = [for node in resource.harvester_virtualmachine.cp_node : "${node.name}: ${join(", ", node.network_interface[*].ip_address)}"]
}