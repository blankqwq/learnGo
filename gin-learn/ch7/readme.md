## 模板

> define

由于使用通配符加载的默认为全部文件，所以导致没有目录名称索引，我们需要使用定义别名来进行使用

```
{{define "index.html"}} 定义别名
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
<h1>{{.title}}</h1>
</body>
</html>
{{end}}
```
