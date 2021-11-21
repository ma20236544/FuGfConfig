package main

import (
	"ConfGenerateGo/FileOperations"
	"fmt"
	"sort"
	"strings"
)

// url
const (
	// inbox ad data
	inboxAdUrl = "https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/Advertising/AdRulesBeta.conf"
	// base data
	baseAdUrl1 = "https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/master/rule/Loon/Advertising/Advertising.list"
	// CustomAdRules
	baseAdUrl2 = "https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/Advertising/AdRules.conf"
	// ios_rule_script QuantumultX Advertising
	// url3 = "https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/master/rule/QuantumultX/Advertising/Advertising.list"
	// ios_rule_script Loon Advertising_Domain
	// url4 = "https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/release/rule/Loon/Advertising/Advertising_Domain.list"
)

var filePath = [...]string{
	"./DataFile/inbox.txt",
	"./DataFile/base.txt",
	"./DataFile/CustomAdRules.txt"}

// "./DataFile/QxAdvertising.txt",
// "./DataFile/LoonDomainAdvertising.txt"}

var policysMap = make(map[string]string)

func main() {
	println("开始")
	fmt.Println("是否要更新or下载远程数据(y or n)")
	var input string
	// fmt.Scanln(&input)
	input = "y"
	// input = "n"
	if input == "y" || input == "Y" {
		// 下载文件
		FileOperations.DownloadFile(inboxAdUrl, filePath[0])
		FileOperations.DownloadFile(baseAdUrl1, filePath[1])
		FileOperations.DownloadFile(baseAdUrl2, filePath[2])
		// FileOperations.DownloadFile(url3, filePath[3])
		// FileOperations.DownloadFile(url4, filePath[4])
		println("更新完成")
	}

	// 处理文件
	//规则分为三个部分
	//匹配类型 ，匹配关键字，策略名称
	//MatchType MatchingKeywords PolicyName
	policyProcessing("REJECT")
	println("处理完成")
	println("结束")
}

// policy processing
func policyProcessing(policyName string) {
	// 循环读取文件 构建 base map
	for i := 1; i <= 2; i++ {
		var ans = FileOperations.ReadFile(filePath[i])
		// 遍历得到的数据
		for _, v := range ans {
			if !strings.HasPrefix(v, "#") &&
				!strings.HasPrefix(v, ";") &&
				!strings.Contains(v, "URL-REGEX") {
				// 忽略注释与 URL-REGEX 规则
				v = formatCorrection(v)

				if (strings.Count(v, "DOMAIN") > 0 && strings.Count(v, ",") >= 1) ||
					(strings.Count(v, "IP-CIDR") > 0 || strings.Count(v, "IP-CIDR6") > 0) ||
					(strings.Count(v, "USER-AGENT") > 0 && strings.Count(v, ",") >= 1) {
					// 如果包含 DOMAIN 或者 IP 或者 USER-AGENT
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
			if !strings.HasPrefix(v, "#") &&
				!strings.HasPrefix(v, ";") &&
				!strings.Contains(v, "URL-REGEX") {
				v = formatCorrection(v)

				if v != "" {
					var str string
					if strings.Contains(v, ",") {
						var a = strings.Split(v, ",")
						if _, ok := policysMap[a[1]]; !ok {
							b1 := []string{a[0], a[1]}
							// b1 := []string{a[0], a[1], policyName}
							b2 := []string{a[0], a[1], "no-resolve"}
							if strings.Contains(v, "IP-CIDR") || strings.Contains(v, "IP-CIDR6") {
								str = strings.Join(b2, ",")
							} else {
								str = strings.Join(b1, ",")
							}
							policysMap[a[1]] = a[0]
						}
					} else {
						if _, ok := policysMap[v]; !ok {
							b := []string{"DOMAIN-SUFFIX", v, policyName}
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

	// 数据结果排序
	sort.Strings(data)

	fmt.Println(len(data))
	// 写入文件
	FileOperations.WriteFile(data, "./DataFile/ans.txt")
	FileOperations.WriteFile(data, "F:\\CodeFile\\Project\\FuGfConfig\\ConfigFile\\Loon\\LoonRemoteRule\\Advertising\\AdRules.conf")
	FileOperations.WriteClashFile(data, "./DataFile/ans1.txt")
	FileOperations.WriteClashFile(data, "F:\\CodeFile\\Project\\FuGfConfig\\ConfigFile\\Clash\\AdRules.txt")
	//清除 bataAd 规则
	var ans1 []string
	ans1 = append(ans1, data[0])
	FileOperations.WriteFile(ans1, "./DataFile/inbox.txt")
	FileOperations.WriteFile(ans1, "F:\\CodeFile\\Project\\FuGfConfig\\ConfigFile\\Loon\\LoonRemoteRule\\Advertising\\AdRulesBeta.conf")
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
	s = strings.Replace(s, "ip-cidr,", "IP-CIDR", 1)
	s = strings.Replace(s, "USER-agent,", "USER-AGENT", 1)
	s = strings.Replace(s, "user-agent,", "USER-AGENT", 1)
	s = strings.Replace(s, "user-AGENT,", "USER-AGENT", 1)

	return s
}
