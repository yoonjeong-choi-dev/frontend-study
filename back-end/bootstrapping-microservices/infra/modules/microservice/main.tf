# variables for each microservice
variable "app_version" {}

variable "service_name" {}

variable "dns_name" {
  default = ""
}

variable "login_server" {}
variable "username" {}
variable "password" {}

variable "service_type" {
  default = "ClusterIP"
}

variable "session_affinity" {
  default = "None"
}

variable "env" {
  default = {}
  type    = map(string)
}

# similar definition in video-streaming.tf
locals {
  image_tag = "${var.login_server}/${var.service_name}:${var.app_version}"

  docker_credential = {
    auths = {
      (var.login_server) = {
        auth = base64encode("${var.username}:${var.password}")
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
    command = "docker buildx build --platform linux/amd64 -t ${local.image_tag} --file ../${var.service_name}/Dockerfile-prod ../${var.service_name}"
  }
}

resource "null_resource" "docker_login" {
  # 도커 이미지 빌드 후 실행
  depends_on = [null_resource.docker_build]

  triggers = {
    always_run = timestamp()
  }

  provisioner "local-exec" {
    command = "docker login ${var.login_server} --username ${var.username} --password ${var.password}"
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
    name = "${var.service_name}-docker-credentials"
  }

  data = {
    ".dockerconfigjson" = jsonencode(local.docker_credential)
  }

  type = "kubernetes.io/dockerconfigjson"
}

# deploy script for each microservice
resource "kubernetes_deployment" "service_deployment" {
  # 이미지 게시 후 배포 가능
  depends_on = [null_resource.docker_push]

  metadata {
    name   = var.service_name
    labels = {
      pod = var.service_name
    }
  }


  spec {
    replicas = 1

    selector {
      match_labels = {
        pod = var.service_name
      }
    }

    template {
      metadata {
        labels = {
          pod = var.service_name
        }
      }

      spec {
        container {
          image = local.image_tag
          name  = var.service_name

          env {
            name  = "PORT"
            value = "80"
          }

          # Set environment in docker-compose.yml
          dynamic "env" {
            for_each = var.env
            content {
              name  = env.key
              value = env.value
            }
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
    name = var.dns_name != "" ? var.dns_name : var.service_name
  }

  spec {
    selector = {
      pod = kubernetes_deployment.service_deployment.metadata[0].labels.pod
    }

    session_affinity = var.session_affinity

    port {
      port        = 80
      target_port = 80
    }

    type = var.service_type
  }
}