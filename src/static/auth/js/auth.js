$('#reg_btn').click(function () {
    window.location.replace("/registration/")
});

$('#login_btn').click(function () {
    let login = $('#login').val();
    let password = $('#password').val();
    $.post("/log/?userData="+JSON.stringify({"LOGIN": login, "PASSWORD": password}), function (response) {

        window.location.replace("/recognition/");

    })
});