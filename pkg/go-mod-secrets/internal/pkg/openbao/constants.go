/*******************************************************************************
 *******************************************************************************/

package openbao

const (
	// NamespaceHeader specifies the header name to use when including Namespace information in a request.
	NamespaceHeader = "X-Vault-Namespace"
	AuthTypeHeader  = "X-Vault-Token"

	HealthAPI              = "/v1/sys/health"
	InitAPI                = "/v1/sys/init"
	UnsealAPI              = "/v1/sys/unseal"
	CreatePolicyPath       = "/v1/sys/policies/acl/%s"
	CreateTokenAPI         = "/v1/auth/token/create"    // nolint: gosec
	CreateTokenByRolePath  = "/v1/auth/token/create/%s" // nolint: gosec
	ListAccessorsAPI       = "/v1/auth/token/accessors" // nolint: gosec
	RevokeAccessorAPI      = "/v1/auth/token/revoke-accessor"
	LookupAccessorAPI      = "/v1/auth/token/lookup-accessor"
	LookupSelfAPI          = "/v1/auth/token/lookup-self"
	RevokeSelfAPI          = "/v1/auth/token/revoke-self"
	TokenRolesByRoleAPI    = "/v1/auth/token/roles/%s"       // nolint: gosec
	RootTokenControlAPI    = "/v1/sys/generate-root/attempt" // nolint: gosec
	RootTokenRetrievalAPI  = "/v1/sys/generate-root/update"  // nolint: gosec
	MountsAPI              = "/v1/sys/mounts"
	namedEntityAPI         = "/v1/identity/entity/name"
	idEntityAPI            = "/v1/identity/entity/id"
	entityAliasAPI         = "/v1/identity/entity-alias"
	oidcKeyAPI             = "/v1/identity/oidc/key"
	oidcRoleAPI            = "/v1/identity/oidc/role"
	oidcGetTokenAPI        = "/v1/identity/oidc/token"      // nolint: gosec
	oidcTokenIntrospectAPI = "/v1/identity/oidc/introspect" // nolint: gosec
	authAPI                = "/v1/sys/auth"
	authMountBase          = "/v1/auth"

	lookupSelfTokenAPI = "/v1/auth/token/lookup-self" // nolint: gosec
	renewSelfTokenAPI  = "/v1/auth/token/renew-self"  // nolint: gosec
)
