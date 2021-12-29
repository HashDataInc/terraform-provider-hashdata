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

resource "hashdata_computing" "example_computing" {
  name = "somebody_test_computing"
  catalog = "29469d36-4331-4988-beaa-2ea729ef62ec"
  image = "img-jgph2cpr"
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
}

