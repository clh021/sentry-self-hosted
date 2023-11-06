# 如何本地搭建

Sentry 是一个开源的非常强大的实时异常收集系统，可以为开发者的提供帮助、诊断，修复和优化其代码的性能的能力，可以用它来监控线上服务的健康状态，实时收集的异常堆栈信息可以帮助我们快速发现、定位和修复问题。支持 Web 前后端、移动应用以及游戏，90 余种主流语言和相关框架，同时还提供了非常友好的管理页面，错误告警、指派、统计等其他丰富的功能。Sentry 提供了商用的服务，可以直接在官网购买。同时也提供了开源版本，可以自行搭建安装。本项目探索本地搭建过程和使用细节。更多介绍可以查看[官方文档](https://docs.sentry.io)。

### 环境依赖
- Docker 19.03.6+
- Compose 2.0.1+
- 4 CPU Cores
- 8 GB RAM
- 20 GB Free Disk Space

### 步骤
1. 初始化项目环境

    `get_latest_release.sh` 下载最新版本的自托管存储库

2. 运行项目中的 ./install.sh

    `./self-hosted/install.sh`

    此脚本将处理您入门所需的所有事情，包括环境检查，镜像构建，基础配置
    运行 ./install.sh 时，需要手动选择是否加入改进计划
    快要运行结束时会询问是否要设置账号

    > 构建过程如果遇到问题，可以参考 [这里的修改记录](https://github.com/clh021/self-hosted/tree/23.10.1.localbuild)。如果可能的话，您应该参考[这个章节的内容](https://develop.sentry.dev/self-hosted/#installing-behind-a-proxy)

3. 在项目目录运行 `docker compose up -d` 启动 `Sentry`

    默认情况下，Sentry 绑定到端口 9000 。


搭建完成。访问 http://127.0.0.1:9000 即可登录。



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

# 如何使用

[如何使用](./HOW.md)