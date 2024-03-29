#!/bin/bash
# 远程服务器IP地址
REMOTE_HOST=$1
# 远程服务器用户名
REMOTE_USER=$2
# 远程服务器密码
REMOTE_PASSWORD=$3
# 本地文件/目录路径
LOCAL_DIR="./"
# 远程服务器目标目录路径
REMOTE_DIR="/usr/local/sd-client"
# 检查远程目录是否存在，如果不存在则创建
sshpass -p "$REMOTE_PASSWORD" ssh $REMOTE_USER@$REMOTE_HOST "mkdir -p $REMOTE_DIR"
# 使用scp命令传输文件
sshpass -p "$REMOTE_PASSWORD" scp -r $LOCAL_DIR $REMOTE_USER@$REMOTE_HOST:$REMOTE_DIR
# 通过ssh远程执行启动程序命令
sshpass -p "$REMOTE_PASSWORD" ssh $REMOTE_USER@$REMOTE_HOST "sudo cd $REMOTE_DIR && sudo chmod 777 sd-install && sudo ./sd-install"