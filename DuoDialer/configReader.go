package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var dirPath string
var redisIp string
var redisPort string
var redisPassword string
var securityIp string
var securityPort string
var redisDb int
var dialerId string
var campaignLimit int
var lbIpAddress string
var lbPort string
var port string
var campaignRequestFrequency time.Duration
var campaignServiceHost string
var campaignServicePort string
var callServerHost string
var callServerPort string
var callRuleServiceHost string
var callRuleServicePort string
var scheduleServiceHost string
var scheduleServicePort string
var callbackServerHost string
var callbackServerPort string
var ardsServiceHost string
var ardsServicePort string
var notificationServiceHost string
var notificationServicePort string
var clusterConfigServiceHost string
var clusterConfigServicePort string
var casServerHost string
var v5_1SecurityToken string
var accessToken string

func GetDirPath() string {
	envPath := os.Getenv("GO_CONFIG_DIR")
	if envPath == "" {
		envPath = "./"
	}
	fmt.Println(envPath)
	return envPath
}

func GetDefaultConfig() Configuration {
	confPath := filepath.Join(dirPath, "conf.json")
	fmt.Println("GetDefaultConfig config path: ", confPath)
	content, operr := ioutil.ReadFile(confPath)
	if operr != nil {
		fmt.Println(operr)
	}

	defconfiguration := Configuration{}
	deferr := json.Unmarshal(content, &defconfiguration)

	if deferr != nil {
		fmt.Println("error:", deferr)
		defconfiguration.RedisIp = "127.0.0.1"
		defconfiguration.RedisPort = "6379"
		defconfiguration.SecurityIp = "127.0.0.1"
		defconfiguration.SecurityPort = "6389"
		defconfiguration.RedisPassword = "DuoS123"
		defconfiguration.RedisDb = 5
		defconfiguration.DialerId = "Dialer2"
		defconfiguration.CampaignLimit = 30
		defconfiguration.LbIpAddress = "192.168.0.15"
		defconfiguration.LbPort = "2226"
		defconfiguration.Port = "2226"
		defconfiguration.CampaignRequestFrequency = 300
		defconfiguration.CampaignServiceHost = "192.168.0.143"
		defconfiguration.CampaignServicePort = "2222"
		defconfiguration.CallServerHost = "192.168.0.53"
		defconfiguration.CallServerPort = "8080"
		defconfiguration.CallRuleServiceHost = "192.168.0.89"
		defconfiguration.CallRuleServicePort = "2220"
		defconfiguration.ScheduleServiceHost = "192.168.3.200"
		defconfiguration.ScheduleServicePort = "2224"
		defconfiguration.CallbackServerHost = "192.168.0.15"
		defconfiguration.CallbackServerPort = "2227"
		defconfiguration.ArdsServiceHost = "192.168.0.15"
		defconfiguration.ArdsServicePort = "2225"
		defconfiguration.NotificationServiceHost = "192.168.0.77"
		defconfiguration.NotificationServicePort = "8086"
		defconfiguration.ClusterConfigServiceHost = "127.0.0.1"
		defconfiguration.ClusterConfigServicePort = "3434"
		defconfiguration.CasServerHost = "localhost:20946"
		defconfiguration.V5_1SecurityToken = ""
		defconfiguration.AccessToken = ""
	}

	return defconfiguration
}

func LoadDefaultConfig() {

	defconfiguration := GetDefaultConfig()

	redisIp = defconfiguration.RedisIp
	redisPort = defconfiguration.RedisPort
	redisPassword = defconfiguration.RedisPassword
	securityIp = defconfiguration.SecurityIp
	securityPort = defconfiguration.SecurityPort
	redisDb = defconfiguration.RedisDb
	dialerId = defconfiguration.DialerId
	campaignLimit = defconfiguration.CampaignLimit
	lbIpAddress = defconfiguration.LbIpAddress
	lbPort = defconfiguration.LbPort
	port = defconfiguration.Port
	campaignRequestFrequency = defconfiguration.CampaignRequestFrequency
	campaignServiceHost = defconfiguration.CampaignServiceHost
	campaignServicePort = defconfiguration.CampaignServicePort
	callServerHost = defconfiguration.CallServerHost
	callServerPort = defconfiguration.CallServerPort
	callRuleServiceHost = defconfiguration.CallRuleServiceHost
	callRuleServicePort = defconfiguration.CallRuleServicePort
	scheduleServiceHost = defconfiguration.ScheduleServiceHost
	scheduleServicePort = defconfiguration.ScheduleServicePort
	callbackServerHost = defconfiguration.CallbackServerHost
	callbackServerPort = defconfiguration.CallbackServerPort
	ardsServiceHost = defconfiguration.ArdsServiceHost
	ardsServicePort = defconfiguration.ArdsServicePort
	notificationServiceHost = defconfiguration.NotificationServiceHost
	notificationServicePort = defconfiguration.NotificationServicePort
	clusterConfigServiceHost = defconfiguration.ClusterConfigServiceHost
	clusterConfigServicePort = defconfiguration.ClusterConfigServicePort
	casServerHost = defconfiguration.CasServerHost
	v5_1SecurityToken = defconfiguration.V5_1SecurityToken
	accessToken = defconfiguration.AccessToken

	redisIp = fmt.Sprintf("%s:%s", redisIp, redisPort)
}

