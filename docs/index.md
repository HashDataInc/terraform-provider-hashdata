# HashData Provider



## Example Usage

```terraform
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


```

