chan:
    1. 协程安全的数据结构
    2. 默认不带缓存，要求发送端和接收端必须分属于两个协程，且同时读写。否则，阻塞
    3. 通过提供带缓存的channel，可以提供系统的性能
