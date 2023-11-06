# sentry-self-hosted

## 本地搭建

### 环境依赖
- Requirements
- Docker 19.03.6+
- Compose 2.0.1+
- 4 CPU Cores
- 8 GB RAM
- 20 GB Free Disk Space

### 步骤
- 初始化项目环境(get_latest_release.sh)
  下载最新版本的自托管存储库

- 运行项目中的 ./install.sh (./self-hosted/install.sh)
  此脚本将处理您入门所需的所有事情，包括环境检查，镜像构建，基础配置
  运行 ./install.sh 时，需要手动选择是否加入改进计划
  快要运行结束时会询问是否要设置账号
  > 构建过程如果遇到问题，可以参考 [这里的修改记录](https://github.com/clh021/self-hosted/tree/23.10.1.localbuild)。如果可能的话，您应该参考[这个章节的内容](https://develop.sentry.dev/self-hosted/#installing-behind-a-proxy)

- 在项目目录运行 docker compose up -d 启动 Sentry
  默认情况下，Sentry 绑定到端口 9000 。

您应该能够访问 http://127.0.0.1:9000 的登录页面。



### 可能有用的输出片段
```
Here's the info we may collect:

  - OS username
  - IP address
  - install log
  - runtime errors
  - performance data

Thirty (30) day retention. No marketing. Privacy policy at sentry.io/privacy.

y or n? n

Understood. To avoid this prompt in the future, use one of these flags:

  --report-self-hosted-issues
  --no-report-self-hosted-issues

or set the REPORT_SELF_HOSTED_ISSUES environment variable:

  REPORT_SELF_HOSTED_ISSUES=1 to send data
  REPORT_SELF_HOSTED_ISSUES=0 to not send data
```

### 快要运行结束时会询问是否要设置账号
```
Created internal Sentry project (slug=internal, id=1)

Would you like to create a user account now? [Y/n]:
Email: clh021@gmail.com
Password:/usr/local/lib/python3.8/getpass.py:91: GetPassWarning: Can not control echo on the terminal.
  passwd = fallback_getpass(prompt, stream)
Warning: Password input may be echoed.
 admin123
Repeat for confirmation:
Warning: Password input may be echoed.
 admin123
Added to organization: sentry
User created: clh021@gmail.com


▶ Setting up GeoIP integration ...
Setting up IP address geolocation ...
Installing (empty) IP address geolocation database ... done.
IP address geolocation is not configured for updates.
See https://develop.sentry.dev/self-hosted/geolocation/ for instructions.
Error setting up IP address geolocation.


-----------------------------------------------------------------

You're all done! Run the following command to get Sentry running:

  docker compose up -d

-----------------------------------------------------------------
```

## 使用

[如何使用](./HOW.md)