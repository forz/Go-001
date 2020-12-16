# 作业思路

整体目录:

1. 采用微服务思想,homework为其中一个子服务.
2. pb与swagger文件放在apis目录,各子服务可调用.
3. 子服务之间并列,都有自己的pkg和internal目录,好处如下:
   1. 某一个微服务开发者就不用看到和自己无关的其他domain的代码.
   2. domain升级为微服务时方便迁移,不用挑代码.

子服务内分层:

1. 微服务框架使用gokit+代码生成器,实现dto的编解码.(类似DDD的assembler和Dto层)
   1. transport层:dto的结构与内存变量的转换.
   2. service层:接口的业务逻辑,编排组装领域服务和外部微服务(类似DDD的应用服务层)
   3. endpoint层:transport层到service层的胶水层.
   4. server与client层:transport层到grpc方法的胶水层.
2. pkg与internal
   1. pkg目录放client,server,endpoint,transport层
      1. 这4层除了编解码和中间件的逻辑,都是代码生成器直接生成,基本没有敏感内容.
      2. 其他使用相同架构的团队,不需要再做一次dto的编解码.
   2. internal目录放service层,是业务逻辑,不允许import.
3. internal内部分层
   * service层:应用服务层
   * domain 领域层
      * foo 聚合层
        * service 领域服务层,编排entity层方法,供应用服务层使用
        * entity 领域实体层,充血模型,调用repository.
        * repository PO层,只存放数据库结构体,剩下的交给orm
4. wire的使用时机
   1. 在domain内部使用,这里是业务的核心层,与业务逻辑解耦,适合用gomock进行测试.
   2. 在main函数使用