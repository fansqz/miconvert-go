# miconvert-go





一个web端的office文件格式的转换网站，极其简单的网站，用于学习



## 底层依赖与实现

1. 使用到的技术

   websocket，需要异步解析，并主动向客户端提供解析进度，websocket需要支持集群

2. go选用框架

   使用gin+gorm

3. 需要安装的依赖

   暂时只用libreoffice,xpdf?



## 如何安装

1. libreoffice

   - 安装

     ~~~shell
     #查看yum是否支持libreoffice
     yum search libreoffice
     #安装
     yum install -y libreoffice
     ~~~

   - 使用到的命令

     ~~~shell
     #window
     
     ~~~

     

2. pdf2doc安装

   - 安装

     ~~~shell
     #这是一个python的库，但是使用命令行进行调用
     pip install pdf2docx
     ~~~

   - 使用命令

     ~~~shell
     pdf2docx convert /path/to/pdf.pdf /path/to/docx.docx
     ~~~

     
