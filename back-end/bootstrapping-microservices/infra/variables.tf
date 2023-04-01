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

// azure 인증에 필요한 변수들
// : 민감한 정보이기 때문에 코드에 설정하지 않음
// => terraform apply 실행 시 입력
variable client_id {
}

variable client_secret {
}