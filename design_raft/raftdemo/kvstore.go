package main

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/hashicorp/raft"
)

// KVStore 是一个简单的键值存储FSM
type KVStore struct {
	mu   sync.Mutex
	data map[string]string
}

// NewKVStore 创建新的KVStore
func NewKVStore() *KVStore {
	return &KVStore{
		data: make(map[string]string),
	}
}

// Apply 处理日志条目，更新状态
func (k *KVStore) Apply(log *raft.Log) interface{} {
	var req struct {
		Op    string // "set" 或 "delete"
		Key   string
		Value string
	}
	if err := json.Unmarshal(log.Data, &req); err != nil {
		return fmt.Errorf("解析日志失败: %v", err)
	}

	k.mu.Lock()
	defer k.mu.Unlock()

	switch req.Op {
	case "set":
		k.data[req.Key] = req.Value
		return nil
	case "delete":
		delete(k.data, req.Key)
		return nil
	default:
		return fmt.Errorf("未知操作: %s", req.Op)
	}
}

// Snapshot 创建当前状态的快照
func (k *KVStore) Snapshot() (raft.FSMSnapshot, error) {
	k.mu.Lock()
	defer k.mu.Unlock()

	// 复制当前数据作为快照
	snapData := make(map[string]string)
	for k, v := range k.data {
		snapData[k] = v
	}
	return &kvsSnapshot{data: snapData}, nil
}

// Restore 从快照恢复状态
func (k *KVStore) Restore(reader io.ReadCloser) error {
	defer reader.Close()

	var snapData map[string]string
	if err := json.NewDecoder(reader).Decode(&snapData); err != nil {
		return err
	}

	k.mu.Lock()
	defer k.mu.Unlock()
	k.data = snapData
	return nil
}

// Get 读取键值（非FSM接口，用于外部查询）
func (k *KVStore) Get(key string) (string, bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	val, ok := k.data[key]
	return val, ok
}

// kvsSnapshot 是KVStore的快照实现
type kvsSnapshot struct {
	data map[string]string
}

func (k *kvsSnapshot) Persist(sink raft.SnapshotSink) error {
	if err := json.NewEncoder(sink).Encode(k.data); err != nil {
		sink.Cancel()
		return err
	}
	return sink.Close()
}

func (k *kvsSnapshot) Release() {}
