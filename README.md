# Soul Web Skeleton

基于Gin + logrus + viper + Gorm的web脚手架, 自用练手项目

## Done:
- [x] 配置文件读取
- [x] 命令行、环境变量指定配置(flag仅-c/-e/-m, env仅RUN_ENV)
- [x] Gin 集成、路由管理、中间件、优雅停止
- [x] logrus 日志库集成
- [x] lumberjack 日志轮转
- [x] Gorm 集成Mysql、Sqlite
- [x] Gorm 日志适配 logrus
- [x] Swagger 集成
- [x] 自定义错误校验
- [x] 数据库迁移
- [x] JWT token 认证
- [x] 跨域中间件
- [x] `Request-Id` 中间件, 为每个请求添加header
- [x] 容器化

## Todo:
- [ ] 分页
- [ ] 规范错误响应
- [ ] Gorm 集成 Postgresql
- [ ] 集成 redis
- [ ] 权限管理
