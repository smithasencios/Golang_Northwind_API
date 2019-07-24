package auth

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}
type PermissionList struct {
	Data []*Permission `json:"data"`
}
type Permission struct {
	Permission_Name      string `json:"permission_name"`
	Description          string `json:"description"`
	Resource_Server_Name string `json:"resource_server_name"`
}
type AccessTokenItem struct {
	Access_Token string `json:"access_token"`
}
