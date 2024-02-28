# kis-flow

基于Golang的流式计算框架, 为保持简单的流动，强调在进行各种活动或工作时保持简洁、清晰、流畅的过程。

## kis-flow整体架构

### 1.1 为什么需要kis-flow

一些大型toB企业级的项目，需要大量的业务数据，多数的数据需要流式实时计算的能力，但是很多公司还不足以承担一个数仓类似，Flink + Hadoop/HBase 等等。但是业务数据的实时计算需求依然存在，所以大多数的企业依然会让业务工程师来消化这些业务数据计算的工作。

而这样只能直接查询业务数据库，这样会间接影响的业务能力，或定时任务/脚本来做定时计算，这些都不是好的办法。本人亲身经历过一个大规模的系统，多达上千个需要计算的业务数据字段，而早期因为没有规划好，最后导致存在1000+的脚本在定时跑，最后导致了脚本之间对数据的影响，数据始终无法准确，导致业务数据经常性的报数据问题错误。

KisFlow就是为了解决当企业不具备数仓平台的计算能力，又依然存在大量数据实时计算的场景，让业务工程师可以投入到数据流式计算的业务中来，并且可以复用常用和通用的计算逻辑。

### 1.2 kis-flow要支持的能力

#### 流式计算

1.分布式批量消费能力（基于上游ODS消费配置：如Binlog、Kafka等）;

2.Stateful Function能力，基于有状态的流式计算节点拼接，流式计算横纵向扩展;

3.数据流监控及修复能力，消费服务监控;

4.多流拼接及第三方中间件存储插件化;

#### 分布式任务调度

5.分布式定时任务调度、日志监控、任务调度状态;

6.可视化调度平台。

### 1.3 kis-flow系统定位

KisFlow为业务上游计算层，上层接数仓/其他业务方ODS层、下游接本业务存储数据中心。

## kis-flow 系统结构

common/ 存放我们一些公用的基础常量和一些枚举参数，还有一些工具类的方法；

config/ 存放flow、function、connector等策略配置信息模块；

conn/ 存放KisConnector的核心代码；

flow/ 存放kis-flow的核心代码；

function/ 存放kis-function的核心代码；

kis/ 存放所有模块的抽象层；

log/ 存放日志代码；

test/ 存放测试代码；

