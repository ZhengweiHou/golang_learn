package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/raft"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("用法: ./raftdemo <节点ID> <地址:端口>")
		os.Exit(1)
	}
	nodeID := os.Args[1]
	addr := os.Args[2]

	// 初始化KVStore
	kv := NewKVStore()

	// 初始化Raft（首次启动单个节点时bootstrap为true，后续节点为false）
	bootstrap := len(os.Args) >= 4 && os.Args[3] == "bootstrap"
	raftNode, err := initRaft(nodeID, addr, kv, bootstrap)
	if err != nil {
		fmt.Printf("初始化Raft失败: %v\n", err)
		os.Exit(1)
	}
	defer raftNode.Shutdown()

	// 等待节点就绪
	time.Sleep(2 * time.Second)

	// 简单交互命令行
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		cmd := scanner.Text()
		switch cmd {
		case "exit":
			return
		case "leader":
			leader := raftNode.Leader()
			fmt.Printf("当前 leader: %s\n", leader)
		case "add":
			// 添加节点示例（格式: add <节点ID> <地址>）
			var newNodeID, newAddr string
			fmt.Scan(&newNodeID, &newAddr)
			future := raftNode.AddVoter(raft.ServerID(newNodeID), raft.ServerAddress(newAddr), 0, 0)
			if err := future.Error(); err != nil {
				fmt.Printf("添加节点失败: %v\n", err)
			} else {
				fmt.Println("添加节点成功")
			}
		default:
			// 处理set命令（格式: set <key> <value>）
			var op, key, value string
			fmt.Sscanf(cmd, "%s %s %s", &op, &key, &value)
			if op == "set" {
				data, _ := json.Marshal(map[string]string{
					"Op":    "set",
					"Key":   key,
					"Value": value,
				})
				future := raftNode.Apply(data, 5*time.Second)
				if err := future.Error(); err != nil {
					fmt.Printf("写入失败: %v\n", err)
				} else {
					fmt.Println("写入成功")
				}
			} else if op == "get" {
				val, ok := kv.Get(key)
				if ok {
					fmt.Printf("%s = %s\n", key, val)
				} else {
					fmt.Printf("键 %s 不存在\n", key)
				}
			} else {
				fmt.Println("未知命令，支持: set <key> <value>, get <key>, add <节点ID> <地址>, leader, exit")
			}
		}
	}
}
