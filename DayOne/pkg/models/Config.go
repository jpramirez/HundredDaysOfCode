package models

//Config Main Configuration File Structure
type Config struct {
	GRPCPort     string   `json:"grpcport"`
	GRPCAddr     string   `json:"grpcaddr"`
	HTTPAddr     string   `json:"httpaddr"`
	HTTPPort     string   `json:"httpport"`
	WebPort      string   `json:"webport"`
	WebAddress   string   `json:"webaddress"`
	IsMaster     string   `json:"ismaster"`
	DatabaseName string   `json:"databasename"`
	JikoModules  []string `json:"jikomodules"`
	AgentRoles   []string `json:"AgentRoles"`
	SystemType   string   `json:"systemtype"`
	APIURL       string   `json:"apiurl"`
	BinURL       string   `json:"binurl"`
	DiffURL      string   `json:"diffurl"`
	UpdateDir    string   `json:"updatedir"`
	AppName      string   `json:"appname"`
	LogFile      string   `json:"logfile"`
	FileSystem   []string `json:"filesystem"`
	Extensions   []string `json:"extensions"`
	CertFolder   string   `json:"certfolder"`
	CaFile       string   `json:"cafile"`
}
