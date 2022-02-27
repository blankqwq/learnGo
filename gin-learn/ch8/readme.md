## 模板

> define

由于使用通配符加载的默认为全部文件，所以导致没有目录名称索引，我们需要使用定义别名来进行使用

```
[GIN-debug] GET    /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
```

使用了默认静态的文件路由
