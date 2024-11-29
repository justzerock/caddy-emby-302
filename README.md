# Caddy Emby 302


> Caddy 反代 Emby，播放 302 直链


## Docker 方式安装
``` sh
# docker-compose的配置
services:
  emby302:
    image: justzerock/emby302:latest
    container_name: emby302
    volumes:
      # 这个目录是配置目录，Caddyfile 将被生成在这个目录下，可以针对请求做各种变化
      - ./config:/config
    environment:
      - MEDIA_SERVER=这里填 emby,jellfin 的服务器地址 ** 必填 **
      - MEDIA_TOKEN=这里填 emby,jellfin 的 API KEY ** 必填 **
      - SERVER_302=这里填302后端服务器 选填
      # 原始路径，即 emby 挂载的路径，
      - ORIGIN_PATH=/mnt/cloud/media 选填
      # 替换路径，即 后端302服务器 的路径，比如 ORIGIN_PATH 最后会被替换成 /media
      - REPLACE_PATH=/media 选填
      # 若部分资源在本地，部分资源在网盘，MATCH_REDIR_302 用来区分网盘路径
      - MATCH_REDIR_302=/mnt/cloud 选填
    ports:
      - "8090:8082"

```


``` sh
# docker run 的命令

docker run -d \
  --name emby302 \
  -v /emby302/config:/config \
  -e MEDIA_SERVER="这里填 emby,jellfin 的服务器地址" \
  -e MEDIA_TOKEN="这里填 emby,jellfin 的 API KEY" \
  -e SERVER_302="这里填302后端服务器 选填" \
  -e ORIGIN_PATH="/mnt/cloud/media" \
  -e REPLACE_PATH="/media" \
  -e MATCH_REDIR_302="/mnt/cloud" \
  -p 8090:8082 \
  justzerock/emby302:latest

```

## Caddy 模式
> 因为是一个 caddy 的插件，所以可以自己构建一个 caddy，然后使用

``` sh
# 主要先去https://github.com/caddyserver/xcaddy下载自己系统需要的 xcaddy，放在项目目录下 
# 然后构建 caddy，如果需要其他插件包，可以自己自行增加
sh ./build-caddy.sh
```


## 其他 env 说明
``` sh
# 这里是302到 115 连接在内存中占用空间大小，默认16M
export CACHE302_SIZE=16
# 这里是302到 115连接的时效性，目前是 1 秒
export CACHE302=1
# 这 2 个是 caddy 的缓存配置
export BADGER_CACHE=/config/badger/cache
export BADGER_CONFIG=/config/badger/config
``` 


> 最后 /config/caddyfile 里面有具体的说明，如果想缓存图片和字幕可以自行取消注释相关的配置