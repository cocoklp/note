[安装harbor](#安装harbor)



# harbor

官方：https://goharbor.io/docs/2.1.0/install-config/

1. 下载离线安装包

2. 配置https证书

   docker 默认通过https访问harbor，所以要配置证书

   1st：CA私钥

   ```sh
   openssl genrsa -out ca.key 4096
   ```

   2nd：CA证书 // 替换yourdomain

   ```sh
   openssl req -x509 -new -nodes -sha512 -days 3650 \
    -subj "/C=CN/ST=Beijing/L=Beijing/O=example/OU=Personal/CN=yourdomain.com" \
    -key ca.key \
    -out ca.crt
   ```

   3rd：server 证书私钥 // 替换yourdomain

   ```sh
   openssl genrsa -out yourdomain.com.key 4096
   ```

   4th：CSR  // 替换yourdomain

   ```sh
   openssl req -sha512 -new \
       -subj "/C=CN/ST=Beijing/L=Beijing/O=example/OU=Personal/CN=yourdomain.com" \
       -key yourdomain.com.key \
       -out yourdomain.com.csr
   ```

   5th：X509 v3文件

   ```sh
   cat > v3.ext <<-EOF
   authorityKeyIdentifier=keyid,issuer
   basicConstraints=CA:FALSE
   keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
   extendedKeyUsage = serverAuth
   subjectAltName = @alt_names
   
   [alt_names]
   DNS.1=yourdomain.com
   DNS.2=yourdomain
   DNS.3=hostname
   EOF
   ```

   6th：产生证书//替换yourdomain

   ```sh
   openssl x509 -req -sha512 -days 3650 \
       -extfile v3.ext \
       -CA ca.crt -CAkey ca.key -CAcreateserial \
       -in yourdomain.com.csr \
       -out yourdomain.com.crt
   ```

   

3. 把证书提供给docker

   ```sh
   cp yourdomain.com.crt /data/cert/
   cp yourdomain.com.key /data/cert/
   openssl x509 -inform PEM -in yourdomain.com.crt -out yourdomain.com.cert
   cp yourdomain.com.cert /etc/docker/certs.d/yourdomain.com/
   cp yourdomain.com.key /etc/docker/certs.d/yourdomain.com/
   cp ca.crt /etc/docker/certs.d/yourdomain.com/
   如果不用443端口，则改为/etc/docker/certs.d/yourdomain.com:port/ 或 /etc/docker/certs.d/ip:port/
   最终结构为
   /etc/docker/certs.d/
       └── yourdomain.com:port
          ├── yourdomain.com.cert  <-- Server certificate signed by CA
          ├── yourdomain.com.key   <-- Server key signed by CA
          └── ca.crt               <-- Certificate authority that signed the registry certificate
          
   重启docker，systemctl restart docker
   ```

4. 修改harbor.yml

   hostname certificate private_key路径

5. 依次 ./prepare ./install 安装  // 重启的话，需要先docker-compose down

6. ./install.sh  --with-chartmuseum  开启chart功能

7. docker ps查看状态，异常的docker logs查看日志，报错的话，去/var/log/harbor/目录下查看

   

## 报错举例

nginx /etc/cert.d/server.key permission denied

需要修改证书权限