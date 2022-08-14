package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

type AppSettingS struct {
	AppName               string
	DefaultPageSize       int
	MaxPageSize           int
	LogSavePath           string
	LogFileName           string
	LogFileExt            string
	LogPrefix             string
	UploadSavePath        string
	UploadServerUrl       string
	UploadImageMaxSize    int
	UploadImageAllowExts  []string
	DefaultContextTimeout time.Duration
}

type DataBaseSettingS struct {
	DBType      string
	UserName    string
	Passwd      string
	Host        string
	DBName      string
	TablePrefix string
	Charset     string
	ParseTime   bool
	MaxIdleCons int
	MaxOpenCons int
	ConTimeOut  string
}

type JWTSetting struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type EmailSetting struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

type DingTalkSetting struct {
	Url string
	To  []string
}

var sections = make(map[string]any)

func (s *Setting) ReadSection(k string, v interface{}) error {

	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
