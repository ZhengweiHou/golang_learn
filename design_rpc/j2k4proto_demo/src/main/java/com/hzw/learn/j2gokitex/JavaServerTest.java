package com.hzw.learn.j2gokitex;


import com.hzw.learn.j2golitex.HelloProtoServiceGrpc;
import com.hzw.learn.j2golitex.Request;
import com.hzw.learn.j2golitex.Response;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import org.junit.Test;

import java.io.IOException;

/**
 * @ClassName JavaServerTest
 * @Description TODO
 * @Author houzw
 * @Date 2025/4/10
 **/
public class JavaServerTest {
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
