# ysfe

一款可在Linux和macos上运行的小巧文件加密工具，利用个人输入的密码，将文件加密后存储在本地。

需要你记住几个加密指令即可流畅使用

## 指令参数说明
```bash
Usage: ysfe [Options]

Options:
  -l  查看加密文件列表, 输入该值其他选项失效
  -d  删除目标加密文件
  -a  动作, e:加密 u:更新 v:查看解密内容 p:终端查看解密内容
    e -- 加密目标文件
	u -- 目标文件解密后放入临时文件, 关闭程序时加密临时文件并覆盖原加密文件
	v -- 目标文件解密后放入临时文件, 120秒后删除临时文件
	p -- 目标文件解密后直接从终端输出
  -f  要操作的目标文件

a为e时f为原始文件, 其他动作f为加密文件(通过-l获取加密文件列表);
当a不为e时, -f后边只需要输入文件名即可;
```

## 初始化说明
使用 ysfe -l命令后工具会自动检测是否需要初始化，如果需要初始化，按照工具一步一步的提示完成即可。

备份文件路径: 存放加密文件的备份文件, 误删除加密文件后, 可手动将该路径下对应加密文件copy到 "软件运行目录/efile" 目录下即可

## 安装
* mac

curl -L -o install_mac.sh https://github.com/yunsonbai/ysfe/releases/download/install-tool/install_mac.sh && sh install_mac.sh && rm -rf install_mac.sh

如果报权限问题请执行:
curl -L -o install_mac.sh https://github.com/yunsonbai/ysab/releases/download/install-tool/install_mac.sh && sudo sh install_mac.sh && rm -rf install_mac.sh

如果安装完后不能输入 ysab 命令，可以重启终端或者执行 source /etc/profile

* linux

curl -L -o install_linux.sh https://github.com/yunsonbai/ysfe/releases/download/install-tool/install_linux.sh && sh install_linux.sh && rm -rf install_linux.sh

如果报权限问题请执行:
curl -L -o install_linux.sh https://github.com/yunsonbai/ysab/releases/download/install-tool/install_linux.sh && sudo sh install_linux.sh && rm -rf install_linux.sh

## 注意
密码不可忘记，忘记文件将无法恢复