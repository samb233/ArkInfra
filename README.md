# ArkInfra

某游戏的基建（种菜）系统的实现。

目前正在实现工厂（左侧9宫格）部分，其他部分设计中

实现思路大体上有两种：

1. 页游思维，存储、计算都在服务端
2. 服务端只做存储和记录，计算在客户端

原游戏毫无疑问是第2种。这种场景确实是2更加适合，服务端开销也远低于1，还更容易实现。

~~但是我又没有客户端，所以就在服务端进行所有逻辑的实现。~~

目前是把逻辑全放在服务端了，或许我应该把逻辑提取出来，做成另一个服务，形成“客户端”



