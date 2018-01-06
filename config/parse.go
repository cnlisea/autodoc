package config

import (
	"strings"
	"strconv"
	"encoding/json"
	cfg "github.com/cnlisea/automation/config"
	"github.com/cnlisea/automation/utils"
)

func ParseConfigFile(path string) ([]interface{}, error) {
	// read all
	b, err := utils.ReadFileAll(path)
	if nil != err {
		return nil, err
	}

	var config []interface{}

	data := make(map[string]interface{})
	if err = json.Unmarshal(b, &data); nil != err {
		return nil, err
	}

	if 0 == len(data) {
		return []interface{}{}, nil
	}

	for k, v := range data {
		switch strings.ToLower(k) {
		case "baseurl":
			fallthrough
		case "base_url":
			fallthrough
		case "url":
			cfg.BaseUrl = utils.ToString(v)
		case "authtype":
			fallthrough
		case "auth_type":
			fallthrough
		case "auth":
			authType, err := strconv.Atoi(utils.ToString(v))
			if nil != err {
				return nil, err
			}
			cfg.AuthType = authType
		case "token_contain":
			fallthrough
		case "tokencontain":
			cfg.TokenContain = utils.ToString(v)
		case "tokenkey":
			fallthrough
		case "token_key":
			cfg.TokenKey = utils.ToString(v)
		case "token":
			cfg.Token = utils.ToString(v)
		default:
			cData, ok := v.(map[string]interface{})
			if !ok {
				continue
			}

			if !strings.Contains(k, "_") {
				for kk, vv := range cData {
					ccData, ok := vv.(map[string]interface{})
					if !ok {
						continue
					}

					var req, res interface{}
					if val, ok := ccData["req"]; ok {
						req = val
					}

					if val, ok := ccData["request"]; ok {
						req = val
					}

					if val, ok := ccData["res"]; ok {
						res = val
					}

					if val, ok := ccData["response"]; ok {
						res = val
					}

					if nil == req || nil == res {
						//TODO logs req or res is nil
						continue
					}

					req, res = cfg.RecoverType(req), cfg.RecoverType(res)
					/*fmt.Println("req:")
					for k, v := range req.(map[string]interface{}){
						t := reflect.TypeOf(v)
						fmt.Println(k, " ", t.String())
					}
					fmt.Println("res:")
					for k, v := range res.(map[string]interface{}){
						t := reflect.TypeOf(v)
						fmt.Println(k, " ", t.String())
					}*/
					config = append(config, k+"_"+kk, req, res)
				}
			} else {
				var req, res interface{}
				if val, ok := cData["req"]; ok {
					req = val
				}

				if val, ok := cData["request"]; ok {
					req = val
				}

				if val, ok := cData["res"]; ok {
					res = val
				}

				if val, ok := cData["response"]; ok {
					res = val
				}

				if nil == req || nil == res {
					//TODO logs req or res is nil
					continue
				}

				if v, ok := req.(map[string]interface{}); ok {
					_ = v
				}

				if v, ok := res.(map[string]interface{}); ok {
					_ = v
				}

				req, res = cfg.RecoverType(req), cfg.RecoverType(res)
				/*fmt.Println("req:")
				for k, v := range req.(map[string]interface{}){
					t := reflect.TypeOf(v)
					fmt.Println(k, " ", t.String())
				}
				fmt.Println("res:")
				for k, v := range res.(map[string]interface{}){
					t := reflect.TypeOf(v)
					fmt.Println(k, " ", t.String())
				}*/
				config = append(config, k, req, res)
			}
		}
	}

	return config, nil
}