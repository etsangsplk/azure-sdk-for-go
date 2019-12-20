module github.com/Azure/azure-sdk-for-go/sdk/storage/blob/2018-11-09/azblob

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v0.0.0-00010101000000-000000000000
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v0.0.0-00010101000000-000000000000
)

replace (
	github.com/Azure/azure-sdk-for-go/sdk/azcore => ../../../../azcore
	github.com/Azure/azure-sdk-for-go/sdk/azidentity => ../../../../azidentity
	github.com/Azure/azure-sdk-for-go/sdk/internal => ../../../../internal
)

go 1.13