resource "azurerm_container_registry" "container_registry" {
  name                = var.app_name
  resource_group_name = azurerm_resource_group.yjtube.name
  location            = var.location
  admin_enabled       = "true"
  sku                 = "Basic"
}

output "registry_hostname" {
  value = azurerm_container_registry.container_registry.login_server
}

output "registry_username" {
  value = azurerm_container_registry.container_registry.admin_username
}

output "registry_password" {
  value     = azurerm_container_registry.container_registry.admin_password
  sensitive = true
}