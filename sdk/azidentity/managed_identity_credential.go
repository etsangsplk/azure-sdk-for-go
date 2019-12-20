// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// ManagedIdentityCredentialOptions contains parameters that can be used to configure a Managed Identity Credential
type ManagedIdentityCredentialOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	// Leave this as nil to use the default HTTP transport.
	HTTPClient azcore.Transport

	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions

	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
}

func (m *ManagedIdentityCredentialOptions) setDefaultValues() *ManagedIdentityCredentialOptions {
	if m == nil {
		m = defaultMSIOpts
	}
	return m
}

// ManagedIdentityCredential attempts authentication using a managed identity that has been assigned to the deployment environment. This authentication type works in Azure VMs,
// App Service and Azure Functions applications, as well as inside of Azure Cloud Shell. More information about configuring managed identities can be found here:
// https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview
type ManagedIdentityCredential struct {
	clientID string
	client   *managedIdentityClient
}

// NewManagedIdentityCredential creates an instance of the ManagedIdentityCredential capable of authenticating a resource with a managed identity.
// clientID: The client id to authenticate for a user assigned managed identity.  More information on user assigned managed identities cam be found here:
// https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview#how-a-user-assigned-managed-identity-works-with-an-azure-vm
// options: Options that allow to configure the management of the requests sent to the Azure Active Directory service.
func NewManagedIdentityCredential(clientID string, options *ManagedIdentityCredentialOptions) *ManagedIdentityCredential {
	return &ManagedIdentityCredential{clientID: clientID, client: newManagedIdentityClient(options)}
}

// GetToken obtains an AccessToken from the Managed Identity service if available.
// scopes: The list of scopes for which the token will have access.
// Returns an AccessToken which can be used to authenticate service client calls, or a default AccessToken if no managed identity is available.
func (c *ManagedIdentityCredential) GetToken(ctx context.Context, opts azcore.TokenRequestOptions) (*azcore.AccessToken, error) {
	return c.client.authenticate(ctx, c.clientID, opts.Scopes)
}

// AuthenticationPolicy implements the azcore.Credential interface on ManagedIdentityCredential.
func (c *ManagedIdentityCredential) AuthenticationPolicy(options azcore.AuthenticationPolicyOptions) azcore.Policy {
	options.Options.Scopes = []string{scopesToResource(options.Options.Scopes[0])}
	return newBearerTokenPolicy(c, options)
}