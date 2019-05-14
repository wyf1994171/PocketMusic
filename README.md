					后端设计文档
1.	http处理

	统一使用Gin来处理。
	
	使用RESTful格式，删除操作使用DELETE，更新操作使用PUT，查询操作使用GET，新建操作使用POST。
	
	url定义使用小写下划线格式，按照路由集群使用分隔符分隔。
2.	数据库操作

	各表统一要有Info结构：id、updater、updated_time、creater、created_at，但不绝对，可以根据需要增删
	
	删除操作使用软删除方式，即将status字段更改为1。(0表示未删除，1表示删除)
	
	查询操作需要过滤掉status为1的字段，查询量大于1000，进行分页，不绝对，根据具体需要更改。
	
3.	代码规范：

	统一使用驼峰命名规则，参数命名要规范。
	
	上传代码之前要用编译器自带的工具对文件进行format
	
4.	Git使用

	自己从master分支创建自己的分支，不可以直接修改master
	
	每次更改要先pull，更改后再push
	
	出现bug，先将master回滚，再进行更改
	
	merge之前review
	
5.	Review：

	Git merge/代码合并之前让至少一个其他后端同学看一遍代码，没问题之后再合并
	
	在群里at要求Review的同学
	


