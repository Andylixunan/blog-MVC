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

function uploadFile(){
    const filedata = $("#album-upload-file").val();
    if (filedata.length <= 0) {
        alert("请选择文件!");
        return
    }
    const formdata = new FormData()
    formdata.append("upload", $("#album-upload-file")[0].files[0]) // formdata stores form data in key-value pairs
    const uploadOptions = {
        url: "/album",
        type: "POST",
        data: formdata,
        processData: false,
        contentType: false,
        dataType: "json", // the server response data type
        success: function(data) {
            alert("data:" + data.message + " code:" + data.code)
            if (data.code == 1) {
                setTimeout(function () {
                    window.location.href = "/album"
                }, 1000)
            }
        },
        error: function (data, status) {
            alert("err:" + data.message + ":" + status)
        }
    }
    $.ajax(uploadOptions)
}

$(document).ready(function(){
    registerValidate();
    loginValidate();
    articleValidate();
    const registerOptions = {
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
    const loginOptions = {
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
    const articleOptions = {
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
    const articleID = $("#write-article-id").val()
    if (articleID > 0) {
        const urlStr = "/article/update/" + articleID
        articleOptions.url = urlStr
    }
    $('#register-form').ajaxForm(registerOptions);
    $('#login-form').ajaxForm(loginOptions);
    $('#write-art-form').ajaxForm(articleOptions);
    $('#album-upload-button').click(uploadFile)
}
);
