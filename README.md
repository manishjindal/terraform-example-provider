# https://www.infracloud.io/blogs/developing-terraform-custom-provider/


# https://www.hashicorp.com/blog/writing-custom-terraform-providers


```
env GOOS=darwin GOARCH=amd64 go build -o terraform-provider-example  

mkdir ~/.terraform.d/plugins/terraform-example.com/exampleprovider/example/1.0.0/darwin_arm64/

cp terraform-provider-example ~/.terraform.d/plugins/terraform-example.com/exampleprovider/example/1.0.0/darwin_arm64/

terraform init

terraform plan

terraform apply
```