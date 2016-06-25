package olilib_lru

import (
        "testing"
        "time"
        "sync"
)

var TESTLRU = NewString(10)

func TestNew(t *testing.T) {
        expectedReturn := &LRUCacheString{
                10,
                make(map[string]string, 10),
                make(map[string]int64, 10),
                &sync.Mutex{},
        }

        result := NewString(10)

        if expectedReturn.KeyCount != result.KeyCount {
                t.Fatalf("Expected %v, got %v", expectedReturn.KeyCount, result.KeyCount)
        }
}

func TestAdd(t *testing.T) {
        key, err := TESTLRU.Add("Key", "Value")

        if key != "Key" {
                t.Fatalf("Failed adding a key to the LRU cache")
        }

        if err != false {
                t.Fatalf("Weird things happened adding a key")
        }
}

func TestAddPreExistingKey(t *testing.T) {
        key, err := TESTLRU.Add("Key", "Value")

        if key != "Value" {
                t.Fatalf("Did not receive the value from the LRU cache.")
        }

        if err != true {
                t.Fatalf("Failed to add pre-existing key.")
        }
}

func TestAddRemoveOldest(t *testing.T) {
        testLRU := NewString(3)
        testLRU.Add("Key1", "value1")
        time.Sleep(time.Nanosecond * 1)
        testLRU.Add("Key2", "value2")
        testLRU.Add("Key3", "value3")

        expectedKeys := []string{"Key1", "Key2", "Key3"}

        for i := range expectedKeys {
                if _, ok := testLRU.Keys[expectedKeys[i]]; !ok {
                        t.Fatalf("Key %v not found in the test LRU (Keys)", expectedKeys[i])
                }

                if _, ok := testLRU.KeyTimeouts[expectedKeys[i]]; !ok {
                        t.Fatalf("Key %v not found in the test LRU (KeyTimeouts)", expectedKeys[i])
                }
        }

        testLRU.Add("Key4", "value4")

        expectedKeys = []string{"Key4", "Key2", "Key3"}

        for i := range expectedKeys {
                if _, ok := testLRU.Keys[expectedKeys[i]]; !ok {
                        t.Fatalf("Key %v not found in the test LRU (Keys)", expectedKeys[i])
                }

                if _, ok := testLRU.KeyTimeouts[expectedKeys[i]]; !ok {
                        t.Fatalf("Key %v not found in the test LRU (KeyTimeouts)", expectedKeys[i])
                }
        }


}

func TestGet(t *testing.T) {
        testLRU := NewString(5)
        testLRU.Add("Key1", "value1")

        originalTime := testLRU.KeyTimeouts["Key1"]
        value, keyExists := testLRU.Get("Key1")

        if keyExists != true {
                t.Fatalf("Key doesn't exist in the LRU Cache")
        }

        if "value1" != value {
                t.Fatalf("Expected value1, got %v", value)
        }

        if testLRU.KeyTimeouts["Key1"] == originalTime {
                t.Fatalf("Time for retrieving a key didnt update, please fix.")
        }
}


func TestGetDoesntExist(t *testing.T) {
        testLRU := NewString(5)
        _, keyExists := testLRU.Get("Key1")

        if keyExists != false {
                t.Fatalf("For whatever reason, the key exists")
        }
}
