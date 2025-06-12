/*******************************************************************************
 * @author: Tingyu Zeng, Dell / Alain Pulluelo, ForgeRock AS
 *******************************************************************************/

package types

// InitResponse contains a Secret Store init response
type InitResponse struct {
	Keys          []string `json:"keys,omitempty"`
	KeysBase64    []string `json:"keys_base64,omitempty"`
	EncryptedKeys []string `json:"encrypted_keys,omitempty"`
	Nonces        []string `json:"nonces,omitempty"`
	RootToken     string   `json:"root_token,omitempty"`
}

// TokenMetadata has introspection data about a token and is the "data" sub-structure for token lookup,
// i.e. TokenLookupResponse, and token self-lookup
type TokenMetadata struct {
	Accessor   string   `json:"accessor"`
	ExpireTime string   `json:"expire_time"`
	Path       string   `json:"path"`
	Policies   []string `json:"policies"`
	Period     int      `json:"period"` // in seconds
	Renewable  bool     `json:"renewable"`
	Ttl        int      `json:"ttl"` // in seconds
}

// BootStrapACLTokenInfo is the key portion of the response metadata from consulACLBootstrapAPI
type BootStrapACLTokenInfo struct {
	SecretID string   `json:"SecretID"`
	Policies []Policy `json:"Policies"`
}

// Alias has introspection data about entity alias
type Alias struct {
	Name string `json:"name"`
}

// EntityMetadata has introspection data about entity
type EntityMetadata struct {
	Aliases  []Alias  `json:"aliases"`
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Policies []string `json:"policies"`
}
