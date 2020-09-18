# Hall_Download
Downloads 268 different writings by Manly P Hall.

All content is free to download at https://manlyphall.info/journals-index-opt.htm

## Docker

BUILD

`docker build -t hall_downloader:latest .`

RUN

```
docker run -d \
--name hall_downloader \
--mount type=bind,source="$(pwd)"/journals,target=/go/src/app \
hall_downloader:latest
```

