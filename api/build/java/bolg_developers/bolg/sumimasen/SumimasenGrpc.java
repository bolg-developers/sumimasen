package bolg_developers.bolg.sumimasen;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.22.0-SNAPSHOT)",
    comments = "Source: api/sumimasen.proto")
public final class SumimasenGrpc {

  private SumimasenGrpc() {}

  public static final String SERVICE_NAME = "bolg.sumimasen.Sumimasen";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<bolg_developers.bolg.sumimasen.SumimasenRequest,
      bolg_developers.bolg.sumimasen.SumimasenResponse> getSumimasenMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Sumimasen",
      requestType = bolg_developers.bolg.sumimasen.SumimasenRequest.class,
      responseType = bolg_developers.bolg.sumimasen.SumimasenResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<bolg_developers.bolg.sumimasen.SumimasenRequest,
      bolg_developers.bolg.sumimasen.SumimasenResponse> getSumimasenMethod() {
    io.grpc.MethodDescriptor<bolg_developers.bolg.sumimasen.SumimasenRequest, bolg_developers.bolg.sumimasen.SumimasenResponse> getSumimasenMethod;
    if ((getSumimasenMethod = SumimasenGrpc.getSumimasenMethod) == null) {
      synchronized (SumimasenGrpc.class) {
        if ((getSumimasenMethod = SumimasenGrpc.getSumimasenMethod) == null) {
          SumimasenGrpc.getSumimasenMethod = getSumimasenMethod = 
              io.grpc.MethodDescriptor.<bolg_developers.bolg.sumimasen.SumimasenRequest, bolg_developers.bolg.sumimasen.SumimasenResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "bolg.sumimasen.Sumimasen", "Sumimasen"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  bolg_developers.bolg.sumimasen.SumimasenRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.lite.ProtoLiteUtils.marshaller(
                  bolg_developers.bolg.sumimasen.SumimasenResponse.getDefaultInstance()))
                  .build();
          }
        }
     }
     return getSumimasenMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static SumimasenStub newStub(io.grpc.Channel channel) {
    return new SumimasenStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static SumimasenBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new SumimasenBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static SumimasenFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new SumimasenFutureStub(channel);
  }

  /**
   */
  public static abstract class SumimasenImplBase implements io.grpc.BindableService {

    /**
     */
    public void sumimasen(bolg_developers.bolg.sumimasen.SumimasenRequest request,
        io.grpc.stub.StreamObserver<bolg_developers.bolg.sumimasen.SumimasenResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getSumimasenMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getSumimasenMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                bolg_developers.bolg.sumimasen.SumimasenRequest,
                bolg_developers.bolg.sumimasen.SumimasenResponse>(
                  this, METHODID_SUMIMASEN)))
          .build();
    }
  }

  /**
   */
  public static final class SumimasenStub extends io.grpc.stub.AbstractStub<SumimasenStub> {
    private SumimasenStub(io.grpc.Channel channel) {
      super(channel);
    }

    private SumimasenStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SumimasenStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new SumimasenStub(channel, callOptions);
    }

    /**
     */
    public void sumimasen(bolg_developers.bolg.sumimasen.SumimasenRequest request,
        io.grpc.stub.StreamObserver<bolg_developers.bolg.sumimasen.SumimasenResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getSumimasenMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class SumimasenBlockingStub extends io.grpc.stub.AbstractStub<SumimasenBlockingStub> {
    private SumimasenBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private SumimasenBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SumimasenBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new SumimasenBlockingStub(channel, callOptions);
    }

    /**
     */
    public bolg_developers.bolg.sumimasen.SumimasenResponse sumimasen(bolg_developers.bolg.sumimasen.SumimasenRequest request) {
      return blockingUnaryCall(
          getChannel(), getSumimasenMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class SumimasenFutureStub extends io.grpc.stub.AbstractStub<SumimasenFutureStub> {
    private SumimasenFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private SumimasenFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SumimasenFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new SumimasenFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<bolg_developers.bolg.sumimasen.SumimasenResponse> sumimasen(
        bolg_developers.bolg.sumimasen.SumimasenRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getSumimasenMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SUMIMASEN = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final SumimasenImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(SumimasenImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SUMIMASEN:
          serviceImpl.sumimasen((bolg_developers.bolg.sumimasen.SumimasenRequest) request,
              (io.grpc.stub.StreamObserver<bolg_developers.bolg.sumimasen.SumimasenResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (SumimasenGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .addMethod(getSumimasenMethod())
              .build();
        }
      }
    }
    return result;
  }
}
