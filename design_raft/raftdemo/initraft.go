package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb/v2"
)

// 初始化Raft节点
func initRaft(nodeID, addr string, fsm *KVStore, bootstrap bool) (*raft.Raft, error) {
	// 创建数据目录
	dataDir := filepath.Join(".", "raft-data", nodeID)
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}

	// 配置
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(nodeID)

	// 日志存储（使用boltdb）
	logStore, err := raftboltdb.NewBoltStore(filepath.Join(dataDir, "raft.log"))
	if err != nil {
		return nil, err
	}

	// 稳定存储（与日志存储共用）
	stableStore := logStore

	// 快照存储
	snapStore, err := raft.NewFileSnapshotStore(dataDir, 3, os.Stderr)
	if err != nil {
		return nil, err
	}

	// 网络传输（TCP）
	trans, err := raft.NewTCPTransport(addr, nil, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return nil, err
	}

	// 初始化Raft
	raftNode, err := raft.NewRaft(config, fsm, logStore, stableStore, snapStore, trans)
	if err != nil {
		return nil, err
	}

	// 首次启动时引导集群（仅在单节点初始化时使用）
	if bootstrap {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      config.LocalID,
					Address: trans.LocalAddr(),
				},
			},
		}
		f := raftNode.BootstrapCluster(configuration)
		if err := f.Error(); err != nil {
			return nil, fmt.Errorf("raft.Raft.BootstrapCluster: %v", err)
		}
		// if err := f.Error(); err != nil && err != raft.ErrAlreadyBootstrapped {
		// 	return nil, err
		// }
	}

	return raftNode, nil
}
