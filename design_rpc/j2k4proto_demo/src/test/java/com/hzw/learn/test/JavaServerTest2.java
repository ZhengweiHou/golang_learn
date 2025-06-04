package com.hzw.learn.test;


import com.hzw.learn.j2golitex.HelloProtoServiceGrpc;
import com.hzw.learn.j2golitex.Request;
import com.hzw.learn.j2golitex.Response;

import io.grpc.Grpc;
import io.grpc.InsecureServerCredentials;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.ServerCredentials;
import io.grpc.TlsServerCredentials;
import io.grpc.TlsServerCredentials.Builder;

import org.junit.Test;
import org.springframework.core.io.ClassPathResource;

import java.io.File;
import java.io.IOException;

/**
 * @ClassName JavaServerTest
 * @Description TODO
 * @Author houzw
 * @Date 2025/4/10
 **/
public class JavaServerTest2 {
    // 启动服务
    @Test
    public void testServer() throws IOException, InterruptedException {
        int port = 8082;
        Server server = ServerBuilder.forPort(port)
                .addService(new HelloGrpcServerImpl())
                .build()
                .start();
        System.out.println("Server started, listening on " + port);

        server.awaitTermination();
    }

    @Test
    public void testServer_https() throws IOException, InterruptedException {
        int port = 8082;

        // TlsServerCredentials.Builder tlsBuilder = TlsServerCredentials.newBuilder();
        // File serverCert = new ClassPathResource("x509/server.crt").getFile();
        // File serverKey = new ClassPathResource("x509/server.key").getFile();
        // File caCert = new ClassPathResource("x509/ca.crt").getFile();
        // ServerCredentials credentials  = tlsBuilder.trustManager(caCert).keyManager(serverCert, serverKey).build();
        ServerCredentials credentials = InsecureServerCredentials.create(); // 不建议使用,非常坑
        Server server = Grpc.newServerBuilderForPort(port, credentials)
        .addService(new HelloGrpcServerImpl()).build();
        server.start().awaitTermination();
        
    }

    // 实现服务接口
    public static class HelloGrpcServerImpl extends HelloProtoServiceGrpc.HelloProtoServiceImplBase {
        public void echo(Request request,
                         io.grpc.stub.StreamObserver<Response> responseObserver) {

            // 处理请求并构建响应
            String msg = request.getMessage();
            System.out.println("get msg:" + msg);
            String message = "Hello, we are java!";

            Response response = Response.newBuilder().setMessage(message).build();
            responseObserver.onNext(response);
            responseObserver.onCompleted();
        }
    }
}
