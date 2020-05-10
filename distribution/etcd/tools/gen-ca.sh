#!/bin/bash
#生成三个文件
#ca-key.pem
#ca.csr
#ca.pem
cfssl gencert -initca ca-csr.json | cfssljson -bare ca -