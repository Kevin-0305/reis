package config

type Alert struct {
	Elastic        Elastic        `yaml:"elastic"`
	FsApp          FsApp          `yaml:"fsapp"`
	Fs             Fs             `yaml:"fs"`
	TencentMessage TencentMessage `yaml:"tencentMessage"`
	TencentPhone   TencentPhone   `yaml:"tencentPhone"`
	HuaWeiMessage  HuaWeiMessage  `yaml:"huweiMessage"`
	AliyunMessage  AliyunMessage  `yaml:"aliyunMessage"`
	AliyunPhone    AliyunPhone    `yaml:"aliyunPhone"`
	AlertEmail     AlertEmail     `yaml:"alertEmail"`
	WorkWeChat     WorkWeChat     `yaml:"workWeChat"`
	DingDing       DingDing       `yaml:"dingDing"`
	WeChat         WeChat         `yaml:"weChat"`
}

type Elastic struct {
	Address   string `yaml:"address"`
	Account   string `yaml:"account"`
	Password  string `yaml:"password"`
	Version   int    `yaml:"version"`
	IndexName string `yaml:"indexName"`
}

type FsApp struct {
	Open              string `mapstructure:"open" json:"open" yaml:"open"`
	AppID             string `mapstructure:"app-id" json:"app-id"" yaml:"app-id""`
	AppSecret         string `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"`
	VerificationToken string `mapstructure:"verification-token" json:"verification-token" yaml:"verification-token"`
	AcceptIds         string `mapstructure:"accept-ids" json:"accept-ids" yaml:"accept-ids"`
	EncryptKey        string `mapstructure:"encrypt-key" json:"encrypt-key" yaml:"encrypt-key"`
}

type Fs struct {
	Open       string `mapstructure:"open" json:"open" yaml:"open"`
	WebHookUrl string `mapstructure:"web-hook-url" json:"web-hook-url" yaml:"web-hook-url"`
}

type WeChat struct {
	Open       string `mapstructure:"open" json:"open" yaml:"open"`
	WebHookUrl string `mapstructure:"web-hook-url" json:"web-hook-url" yaml:"web-hook-url"`
}

type DingDing struct {
	Open       string `mapstructure:"open" json:"open" yaml:"open"`
	WebHookUrl string `mapstructure:"web-hook-url" json:"web-hook-url" yaml:"web-hook-url"`
}

type WorkWeChat struct {
	Open    string `mapstructure:"open" json:"open" yaml:"open"`
	CropID  string `mapstructure:"crop-id" json:"crop-id" yaml:"crop-id"`
	Secret  string `mapstructure:"secret" json:"secret" yaml:"secret"`
	AgentID int64  `mapstructure:"agent-id" json:"agent-id" yaml:"agent-id"`
	ToUser  string `mapstructure:"to-user" json:"to-user" yaml:"to-user"`
	ToParty string `mapstructure:"to-party" json:"to-party" yaml:"to-party"`
	ToTag   string `mapstructure:"to-tag" json:"to-tag" yaml:"to-tag"`
}

type TencentMessage struct {
	Open       string `mapstructure:"open" json:"open" yaml:"open"`
	AppKey     string `mapstructure:"app-key" json:"app-key" yaml:"app-key"`
	TemplateID int    `mapstructure:"template-id" json:"template-id" yaml:"template-id"`
	Signature  string `mapstructure:"signature" json:"signature" yaml:"signature"`
	SdkAppID   string `mapstructure:"sdk-app-id" json:"sdk-app-id" yaml:"sdk-app-id"`
}

type TencentPhone struct {
	Open       string `mapstructure:"open" json:"open" yaml:"open"`
	AppKey     string `mapstructure:"app-key" json:"app-key" yaml:"app-key"`
	TemplateID string `mapstructure:"template-id" json:"template-id" yaml:"template-id"`
	SdkAppID   string `mapstructure:"sdk-app-id" json:"sdk-app-id" yaml:"sdk-app-id"`
}

type HuaWeiMessage struct {
	Open       string `mapstructure:"open" json:"open" yaml:"open"`
	AppKey     string `mapstructure:"app-key" json:"app-key" yaml:"app-key"`
	AppSecret  string `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"`
	TemplateID string `mapstructure:"template-id" json:"template-id" yaml:"template-id"`
	AppUrl     string `mapstructure:"app-url" json:"app-url" yaml:"app-url"`
	Signature  string `mapstructure:"signature" json:"signature" yaml:"signature"`
	Sender     string `mapstructure:"sender" json:"sender" yaml:"sender"`
}

type AliyunMessage struct {
	Open            string `mapstructure:"open" json:"open" yaml:"open"`
	AccessKeyID     string `mapstructure:"access-key-id" json:"access-key-id" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"access-key-secret" yaml:"access-key-secret"`
	SignName        string `mapstructure:"sign-name" json:"sign-name" yaml:"sign-name"`
	TemplateCode    string `mapstructure:"template-code" json:"template-code" yaml:"template-code"`
}

type AliyunPhone struct {
	Open            string `mapstructure:"open" json:"open" yaml:"open"`
	AccessKeyID     string `mapstructure:"access-key-id" json:"access-key-id" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"access-key-secret" yaml:"access-key-secret"`
	ShowNumber      string `mapstructure:"show-number" json:"show-number" yaml:"show-number"`
	TemplateCode    string `mapstructure:"template-code" json:"template-code" yaml:"template-code"`
}

type AlertEmail struct {
	Open         string `mapstructure:"open" json:"open" yaml:"open"`
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	FromEmail    string `mapstructure:"from-email" json:"from-email" yaml:"from-email"`
	AcceptEmails string `mapstructure:"accept-emails" json:"accept-emails" yaml:"accept-emails"`
	Title        string `mapstructure:"title" json:"title" yaml:"title"`
}
