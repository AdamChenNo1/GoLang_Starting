#!/bin/bash
#输入peer数量，为每个peer生成三个文件
#member1-key.pem
#member1.csr
#member1.pem 
num=${1:-1}
for((i=1;i<=$num;i++));do
    echo '{"CN":"client","hosts":["127.0.0.1"],"key":{"algo":"rsa","size":2048 }}'\
    | cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json\
    -profile=peer -hostname="127.0.0.1,server,member$i" - | cfssljson -bare member$i; 
done;
