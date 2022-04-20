# GenPerfTest

批量生成pjsua和sipp测试配置，用于呼叫中心的压力测试

- 构建

  make

  

- 使用
  - Usage of genpjtest.exe:
      -group string
            agent group
      -idip string
            sip id ip
      -password string
            sip password，各个账号需统一密码
      -port int
            起始的本地端口，默认10000,每新增一个配置会自动加1 (default 10000)
      -realm string
            sip realm (default "*")
      -reg string
            sip registrar
      -set string
            设置项，分别是 主叫开始项(忽略，仅为兼容pjtest而设),被叫开始项,数量，设置项之间用逗号隔
    开，多个设置用竖线隔开。例如：6101,6001,99|6300,6200,100
  - Usage of gensipptest.exe:
      -group string
            agent group
      -head string
            文件头部设置，可以为SEQUENTIAL或RANDOM (default "SEQUENTIAL")
      -set string
            设置项，分别是 主叫开始项,被叫开始项,数量，设置项之间用逗号隔开，多个设置用竖线隔开。例
    如："6101,6001,99|6300,6200,100"

- 例子
  - 详见 Makefile

- 参考
  - pjsua
  - sipp
