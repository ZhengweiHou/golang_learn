package com.hzw.learn.j2gokitex;

import com.hzw.learn.j2golitex.HelloProtoServiceGrpc;
import com.hzw.learn.j2golitex.Request;
import com.hzw.learn.j2golitex.Response;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.stub.StreamObserver;
import org.junit.Test;
import sun.misc.GThreadHelper;

/**
 * @ClassName Java2GoKitexTest
 * @Description TODO
 * @Author houzw
 * @Date 2025/4/10
 **/
public class Java2GoKitexTest {

    @Test
    public void tojava() throws InterruptedException {
        call(8082);
    }

    @Test
    public void togokitex() throws InterruptedException {
        call(8888);
//        call(8001);
//        callFuture(8888);
    }

    public void call( int port) throws InterruptedException {
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", port)
                .usePlaintext()
                .build();
        Request request = Request.newBuilder().setMessage("Hello from Java client").build();
        HelloProtoServiceGrpc.HelloProtoServiceBlockingStub stub
                = HelloProtoServiceGrpc.newBlockingStub(channel);
        for (int i = 0; i < 1; i++) {
            Response response = stub.echo(request);
            System.out.println("Response from server: " + response.getMessage());
            Thread.sleep(1000);
        }
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
}
