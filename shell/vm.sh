#!/bin/bash


sudo mv /etc/ssh/sshd_config /etc/ssh/sshd_config.backup

wget http://172.17.20.4/releases/vm/sshd_config

sudo mv ./sshd_config /etc/ssh

sudo systemctl daemon-reload
sudo systemctl restart ssh
sudo systemctl restart sshd


