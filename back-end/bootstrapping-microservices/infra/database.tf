# Kubernetes yml 설정 파일과 비슷
resource "kubernetes_deployment" "database" {
  metadata {
    name   = "database"
    labels = {
      pod = "database"
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        pod = "database"
      }
    }

    template {
      metadata {
        labels = {
          pod = "database"
        }
      }

      # container 정의
      spec {
        container {
          image = "mongo:4.2.8"
          name  = "database"

          port {
            container_port = 27017
          }
        }
      }
    }
  }
}

# service: 클러스터 내 다른 컨테이너가 접근 가능하도록 DNS 레코드를 생성하는 역할
resource "kubernetes_service" "database" {
  metadata {
    name = "database"
  }

  spec {
    selector = {
      pod = kubernetes_deployment.database.metadata[0].labels.pod
    }

    port {
      port = 27017
    }

    # azure 로드밸런서를 이용하여 외부에 서비스 노출
    # i.e public IP 할당
    # => 배포 테스트용(운영에서는 노출X)
    type = "LoadBalancer"
  }
}