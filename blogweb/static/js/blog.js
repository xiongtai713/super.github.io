
$(document).ready(function () {
$("#write-art-form").validate({ //验证表单

    rules:{  //关键字
        title:"required" , //必填required
        tags:"required",
        short:{
            required:true,
            minlength: 2

        },
        content:{
            required:true,
            minlength: 2
        }

    },
    messages:{
        title:"请输入标题",
        tags:"请输入标签",
        short:{
            required: "请输入简介",
            minlength:"简介内容最少2个字符"
        },
        content:{
            required:"请输入文章内容",
            minlength:"文章内容最少两个字符",
        },

    },
    submitHandler: function(form)
    {
       var urlstr="/add"  //默认是添加
        //判断文章id确定提交的表单的服务器地址
        var artid=$("#write-article-id").val()
        if(artid>0){
            urlstr="/article/update"   //如果文章id大于0 就说明是修改
        }

        $(form).ajaxSubmit({
            url:urlstr,
            type:"post",
            dataType:"json",
            success:function (data,status) {   // 登录成功 data 是beego的data 成功之后的回调
                alert("data:"+data.message)//status是状态
                setTimeout(function () {
                    window.location.href = "/"
                },1000)

            },
            error:function (data,status) {  //登录失败  data  和 状态
                alert("err"+data+status)

            }
        });

    }


    })

    //上传相册
    $("#album-upload-button").click(function () {
       var filedata= $("#album-upload-file").val()
        if (filedata.length<=0){
           alert("请选择文件")
            return
        }
        //文件上传通过Formdata去储存文件数据
        var data = new FormData()
        data.append("upload",$("#album-upload-file")[0].files[0])  //把
        //文件提交
        alert(data)
        var urlstr="/upload"
        $.ajax({
            url:urlstr,
            type:"post",
            dataType:"json",
            contentType:false,
            processData:false,
            data:data,
            success:function (data,status) {   // 登录成功 data 是beego的data 成功之后的回调
                alert("data:"+data.message)//status是状态
                setTimeout(function () {
                    window.location.href = "#"
                },1000)

            },
            error:function (data,status) {  //
                alert("错误"+data.message+status)

            }
        });


    })
























    // $("#write-submit").click(function () {
    //     //serialize将该表单里面全部要提交的数据序列化位json字符串
    //    var jsonStr= $("#write-art-form").serialize()
    //                 服务器路径                json表单
    //
    // $.post("http://localhost:8080/add",jsonStr,postCompelte)
    //     function postCompelte(data,status) {
    //        if (data.code ==1){
    //            alert(data.massage)
    //            //等待一秒会自动触发第一个参数里面的方法 ,会
    //            setTimeout(function () {
    //                window.location.href="/"
    //
    //            },1000)
    //        }else {
    //            alert(status)
    //        }
    // //
    //
    //     }

    // })




})