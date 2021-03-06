# 五子棋
> 在线五子棋

## 前言
该游戏是基于Websocket开发的，采用Go + Bootstrap5   
你可以了解到:
- 五子棋盘的实现 --> internal/chessboard
- 如何优雅地处理通信消息 --> internal/httpserver/httpserver.go
- 简单的断线重连

### 特性
- [x] 创建房间、加入房间、离开房间
- [x] 随机先手
- [x] 游戏重开(仅房主)
- [x] 自动转移房主(对方退出房间)
- [x] 断线重连
- [x] 多人观战
- [x] 悔棋

### 开始
```
go run main.go
```

### nginx 配置
```
server {
       listen 80;
       listen [::]:80;

       server_name gk.sadeye.cn;

       index index.html;

       location / {
            proxy_pass http://127.0.0.1:8000;
       }
       location /ws {
           proxy_pass http://127.0.0.1:8000/ws;
           proxy_http_version 1.1;
           proxy_set_header Upgrade $http_upgrade;
           proxy_set_header Connection "upgrade";
       }
}
```
### 参考资料

- UI : https://github.com/ccforward/cc/issues/51
