$(document).ready(function () {
    $("#loginform").validate({ //验证表单

        rules: {  //关键字
            username: {
                required:true,
                rangelength:[2,16],
            },
            password: {
                required:true,
                rangelength:[5,10],
            },
            repassword: {
                required:true,
                rangelength:[5,10],
                equalTo:"#register_password"
            },

        },
        messages: {
            username: {
                required:"请输入用户名",
                rangelength:"用户名必须是5-10位",
            },
            password: {
                required:"请输入密码",
                rangelength:"密码必须是5-10位",
            },
            repassword: {
                required:"请确认密码",
                rangelength:"密码必须是5-10位",
                equalTo:"密码必须一样"
            },



        },
        submitHandler: function(form){

            $(form).ajaxSubmit({
                url:"/register",
                type:"post",
                dataType:"json",
                success:function (data,status) {   // 登录成功 data 是beego的data 成功之后的回调
                    alert("data:添加成功"+data.message)//status是状态
                    if(data.code==1){  //判断状态码 如果是1 说明成功 跳转主页 否则 还在原界面
                        setTimeout(function () {
                            window.location.href = "/"
                        },1000)
                    }


                },
                error:function (data,status) {  //登录失败  data  和 状态
                    alert("err"+data+status)

                }
            });
        }


    })
})



//登录------------------------------------------------------------------------------


$(document).ready(function () {
    $("#login-form").validate({ //验证表单

        rules: {  //关键字
            username: {
                required:true,
                rangelength:[2,16],
            },
            password: {
                required:true,
                rangelength:[5,10],
            },


        },
        messages: {
            username: {
                required:"请输入用户名",
                rangelength:"用户名必须是5-10位",
            },
            password: {
                required:"请输入密码",
                rangelength:"密码必须是5-10位",
            },



        },
        submitHandler: function(form){

            $(form).ajaxSubmit({
                url:"/login",
                type:"post",
                dataType:"json",
                success:function (data,status) {   // 登录成功 data 是beego的data 成功之后的回调
                    alert("data:"+data.message)//status是状态
                    if(data.code==1){  //判断状态码 如果是1 说明成功 跳转主页 否则 还在原界面
                        setTimeout(function () {
                            window.location.href = "/"
                        },1000)
                    }


                },
                error:function (data,status) {  //登录失败  data  和 状态
                    alert("err"+data.message+status)

                }
            });
        }


    })
})