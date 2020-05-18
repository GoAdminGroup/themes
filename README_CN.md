# GoAdmin 官方主题

- [adminlte](https://github.com/GoAdminGroup/themes/tree/master/adminlte)
- [sword](https://github.com/GoAdminGroup/themes/tree/master/sword)

## 如何使用

- 导入主题
- 在全局配置中设置

```go

package main

import (
	...
	_ "github.com/GoAdminGroup/themes/adminlte"
	...
)

func main()  {
	
	...
	
	cfg := config.Config{
    		...
    		
    		Theme: "adminlte",
    		
    		...
    	}
	
	...
 
}

```

## 如何修改，自定义

使用每个主题下面的 Makefile 命令