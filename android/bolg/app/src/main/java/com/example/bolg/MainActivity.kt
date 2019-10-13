package com.example.bolg

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import kotlinx.android.synthetic.main.activity_main.*
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.async
import kotlinx.coroutines.launch

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        bt.setOnClickListener{
            func()
        }
    }

    fun func() = GlobalScope.launch(Dispatchers.Main){
        try {
            val ss = SumimasenGrpctask()
            async(Dispatchers.Default) {
                val host = host_edit.text
                val port = port_edit.text
                ss.greet(host.toString(), port.toString())
            }.await().let {
                tv.text = it.message
            }
        }catch (e:Exception){
            tv.setText(e.toString())
        }
    }
}
