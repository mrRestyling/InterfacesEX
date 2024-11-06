package main

import "unsafe"

type iface struct {
	// это указатель на Interface Table или itable
	// - структуру, которая хранит некоторые метаданные о типе и список методов,
	//  используемых для удовлетворения интерфейса
	tab *itab

	// хранимые данные (указатель на значение)
	data unsafe.Pointer
}

type itab struct { // 40 bytes on a 64bit arch
	inter *interfacetype // тип интерфейса
	_type *_type         // все, что мы знаем про тип из которого образован элемент интерфейса
	hash  uint32         // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr // методы, которые должна описывать структура, чтобы релизовывать интерфейс
}
