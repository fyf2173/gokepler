package gokepler

import (
	"github.com/fyf2173/gokepler/auth"
	"github.com/fyf2173/gokepler/jcq"
)

type Config struct {
	Pin        string
	ChannelId  int64
	CustomerID int64
	AppID      string
	AppKey     string
	AppSecret  string
	Token      string

	MqAccessKey    string
	MqAccessSecret string
	MqTenantId     int64
	MqGroupId      string
}

type Option func(conf *Config)

func OptPin(pin string) Option {
	return func(conf *Config) {
		conf.Pin = pin
	}
}

func OptChannelId(channelId int64) Option {
	return func(conf *Config) {
		conf.ChannelId = channelId
	}
}

func OptCustomerId(customerId int64) Option {
	return func(conf *Config) {
		conf.CustomerID = customerId
	}
}

func OptToken(token string) Option {
	return func(conf *Config) {
		conf.Token = token
	}
}

func OptMqAccessKey(accessKey string) Option {
	return func(conf *Config) {
		conf.MqAccessKey = accessKey
	}
}

func OptMqAccessSecret(accessSecret string) Option {
	return func(conf *Config) {
		conf.MqAccessSecret = accessSecret
	}
}

func OptMqTenantId(tenantId int64) Option {
	return func(conf *Config) {
		conf.MqTenantId = tenantId
	}
}

func OptMqGroupId(groupId string) Option {
	return func(conf *Config) {
		conf.MqGroupId = groupId
	}
}

func NewKepler(appKey, appSecret string, opts ...Option) *Config {
	conf := &Config{AppKey: appKey, AppSecret: appSecret}
	for _, f := range opts {
		f(conf)
	}
	return conf
}

func (conf *Config) NewJcqClient() *jcq.Client {
	if conf.MqAccessKey == "" || conf.MqAccessSecret == "" || conf.MqGroupId == "" {
		return nil
	}
	return jcq.NewClient(conf.AppKey, conf.MqTenantId, conf.MqAccessKey, conf.MqAccessSecret).WithGroupId(conf.MqGroupId)
}

func (conf *Config) NewAuthClient() *auth.AccessClient {
	if conf.AppKey == "" || conf.AppSecret == "" {
		return nil
	}
	return auth.NewAccessClient(conf.AppKey, conf.AppSecret)
}
