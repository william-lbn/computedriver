#!/bin/bash


mkdir  computefirmware
cd computefirmware
wget http://172.17.20.4/releases/computefirmware/install.sh
chmod +x install.sh
./install.sh runc


runc --version