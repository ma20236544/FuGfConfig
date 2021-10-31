package main

import (
	"ConfGenerateGo/FileOperations"
	"fmt"
	"strings"
)

// url
const (
	// ios_rule_script Loon Advertising
	url1 = "https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/master/rule/Loon/Advertising/Advertising.list"
	// CustomAdRules
	url2 = "https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/CustomAdRules.conf"
	// ios_rule_script QuantumultX Advertising
	url3 = "https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/master/rule/QuantumultX/Advertising/Advertising.list"
	// ios_rule_script Loon Advertising_Domain
	url4 = "https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/release/rule/Loon/Advertising/Advertising_Domain.list"
)

var filePath = [...]string{"./DataFile/inbox.txt", "./DataFile/base.txt", "./DataFile/CustomAdRules.txt", "./DataFile/QxAdvertising.txt", "./DataFile/LoonDomainAdvertising.txt"}

var policysMap = make(map[string]string)

func main() {
	println("开始")
	fmt.Println("是否要更新or下载远程数据(y or n)")
	var input string
	// fmt.Scanln(&input)
	input = "n"
	if input == "y" {
		// 下载文件
		FileOperations.DownloadFile(url1, filePath[1])
		FileOperations.DownloadFile(url2, filePath[2])
		FileOperations.DownloadFile(url3, filePath[3])
		FileOperations.DownloadFile(url4, filePath[4])
		println("更新完成")
	}

	// 处理文件
	//规则分为三个部分
	//匹配类型 ，匹配关键字，策略名称
	//MatchType MatchingKeywords PolicyName
	policyProcessing()
	println("处理完成")
	// 合并文件
	println("合并完成")
	// 写入文件
	println("写入完成")
	println("结束")
}

// policy processing
func policyProcessing() {
	for i := 2; i <= 4; i++ {
		// 循环读取文件
		var ans = FileOperations.ReadFile(filePath[i])
		// 遍历得到的数据
		for _, v := range ans {
			// var str = v
			if !strings.HasPrefix(v, "#") {
				// 规则命名统一一下
				v = strings.Replace(v, "\n", "", -1)
				v = strings.Replace(v, "HOST", "DOMAIN", 1)
				v = strings.Replace(v, "host", "DOMAIN", 1)
				v = strings.Replace(v, "domain", "DOMAIN", 1)
				v = strings.Replace(v, "IP6-CIDR", "IP-CIDR6", 1)
				v = strings.Replace(v, "ip6-cidr", "IP-CIDR6", 1)
				v = strings.Replace(v, "ip-cidr6", "IP-CIDR6", 1)

				if strings.Count(v, "DOMAIN") > 0 && strings.Count(v, ",") >= 1 {
					// 如果包含 DOMAIN
					var data = strings.Split(v, ",")
					policysMap[data[1]] = data[0]
				} else if strings.Count(v, "IP-CIDR") > 0 || strings.Count(v, "IP-CIDR6") > 0 {
					var data = strings.Split(v, ",")
					policysMap[data[1]] = data[0]
				} else {
					v = strings.Replace(v, "\n", "", -1)
					v = strings.TrimPrefix(v, ".")
					policysMap[v] = "DOMAIN"
				}
			}
		}
	}
}
