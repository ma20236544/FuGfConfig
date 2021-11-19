# FuGfConfig

## 特别声明

本项目中所有内容只供学习和研究使用，不得将本项目中任何内容用于违反任何国家/地区/组织等的法律法规或相关规定的其他用途。

## 支持

本项目对 Quan X、Loon 提供完全支持

对小火箭提供能用的支持（仅仅是能用

优先级： Loon > Quan X > 小火箭

## 使用方法

### Quan X

建议配合 ios_rule_script 项目中的 [Quan X 去广告](https://github.com/blackmatrix7/ios_rule_script/tree/master/rule/QuantumultX/Advertising) 一起使用（已去重

优先级从高到低：

```
CustomAdRules
AdRules
CustomRules
AppleRules
GFWRules
TelegramRules
BasicRules
```

### Loon

#### Loon 分流规则

优先级从高到低：

```
CustomNoAdProxy
CustomNoAdDirect
CustomAdRules
FuckRogueSoftware
CustomProxy
TelegramRules
AppleRules
AppleAPIRules
AppleCDNRules
AppleNoChinaCDNRules
GFWRules
CustomDirect
BaseRules
```

```
https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/CustomNoAdProxy.conf, policy=PROXY, tag=CustomNoAdProxy, enabled=true

https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/CustomNoAdDirect.conf, policy=DIRECT, tag=CustomNoAdDirect, enabled=true

https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/CustomAdRules.conf, policy=Advertising, tag=CustomAd, enabled=true

https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/CustomProxy.conf, policy=PROXY, tag=CustomProxy, enabled=true

https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/TelegramRules.conf, policy=PROXY, tag=TelegramRules, enabled=true

https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/AppleRules.conf, policy=Apple, tag=Apple, enabled=true

https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/AppleAPIRules.conf, policy=AppleAPI, tag=AppleAPI, enabled=true

https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/AppleCDNRules.conf, policy=AppleCDN, tag=AppleCDN, enabled=true

https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/AppleNoChinaCDNRules.conf, policy=AppleNoChinaCDN, tag=AppleNoChinaCDN, enabled=true

https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/GFWRules.conf, policy=PROXY, tag=FuckGFW, enabled=true

https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/CustomDirect.conf, policy=DIRECT, tag=CustomDirect, enabled=true

https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/LoonRemoteRule/BaseRules.conf, policy=DIRECT, tag=BaseRules, enabled=true

```

### 对于 FuckRogueSoftware 规则的说明

此规则极其激进，对某些国内软件强屏蔽，包括但不限于广告，跟踪，数据分析，仅保证软件最低程度功能的正常使用，使用需谨慎

```
https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/FuckRogueSoftware.conf, policy=Advertising, tag=FuckRogueSoftware, enabled=true
```

### 对于 Apple 解锁功能

请查看 [iRingo 解锁完整的 Apple 功能和集成服务](https://github.com/VirgilClyne/iRingo) 仓库

**建议优先采用上文仓库的规则**

### 对于 Apple 分流规则

请参考 [这篇文章](https://blog.royli.dev/2019/better-proxy-rules-for-apple-services)

对于本项目
AppleRules 是与本地化息息相关的规则，比如地图、天气、查找、Facetime、Apple Pay
( iCloud 上传与下载也归于此规则集
AppleCDNRules 是苹果的全球 CDN
AppleNoChinaCDNRules 是大陆没有的 CDN 节点
AppleAPIRules 是苹果的 API 域名

#### 使用中国区账号（App Store + iCloud）

AppleRules 直连
AppleCDNRules 直连
AppleNoChinaCDNRules 代理
AppleAPIRules 直连

#### 使用美国区账号（App Store + iCloud）

AppleRules 直连
AppleCDNRules 直连
AppleNoChinaCDNRules 代理
AppleAPIRules 代理

建议 AppleAPIRules 依然直连，上文是根据上述文章给出的建议，请结合自身情况使用

### Loon 插件

#### DNSMap

对常用的国内域名进行 DNS 解析分流 ，国内走国内的各大 Doh

```
# DNS 映射
https://raw.githubusercontent.com/dunLan0/FuGfConfig/main/ConfigFile/Loon/LoonPlugin/DNSMap.plugin, tag=DNS Map, enabled=true
```

## 感谢

本项目的数据大部分来自一下项目，在此对他们表示感谢（以下排名不分先后

[ios_rule_script](https://github.com/blackmatrix7/ios_rule_script)

[NextDNS 维护的系统级跟踪列表](https://github.com/nextdns/metadata/tree/master/privacy/native)

[Shadowrocket-ADBlock-Rules](https://github.com/h2y/Shadowrocket-ADBlock-Rules)

[neohosts](https://github.com/neoFelhz/neohosts)

[gfwlist](https://github.com/gfwlist/gfwlist)

[SS-Rule-Snippet](https://github.com/Hackl0us/SS-Rule-Snippet#%E5%85%B3%E4%BA%8E%E9%A1%B9%E7%9B%AE)

[ios_rule_script](https://github.com/blackmatrix7/ios_rule_script)
