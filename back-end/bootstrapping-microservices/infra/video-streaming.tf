locals {
  service_name = "sample-video-streaming"
  login_server = azurerm_container_registry.container_registry.login_server
  username     = azurerm_container_registry.container_registry.admin_username
  password     = azurerm_container_registry.container_registry.admin_password

  # 빌드된 도커 이미지를 레지스트리에 게시하기 위한 이미지 태그
  image_tag = "${local.login_server}/${local.service_name}:${var.app_version}"

  docker_credential = {
    auths = {
      (local.login_server) = {
        auth = base64encode("${local.username}:${local.password}")
      }
    }
  }
}

# null_resource: 특정 리소스 타입이 없는 테라폼 리소스
# => 도커 명령을 호출하여 이미지 빌드 및 게시를 위해 사용
resource "null_resource" "docker_build" {
  triggers = {
    always_run = timestamp()
  }

  # build docker image
  provisioner "local-exec" {
    # M1 환경에서 이미지 빌드 시, 노드의 운영체제에 맞게 빌드를 해줘야 함
    command = "docker buildx build --platform linux/amd64 -t ${local.image_tag} --file ../${local.service_name}/Dockerfile-prod ../${local.service_name}"
  }
}

resource "null_resource" "docker_login" {
  # 도커 이미지 빌드 후 실행
  depends_on = [null_resource.docker_build]

  triggers = {
    always_run = timestamp()
  }

  provisioner "local-exec" {
    command = "docker login ${local.login_server} --username ${local.username} --password ${local.password}"
  }
}

resource "null_resource" "docker_push" {
  # 도커 빌드 -> 로그인 이후 실행
  depends_on = [null_resource.docker_login]

  triggers = {
    always_run = timestamp()
  }

  provisioner "local-exec" {
    command = "docker push ${local.image_tag}"
  }
}


# secret: 쿠버네티스에 민감한 정보 저장
# => 컨테이너 레지스트리에서 이미지를 가져오기 위한 권한 정보
resource "kubernetes_secret" "docker_credentials" {
  metadata {
    name = "docker-credentials"
  }

  data = {
    ".dockerconfigjson" = jsonencode(local.docker_credential)
  }

  type = "kubernetes.io/dockerconfigjson"
}

resource "kubernetes_deployment" "service_deployment" {
  # 이미지 게시 후 배포 가능
  depends_on = [null_resource.docker_push]

  metadata {
    name   = local.service_name
    labels = {
      pod = local.service_name
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        pod = local.service_name
      }
    }

    template {
      metadata {
        labels = {
          pod = local.service_name
        }
      }

      spec {
        container {
          image = local.image_tag
          name  = local.service_name

          env {
            name  = "PORT"
            value = "80"
          }
        }

        image_pull_secrets {
          name = kubernetes_secret.docker_credentials.metadata[0].name
        }
      }
    }
  }
}

resource "kubernetes_service" "service" {
  metadata {
    name = local.service_name
  }

  spec {
    selector = {
      pod = kubernetes_deployment.service_deployment.metadata[0].labels.pod
    }

    session_affinity = "ClientIP"

    port {
      port=80
      target_port = "80"
    }

    type = "LoadBalancer"
  }
}