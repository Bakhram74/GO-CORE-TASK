package main

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	hashMap := NewHashMap(7)

	key := "anyKey"
	value := 10
	index := hashFunction(key, hashMap.size)

	if hashMap.buckets[index] != nil {
		t.Error("hashMap must be nil")
	}

	hashMap.Add(key, value)

	if hashMap.buckets[index] == nil {
		t.Error("hashMap equal to nil")
	}

}

func TestGet(t *testing.T) {
	hashMap := NewHashMap(7)

	key := "anyKey"
	value := 10

	hashMap.Add(key, value)

	val, ok := hashMap.Get(key)

	if ok != true {
		t.Errorf("invalid bool value had gone: %t", ok)
	}

	if val != value {
		t.Errorf("invalid value had gone: %d", val)
	}
}
func TestExists(t *testing.T) {
	hashMap := NewHashMap(7)

	key := "anyKey"
	value := 10

	ok := hashMap.Exists(key)

	if ok != false {
		t.Errorf("bool value must be false: %t", ok)
	}

	hashMap.Add(key, value)

	ok = hashMap.Exists(key)

	if ok != true {
		t.Errorf("bool value must be true: %t", ok)
	}

}

func TestRemove(t *testing.T) {
	hashMap := NewHashMap(7)

	key := "anyKey"
	value := 10

	hashMap.Add(key, value)

	hashMap.Remove(key)

	val, ok := hashMap.Get(key)

	if ok != false {
		t.Errorf("invalid bool value had gone: %t", ok)
	}

	if val != 0 {
		t.Errorf("invalid value had gone: %d", val)
	}
}

func TestCopy(t *testing.T) {
	hashMap := NewHashMap(7)

	Keys := []string{"anyKey", "apple", "lemon", "some"}
	expected := make(map[string]int, len(Keys))

	for i := range Keys {
		expected[Keys[i]] = i
	}

	for i := range Keys {
		hashMap.Add(Keys[i], i)
	}

	gotMap := hashMap.Copy()

	if !reflect.DeepEqual(expected, gotMap) {
		t.Errorf("gotMap :%v and expected: %v are not equal", gotMap, expected)
	}

}
