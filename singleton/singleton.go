package singleton

/**
设计模式：单例模式  ======= 保证一个类全局永远只能有一个对象，这个对象还能被系统的其他模块使用
*/

import (
	"sync"
	"sync/atomic"
)

type singleton struct{}

// 饿汉式单例，在编译期就会创建对象分配内存
// var instance *singleton = new(singleton)

// 懒汉式单例，在没有创建时才创建对象分配内存
var instance *singleton

// 原子操作
var atomicInt32 *int32

// 定义一个互斥锁
var lock sync.Mutex

// GetInstance 全局方法获取实例对象
func GetInstance() *singleton {
	if atomic.LoadInt32(atomicInt32) == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()
	if instance != nil {
		instance = new(singleton)
		atomic.StoreInt32(atomicInt32, 1)
	}
	return instance
}

// ======== go sync库提供的Once实现方式
var once sync.Once

func GetInstance2() *singleton {
	once.Do(func() {
		if instance != nil {
			instance = new(singleton)
		}
	})
	return instance
}
