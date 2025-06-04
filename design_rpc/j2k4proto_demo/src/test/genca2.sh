#!/bin/bash

# 指定生成证书的目录
dir=$(dirname "$0")/resources/x509
[ -d "$dir" ] && find "$dir" -type f -exec rm -rf {} \;
mkdir -p "$dir"
pushd "$dir" || exit

# 生成.key 私钥文件 和 csr 证书签名请求文件
openssl req -new -nodes -sha256 -newkey rsa:2048 -keyout ca.key -out ca.csr \
-subj "/C=CN/ST=Zhejiang/L=Hangzhou/O=Ghimi Technology/OU=Ghimi Cloud/CN=ghimi.top"

# 生成自签名 .crt 证书文件
openssl x509 -req -in ca.csr -key ca.key -out ca.crt -days 3650

# 生成服务器私钥文件 和 csr 证书请求文件(私钥签名文件)
openssl req -new -nodes -sha256 -newkey rsa:2048 -keyout server.key -out server.csr \
-subj "/C=CN/ST=Zhejiang/L=Hangzhou/O=Ghimi Technology/OU=Ghimi Blog/CN=blog.ghimi.top"

# 生成 server 证书,由 ca证书颁发
openssl x509 -req -in server.csr -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 -extensions SAN \
-extfile <(cat /etc/ssl/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:dns.ghimi.top,IP:127.0.0.1,IP:::1"))

# 将 crt 证书转换为 pkcs12 格式,生成 server.p12 文件,密码 123456
openssl pkcs12 -export -in server.crt -inkey server.key -CAfile ca.crt -out server.p12 -passout pass:123456

# 验证 PKCS12 文件的完整性
openssl pkcs12 -info -in server.p12 -noout -passin pass:123456

# 将 PKCS12 文件转换为 PEM 格式
openssl pkcs12 -in server.p12 -out server.pem -nodes -passin pass:123456

# 从 PEM 文件中提取私钥和证书
openssl rsa -in server.pem -out server.key
openssl x509 -in server.pem -out server.crt

# 使用 keytool 导入到 JKS 中
keytool -import -alias server -file server.crt -keystore server.jks -storepass 123456

# 删除已有的别名
# keytool -delete -alias ca -keystore server.jks -storepass 123456

# keytool -import -alias ca -file ca.crt -keystore server.jks -storepass 123456

# 将 ca 证书导入到 server.jks 中
keytool -importcert -keystore server.jks -file ca.crt -alias ca -storepass 123456 -noprompt

popd || exit