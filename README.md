# Test provider locally

```
env GOOS=darwin GOARCH=amd64 go build -o terraform-provider-example  

mkdir ~/.terraform.d/plugins/terraform-example.com/exampleprovider/example/1.0.0/darwin_arm64/

cp terraform-provider-example ~/.terraform.d/plugins/terraform-example.com/exampleprovider/example/1.0.0/darwin_arm64/

terraform init

terraform plan

terraform apply
```

# Publishing the provider

## Create .goreleaser.yml file

Create the `.goreleaser.yml` file from https://github.com/hashicorp/terraform-provider-scaffolding-framework/blob/main/.goreleaser.yml repository


## Create `terraform-registry-manifest.json` file
Create the `terraform-registry-manifest.json` file with following content
```
{
    "version": 1,
    "metadata": {
      "protocol_versions": ["5.0"]
    }
}
```

## Setting up the gpg key

```
gpg --full-generate-key
```
```
gpg --list-secret-keys --keyid-format=long
gpg --list-keys
```

## Set environment variable for GPG_FINGERPRINT
```
export GPG_FINGERPRINT=<fingerprint from previus command>
```


## Export public key
```
gpg --armor --export <fingerprint from previus command>
```

Now add this key to hashicorp so that your provider checksum can be verified by hashicorp.

## Github PAT token
Create the github PAT token with the right access and export it as en environment variable

```
export GITHUB_TOKEN=ghp_............ 
```

## Make a release
```
git tag v0.0.5
```

```
git push --tags
```

```
goreleaser release --clean
```


# References

* https://www.infracloud.io/blogs/developing-terraform-custom-provider/
* https://www.hashicorp.com/blog/writing-custom-terraform-providers
* https://thekevinwang.com/2023/10/05/build-publish-terraform-provider