# 迁移数据库
    bee generate migration 表名
    
    bee migrate -driver=mysql -conn="root:root@tcp(192.168.3.66:3306)/zy-o2o-server"
    
    
 #自动迁移数据库配置
    //postgres
    dsn := "host=" + dbhost + " port=" + dbport + " user=" + dbuser + " password=" + dbpassword + " dbname=" + dbname + " sslmode=" + dbssl
    //注册数据库驱动
    orm.RegisterDriver("postgres", orm.DRPostgres)
    //注册数据库 ORM 必须注册一个别名为 default 的数据库作为默认使用
    orm.RegisterDataBase("default", "postgres", dsn)
    // 注册模型
    orm.RegisterModel(new(User), new(Profile))
    //添加表名前缀注册模型
    orm.RegisterModelWithPrefix(dbprefix, new(User), new(Profile), new(Object))
    // 自动创建表 第二个参数为 true 时 强制重新建表，为 false 时 可以添加字段，不能删除字段
    orm.RunSyncdb("default", true, true)
