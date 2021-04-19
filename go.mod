module github.com/ontariosystems/terraform-provider-cloudranger

go 1.15

require (
	github.com/hashicorp/terraform v0.14.10
	github.com/jaxi/terraform-provider-cloudranger v0.0.0-20181207180518-bf19d24b2492
	golang.org/x/net v0.0.0-20210415231046-e915ea6b2b7d
	golang.org/x/oauth2 v0.0.0-20210413134643-5e61552d6c78
)

replace github.com/jaxi/terraform-provider-cloudranger => github.com/ontariosystems/terraform-provider-cloudranger v1.0.2-0.20210419193140-188de7c8975d
