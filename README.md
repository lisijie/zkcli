# zkcli - 一个zookeeper客户端工具

## 使用

使用前需要先初始化zk配置：

```bash
$ zkcli init 192.168.1.1:2181,192.168.1.2:2181,192.168.1.3:2181
```

列出子节点信息：
```bash
$ zkcli ls /
Node                       Children  Data                                                Ctime                Mtime                Version  
--------------------------------------------------------------------------------------------------------------------------------------------
/cluster                   1                                                             2017-11-08 14:17:09  2017-11-08 14:17:09  0        
/controller                0         {"version":1,"brokerid":1001,"timestamp":"15104...  2017-11-12 13:14:47  2017-11-12 13:14:47  0        
/brokers                   3                                                             2017-11-08 14:17:09  2017-11-08 14:17:09  0        
/zookeeper                 1                                                             1970-01-01 08:00:00  1970-01-01 08:00:00  0        
/test                      0         a                                                   2017-11-09 15:34:44  2017-11-09 15:37:40  2        
/admin                     1                                                             2017-11-08 14:17:09  2017-11-08 14:17:09  0        
/isr_change_notification   0                                                             2017-11-08 14:17:09  2017-11-08 14:17:09  0        
/controller_epoch          0         47                                                  2017-11-08 14:17:10  2017-11-12 13:14:47  46       
/kafka-manager             4         172.18.0.7                                          2017-11-08 14:27:26  2017-11-08 14:27:26  0        
/consumers                 1                                                             2017-11-08 14:00:34  2017-11-08 14:00:34  0        
/latest_producer_id_block  0         {"version":1,"broker":1001,"block_start":"1000"...  2017-11-08 14:17:09  2017-11-08 14:18:18  2        
/config                    3                                                             2017-11-08 14:17:09  2017-11-08 14:17:09  0    
```

创建节点：
```bash
$ zkcli create /foo
Created /foo

STAT:
Czxid = 51539607592
Mzxid = 51539607592
Ctime = 1510496288475
Mtime = 1510496288475
Version = 0
Cversion = 0
Aversion = 0
EphemeralOwner = 0
DataLength = 0
NumChildren = 0
Pzxid = 51539607592
```

设置值：
```bash
$ zkcli set /foo hello
Set /foo

STAT:
Czxid = 51539607612
Mzxid = 51539607615
Ctime = 1510496524776
Mtime = 1510496530355
Version = 1
Cversion = 0
Aversion = 0
EphemeralOwner = 0
DataLength = 5
NumChildren = 0
Pzxid = 51539607612
```

查看节点信息：
```bash
$ zkcli get /foo
DATA:
hello

STAT:
Czxid = 51539607612
Mzxid = 51539607615
Ctime = 1510496524776
Mtime = 1510496530355
Version = 1
Cversion = 0
Aversion = 0
EphemeralOwner = 0
DataLength = 5
NumChildren = 0
Pzxid = 51539607612
```

以树形结构列出子节点信息：
```bash
$ zkcli tree /foo
/foo
├── bar
│   ├── a
│   ├── b
│   └── c
└── languages
    ├── php
    ├── java
    └── golang
```

删掉节点及其所有子节点：
```bash
$ zkcli rm -r /foo
Deleted /foo/bar/a
Deleted /foo/bar/b
Deleted /foo/bar/c
Deleted /foo/bar
Deleted /foo/languages/php
Deleted /foo/languages/java
Deleted /foo/languages/golang
Deleted /foo/languages
Deleted /foo
```
