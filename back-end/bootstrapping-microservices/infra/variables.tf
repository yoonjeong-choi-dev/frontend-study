variable "app_name" {
  default = "yjtube"
}

variable "location" {
  #default = "Korea South"
  default = "West US"
}

variable admin_username {
  default = "linux_admin"
}

// 배포 시, 수동으로 설정 필요
variable app_version {}

// azure 인증에 필요한 변수들
// : 민감한 정보이기 때문에 코드에 설정하지 않음
// => terraform apply 실행 시 입력
variable client_id {
}

variable client_secret {
}