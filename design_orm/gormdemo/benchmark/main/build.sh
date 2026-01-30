#!/bin/bash

# 性能基准测试构建脚本
# 用于在服务器上构建和运行性能基准测试

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 项目根目录
PROJECT_ROOT="../../.."
BENCHMARK_DIR="."

# 输出函数
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 检查依赖
check_dependencies() {
    log_info "检查依赖..."
    
    # 检查 Go
    if ! command -v go &> /dev/null; then
        log_error "Go 未安装，请先安装 Go"
        exit 1
    fi
    
    # 检查 Go 版本
    GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
    log_info "Go 版本: $GO_VERSION"
    
    # 检查是否在正确的目录
    if [ ! -f "$PROJECT_ROOT/go.mod" ]; then
        log_error "未找到 go.mod 文件，请确保在正确的目录中运行"
        exit 1
    fi
    
    log_success "依赖检查通过"
}

# 构建基准测试程序
build_benchmark() {
    log_info "构建性能基准测试程序..."
    
    cd "$PROJECT_ROOT"
    
    # 下载依赖
    log_info "下载依赖..."
    go mod download
    
    # 构建基准测试程序
    log_info "编译基准测试程序..."
    go build -o benchmark_main ./gormdemo/benchmark/main/benchmark_main.go
    
    if [ $? -eq 0 ]; then
        log_success "基准测试程序构建成功"
        mv benchmark_main gormdemo/benchmark/main/
    else
        log_error "构建失败"
        exit 1
    fi
}

# 运行基准测试
run_benchmark() {
    log_info "运行性能基准测试..."
    
    cd "$BENCHMARK_DIR"
    
    # 检查可执行文件是否存在
    if [ ! -f "benchmark_main" ]; then
        log_error "未找到 benchmark_main 可执行文件，请先运行构建"
        exit 1
    fi
    
    # 创建输出目录
    mkdir -p reports
    
    # 获取当前时间戳
    TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
    REPORT_FILE="reports/benchmark_report_${TIMESTAMP}.json"
    
    # 运行基准测试
    log_info "开始性能基准测试..."
    log_info "报告将保存到: $REPORT_FILE"
    
    # 使用不同的配置运行测试
    echo "=== 快速测试 (10线程, 30秒) ==="
    ./benchmark_main -threads 10 -test-time 30s -report "$REPORT_FILE"
    
    echo ""
    echo "=== 标准测试 (50线程, 60秒) ==="  
    ./benchmark_main -threads 50 -test-time 60s -report "reports/benchmark_standard_${TIMESTAMP}.json"
    
    echo ""
    echo "=== 压力测试 (100线程, 120秒) ==="
    ./benchmark_main -threads 100 -test-time 120s -report "reports/benchmark_stress_${TIMESTAMP}.json"
    
    log_success "所有基准测试完成"
}

# 清理
cleanup() {
    log_info "清理构建文件..."
    
    cd "$BENCHMARK_DIR"
    
    if [ -f "benchmark_main" ]; then
        rm benchmark_main
        log_success "清理完成"
    fi
}

# 显示帮助
show_help() {
    echo "性能基准测试管理脚本"
    echo ""
    echo "用法: $0 [命令]"
    echo ""
    echo "命令:"
    echo "  build     构建基准测试程序"
    echo "  run       运行基准测试"
    echo "  all       构建并运行基准测试"
    echo "  clean     清理构建文件"
    echo "  help      显示此帮助信息"
    echo ""
    echo "示例:"
    echo "  $0 build    # 仅构建"
    echo "  $0 run      # 仅运行（需要先构建）"
    echo "  $0 all      # 构建并运行"
    echo "  $0 clean    # 清理"
}

# 主函数
main() {
    case "$1" in
        "build")
            check_dependencies
            build_benchmark
            ;;
        "run")
            check_dependencies
            run_benchmark
            ;;
        "all")
            check_dependencies
            build_benchmark
            run_benchmark
            ;;
        "clean")
            cleanup
            ;;
        "help"|"-h"|"--help")
            show_help
            ;;
        *)
            log_error "未知命令: $1"
            echo ""
            show_help
            exit 1
            ;;
    esac
}

# 执行主函数
main "$1"
