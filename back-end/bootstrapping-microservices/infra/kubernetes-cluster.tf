resource "azurerm_kubernetes_cluster" "cluster" {
  location            = var.location
  name                = var.app_name
  resource_group_name = azurerm_resource_group.yjtube.name
  dns_prefix = var.app_name
  kubernetes_version = "1.25.4"

  linux_profile {
    admin_username = var.admin_username
    ssh_key {
      key_data = "${trimspace(tls_private_key.key.public_key_openssh)} ${var.admin_username}@azure.com"
    }
  }

  default_node_pool {
    name    = "default"
    node_count = 1
    vm_size = "Standard_B2ms"
  }

  # azure 인증 방식
  service_principal {
    client_id     = var.client_id
    client_secret = var.client_secret
  }
}

output "cluster_client_key" {
  value = azurerm_kubernetes_cluster.cluster.kube_config[0].client_key
  sensitive = true
}

output "cluster_client_certificate" {
  value = azurerm_kubernetes_cluster.cluster.kube_config[0].client_certificate
  sensitive = true
}

output "cluster_cluster_ca_certificate" {
  value = azurerm_kubernetes_cluster.cluster.kube_config[0].cluster_ca_certificate
  sensitive = true
}

output "cluster_cluster_username" {
  value = azurerm_kubernetes_cluster.cluster.kube_config[0].username
  sensitive = true
}

output "cluster_cluster_password" {
  value = azurerm_kubernetes_cluster.cluster.kube_config[0].password
  sensitive = true
}

output "cluster_kube_config" {
  value = azurerm_kubernetes_cluster.cluster.kube_config_raw
  sensitive = true
}

output "cluster_host" {
  value = azurerm_kubernetes_cluster.cluster.kube_config[0].host
  sensitive = true
}
