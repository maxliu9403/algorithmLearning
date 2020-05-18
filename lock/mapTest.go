package main

//type safeMap struct {
//	// 封装读写锁
//	sync.RWMutex
//	Map map[int]int
//}
//
//func main() {
//	safeMap := newSafeMap()
//	for i := 0; i < 10; i++ {
//		go safeMap.writeMap(i, i)
//		go safeMap.readMap(i)
//	}
//
//	time.Sleep(time.Second)
//	safeMap.outPut(*safeMap)
//}
//
//func newSafeMap() *safeMap {
//	sm := new(safeMap)
//	sm.Map = make(map[int]int)
//	return sm
//
//}
//
//func (sm *safeMap) readMap(key int) int {
//	sm.RLock()
//	value := sm.Map[key]
//	sm.RUnlock()
//	return value
//}
//
//func (sm *safeMap) writeMap(key int, value int) {
//	sm.Lock()
//	sm.Map[key] = value
//	sm.Unlock()
//}
//
//func (sm *safeMap) outPut(safeMap safeMap) {
//	for k, v := range safeMap.Map {
//		fmt.Printf("%d, %d\n", k, v)
//	}
//}

func main() {
	Map := make(map[int]int)

	for i := 0; i < 10; i++ {
		go writeMaps(Map, i, i)
		go readMaps(Map, i)
	}
}

func readMaps(Map map[int]int, key int) int {
	return Map[key]
}

func writeMaps(Map map[int]int, key int, value int) {
	Map[key] = value
}

//Output：fatal error: concurrent map writes