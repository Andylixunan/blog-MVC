function registerValidate() {
    $('#register-form').validate({
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
        }
    })
}

function loginValidate(){
    $('#login-form').validate({
        rules: {
            username: {
                required: true,
                rangelength: [5, 10]
            },
            password: {
                required: true,
                rangelength: [5, 10]
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
            }
        }
    })
}

function articleValidate(){
    $('#write-art-form').validate({
        rules:{
            title: "required",
            tags: "required",
            short: {
                required: true,
                minlength: 2
            },
            content: {
                required: true,
                minlength: 2
            }
        },
        messages:{
            title: "请输入标题",
            tags: "请输入标签",
            short: {
                required: "请输入简介",
                minlength: "简介内容最少两个字符"
            },
            content: {
                required: "请输入文章内容",
                minlength: "文章内容最少两个字符"
            }
        }
    })
}

$(document).ready(function(){
    registerValidate();
    loginValidate();
    articleValidate();
    var registerOptions = {
        url: "/register",
        type: "post",
        dataType: "json",
        success: function (data) {
            alert("data:" + data.message + " code:" + data.code)
            if (data.code == 1) {
                setTimeout(function () {
                    window.location.href = "/login"
                }, 1000)
            }
        },
        error: function (data, status) {
            alert("err:" + data.message + ":" + status)
        }
    }
    var loginOptions = {
        url: "/login",
        type: "post",
        dataType: "json",
        success: function (data) {
            alert("data:" + data.message + " code:" + data.code)
            if (data.code == 1) {
                setTimeout(function () {
                    window.location.href = "/"
                }, 1000)
            }
        },
        error: function (data, status) {
            alert("err:" + data.message + ":" + status)
        }
    }
    var articleOptions = {
        url: "/article/add",
        type: "post",
        dataType: "json",
        success: function (data) {
            alert("data:" + data.message + " code:" + data.code)
            if (data.code == 1) {
                setTimeout(function () {
                    window.location.href = "/"
                }, 1000)
            }
        },
        error: function (data, status) {
            alert("err:" + data.message + ":" + status)
        }
    }
    $('#register-form').ajaxForm(registerOptions);
    $('#login-form').ajaxForm(loginOptions);
    $('#write-art-form').ajaxForm(articleOptions);
}
);
