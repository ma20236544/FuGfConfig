package main

import (
	"ConfGenerateGo/FileOperations"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// url
const (
	// inbox ad data
	// inboxAdUrl = "https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/Advertising/AdRulesBeta.conf"
	// base data
	baseAdUrl1 = "https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/master/rule/Loon/Advertising/Advertising.list"
	// CustomAdRules
	// baseAdUrl2 = "https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/Advertising/AdRules.conf"
	// ios_rule_script QuantumultX Advertising
	inbox1AdUrl = "https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/master/rule/QuantumultX/Advertising/Advertising.list"
	// ios_rule_script Loon Advertising_Domain
	// url4 = "https://raw.githubusercontent.com/blackmatrix7/ios_rule_script/release/rule/Loon/Advertising/Advertising_Domain.list"
)

var inboxFilePath = [...]string{
	"F:\\CodeFile\\Project\\FuGfConfig\\ConfigFile\\Loon\\LoonRemoteRule\\Advertising\\AdRulesBeta.conf",
	"./DataFile/inbox1.txt"}

var baseFilePath = [...]string{
	"./DataFile/base.txt",
	"F:\\CodeFile\\Project\\FuGfConfig\\ConfigFile\\Loon\\LoonRemoteRule\\Advertising\\AdRules.conf"}

var policysMap = make(map[string]string)

func main() {
	println("开始")
	fmt.Println("是否要更新or下载远程数据(y or n)")
	var input string
	// fmt.Scanln(&input)
	input = "y"
	input = "n"
	if input == "y" || input == "Y" {
		// 下载文件
		// FileOperations.DownloadFile(inboxAdUrl, inboxFilePath[0])
		FileOperations.DownloadFile(baseAdUrl1, baseFilePath[0])
		// FileOperations.DownloadFile(baseAdUrl2, baseFilePath[1])
		FileOperations.DownloadFile(inbox1AdUrl, inboxFilePath[1])
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
	for i := 0; i < len(baseFilePath); i++ {
		var ans = FileOperations.ReadFile(baseFilePath[i])
		fmt.Println("base map", i, "共", len(ans), "条数据")
		// 遍历得到的数据
		for _, v := range ans {
			if !isNote(v) {
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

	fmt.Println("基础数据库构建完成，共", len(policysMap), "条数据")

	// 循环读取待处理的数据文件
	var data []string
	for i := 0; i < len(inboxFilePath); i++ {
		var ans = FileOperations.ReadFile(inboxFilePath[i])
		fmt.Println("读取待处理数据", i, "，共", len(ans), "条数据")
		for _, v := range ans {
			if !isNote(v) {
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
						// 仅域名或IP
						if _, ok := policysMap[v]; !ok {
							// if isIPV4(v) || isIPV6(v) {
							if isIPV4(v) {
								b := []string{"IP-CIDR", v}
								// b := []string{"IP-CIDR", v, policyName}
								str = strings.Join(b, ",")
								policysMap[v] = "IP-CIDR"
							} else {
								b := []string{"DOMAIN-SUFFIX", v}
								// b := []string{"DOMAIN-SUFFIX", v, policyName}
								str = strings.Join(b, ",")
								policysMap[v] = "DOMAIN-SUFFIX"
							}
						}
					}
					if str != "" {
						data = append(data, str)
					}
				}

			}
		}
	}

	fmt.Println("处理后共有", len(data), "条 new 数据")

	// 新数据与老数据合并
	var ans = FileOperations.ReadFile("F:\\CodeFile\\Project\\FuGfConfig\\ConfigFile\\Loon\\LoonRemoteRule\\Advertising\\AdRules.conf")
	data = append(data, ans...)

	// 数据结果排序
	sort.Strings(data)

	fmt.Println("更新后去广告规则共有", len(data))
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
	s = strings.Replace(s, "\r", "", -1)
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

func isNote(s string) bool {
	// 忽略注释与 URL-REGEX 规则和空行
	if strings.HasPrefix(s, "#") ||
		strings.HasPrefix(s, ";") ||
		strings.HasPrefix(s, "\n") ||
		strings.HasPrefix(s, "//") ||
		strings.Contains(s, "URL-REGEX") {
		return true
	}
	return false
}

func isIPV4(s string) bool {
	// 判断是否为IPV4
	partIp := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammer := partIp + "\\." + partIp + "\\." + partIp + "\\." + partIp
	matchMe := regexp.MustCompile(grammer)

	return matchMe.MatchString(s)
}

func isIPV6(s string) bool {
	// 判断是否为IPV4
	partIp := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammer := partIp + "\\." + partIp + "\\." + partIp + "\\." + partIp
	matchMe := regexp.MustCompile(grammer)

	return matchMe.MatchString(s)
}
