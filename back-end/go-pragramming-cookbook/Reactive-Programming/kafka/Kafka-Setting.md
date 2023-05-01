# 카프카 서버 설치
* https://downloads.apache.org/kafka/3.4.0/kafka_2.13-3.4.0.tgz 설치
  * 압축 해제 필요
  * 압축 해제하여 /Users/yjchoi/kafka/ 에 복사
    * `/Users/yjchoi/kafka/kafka_2.13-3.4.0`
* zookeeper & kafka 실행
  * `cd /Users/yjchoi/kafka/kafka_2.13-3.4.0`
  * `bin/zookeeper-server-start.sh config/zookeeper.properties`
  * `bin/kafka-server-start.sh config/server.properties`