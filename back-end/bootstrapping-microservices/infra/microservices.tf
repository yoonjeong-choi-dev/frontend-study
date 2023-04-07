locals {
  login_server = azurerm_container_registry.container_registry.login_server
  username     = azurerm_container_registry.container_registry.admin_username
  password     = azurerm_container_registry.container_registry.admin_password
  rabbit_host  = "amqp://guest:guest@rabbit:5672"
  db_host      = "mongodb://db:27017"
}

module "gateway-service" {
  source           = "./modules/microservice"
  service_name     = "gateway"
  service_type     = "LoadBalancer"
  session_affinity = "ClientIP"
  env              = {
    METADATA_HOST  = "metadata"
    STREAMING_HOST = "video-streaming"
    HISTORY_HOST   = "history"
    UPLOAD_HOST    = "video-upload"
  }
  app_version  = var.app_version
  login_server = local.login_server
  username     = local.username
  password     = local.password
}

module "video-streaming-service" {
  source       = "./modules/microservice"
  service_name = "video-streaming"
  env          = {
    VIDEO_STORAGE_HOST = "video-storage"
    VIDEO_STORAGE_PORT = 80
    RABBIT_HOST        = local.rabbit_host
  }
  app_version  = var.app_version
  login_server = local.login_server
  username     = local.username
  password     = local.password
}

module "video-upload-service" {
  source       = "./modules/microservice"
  service_name = "video-upload"
  env          = {
    METADATA_HOST      = "metadata"
    VIDEO_STORAGE_HOST = "video-storage"
    RABBIT_HOST        = local.rabbit_host
  }
  app_version  = var.app_version
  login_server = local.login_server
  username     = local.username
  password     = local.password
}

module "azure-storage-service" {
  source       = "./modules/microservice"
  service_name = "azure-storage"
  dns_name     = "video-storage"
  env          = {
    AZURE_STORAGE_ACCOUNT_NAME   = var.storage_account_name
    AZURE_STORAGE_ACCESS_KEY     = var.storage_access_key
    AZURE_STORAGE_CONTAINER_NAME = "videos"
  }
  app_version  = var.app_version
  login_server = local.login_server
  username     = local.username
  password     = local.password
}

module "history-service" {
  source       = "./modules/microservice"
  service_name = "history"
  env          = {
    DB_HOST     = local.db_host
    DB_NAME     = "history"
    RABBIT_HOST = local.rabbit_host
  }
  app_version  = var.app_version
  login_server = local.login_server
  username     = local.username
  password     = local.password
}

module "metadata-service" {
  source       = "./modules/microservice"
  service_name = "metadata"
  env          = {
    DB_HOST     = local.db_host
    DB_NAME     = "metadata"
    RABBIT_HOST = local.rabbit_host
  }
  app_version  = var.app_version
  login_server = local.login_server
  username     = local.username
  password     = local.password
}