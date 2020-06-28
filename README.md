# webdavGo
this is a webdav service implemented in docker using golang/这是一个安装在docker中用go语言编写的webdav服务

## preconditions
- already installed docker and docker-compose
## usage
### configuration 
edit your local configuration **docker-compose.yml**
```yaml
version: '3.7'
services:
  webdav:
    build: ./
    ports:
      - webdav access port:8080
    restart: unless-stopped
    volumes:
      - your_directory:/webdav_files #
    environment:
      UPLOAD_RATE_LIMIT: 1024 # upload speed limit:bit, example:1024=1kb/s

```
### install & run
```shell
git clone https://github.com/earth-hotfix/webdavGo.git
cd webdavGo
docker-compose up  # -d
```
## feature
- ...