#!/bin/bash

mkdir computedriver
wget http://172.17.20.4/releases/computedriver/install.sh
chmod +x install.sh
./install.sh http://172.17.20.4/releases v1.1
