package v1

type Product struct {
	Product     string `json:"product"`
	GitLoc      string `json:"gitLoc"`
	ClusterName string `json:"clusterName"`
	Cloud       string `json:"cloud"`
	Account     string `json:"account"`
	Env         string `json:"env"`
	Region      string `json:"region"`
}
