#!/bin/bash

# ent 性能基准测试构建脚本

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 函数定义
print_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 构建函数
build() {
    print_info "开始构建 ent 性能基准测试程序..."
    
    if go build -o benchmark_main benchmark_main.go; then
        print_info "构建成功！可执行文件: benchmark_main"
        ls -la benchmark_main
    else
        print_error "构建失败！"
        exit 1
    fi
}

# 运行函数
run() {
    print_info "运行 ent 性能基准测试..."
    
    if [ -f "./benchmark_main" ]; then
        ./benchmark_main "$@"
    else
        print_error "可执行文件不存在，请先运行 './build.sh build'"
        exit 1
    fi
}

# 清理函数
clean() {
    print_info "清理构建文件..."
    
    if [ -f "./benchmark_main" ]; then
        rm -f ./benchmark_main
        print_info "已删除 benchmark_main"
    fi
    
    if [ -f "./benchmark_report.json" ]; then
        rm -f ./benchmark_report.json
        print_info "已删除 benchmark_report.json"
    fi
}

# 帮助函数
help() {
    echo "ent 性能基准测试工具"
    echo ""
    echo "用法: $0 {build|run|clean|all|help} [参数...]"
    echo ""
    echo "命令:"
    echo "  build       构建基准测试程序"
    echo "  run         运行基准测试 (可传递参数给基准测试程序)"
    echo "  clean       清理构建文件"
    echo "  all         构建并运行基准测试"
    echo "  help        显示此帮助信息"
    echo ""
    echo "基准测试参数示例:"
    echo "  ./build.sh run -threads 10 -test-time 30s -batch-size 50"
    echo ""
    echo "可用参数:"
    echo "  -threads: 并发线程数 (默认: 1)"
    echo "  -iterations: 每个线程的迭代次数 (默认: 1)"
    echo "  -warmup: 预热时间 (默认: 5s)"
    echo "  -test-time: 测试时间 (默认: 5s)"
    echo "  -batch-size: 批量操作大小 (默认: 100)"
    echo "  -report: 报告文件名 (默认: benchmark_report.json)"
    echo "  -dsn: 数据库连接字符串"
    echo ""
}

# 构建并运行
all() {
    build
    echo ""
    run "$@"
}

# 主逻辑
case "$1" in
    build)
        build
        ;;
    run)
        shift
        run "$@"
        ;;
    clean)
        clean
        ;;
    all)
        shift
        all "$@"
        ;;
    help|*)
        help
        ;;
esac
