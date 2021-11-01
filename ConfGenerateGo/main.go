package main

import (
	"ConfGenerateGo/FileOperations"
	"fmt"
	"sort"
	"strings"
)

// url
const (
	// inbox data
	url0 = "https://raw.githubusercontent.com/DivineEngine/Profiles/master/Surge/Ruleset/Guard/Advertising.list"
	// ios_rule_script Loon Advertising
	url1 = "https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/master/rule/Loon/Advertising/Advertising.list"
	// CustomAdRules
	url2 = "https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/CustomAdRules.conf"
	// ios_rule_script QuantumultX Advertising
	url3 = "https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/master/rule/QuantumultX/Advertising/Advertising.list"
	// ios_rule_script Loon Advertising_Domain
	url4 = "https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/release/rule/Loon/Advertising/Advertising_Domain.list"
)

var filePath = [...]string{
	"./DataFile/inbox.txt",
	"./DataFile/base.txt",
	"./DataFile/CustomAdRules.txt",
	"./DataFile/QxAdvertising.txt",
	"./DataFile/LoonDomainAdvertising.txt"}

var policysMap = make(map[string]string)

func main() {
	println("开始")
	fmt.Println("是否要更新or下载远程数据(y or n)")
	var input string
	// fmt.Scanln(&input)
	input = "n"
	if input == "y" {
		// 下载文件
		FileOperations.DownloadFile(url0, filePath[0])
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
	println("结束")
}

// policy processing
func policyProcessing() {
	// 循环读取文件 构建 base map
	for i := 1; i <= 2; i++ {
		var ans = FileOperations.ReadFile(filePath[i])
		// 遍历得到的数据
		for _, v := range ans {
			if !strings.HasPrefix(v, "#") && !strings.Contains(v, "URL-REGEX") {
				v = formatCorrection(v)

				if strings.Count(v, "DOMAIN") > 0 && strings.Count(v, ",") >= 1 {
					// 如果包含 DOMAIN
					var data = strings.Split(v, ",")
					policysMap[data[1]] = data[0]
				} else if strings.Count(v, "IP-CIDR") > 0 || strings.Count(v, "IP-CIDR6") > 0 {
					var data = strings.Split(v, ",")
					policysMap[data[1]] = data[0]
				} else {
					policysMap[v] = "DOMAIN"
				}
			}
		}
	}

	// 循环读取待处理的数据文件
	var data []string
	for i := 0; i <= 0; i++ {
		var ans = FileOperations.ReadFile(filePath[i])
		for _, v := range ans {
			if !strings.HasPrefix(v, "#") && !strings.Contains(v, "URL-REGEX") {
				v = formatCorrection(v)

				// if v == "USER-AGENT,AVOS*" {
				// 	fmt.Println(v)
				// }

				if v != "" {
					var str string
					if strings.Contains(v, ",") {
						var a = strings.Split(v, ",")
						if _, ok := policysMap[a[1]]; !ok {
							b1 := []string{a[0], a[1], "REJECT"}
							b2 := []string{a[0], a[1], "REJECT", "no-resolve"}
							if strings.Contains(v, "IP-CIDR") || strings.Contains(v, "IP-CIDR6") {
								str = strings.Join(b2, ",")
							} else {
								str = strings.Join(b1, ",")
							}
							policysMap[a[1]] = a[0]
						}
					} else {
						if _, ok := policysMap[v]; !ok {
							b := []string{"DOMAIN-SUFFIX", v, "REJECT"}
							str = strings.Join(b, ",")
							policysMap[v] = "DOMAIN-SUFFIX"
						}
					}
					if str != "" {
						data = append(data, str)
					}
				}

			}
		}
	}

	// 新数据与老数据合并
	var ans = FileOperations.ReadFile(filePath[2])
	data = append(data, ans...)

	// 结果排序
	sort.Strings(data)

	fmt.Println(len(data))
	// 写入文件
	// FileOperations.WriteFile(data, "F:\\CodeFile\\Project\\FuGfConfig\\ConfigFile\\Loon\\CustomAdRules.conf")
	FileOperations.WriteFile(data, "./DataFile/ans.txt")
}

// 规则格式统一
func formatCorrection(s string) string {
	s = strings.TrimPrefix(s, ".")
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "HOST", "DOMAIN", 1)
	s = strings.Replace(s, "host", "DOMAIN", 1)
	s = strings.Replace(s, "domain", "DOMAIN", 1)
	s = strings.Replace(s, "DOMAIN-suffix", "DOMAIN-SUFFIX", 1)
	s = strings.Replace(s, "IP6-CIDR", "IP-CIDR6", 1)
	s = strings.Replace(s, "ip6-cidr", "IP-CIDR6", 1)
	s = strings.Replace(s, "ip-cidr6", "IP-CIDR6", 1)

	return s
}
