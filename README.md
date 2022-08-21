## What is Farseer.Go?

[English](https://github.com/farseer-go/fs) | [中文](https://github.com/farseer-go/fs/blob/main/README.zh-cn.md)

A standard set of frameworks developed for the **golang** platform.

We have selected the most popular components for you and provide the use of these components in a modular way.

The framework perfectly supports **DDD domain-driven** technical implementations such as `repository`, `application-layer transactions`, `domain events`, `application-layer`, `dynamic WebAPI`.

We hope that in the daily development, only need to directly rely on this set of frameworks can cope with the common technical components

It has a [.net core](https://github.com/FarseerNet/Farseer.Net/tree/dev/Doc) Mature version, has been used for more than 10 years, very good

> *The current version of golang is `Alpha`, which means the features are not perfect yet, we are working on adding new features to the framework gradually. You can follow us first*.

## What are the features?

**Elegant**

We use `IOC` technology throughout the framework and your business systems.

**Simple**

We use `AOP` technology so that you don't have to write additional non-business functional code such as transactions, caching, exception catching, logging, linking Track

**Lightweight**

The framework makes extensive use of `collection pooling` technology to make your application take up less memory.

**Tracking**

If you use Orm, Redis, Http, Grpc, Elasticsearch, MQ (Rabbit, RedisStream, Rocker, local Queue), EventBus, Task, FSS, etc. that we provide, you don't need to do anything, the system will implicitly implement link tracking for you and provide API request logs, slow queries (all of the previously mentioned will be logged).

[FOPS](https://github.com/FarseerNet/FOPS) Project (automatic build, link trace console, K8S cluster log collection) supports code non-intrusive full link real-time monitoring.

## What packages are available?

| Component                                                    | Description                                                          |
|--------------------------------------------------------------|----------------------------------------------------------------------|
| [cache](https://github.com/farseer-go/cache)                 | Multi-level cache                                                    |
| [cacheMemory](https://github.com/farseer-go/cacheMemory)     | Memory Cache                                                         |
| [collections](https://github.com/farseer-go/collections)     | Support for List collections and linq syntax                         |
| [data](https://github.com/farseer-go/data)                   | Database ORM                                                         |
| [elasticSearch](https://github.com/farseer-go/elasticSearch) | elasticSearch client                                                 |
| [eventBus](https://github.com/farseer-go/eventBus)           | Publish subscription for events                                      |
| [fs](https://github.com/farseer-go/fs)                       | Farseer Basic                                                        |
| [fss](https://github.com/farseer-go/fss)                     | fss client                                                           |
| [linq](https://github.com/farseer-go/linq)                   | Support linq methods                                                 |
| [mapper](https://github.com/farseer-go/mapper)               | Conversions between objects, such as DO to DTO                       |
| [queue](https://github.com/farseer-go/queue)                 | Local queue, multiple writes, bulk consumption, multiple subscribers |
| [rabbit](https://github.com/farseer-go/rabbit)               | rabbit client                                                        |
| [redis](https://github.com/farseer-go/redis)                 | redis client                                                         |
| [tasks](https://github.com/farseer-go/tasks)                 | Local job                                                            |
| [utils](https://github.com/farseer-go/utils)                 | General Tools                                                        |

## What are the functions?
* [fs（框架初始化）](#how-to-start)
    * Initialize （初始化框架）
* [modules](modules/)
    * StartModules （启动模块）
* [configure（配置读写）](configure/)
    * GetString （获取配置）
    * SetDefault （设置配置的默认值）
* [container（容器IOC）](container/)
    * func
        * Use （自定义注册）
          * Transient（临时模式（默认为单例模式））
          * Name（Ioc别名）
          * Register （注册到容器）
        * Register （单例且没有别名注册到容器）
        * Resolve （从容器中获取实例）
        * ResolveName （指定ioc别名从容器中获取实例）
* [exception（异常处理）](exception/)
    * struct
        * RefuseException
    * func
        * ThrowRefuseException （抛出RefuseException异常）
        * ThrowRefuseExceptionf （抛出RefuseException异常）
        * Catch（捕获异常）
            * .RefuseException（捕获RefuseException异常）
                * .ContinueRecover（是否继续让下一个捕获继续处理）
            * .String（捕获String异常）
            * .Any（捕获Any异常）
* [parse（类型转换）](parse/)
    * Convert （通用的类型转换）
    * IsInt （是否为int类型）
    * IsEqual（两个any值是否相等）
* [stopwatch](stopwatch/)
    * func
        * StartNew（创建计时器，并开始计时）
    * struct
        * Stopwatch
            * Restart（重置计时器）
            * Start（继续计时）
            * Stop（停止计时）
            * ElapsedMilliseconds（返回当前已计时的时间（毫秒））
* [core（通用类型）](core/)
    * struct
        * ApiResponseString （标准的API输出（默认string值））
        * ApiResponseInt （标准的API输出（默认int值））
        * ApiResponseLong （标准的API输出（默认int64值））
        * ApiResponse （标准的API输出（泛型））
            * .SetData （设置Data字段的值）
    * func
        * Success （接口调用成功后返回的Json）
        * Error （接口调用失时返回的Json）
        * Error403 （接口调用失时返回的Json）
* core/eumLogLevel
    * Enum （日志等级）
* [net](net/)
    * LocalIPv4s （获取本机IP地址）

## How to start？
StartupModule is the startup module you define, as detailed in：[modules](modules/)
```go
// The framework starts and performs module initialization
fs.Initialize[StartupModule]("FOPS")
```