func LoadConfiguration() {
	dirPath = GetDirPath()
	confPath := filepath.Join(dirPath, "custom-environment-variables.json")
	fmt.Println("InitiateRedis config path: ", confPath)

	content, operr := ioutil.ReadFile(confPath)
	if operr != nil {
		fmt.Println(operr)
	}

	envconfiguration := EnvConfiguration{}
	enverr := json.Unmarshal(content, &envconfiguration)
	if enverr != nil {
		fmt.Println("error:", enverr)
		LoadDefaultConfig()
	} else {
		var converr error
		defConfig := GetDefaultConfig()

		redisIp = os.Getenv(envconfiguration.RedisIp)
		redisPort = os.Getenv(envconfiguration.RedisPort)
		redisPassword = os.Getenv(envconfiguration.RedisPassword)
		securityIp = os.Getenv(envconfiguration.SecurityIp)
		securityPort = os.Getenv(envconfiguration.SecurityPort)
		redisDb, converr = strconv.Atoi(os.Getenv(envconfiguration.RedisDb))
		dialerId = os.Getenv(envconfiguration.DialerId)
		campaignLimit, converr = strconv.Atoi(os.Getenv(envconfiguration.CampaignLimit))
		lbIpAddress = os.Getenv(envconfiguration.LbIpAddress)
		lbPort = os.Getenv(envconfiguration.LbPort)
		port = os.Getenv(envconfiguration.Port)
		campaignRequestFrequencytemp := os.Getenv(envconfiguration.CampaignRequestFrequency)
		campaignServiceHost = os.Getenv(envconfiguration.CampaignServiceHost)
		campaignServicePort = os.Getenv(envconfiguration.CampaignServicePort)
		callServerHost = os.Getenv(envconfiguration.CallServerHost)
		callServerPort = os.Getenv(envconfiguration.CallServerPort)
		callRuleServiceHost = os.Getenv(envconfiguration.CallRuleServiceHost)
		callRuleServicePort = os.Getenv(envconfiguration.CallRuleServicePort)
		scheduleServiceHost = os.Getenv(envconfiguration.ScheduleServiceHost)
		scheduleServicePort = os.Getenv(envconfiguration.ScheduleServicePort)
		callbackServerHost = os.Getenv(envconfiguration.CallbackServerHost)
		callbackServerPort = os.Getenv(envconfiguration.CallbackServerPort)
		ardsServiceHost = os.Getenv(envconfiguration.ArdsServiceHost)
		ardsServicePort = os.Getenv(envconfiguration.ArdsServiceHost)
		notificationServiceHost = os.Getenv(envconfiguration.NotificationServiceHost)
		notificationServicePort = os.Getenv(envconfiguration.NotificationServicePort)
		clusterConfigServiceHost = os.Getenv(envconfiguration.ClusterConfigServiceHost)
		clusterConfigServicePort = os.Getenv(envconfiguration.ClusterConfigServicePort)
		casServerHost = os.Getenv(envconfiguration.CasServerHost)
		v5_1SecurityToken = os.Getenv(envconfiguration.V5_1SecurityToken)
		accessToken = os.Getenv(envconfiguration.AccessToken)

		if redisIp == "" {
			redisIp = defConfig.RedisIp
		}
		if redisPort == "" {
			redisPort = defConfig.RedisPort
		}
		if redisPassword == "" {
			redisPassword = defConfig.RedisPassword
		}
		if redisDb == 0 || converr != nil {
			redisDb = defConfig.RedisDb
		}
		if securityIp == "" {
			securityIp = defConfig.SecurityIp
		}
		if securityPort == "" {
			securityPort = defConfig.SecurityPort
		}
		if dialerId == "" {
			dialerId = defConfig.DialerId
		}
		if campaignLimit == 0 || converr != nil {
			campaignLimit = defConfig.CampaignLimit
		}
		if lbIpAddress == "" {
			lbIpAddress = defConfig.LbIpAddress
		}
		if lbPort == "" {
			lbPort = defConfig.LbPort
		}
		if port == "" {
			port = defConfig.Port
		}
		if campaignRequestFrequencytemp == "" {
			campaignRequestFrequency = defConfig.CampaignRequestFrequency
		} else {
			campaignRequestFrequency, _ = time.ParseDuration(campaignRequestFrequencytemp)
		}
		if campaignServiceHost == "" {
			campaignServiceHost = defConfig.CampaignServiceHost
		}
		if campaignServicePort == "" {
			campaignServicePort = defConfig.CampaignServicePort
		}
		if callServerHost == "" {
			callServerHost = defConfig.CallServerHost
		}
		if callServerPort == "" {
			callServerPort = defConfig.CallServerPort
		}
		if callRuleServiceHost == "" {
			callRuleServiceHost = defConfig.CallRuleServiceHost
		}
		if callRuleServicePort == "" {
			callRuleServicePort = defConfig.CallRuleServicePort
		}
		if scheduleServiceHost == "" {
			scheduleServiceHost = defConfig.ScheduleServiceHost
		}
		if scheduleServicePort == "" {
			scheduleServicePort = defConfig.ScheduleServicePort
		}
		if callbackServerHost == "" {
			callbackServerHost = defConfig.CallbackServerHost
		}
		if callbackServerPort == "" {
			callbackServerPort = defConfig.CallbackServerPort
		}
		if ardsServiceHost == "" {
			ardsServiceHost = defConfig.ArdsServiceHost
		}
		if ardsServicePort == "" {
			ardsServicePort = defConfig.ArdsServicePort
		}
		if notificationServiceHost == "" {
			notificationServiceHost = defConfig.NotificationServiceHost
		}
		if notificationServicePort == "" {
			notificationServicePort = defConfig.NotificationServicePort
		}
		if clusterConfigServiceHost == "" {
			clusterConfigServiceHost = defConfig.ClusterConfigServiceHost
		}
		if clusterConfigServicePort == "" {
			clusterConfigServicePort = defConfig.ClusterConfigServicePort
		}
		if casServerHost == "" {
			casServerHost = defConfig.CasServerHost
		}
		if v5_1SecurityToken == "" {
			v5_1SecurityToken = defConfig.V5_1SecurityToken
		}
		if accessToken == "" {
			accessToken = defConfig.AccessToken
		}

		redisIp = fmt.Sprintf("%s:%s", redisIp, redisPort)
		securityIp = fmt.Sprintf("%s:%s", securityIp, securityPort)
	}

	fmt.Println("redisIp:", redisIp)
	fmt.Println("redisDb:", redisDb)
	fmt.Println("securityIp:", securityIp)
	fmt.Println("dialerId:", dialerId)
	fmt.Println("campaignLimit:", campaignLimit)
}

func LoadCallbackConfiguration() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in LoadCallbackConfiguration", r)
		}
	}()
	//Request campaign callback reaseons from Campaign Manager service
	jwtToken := fmt.Sprintf("Bearer %s", accessToken)
	client := &http.Client{}

	request := fmt.Sprintf("http://%s/DVP/API/1.0.0.0/CampaignManager/Campaign/Configuration/callback/Reasons", CreateHost(campaignServiceHost, campaignServicePort))
	fmt.Println("Start RequestCampaignCallbackReason request: ", request)
	req, _ := http.NewRequest("GET", request, nil)
	req.Header.Set("authorization", jwtToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("LoadCallbackConfiguration Failed::", err.Error())
	}
	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(response))

	callbackConf := CallbackConfiguration{}
	err = json.Unmarshal(response, &callbackConf)
	if err != nil {
		fmt.Println("error in LoadCallbackConfiguration::", err)
	} else {
		for _, conf := range callbackConf.Result {
			for _, hangCause := range conf.HangupCause {
				confKey := fmt.Sprintf("CallbackReason:%s", hangCause)
				fmt.Println(confKey, "::", conf.Reason)
				RedisSetNx(confKey, conf.Reason)
			}
		}
	}
}
