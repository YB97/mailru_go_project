$('#registration_btn').click(function () {
    let login = $('#login').val();
    let password = $('#password').val();
    $.post("/registration/?userData="+JSON.stringify({"LOGIN": login, "PASSWORD": password}), function (response) {
        console.log(response);
        window.location.replace("/recognition/");
    })
});

$('#login_btn').click(function () {
    let login = $('#login').val();
    let password = $('#password').val();
    $.post("/login/?userData="+JSON.stringify({"LOGIN": login, "PASSWORD": password}), function (response) {
        window.location.replace("/recognition/");

    })
});