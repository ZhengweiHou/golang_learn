package com.hzw.learn.test;

import java.io.File;
import java.io.IOException;
import java.util.concurrent.TimeUnit;

import org.junit.Test;
//import sun.misc.GThreadHelper;
import org.springframework.core.io.ClassPathResource;

import com.hzw.learn.j2golitex.HelloProtoServiceGrpc;
import com.hzw.learn.j2golitex.Request;
import com.hzw.learn.j2golitex.Response;

import io.grpc.ChannelCredentials;
import io.grpc.Grpc;
import io.grpc.InsecureChannelCredentials;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.TlsChannelCredentials;
import io.grpc.stub.StreamObserver;

/**
 * @ClassName Java2GoKitexTest
 * @Description TODO
 * @Author houzw
 * @Date 2025/4/10
 **/
public class Java2GoKitexTest2 {


    @Test
    public void togokitex() throws InterruptedException {
	    call(8886, false); // golang
    }

    @Test
    public void testclient() {
	    try {
	        call(8886, true); // golang
        } catch (InterruptedException e) {e.printStackTrace();} 
   }
    @Test
    public void t1() {
	    try {
	        call(1111, false); // golang
        } catch (InterruptedException e) {e.printStackTrace();} 
   }
    @Test
    public void t2() {
	    try {
	        call(7777, false); // golang
        } catch (InterruptedException e) {e.printStackTrace();} 
   }
    @Test
    public void t3() {
	    try {
	        call(8082, false); // golang
        } catch (InterruptedException e) {e.printStackTrace();} 
   }
    @Test
    public void t4() {
	    try {
	        call(8886, false); // golang
        } catch (InterruptedException e) {e.printStackTrace();} 
   }
    

    public void call(int port, boolean e) throws InterruptedException {
        // 创建一个非安全的 gRPC 通道，连接到指定端口的本地服务器
        ManagedChannelBuilder<?> build = ManagedChannelBuilder.forAddress("localhost", port);
        build.usePlaintext();
        ManagedChannel channel = build.build();

        // 构建一个请求对象，包含消息 "Hello from Java client"

        // 创建一个阻塞的 gRPC 存根，用于发送请求
        HelloProtoServiceGrpc.HelloProtoServiceBlockingStub stub
                = HelloProtoServiceGrpc.newBlockingStub(channel);

        String msg = "Hello from Java client";
        if (e) {
            msg = "err" + msg;
        }

        // 发送请求并接收响应
        for (int i = 0; i < 5; i++) {
            Request request = Request.newBuilder().setMessage(msg + i).build();
            Response response = stub.echo(request);
            System.out.println("Response from server: " + response.getMessage());
            Thread.sleep(500);
        }
        // 关闭通道
        channel.shutdown();
    }


    public void callFuture(int port) throws InterruptedException {
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", port)
                .usePlaintext()
                .build();
        Request request = Request.newBuilder().setMessage("Hello from Java client").build();

        HelloProtoServiceGrpc.HelloProtoServiceStub stub
                = HelloProtoServiceGrpc.newStub(channel);
        stub.echo(request, new StreamObserver<Response>() {
            @Override
            public void onNext(Response response) {
                System.out.println("Response from server: " + response.getMessage());
            }
            @Override
            public void onError(Throwable t) {
                System.err.println("Error: " + t.getMessage());
            }
            @Override
            public void onCompleted() {
                System.out.println("Request completed.");
                channel.shutdown();
            }
        });

        Thread.sleep(1000);
    }

    public void calltsl( int port) throws InterruptedException, IOException {
        // 当服务器配置了证书时需要指定 ca 证书
        // TlsChannelCredentials.Builder tlsBuilder = TlsChannelCredentials.newBuilder();
        // File caCert = new ClassPathResource("x509/ca.crt").getFile();
        // ChannelCredentials credentials = tlsBuilder.trustManager(caCert).build();
        // 不做服务器证书验证时使用这个
        ChannelCredentials credentials = InsecureChannelCredentials.create();
        ManagedChannelBuilder<?> builder = Grpc.newChannelBuilderForAddress("localhost", port, credentials);
        // builder.usePlaintext();
        ManagedChannel channel = builder.build();

        Request request = Request.newBuilder().setMessage("Hello from Java client").build();
        HelloProtoServiceGrpc.HelloProtoServiceBlockingStub stub = HelloProtoServiceGrpc.newBlockingStub(channel);

        for (int i = 0; i < 1; i++) {
            Response response = stub.echo(request);
            System.out.println("Response from server: " + response.getMessage());
            Thread.sleep(1000);
        }
        channel.shutdown();
    }
}
