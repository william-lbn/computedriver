## 安装

安装方式分为两种：

- 1、手动安装

- 2、上游组件通过启动脚本拉起

手动安装





## 上游组件安装

### 启动脚本语言

shell

### 脚本下发方式

​		--待上游解析

### 脚本接口参数

- 1、文件路径

- 2、端口号

### 过程

1、下载并解压computedriver文件到/usr/local/bin/

```shell
wget -c $1/computedriver-$2 -O - | sudo tar -xz -C /usr/local
```

2、拷贝computedriver.service文件到/etc/systemd/system/

```shell
mv /usr/local/bin/computedriver.service /etc/systemd/system/
```

3、启动computefirmware

```shell
sudo systemctl daemon-reload
sudo systemctl start computedriver
sudo systemctl status computedriver
```



## 手动安装

同上，区别是手工执行命令