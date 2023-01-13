# hertz-common
some common component of hertz

## use
### 1. add ssh of github.com and set local ssh for git

### 2. add go env
``` shell
# 配置私有项目地址
go env -w GOPRIVATE="github.com/chenke1115/*"
```

### 3. add git config
```shell
git config --global url."git@github.com:chenke1115/".insteadof "https://github.com/chenke1115/"
```

### 4. go get 
```shell
go get -u -v github.com/chenke1115/hertz-common@latest
```