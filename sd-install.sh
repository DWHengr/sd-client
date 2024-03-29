
#1 将项目传的服务器上

#创建文件夹
sudo mkdir /usr/local/sd-client

# 将项目拷贝到/usr/local/sd-client目录下
sudo cp -r ./* /usr/local/sd-client

# 修改程序权限
sudo chmod 777 /usr/local/sd-client/sd-client

#在systemd文件夹添加server服务
sudo cat > /etc/systemd/system/sd-client.service <<EOF
[Unit]
Description=sd client local
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/sd-client/sd-client -dir=usr/local/sd-client/
Restart=on-failure
PrivateTmp=true
RestartSec=10s

[Install]
WantedBy=multi-user.target
EOF

sudo chmod 777 /etc/systemd/system/sd-client.service

#重新加在systemctl配置
sudo systemctl daemon-reload
#开机自启sd-client服务
sudo systemctl enable sd-client.service
#开启
sudo systemctl start sd-client.service
#查看服务状态
sudo systemctl status sd-client.service