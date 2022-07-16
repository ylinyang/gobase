# 1. 默认登录没有密码
mongo 127.0.0.1/admin

# 2. 创建管理员
use admin
db.createUser({user:"admin",pwd:"admin",roles:[{role:"root",db:"admin"}]});
db.auth('admin','admin')


