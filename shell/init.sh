#!/bin/bash
#author fzw

#保存配置路径
shell_path=`pwd`
mkdir /var/miconvert
cd /var/miconvert
#更新yum
yum -y update
#更新pip
python -m pip install --upgrade pip
#安装pdf2docx
pip install pdf2docx
#安装libreoffice
mkdir /var/miconvert/file
cd /var/miconvert/file
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

#安装docker，一键安装
curl -sSL https://get.daocloud.io/docker | sh
systemctl start docker


#安装mysql
mkdir /var/miconvert/mysql
cd /var/miconvert/mysql
docker pull mysql
#启动容器
docker run -id \
-p 3306:3306 \
--name=miconvert_mysql \
-v $PWD/conf:/etc/mysql/conf.d \
-v $PWD/logs:/logs \
-v $PWD/data:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=miconvert*.* \
mysql:latest
cd /var/miconvert

#安装nginx
mkdir /var/miconvert/nginx
cd /var/miconvert/nginx
docker pull nginx
#拷贝配置文件
cp $shell_path/nginx.conf /var/miconvert/nginx/nginx.conf
#启动nginx容器
docker run -id --name=miconvert_nginx \
-p 80:80 \
-v $PWD/nginx.conf:/etc/nginx/nginx.conf \
-v $PWD/conf:/etc/nginx/conf.d \
-v $PWD/logs:/var/log/nginx \
-v $PWD/html:/usr/share/nginx/html \
nginx:latest

#创建go目录
mkdir /var/miconvert/go
