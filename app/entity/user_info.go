package entity

type UserInfo struct {
	Name      string `json:"name"`
	Avatar    string `json:"avatar,omitempty"`
	Role      string `json:"role"`
	Access    Access `json:"access"`
	LoginInfo `json:"loginInfo,omitempty"`
}

type Access struct {
	Query      bool `json:"query"`
	New        bool `json:"new"`
	Edit       bool `json:"edit"`
	ViewDetail bool `json:"viewDetail"`
	View       bool `json:"view"`
	Delete     bool `json:"delete"`
	BulkDelete bool `json:"bulkDelete"`
}
type LoginInfo struct {
	Token    string    `json:"token,omitempty"`
	Projects []Project `json:"projects,omitempty"`
}
type Project struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}
