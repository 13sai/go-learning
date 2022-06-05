# boltdb 

bolt 是一个简单的kv数据库，使用及其简单，目前github项目处于只读状态。

使用很简单，但需要注意以下几点：
- 只读View中，不能使用编辑、删除、新增等写操作，会产生错误
- 因为底层使用了读写锁，进行写操作，要尽可能快，更不要开启长事务，会造成阻塞，影响性能