resource "azurerm_resource_group" "yjtube" {
  name = var.app_name
  location = var.location

  # Legacy before applying variable
#  location = "Korea South"
#  name     = "yj-tube"
}