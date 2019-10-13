package com.example.bolg

import android.util.Log
import bolg_developers.sumimasen.api.SumimasenGrpc
import bolg_developers.sumimasen.api.SumimasenRequest
import bolg_developers.sumimasen.api.SumimasenResponse
import io.grpc.ManagedChannelBuilder

class SumimasenGrpctask {

    fun greet(vararg params: String): SumimasenResponse {
        Log.d(javaClass.simpleName, "call greet")
        val host = params[0]
        val port = Integer.parseInt(params[1])
        Log.d(javaClass.simpleName, "host:${host},port:${port}")

        try {
            val channel = ManagedChannelBuilder
                // 127.0.0.1 はダメ！　local ipを直接指定する。
                .forAddress(host, port)
                .usePlaintext()
                .build()
            Log.d(javaClass.simpleName,"chammel" + channel.toString())
            val stub = SumimasenGrpc.newBlockingStub(channel)
            Log.d(javaClass.simpleName,"stub" + stub.toString())
            val request = SumimasenRequest.newBuilder().build()
            Log.d(javaClass.simpleName,"request" + request.toString())
            return try {
                //stub.sumimasen(request)
                SumimasenResponse.newBuilder().setMessage(stub.sumimasen(request).message).build()
            } catch (e: Exception) {
                Log.d(javaClass.simpleName, "Failed to get message", e)
                SumimasenResponse.newBuilder().setMessage(e.toString()).build()
            }
        }catch (e:Exception){
            Log.d(javaClass.simpleName,e.toString())
            return SumimasenResponse.newBuilder().setMessage("失敗").build()
        }
    }
}