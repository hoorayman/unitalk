# unitalk

unitalk is a distributed chat system which can be used as chat rooms or state synchronization.
unitalk registers itself on zookeeper when it start up, use redis cluster to broadcast messages or state, and permanent messages or state to kafka and then save to db.

## architecture

![image](https://github.com/hoorayman/unitalk/blob/main/architecture.png)

