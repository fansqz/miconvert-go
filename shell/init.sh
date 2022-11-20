#!/bin/bash
#author fzw

#更新pip
python -m pip install --upgrade pip
#安装pdf2docx
pip install pdf2docx
#安装libreoffice
wget https://mirrors.cloud.tencent.com/libreoffice/libreoffice/stable/7.4.2/rpm/x86_64/LibreOffice_7.4.2_Linux_x86-64_rpm.tar.gz
tar -zxvf LibreOffice_7.4.2_Linux_x86-64_rpm.tar.gz
cd ./LibreOffice_7.4.2.3_Linux_x86-64_rpm/RPMS
yum -y localinstall *.rpm
libreoffice7.4 --version
#安装libreoffice-headle
yum install -y libreoffice-headless
#安装libreoffice
yum install libreoffice-writer
#软链接
ln -s /opt/libreoffice7.4/program/soffice /usr/bin/soffice