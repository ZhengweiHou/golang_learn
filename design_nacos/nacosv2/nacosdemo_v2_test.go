package nacosv1_test

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func TestNacosConfigNamingClient(t *testing.T) {
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "e525eafa-f7d7-4029-83d9-008937f9d468", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建clientConfig的另一种方式
	clientConfig = *constant.NewClientConfig(
		constant.WithNamespaceId("e525eafa-f7d7-4029-83d9-008937f9d468"), //当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "console1.nacos.io",
			ContextPath: "/nacos",
			Port:        80,
			Scheme:      "http",
		},
		{
			IpAddr:      "console2.nacos.io",
			ContextPath: "/nacos",
			Port:        80,
			Scheme:      "http",
		},
	}

	// 创建serverConfig的另一种方式
	serverConfigs = []constant.ServerConfig{
		*constant.NewServerConfig(
			"console1.nacos.io",
			80,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
		*constant.NewServerConfig(
			"console2.nacos.io",
			80,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
	}

	// 创建服务发现客户端
	_, _ = clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	// 创建动态配置客户端
	_, _ = clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v  %v", namingClient, configClient)

}

// TestNacosConfigNamingClient ###### 测试nacos配置中心 ######
func TestNacosConfig(t *testing.T) {
	fmt.Println("###### 测试nacos配置中心 ######")
	hosts := []string{
		"localhost:8848",
		// "localhost:8813",
	}
	configClient, err := newNacosConfigClient(hosts, "test")
	if err != nil {
		t.Fatal(err)
	}

	// 发布配置 PublishConfig
	fmt.Println("==== 发布配置 PublishConfig =====")
	success, err := configClient.PublishConfig(vo.ConfigParam{
		DataId:  "golangNacosConfigTest",
		Group:   "hzw_gotest",
		Content: "hello world!222222",
	})
	fmt.Printf("PublishConfig success:%v err:%v\n", success, err)

	// 获取配置 GetConfig
	fmt.Println("==== 获取配置 GetConfig =====")
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "golangNacosConfigTest",
		Group:  "hzw_gotest",
	})
	fmt.Printf("GetConfig content:%v err:%v\n", content, err)

	// 监听配置 ListenConfig
	fmt.Println("==== 监听配置 ListenConfig =====")
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "golangNacosConfigTest",
		Group:  "hzw_gotest",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Printf("--- config changed group:%s dataId:%s content:%s\n", group, dataId, data)
		},
	})
	fmt.Printf("ListenConfig err:%v\n", err)
	time.Sleep(time.Second)

	// 修改 PublishConfig
	fmt.Println("==== 修改 PublishConfig =====")
	success, err = configClient.PublishConfig(vo.ConfigParam{
		DataId:  "golangNacosConfigTest",
		Group:   "hzw_gotest",
		Content: "hello world!33333",
	})
	fmt.Printf("[update] PublishConfig success:%v err:%v\n", success, err)

	time.Sleep(time.Second)

	// 取消监听配置 CancelListenConfig
	fmt.Println("==== 取消监听配置 CancelListenConfig =====")
	err = configClient.CancelListenConfig(vo.ConfigParam{
		DataId: "golangNacosConfigTest",
		Group:  "hzw_gotest",
	})
	fmt.Printf("CancelListenConfig err:%v\n", err)

	// 搜索配置 SearchConfig
	fmt.Println("==== 搜索配置 SearchConfig =====")
	configPage, err := configClient.SearchConfig(vo.SearchConfigParam{
		// Search:   "accurate", //精确搜索
		Search:   "blur", // 模糊搜索
		Group:    "hzw_gotest",
		DataId:   "golangNacosConfigTes*",
		PageNo:   1,
		PageSize: 10,
	})
	fmt.Printf("SearchConfig configPage:%v err:%v\n", configPage, err)
	item0, _ := json.Marshal(configPage.PageItems[0])
	fmt.Printf("item[0]:%s\n", item0)

	// 删除配置 DeleteConfig
	fmt.Println("==== 删除配置 DeleteConfig =====")
	success, err = configClient.DeleteConfig(vo.ConfigParam{
		DataId: "golangNacosConfigTest",
		Group:  "hzw_gotest",
	})
	fmt.Printf("DeleteConfig success:%v err:%v\n", success, err)
}

// TestNacosNamingClient ###### 测试nacos服务发现 ######
func TestNacosNaming(t *testing.T) {
	fmt.Println("###### 测试nacos服务发现 ######")
	hosts := []string{
		"localhost:8848",
		// "localhost:8813",
	}
	namingClient, err := newNacosNamingClient(hosts, "test")
	if err != nil {
		t.Fatal(err)
	}

	// 注册实例 RegisterInstance
	fmt.Println("==== 注册实例 RegisterInstance =====")
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "hzwHost",
		ServiceName: "hzw_go_test",
		Port:        8080,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		ClusterName: "hzwcluster",
		GroupName:   "hzw_group",
		Metadata:    map[string]string{"h": "hzw"},
	})
	fmt.Printf("RegisterInstance success:%v err:%v\n", success, err)

	// 获取服务信息 GetService
	fmt.Println("==== 获取服务信息 GetService =====")
	services, err := namingClient.GetService(vo.GetServiceParam{
		ServiceName: "hzw_go_test",
		Clusters:    []string{"hzwcluster"}, // 默认值DEFAULT
		GroupName:   "hzw_group",             // 默认值DEFAULT_GROUP
	})
	// sstr, err := json.MarshalIndent(services, "", " ")
	sstr, _ := json.Marshal(services)
	fmt.Printf("GetService services:%s err:%v\n", sstr, err)

	time.Sleep(time.Second * 20)

	// 注销实例 DeregisterInstance
	fmt.Println("==== 注销实例 DeregisterInstance =====")
	success, err = namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          "hzwHost",
		ServiceName: "hzw_go_test",
		Port:        8080,
		Ephemeral:   true,
		GroupName:   "hzw_group",
		Cluster:     "hzwcluster",
	})
	fmt.Printf("DeregisterInstance success:%v err:%v\n", success, err)

}

// newNacosConfigClient 创建nacos配置客户端
func newNacosConfigClient(hosts []string, ns string) (config_client.IConfigClient, error) {
	clientConfig := newNacosClientConfig(ns)
	serverConfigs := newNacosServerConfig(hosts)

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return nil, err
	}
	return configClient, nil
}

// newNacosNamingClient 创建nacos服务发现客户端
func newNacosNamingClient(hosts []string, ns string) (naming_client.INamingClient, error) {
	clientConfig := newNacosClientConfig(ns)
	serverConfigs := newNacosServerConfig(hosts)

	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return nil, err
	}
	return namingClient, nil
}
func newNacosClientConfig(ns string) constant.ClientConfig {

	// 方式1
	clientConfig := constant.ClientConfig{
		NamespaceId:         ns,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 方式2
	clientConfig = *constant.NewClientConfig(
		constant.WithNamespaceId(ns),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	return clientConfig
}

func newNacosServerConfig(hostList []string) []constant.ServerConfig {
	serverConfigs := []constant.ServerConfig{}
	// if hosts == "" {
	// 	return serverConfigs
	// }

	// hostList := strings.Split(hosts, ",")
	for _, host := range hostList {
		parts := strings.Split(host, ":")
		if len(parts) != 2 {
			continue
		}
		port, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		serverConfig := constant.NewServerConfig(
			parts[0],
			uint64(port),
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		)
		serverConfigs = append(serverConfigs, *serverConfig)
	}
	return serverConfigs
}
