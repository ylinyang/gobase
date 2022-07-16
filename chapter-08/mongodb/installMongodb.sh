#!/bin/bash
echo  "1. 解压移动配置path"
mkdir -p /tmp/install/
mv $1 /tmp/install && cd /tmp/install && tar  zxvf $1 && mv ./* /usr/local/mongodb
echo "export MONGODB=/usr/local/mongodb" >> /etc/profile
echo "export PATH=\$MONGODB/bin:\$PATH" >> /etc/profile
source /etc/profile

echo "2. 创建配置文件"
mkdir -p /data1/mongodb/data
echo "202d604af9d34e0bb77d382c55ef6839" >/data1/mongodb/auth.key
chmod 700 /data1/mongodb/auth.key
cat > /data1/mongodb/mongodb.yaml << EOF
systemLog:
  destination: file
  path: /data1/mongodb/mongod.log
  logAppend: true
processManagement:
  pidFilePath: /data1/mongodb/mongod.pid
  fork: true
net:
  unixDomainSocket:
    enabled: true
    pathPrefix: /data1/mongodb/data
  bindIp: 0.0.0.0
  port: 27017
  maxIncomingConnections: 10000
storage:
  dbPath: /data1/mongodb/data
  engine: wiredTiger
  wiredTiger:
    engineConfig:
      directoryForIndexes: true
    collectionConfig:
      blockCompressor: zlib
    indexConfig:
      prefixCompression: true
  journal:
    enabled: true
  directoryPerDB: true
operationProfiling:
  slowOpThresholdMs: 300
  mode: slowOp
security:
  clusterAuthMode: keyFile
  authorization: enabled
  keyFile: /data1/mongodb/auth.key
replication:
  oplogSizeMB: 102400
# replSetName: srcb_d5ee5a48
EOF

echo "3. 启动"

mongod -f /data1/mongodb/mongodb.yaml