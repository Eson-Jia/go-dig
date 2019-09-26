
# 开发日志

## gin

### 模型绑定和校验

为了将请求体绑定一个类型，我们使用类型绑定。现在支持 JSON,XML,YAML 和标准 form(foo=bar&boo=baz)

gin使用[go-playground/validator.v8](https://github.com/go-playground/validator)来进行合法性校验。查看`tags`使用的完整文档[点击](http://godoc.org/gopkg.in/go-playground/validator.v8#hdr-Baked_In_Validators_and_Tags)

注意你需要为所有你想要绑定的字段设置相应的绑定标签。例如，当绑定 JSON 的时候需要将标签设置为`json:"fieldname"`.

gin 提供了两种绑定的方法集：

- Type - 必须绑定

- - 方法- Bind,BindJson,BindXML,BindQuery,BindYAML,BindHeader
- - 行为- 这些方法底层使用`MustBindWith`函数。如果出现绑定错误，`c.AbortWithError(400,err).SetType(ErrorTypeBind)`会将请求中断。它会将状态码设置为 400 并且`Content-Type`头设置为`text/plain;charset=utf-8`.注意如果在这之后你尝试设置回复状态码，就会导致警告`[GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422.`如果你想要比这更多的控制自由度，考虑使用`ShouldBind`对应的函数。

- Type - 应该绑定

- - 方法 - ShouldBind,ShouldBindJSON,ShouldBindXML,ShoulBindQuery,ShouldBindYAML,ShouldBindHeader
- - 行为 - 这些方法底层使用`ShouldBindWith`方法。如果出现绑定错误，这个`error`会被返回，开发者自己决定如何处理这个`error`。
当你使用`Bind-method(Bind,MustBind)`的时候，gin会尝试根据`Content-Type`推断绑定的类型，如果你确定你绑定的类型你可以使用`MustBindWith`或者`BindWith`.

### http 跨域问题 CORS(cross-origin resource sharing)

首先要先区分简单请求和非简单请求的区别

简单请求有两个条件：

1. 请求方法必须是其中之一:

- HEAD
- GET
- POST

2. http的头消息不超过一下几种字段：

- Accept
- Accept-Language
- Content-Language
- Last-Event-ID
- Content-Type:只限于三个值 application/x-www-form-urlencoded,multipart/form-data,text/plain

凡是不能同时满足上面条件的请求都是非简单请求。
浏览器会区分对待这两种请求：


参考阮一峰[博客](http://www.ruanyifeng.com/blog/2016/04/cors.html)

## test

**go test**有两种模式

- 本地路径模式：本模式没有参数`go test`或者`go test -v`只会测试当前路径下的符合规则的测试文件

- 包列表模式：本模式会测试参数中列出的所有包 例如：`go test math strings`会测试 math 和 strings 这两包。`go test ./...`会测试当前路径下以及自路径下的所有测试包

### test in CI

在 CI 中运行`go test`于在开发环境别无二致，但是如果测试中需要连接其他服务（如: mysql,redis）情况就不一样了，我们可以通过在`job`中配置`services`提供`mysql`等服务。
如下：
测试需连接`mongo`数据库服务，通过`services`我们开启一个`mongo`的镜像，该`mongo`镜像与`test-job`所运行的镜像使用docker的[link](https://docs.docker.com/network/links/)连在一起，`mongo`容器可以通过主机名`mongo`进行访问。

```yml
test-job:
  extends: .job
  image: dockerhub.bmi:5000/golang:1.13rc1-buster
  stage: test
  services:
  - mongo:3.4-xenial # 此镜像可以通过主机名 mongo 进行访问
  variables:
    HLS_MONGODB_URI: mongodb://mongo # 通过主机名 mongo 进行访问
  script:
    - go test ./...
    - echo 更多测试

```

## viper config

### environment

- AutomaticEnv()
- BindEnv(string...) : error
- SetEnvPrefix(string)
- SetEnvKeyReplacer(string...) *strings.Replacer
- AllowEmptyEnv(bool)

当使用环境变量的时候，你需要意识到 `viper`对环境变量是区别大小写的。

`viper`提供了一个机制来保证环境变量的唯一性，通过调用`SetEnvPrefix`函数，你可告诉`viper`在读取环境变量的时候使用一个前缀。`BindEnv`和`AutomaticEnv`都会使用这个前缀。

`BindEnv`需要一到两个参数，第一个参数是键的名字，第二个参数是环境变量的名字，环境变量的名字是区分大小写的。如果未提供环境变量的名字，`viper`会自动假设环境变量的名字按照以下格式：`前缀（如果设置了的话）+"_"键名字全部大写`。但是如果你在第二个参数中明确指定了环境变量的名字，`viper`则不会自动添加前缀。例如：如果第二个参数是`"id"`,`viper`会查找环境变量`"ID"`。

一个重点是当使用环境变量的时候，每次访问该键都会重新读取环境变量。当调用`BindEnv`的时候`viper`不会固定存储该值。

`AutomaticEnv`是一个得力助手，尤其是在与`SetEnvPrefix`一起使用的时候。调用`AutomaticEnv`函数后，每次`viper.Get`被调用的时候`viper`都会去检查一个环境变量，规则如下：环境变量是键的全大写，如果`EnvPrefix`设置了前缀的话就加上该前缀。

`SetEnvKeyReplacer`为你提供一个`strings.Replacer`对象，从而使你可以在一定程度上重写环境键。例如，在调用`Get`的时候你想使用`-`或其他字符，但是又想你的环境变量使用`_`分隔符的时候，这时候就特别有用。

默认情况下空的环境变量会被认为是该环境变量未设置，然后会回退使用下一个配置资源，为了把空环境变量当做已设置，需要调用`AllowEmptyEnv`函数。

## golang 语法

### slice

python 对于数组有翻转的语法糖

```python
# python 翻转数据语法
a = [1, 2, 3, 4, 5]
b = a[::-1]
print(b)
# output: [5, 4, 3, 2, 1]
```

go 不支持这样的语法糖

```golang
func TestSliceReverse(t *testing.T) {

    s := []int{1, 2, 3, 4}
    b = s[:-1] // 报错 index must not be negative
    fmt.Println(s)
}
```

## ffmpeg

### HLS

ffmpeg 将文件转换成多码率适配的`HLS`格式:

```bash
ffmpeg -y -i sintel_trailer-1080p.mp4 \
  -preset slow -g 48 -sc_threshold 0 \
  -map 0:0 -map 0:1 -map 0:0 -map 0:1 \
  -s:v:0 640x360 -c:v:0 libx264 -b:v:0  2000k \
  -s:v:1 960x540 -c:v:1 libx264 -b:v:1 365k  \
  -c:a copy \
  -var_stream_map "v:0,a:0 v:1,a:1" \
  -master_pl_name master.m3u8 \
  -f hls -hls_time 6 -hls_list_size 0 \
  -hls_segment_filename "v%v/fileSequence%d.ts" \
  v%v/prog_index.m3u8
```

参考文档：

- <https://hlsbook.net/creating-a-master-playlist-with-ffmpeg/>

## JOSN

- `json`对`struct`序列化和反序列化时，会忽略不可导出（首字母小写）的成员
- `json`对`struct`进行反序列化的时候会尽量匹配大小写，但是如果大小写不匹配也能解析
- `json`将一个`struct`序列化成字符串之后如果将结果反序列化时传入一个`map`

### interface conversion: interface {} is float64, not int

json 中的数据反序列化后是`float64`
