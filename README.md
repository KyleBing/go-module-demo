Go Module 学习

这个就需要看官方的教程了，是英文的，我也懒得翻译了。但是教程写的真的特别好，不明白的，看完就明白了。
[https://golang.org/doc/tutorial/create-module](https://golang.org/doc/tutorial/create-module)
它里面的教程内容是：

1. 创建一个 `module`
2. 从另外一个 `module` 中调用刚才创建的 `module`
3. 返回异常和处理异常
4. 返回随机的问候语（例子中创建的是一个返回问候语的 `Module`）
5. 返回多人的问候语
6. 添加测试
7. 编译并安装应用

看完之后你会了解：

- `module` 如何创建，
- 如何从一个 `module` 中调用另一个 `module`，
- 如何安装 `module` 到系统路径中，可以直接从 `terminal` 中执行这个指令