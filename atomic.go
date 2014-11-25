package atomic

import (
	"sync"
	old_atomic "sync/atomic"
	"unsafe"
)

var myMutex *sync.Mutex = &sync.Mutex{}

// BUG(rsc): On x86-32, the 64-bit functions use instructions unavailable before the Pentium MMX.
//
// On non-Linux ARM, the 64-bit functions use instructions unavailable before the ARMv6k core.
//
// On both ARM and x86-32, it is the caller's responsibility to arrange for 64-bit
// alignment of 64-bit words accessed atomically. The first word in a global
// variable or in an allocated struct or slice can be relied upon to be
// 64-bit aligned.

// SwapInt32 atomically stores new into *addr and returns the previous *addr value.
func SwapInt32(addr *int32, new int32) (old int32) {
	return old_atomic.SwapInt32(addr, new)
}

// SwapInt64 atomically stores new into *addr and returns the previous *addr value.
func SwapInt64(addr *int64, new int64) (old int64) {
	myMutex.Lock()
	defer myMutex.Unlock()

	savedValue := *addr
	*addr = new
	return savedValue
}

// SwapUint32 atomically stores new into *addr and returns the previous *addr value.
func SwapUint32(addr *uint32, new uint32) (old uint32) {
	return old_atomic.SwapUint32(addr, new)
}

// SwapUint64 atomically stores new into *addr and returns the previous *addr value.
func SwapUint64(addr *uint64, new uint64) (old uint64) {
	myMutex.Lock()
	defer myMutex.Unlock()

	savedValue := *addr
	*addr = new
	return savedValue
}

// SwapUintptr atomically stores new into *addr and returns the previous *addr value.
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr) {
	return old_atomic.SwapUintptr(addr, new)
}

// SwapPointer atomically stores new into *addr and returns the previous *addr value.
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer) {
	return old_atomic.SwapPointer(addr, new)
}

// CompareAndSwapInt32 executes the compare-and-swap operation for an int32 value.
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool) {
	return old_atomic.CompareAndSwapInt32(addr, old, new)
}

// CompareAndSwapInt64 executes the compare-and-swap operation for an int64 value.
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool) {

	return old_atomic.CompareAndSwapInt64(addr, old, new)
}

// CompareAndSwapUint32 executes the compare-and-swap operation for a uint32 value.
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool) {
	return old_atomic.CompareAndSwapUint32(addr, old, new)
}

// CompareAndSwapUint64 executes the compare-and-swap operation for a uint64 value.
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool) {
	return old_atomic.CompareAndSwapUint64(addr, old, new)
}

// CompareAndSwapUintptr executes the compare-and-swap operation for a uintptr value.
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool) {
	return old_atomic.CompareAndSwapUintptr(addr, old, new)
}

// CompareAndSwapPointer executes the compare-and-swap operation for a unsafe.Pointer value.
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool) {
	return old_atomic.CompareAndSwapPointer(addr, old, new)
}

// AddInt32 atomically adds delta to *addr and returns the new value.
func AddInt32(addr *int32, delta int32) (new int32) {
	return old_atomic.AddInt32(addr, delta)
}

// AddUint32 atomically adds delta to *addr and returns the new value.
// To subtract a signed positive constant value c from x, do AddUint32(&x, ^uint32(c-1)).
// In particular, to decrement x, do AddUint32(&x, ^uint32(0)).
func AddUint32(addr *uint32, delta uint32) (new uint32) {
	return old_atomic.AddUint32(addr, delta)
}

// AddInt64 atomically adds delta to *addr and returns the new value.
func AddInt64(addr *int64, delta int64) (new int64) {
	myMutex.Lock()
	defer myMutex.Unlock()

	*addr = *addr + delta
	return *addr
}

// AddUint64 atomically adds delta to *addr and returns the new value.
// To subtract a signed positive constant value c from x, do AddUint64(&x, ^uint64(c-1)).
// In particular, to decrement x, do AddUint64(&x, ^uint64(0)).
func AddUint64(addr *uint64, delta uint64) (new uint64) {
	myMutex.Lock()
	defer myMutex.Unlock()

	*addr = *addr + delta
	return *addr
}

// AddUintptr atomically adds delta to *addr and returns the new value.
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr) {
	return old_atomic.AddUintptr(addr, delta)
}

// LoadInt32 atomically loads *addr.
func LoadInt32(addr *int32) (val int32) {
	return old_atomic.LoadInt32(addr)
}

// LoadInt64 atomically loads *addr.
func LoadInt64(addr *int64) (val int64) {
	myMutex.Lock()
	defer myMutex.Unlock()

	return *addr
}

// LoadUint32 atomically loads *addr.
func LoadUint32(addr *uint32) (val uint32) {
	return old_atomic.LoadUint32(addr)
}

// LoadUint64 atomically loads *addr.
func LoadUint64(addr *uint64) (val uint64) {
	myMutex.Lock()
	defer myMutex.Unlock()

	return *addr
}

// LoadUintptr atomically loads *addr.
func LoadUintptr(addr *uintptr) (val uintptr) {
	return old_atomic.LoadUintptr(addr)
}

// LoadPointer atomically loads *addr.
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer) {
	myMutex.Lock()
	defer myMutex.Unlock()

	val = old_atomic.LoadPointer(addr)

	return val
}

// StoreInt32 atomically stores val into *addr.
func StoreInt32(addr *int32, val int32) {
	myMutex.Lock()
	defer myMutex.Unlock()

	old_atomic.StoreInt32(addr, val)
	return
}

// StoreInt64 atomically stores val into *addr.
func StoreInt64(addr *int64, val int64) {
	myMutex.Lock()
	defer myMutex.Unlock()

	*addr = val
	return
}

// StoreUint32 atomically stores val into *addr.
func StoreUint32(addr *uint32, val uint32) {
	myMutex.Lock()
	defer myMutex.Unlock()

	old_atomic.StoreUint32(addr, val)
	return
}

// StoreUint64 atomically stores val into *addr.
func StoreUint64(addr *uint64, val uint64) {
	myMutex.Lock()
	defer myMutex.Unlock()

	*addr = val
	return
}

// StoreUintptr atomically stores val into *addr.
func StoreUintptr(addr *uintptr, val uintptr) {
	old_atomic.StoreUintptr(addr, val)
	return
}

// StorePointer atomically stores val into *addr.
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer) {
	old_atomic.StorePointer(addr, val)
	return
}

// Helper for ARM.  Linker will discard on other systems
func panic64() {
	panic("sync/atomic: broken 64-bit atomic operations (buggy QEMU)")
}
