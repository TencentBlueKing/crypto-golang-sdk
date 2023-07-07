BlueKing crypto-golang-sdk
---------------------------

[(English Documents Available)](https://github.com/TencentBlueKing/crypto-golang-sdk/blob/master/readme_en.md)

## Overview

️🔧 BlueKing crypto-golang-sdk 是一个基于 [铜锁](https://github.com/Tongsuo-Project/Tongsuo/) 的轻量级密码学工具包，为
Golang 应用提供 SM4 的加解密实现

## Features

* [Basic] 支持中国商用密码学算法：SM4
* [Basic] 支持国际主流密码学算法：AES

## Getting started

### Installation

#### 铜锁编译安装

1. 下载铜锁安装包，推荐使用 8.3.2 版本
   https://github.com/Tongsuo-Project/Tongsuo/releases

2. 铜锁库编译和安装
   目前铜锁支持的操作系统有：各种Linux发行版、macOS、Android、iOS和Windows。在这些操作系统上，还需要事先准备好对应的环境：
    - make
    - Perl 5，以及Text::Template模块
    - C编译器
    - C库
      铜锁对第三方库的依赖很少，但是目前依然对Perl依赖较大。

下面以 linux 环境为例，假设铜锁的安装路径为：`/opt/tongsuo`：

``` bash
# --prefix 选项可以自定义安装目录，不设置则默认为 /usr/local/ 下
./config --prefix=/opt/tongsuo
# 编译
make -j
# 安装
make install
```

详细编译安装流程可参考：
https://www.yuque.com/tsdoc/ts/rp7ul8a4ttav8ql9

#### 代码中引用 crypto-golang-sdk

``` golang
import "github.com/TencentBlueKing/crypto-golang-sdk"
```

#### 编译引用了 crypto-golang-sdk 的代码

1. 设置 CGO_ENABLED=1 环境变量，开启 CGO
2. 设置 CGO_CFLAGS 环境变量为 `-I${第一步设置的铜锁库的安装路径}/include -Wno-deprecated-declarations`
3. 设置 CGO_LDFLAGS 环境变量为 `-L/${第一步设置的铜锁库的安装路径}/lib -lssl -lcrypto`

下面以 linux 环境为例，假设铜锁的安装路径为：`/opt/tongsuo`：

``` bash
CGO_ENABLED=1 CGO_CFLAGS="-I/opt/tongsuo/include -Wno-deprecated-declarations" CGO_LDFLAGS="-L/opt/tongsuo/lib -lssl -lcrypto" go build -o main .
```

### Usage

下面以 SM4 加解密为例：

```golang
package main

import (
	"fmt"

	"github.com/TencentBlueKing/crypto-golang-sdk"
)

func main() {
	// 原文
	plaintext := []byte("origin text")

	// 配置的key和iv
	key := []byte("kFkfZiCTSTz0iar2")
	iv := []byte("X85xBwRMfmiG11QP")

	// 创建加解密器
	crypto, err := bkcrypto.NewSm4(key, iv)
	if err != nil {
		fmt.Printf("new crypto failed, err: %v\n", err)
		return
	}

	// 加密操作
	ciphertext, err := crypto.Encrypt(plaintext)
	if err != nil {
		fmt.Printf("encrypt failed, err: %v\n", err)
		return
	}

	// 解密操作
	decrypted, err := crypto.Decrypt(ciphertext)
	if err != nil {
		fmt.Printf("decrypt failed, err: %v\n", err)
		return
	}

	fmt.Printf("encrypted text: %s, decrypted text: %s\n", ciphertext, decrypted)
}
```

## Roadmap

- [版本日志](https://github.com/TencentBlueKing/crypto-golang-sdk/blob/master/RELEASE.md)

## Support

- [蓝鲸论坛](https://bk.tencent.com/s-mart/community)
- [蓝鲸 DevOps 在线视频教程](https://bk.tencent.com/s-mart/video/)
- [蓝鲸社区版交流群](https://jq.qq.com/?_wv=1027&k=5zk8F7G)

## BlueKing Community

- [BK-CMDB](https://github.com/TencentBlueKing/bk-cmdb)：蓝鲸配置平台（蓝鲸 CMDB）是一个面向资产及应用的企业级配置管理平台。
- [BK-CI](https://github.com/TencentBlueKing/bk-ci)：蓝鲸持续集成平台是一个开源的持续集成和持续交付系统，可以轻松将你的研发流程呈现到你面前。
- [BK-BCS](https://github.com/TencentBlueKing/bk-bcs)：蓝鲸容器管理平台是以容器技术为基础，为微服务业务提供编排管理的基础服务平台。
- [BK-PaaS](https://github.com/TencentBlueKing/bk-paas)：蓝鲸 PaaS 平台是一个开放式的开发平台，让开发者可以方便快捷地创建、开发、部署和管理
  SaaS 应用。
- [BK-SOPS](https://github.com/TencentBlueKing/bk-sops)：标准运维（SOPS）是通过可视化的图形界面进行任务流程编排和执行的系统，是蓝鲸体系中一款轻量级的调度编排类
  SaaS 产品。
- [BK-JOB](https://github.com/TencentBlueKing/bk-job) 蓝鲸作业平台(Job)是一套运维脚本管理系统，具备海量任务并发处理能力。

## Contributing

如果你有好的意见或建议，欢迎给我们提 Issues 或 Pull Requests，为蓝鲸开源社区贡献力量。   
[腾讯开源激励计划](https://opensource.tencent.com/contribution) 鼓励开发者的参与和贡献，期待你的加入。

## License

基于 MIT 协议， 详细请参考 [LICENSE](https://github.com/TencentBlueKing/crypto-golang-sdk/blob/master/LICENSE.txt)
