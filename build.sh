rm go.mod
rm go.sum
rm terraform-provider-dmprovider
go mod init dorpm.sbs/m/1.0.0
go fmt
go mod tidy
go build -o terraform-provider-dmprovider
rm /Users/olishchuk/.terraform.d/plugins/dorpm.sbs/terraform-provider-dmprovider/dmprovider/1.0.0/darwin_arm64/terraform-provider-dmprovider
cp terraform-provider-dmprovider /Users/olishchuk/.terraform.d/plugins/dorpm.sbs/terraform-provider-dmprovider/dmprovider/1.0.0/darwin_arm64/terraform-provider-dmprovider