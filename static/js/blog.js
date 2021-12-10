$(document).ready(function () {
    //注册
    $("#register-form").validate({
        rules: {
            username: {
                required: true,
                rangelength: [5, 10]
            },
            password: {
                required: true,
                rangelength: [5, 10]
            },
            repassword: {
                required: true,
                rangelength: [5, 10],
                equalTo: "#register-password"
            }
        },
        messages: {
            username: {
                required: "请输入用户名",
                rangelength: "用户名必须是5-10位"
            },
            password: {
                required: "请输入密码",
                rangelength: "密码必须是5-10位"
            },
            repassword: {
                required: "请确认密码",
                rangelength: "密码必须是5-10位",
                equalTo: "两次输入的密码必须相等"
            }
        },
        submitHandler: function (form) {
            const urlStr = "/register";
            // alert("urlStr:"+urlStr)
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message)
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/login"
                        }, 1000)
                    }
                },
                err: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            })
        }
    });
});