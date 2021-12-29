terraform {
  required_providers {
    hashdata = {
      source = "HashDataInc/hashdata"
      version = "0.0.1"
    }
  }
  required_version = ">= 0.12.0"
}

provider "hashdata" {
  username      = "your username"
  password      = "your password"
  client_id     = "your client id"
  client_secret = "your client secret"
  endpoint      = "https://console.hashdata.xyz"
}

resource "hashdata_warehouse" "example_warehouse" {
  name = "somebody_test_warehouse"
  image = "img-xlwrvrsq"
  master  {
    instance_type = "qingcloud-c2m4-basic"
    volume_type = "qingcloud-basic"
    volume_size = 20
    zone = "qingcloud-pek3b"
  }
  segment  {
    count = 2
    instance_type = "qingcloud-c2m4-basic"
    volume_type = "qingcloud-basic"
    volume_size = 20
    zone = "qingcloud-pek3b"
  }
  extra  {
    vpc = "vpc-4nwv90w2"
    subnet = "vxnet-5tjdylj"
    keypair = "kp-o07unn26"
  }
  metadata  {
    default_database = "warehouse"
    default_password = "warehouse"
    default_user = "warehouse"
    logic_part = 16
  }
  feature  {
    local_storage =  false
    mirror_standby = false
  }
}

