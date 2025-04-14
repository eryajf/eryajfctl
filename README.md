<div align="center">
<h1>eryajfctl</h1>

[![Auth](https://img.shields.io/badge/Auther--eryajf-ff69b4.svg?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIAAAAyCAYAAAAeP4ixAAAACXBIWXMAAAsTAAALEwEAmpwYAAADZElEQVR4nO2ZX2iPURjHP/7/aZN/E21DaZvtwoVYyQUuGXLB/LtkLRcUhSJMSVwg3KCUJPJvLmRZtMQFLvwZhUJk/saGLWaYV6eet06n9/3tfd+9531/sW89td9z3vOc8z3nPOc8zzPoQQ+yAnlAA/AbcCzKR2ClTSL7LRNwNPkB5Noi0pggEQeYbovIh4SJLLBBojfwK2EiVTaIjEyYhANsskGkNAUi+2wQmZECkRM2iCxMgUi9DSKrxPhS7GOZjHXPhvGtYlwdMduYKWO9tmH8oBgvxj6KZawOoFfcxk+L8RzsI0fzkyFhO28EvqXg0FHlK1DtReRzFkzOCSmfvIhczIKJOSHlmheRgcAK4E4WTNDpQp4Ba4P4zxRgO3ArgSQqqDQDx4AKCWBDYwSwRDP4JWECG4CpQB9iwCTNsBsJl8srvxk4CtQBN4Enkre0GKQ7RaekCXgAXAcuAHuB1bLaZcAgrd+2uK78ai0PsfLSBshG70u9IDJmyUq2AYvlbD5N4Dg9l/EnArdFV9sdIlfEiLrNzFAlqPwE2kP2OaKNlyfH8Q9QFJVIixgepukqIuTdRRI3Be1TaczjnOgXRSXyTgwUarrBIVb4htbvcMA+nR7+UCdtc6MSOS8G9hj6+gjlnHEBd+WuMVaZ9FMyJiqRydrgB4B80VcFmFBtxALfFm3nlwPv48rfK7WoWDkc8jgpYq0eE/kOnAGG+4RBh3z6qTGOyzfIJeG2nQL6EQMmaEaTgtNdv+jKcFJwbI33XxEpAM76+IArrXIbFqVFpE0MuzeXF4nmEI9ecwZbhVo6GzsaxLhnniw74YQUVdDwwjppv2yzaPbIJ7HJdJz8RIX5JgYAL6V9vg0ifYEXMsAaj3Z3cpfkmPmhQAs5vHxgh+gbbdS0XMyTQdrkbdHhTiwTCdMHTCLl8hCqeGsalnFSO2JDNb3XxOZINqhkttFmfq8W4I1PbGcFuVr21qCloe7EarRvmzT9K01fYxAZJVmgA1yNKxwh4NF4q0Wr+UZs5JIxndskoVLoEskI1e/HRu6TCEpklR1ZebPcqk/YT9cu1UL190NgNCmhQGpfTjelLo2dMNEf2BkylXVF7eL6qAU3WyiVvCHIv7A7pBY2nizGWGC3FOr0XeoQZ96VIcbqAf8K/gLNGaTJ3vwbFgAAAABJRU5ErkJggg==)](https://github.com/eryajf)
[![Eryajf HitCount](https://views.whatilearened.today/views/github/eryajf/eryajf.svg)](https://github.com/eryajf)
[![GitHub license](https://img.shields.io/github/license/eryajf/eryajfctl)](https://github.com/eryajf/eryajfctl/blob/main/LICENSE)
[![GitHub Stars](https://img.shields.io/github/stars/eryajf/eryajfctl)](https://github.com/eryajf/eryajfctl/stargazers)
[![Eryajf Blog](https://img.shields.io/badge/%E5%8D%9A%E5%AE%A2-%E4%BA%8C%E4%B8%AB%E8%AE%B2%E6%A2%B5-d7b1bf?logo=Blogger)](https://wiki.eryajf.net)
[![Eryajf WeChat](https://img.shields.io/badge/%E5%85%AC%E4%BC%97%E5%8F%B7-%E8%BF%90%E7%BB%B4%E8%89%BA%E6%9C%AF-71f9fe?logo=WeChat)](https://y.gtimg.cn/music/photo_new/T053M000003iCCnF30PTi3.jpg)
[![Eryajf Awesome Stars](https://img.shields.io/badge/Awesome-MyStarList-c780fa?logo=Awesome-Lists)](https://github.com/eryajf/awesome-stars-eryajf#readme)

<p> 🌉 基于Cobra库快速开发类似kubectl一样的命令行工具框架 🌉</p>
<img src="https://cnb.cool/66666/resource/-/git/raw/main/img/hengtiao.gif" width="100%"  height="3">
</br>
</div>

运维也可以如此优雅！快用这个框架打造一个专属于你的工具箱吧！

通过这个框架，你可以快速上手，直接构建你想要的运维工具，而不必再考虑配置，框架设计等内容。

## 如何使用

先拷贝配置文件:

```sh
cp config.example.yml config.yml
```

执行如下指令，运行示例参数：

```sh
$ go run main.go ex getconfig -w "hello, This is eryajfctl"
通过配置文件获取到的用户名: eryajf
通过配置文件获取到的密码: 123456
通过命令行获取到的内容是: hello, This is eryajfctl
```

> 其中ex为一级参数，getconfig为二级参数，大多数场景下，分两个层级就够用了，你可以把一级参数当做归类，比如 jenkins, gitlab，二级参数当做功能参数，再往后的参数则是该二级参数所需要的运行时参数。

也可以编译成二进制，然后通过如下方式查看帮助信息：

```
# 编译
$ make build

#运行测试
$ ./eryajfctl ex getconfig -h
通过命令行获取配置信息

Usage:
  eryajfctl ex getconfig [flags]

Flags:
  -h, --help          help for getconfig
  -w, --word string   测试参数 (default "你好，这是测试")
```

## 开始开发

你可以直接参考ex参数的流程，开发新的参数，从而满足实际使用需求。

如果你的配置文件中有敏感数据，可以考虑结合go-bindata来使用，通过执行：

```
go-bindata -o=./public/bindata_config.go -pkg=public config.yml
```

然后更改 [public/config.go](https://github.com/eryajf/eryajfctl/blob/4cd30714062e5b65746bdb5f100f19bfe38ed52e/public/config.go#L28) 中的配置信息读取方式，接着进入开发即可。

如果后续本地的config.yml配置有更新，则再次执行上边的命令，将配置文件注入到 `bindata_config.go` 即可。

这样做的一个好处是，你的二进制放到服务器等地运行的时候，既不需要添加config.yml文件，也能防止文件中的敏感信息暴漏出去。

## 生成文档

执行如下命令，会自动生成该工具的文档，且生成子命令的文档。

```
$ ./eryajfctl --md-docs
```

在这里你可以看到所有的文档：[点我查看](./docs/eryajfctl.md)

## 感谢开源

此框架建立在如下几个优秀的开源项目之上：

- [gopkg.in/yaml.v3 v3.0.1](https://github.com/go-yaml/yaml)
- [github.com/spf13/cobra v1.9.1](https://github.com/spf13/cobra)
- [github.com/spf13/viper v1.20.1](https://github.com/spf13/viper)

## 其他参考

如果你想熟悉了解此框架的详细用法，还可以参考我的如下三篇文章：

- [使用go-bindata将文件编译进二进制](https://wiki.eryajf.net/pages/2bf6c3/)
- [利用cobra库快速开发类似kubectl一样的命令行工具](https://wiki.eryajf.net/pages/5c4163/)
- [近期关于cobra库的一些实践心得总结](https://wiki.eryajf.net/pages/7b8eff/)

## 项目源码

可选择你熟悉的平台浏览源码：

|   服务商   |                   地址                   |
| :------: | :------------------------------------------: |
|  `CNB`  | <https://cnb.cool/eryajf/eryajfctl>  |
| `GitHub` | <https://github.com/eryajf/eryajfctl> |
