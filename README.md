# my_rename

## 一、项目介绍

由于`linux`自带的`rename`不够强大，使用`Golang`编写一个小工具。

将执行目录中的文件，按照文件修改日期进行批量重命名。

将前缀为`{{ 前缀 }}`扩展名为`{{ 扩展名 }}`的文件，批量命名为

```yaml
{{ 前缀 }}yyyyMMdd-hhmm.{{ 扩展名 }}
```

## 二、使用方法

### 1. 编写配置文件

编写配置文件 `conf.yml` 当然可以自己定义`yml`文件

```yaml
# vim /data/sh/conf.yml or {{ confFile }}
path: {{ 执行目录 }}
fileExtensions: {{ .扩展名 }} #记得前面要加.
prefixName: {{ 前缀 }}
```

### 2. 执行

```shell
chmod +x rename
./rename

# 如果是自定义的conf文件

./rename -path {{ confFile }}
```
