# BlueKing crypto-golang-sdk

---

## Overview

️🔧 BlueKing crypto-golang-sdk is a lightweight cryptography toolkit based on encryption libraries such as tongsuo, 
providing a unified encryption and decryption implementation for Golang applications, making it easy for
projects to switch between different encryption methods without intrusion.

## Features

* [Basic] Supports Chinese commercial cryptography algorithms: SM4
* [Basic] Supports mainstream international cryptography algorithms: AES

## Getting started

### Installation

#### tongsuo compilation and installation

1. Download the tongsuo installation package, version 8.3.2 is recommended
   https://github.com/Tongsuo-Project/Tongsuo/releases

2. tongsuo library compilation and installation
   At present, the operating systems supported by tongsuo are: various Linux distributions, macOS, Android, iOS and Windows. On these operating systems, the corresponding environment needs to be prepared in advance:
    - make
    - Perl 5, and the Text::Template module
    - C compiler
    - C library
      Tongsuo has little dependence on third-party libraries, but it still relies heavily on Perl.

Let's take the linux environment as an example, assuming that the tongsuo installation path is: `/opt/tongsuo`:

``` bash
# --prefix option can customize the installation directory, if not set, the default is /usr/local/
./config --prefix=/opt/tongsuo
# compile
make -j
# Install
make install
```

For detailed compilation and installation process, please refer to:
https://www.yuque.com/tsdoc/ts/rp7ul8a4ttav8ql9

#### Reference crypto-golang-sdk in the code

```golang
import "github.com/TencentBlueKing/crypto-golang-sdk"
```

#### Compile code that references crypto-golang-sdk

1. Install the pkg-config tool
2. Set the PKG_CONFIG_PATH environment variable to the pkg-config path of the tongsuo library compiled in the first step (usually the lib/pkgconfig directory under the installation path), and compile

Let's take the Linux environment as an example, assuming that the installation path of tongsuo is: `/opt/tongsuo`, and the path of pkg-config is: `/opt/tongsuo/lib/pkgconfig`:

``` bash
PKG_CONFIG_PATH=/opt/tongsuo/lib/pkgconfig CGO_ENABLED=1 go build -o main .
```

### Usage

Let's take SM4 encryption and decryption as an example:

```golang
package main

import (
	"fmt"

	"github.com/TencentBlueKing/crypto-golang-sdk"
)

func main() {
	// 原文
	plaintext := []byte("origin text")

	// 配置的key
	key := []byte("key\\0")

	// 创建加解密器
	crypto, err := bkcrypto.NewSm4(key)
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

- [Version Log](https://github.com/TencentBlueKing/crypto-golang-sdk/blob/main/release.md)

## Support

- [BK Forum](https://bk.tencent.com/s-mart/community)
- [BK DevOps Online Video Tutorial (In Chinese)](https://bk.tencent.com/s-mart/video/)
- [Technical Exchange QQ Group](https://jq.qq.com/?_wv=1027&k=5zk8F7G)

## BlueKing Community

- [BK-CMDB](https://github.com/Tencent/bk-cmdb): BlueKing CMDB is an enterprise-level management platform designed for
  assets and applications.
- [BK-CI](https://github.com/Tencent/bk-ci): BlueKing Continuous Integration platform is a free, open source CI service,
  which allows developers to automatically create - test - release workflow, and continuously, efficiently deliver their
  high-quality products.
- [BK-BCS](https://github.com/Tencent/bk-bcs): BlueKing Container Service is a container-based basic service platform
  that provides management service to microservice businesses.
- [BK-PaaS](https://github.com/Tencent/bk-paas): BlueKing PaaS is an open development platform that allows developers to
  efficiently create, develop, set up, and manage SaaS apps.
- [BK-SOPS](https://github.com/Tencent/bk-sops): BlueKing SOPS is a system that features workflow arrangement and
  execution using a graphical interface. It's a lightweight task scheduling and arrangement SaaS product of the Blueking
  system.
- [BK-JOB](https://github.com/Tencent/bk-job):BlueKing JOB is a set of operation and maintenance script management
  platform with the ability to handle a large number of tasks concurrently.

## Contributing

If you have good ideas or suggestions, please let us know by Issues or Pull Requests and contribute to the Blue Whale
Open Source Community.      
[Tencent Open Source Incentive Program](https://opensource.tencent.com/contribution) welcome developers from all over
the globe to contribute to Tencent Open Source projects.

## License

Based on the MIT protocol. Please refer to [LICENSE](LICENSE.txt)