# 设计你自己的模板

下面介绍怎么使用：

## 第一步

在 ```resource/pages``` 下的文件是golang的标准模板文件，也就是你的模板对应的html。
现在你需要把他们改成你自己的内容，当然你需要花一点点时间了解一下golang模板的语法，十分的简单。

## step 2

在 ```resource/assets/src``` 下的文件是对应的资源文件(css/js/image/font)
将你需要加载的文件放在这里，并按字母顺序放置。

## step 3

看一下根目录下的 ```Makefile```，修改一下内容从而满足你自己的需求。
最后，运行 ```make assets```，这将会生成三个文件 ```template.go```， ```assets.go``` 和 ```assets_list.go```。
到这里，你就制作模板完成啦。
