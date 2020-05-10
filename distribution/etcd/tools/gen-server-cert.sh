#!/bin/bash
#生成三个文件
#server-key.pem
#server.csr
#server.pem 
echo '{"CN":"client","hosts":["35.234.29.146","127.0.0.1"],"key":{"algo":"rsa","size":2048 }}'\
 | cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json\
 -profile=server -hostname="127.0.0.1,35.234.29.146,server" - | cfssljson -bare server 