# 基建业务流程

## 阅读材料

- DDD 实践手册 (番外篇: 事件风暴 - 实践)：https://xie.infoq.cn/article/e36e68f9c6516fa10cf546f32
- 微服务架构下的统一身份认证和授权：https://www.jianshu.com/p/2571f6a4e192
- 认证、授权、鉴权和权限控制：http://www.hyhblog.cn/2018/04/25/user_login_auth_terms/

## 基建工厂业务流程

简单分析一下一些业务的流程，先把工厂部分的考虑清楚，之后再考虑电力系统、宿舍装修系统、右侧房间

1. 初始化
2. 首页页面的产出、问题提醒
3. 基建详情
4. 获取产出物品
5. 批量获取产出物品
6. 升级建筑
7. 修改建筑产出物品
8. 修改建筑工作人员

以下“**系统**”是对其他模块的一个抽象，指代基建模块的调用方。正常来讲肯定要理清调用者是谁，但由于我这里只涉及一个基建，就偷点懒不管了。



### 初始化

1. 玩家通过了某一关，客户端将信息发送给系统
2. 系统发现玩家通过了某一关，调用接口，传递玩家ID，通知基建模块初始化
3. 基建模块收到消息，进行初始化，包括设定好的比如一层一级建筑等，将相关信息写入数据库/缓存。失败返回err，成功则不返回。

```protobuf
// api/arkinfra/v1/arkinfra.proto

service InfraService {
    rpc CreateInfra(User) returns (CreateInfraReply);
}

// 一般来讲应该是调用user服务的pb文件
message User {
    string user_id = 1;
}

message CreateInfraReply {}
```



***问题零：微服务间的调用，用户信息是通过显示传递还是用meta data带token？我认为是应该显示传递，但还是得找找有没有这部分的最佳实践，学习学习。***



### 首页产出、问题提醒

1. 玩家到达首页，客户端调用系统接口拉取各种信息，包括基建信息
2. 系统调用基建模块提供的接口，传递玩家ID，拉取基建中的产出与问题数量信息
3. 基建系统拉取玩家基建信息，进行产出与问题的计算，如果生产物品中存在需要消耗其他物品的物品，则调用物品服务接口；更新基建信息，将数量返回给系统

```protobuf
rpc GetInfo(User) returns (GetInfoReply);

message GetInfoResp {
    int32 product_amount = 1;
    int32 error_amount = 2;
}
```



***问题一：进入到详情页也会进行一次这个计算，有点浪费，考虑是否建立缓存？***



### 基建详情

1. 玩家进入基建页面，客户端调用系统接口拉取基建详细信息
2. 系统调用基建模块提供的接口，传递玩家ID，拉取详细信息
3. 基建系统拉取玩家基建信息，进行产出与问题的计算，如果生产物品中存在需要消耗其他物品的物品，则调用物品服务接口；更新基建信息，将详细信息返回给系统

这里需要返回每个房间的详细信息，详细信息包括：

1. 房间ID
2. 房间类型
3. 房间等级
4. 进驻角色ID
5. 进驻角色当前疲劳
6. 房间库存
7. 房间已使用库存（已生产出的物品数量）
8. 房间生产的物品
9. 房间正在生产物品的进度

图方便，先假设上述所有信息都传；角色技能信息等的数据可以从其他渠道获得，暂定不传。

```protobuf
rpc GetDetail(User) returns (GetDetailReply);

message GetDetailReply {
    repeated Room rooms = 1;
}

message Room {
    int32 room_id = 1;
    int32 room_type = 2;
    int32 room_level = 3;

    repeated string character_id = 4;
    repeated int32 character_ftg = 5;

    int32 storage = 6;
    int32 storage_used = 7;

    int32 item_id = 8;
    int32 item_progress = 9;
}
```



### 获取产出物品

1. 玩家在基建页面点击“收取”，客户端发送收取消息给系统
2. 系统调用基建系统接口，传递玩家ID、房间ID
3. 基建系统处理，返回物品ID，物品数量

### 批量获取产出物品

1. 玩家在基建点击批量获取，批量获取某一种类的产出，客户端发送批量获取消息给系统
2. 系统调用基建服务接口，传递玩家ID，房间类型
3. 基建系统处理，返回物品ID，物品数量

```protobuf
// 获取物品
rpc GetProduction(GetProductionReq) returns (GetProductionReply);

message GetProductionReq {
    User user = 1;
    repeated int32 room_id = 2;
}

message GetProductionReply {
    repeated Production productions = 1;
}

message Production {
    int32 item_id = 1;
    int32 item_amount = 2;
}
```



### 升级某个建筑

1. 玩家点击升级，客户端将相关信息给到系统
2. 系统首先调用相关接口核对升级合不合法，之后调用基建服务接口，传递玩家ID，房间ID
3. 基建系统处理，升级该房间，并把房间内容清空，失败返回err，成功返回空消息

```protobuf
// 升级建筑
rpc UpdateInfra(UpdateInfraReq) returns (UpdateInfraReply);

message UpdateInfraReq {
    User user = 1;
    int32 room_id = 2;
}

message UpdateInfraReply {}
```



### 修改建筑产出物品

1. 玩家点击修改物品，选择物品、数量，客户端将相关信息给到系统
2. 如果是勋章等需要消耗其他物品的物品，系统会先调用相关接口，检验请求是否合法；之后系统调用基建服务接口，传递玩家ID、房间ID、物品ID、物品数量
3. 基建服务收到消息，首先判断该房间类型、房间等级能否生产这个物品；如果可以，先判断当前房间有没有已经生产完的物品，有则等会返回；之后基建服务修改生产物品的ID、最后操作时间，根据角色buff计算每个物品需要多长时间，根据角色buff重新计算角色每小时消耗疲劳值；返回生产完的物品ID、物品数量

***问题三：这里要不要返回当前每个物品需要多长时间，当前每个角色消耗的疲劳值***

```protobuf
// 更改生产物品
rpc ChangeProduction(ChangeProductionReq) returns (ChangeProductionReply);

message ChangeProductionReq {
    User user = 1;
    int32 room_id = 2;
    int32 item_id = 3;
    int32 item_amount = 4;
}

message ChangeProductionReply {
    // 修改物品前生产完毕的物品和数量
    int32 item_producted_id = 1;
    int32 item_producted_amount = 2;
}
```



### 修改建筑工作人员

1. 玩家点击修改工作人员，选择工作人员，客户端将相关信息给到系统
2. 系统进行校验人物是否存在等，之后调用基建接口，传入玩家ID、房间ID、人物ID数组
3. 基建收到消息，根据角色buff计算每个物品需要多长时间，根据角色buff重新计算物品生产时间，以及角色每小时消耗疲劳值；失败返回err，成功则不返回。

```protobuf
// 更改工作人员
rpc ChangeWorker(ChangeWorkerReq) returns (ChangeWorkerReply);

message ChangeWorkerReq {
    User user = 1;
    int32 room_id = 2;
    repeated int32 character_id = 3;
}

message ChangeWorkerReply {}
```

