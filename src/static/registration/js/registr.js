$('#reg_btn').click(function () {
    let login = $('#login').val();
    let password = $('#password').val();
    $.post("/reg/?userData="+JSON.stringify({"login": login, "password": password}), function (response, status) {

        window.location.replace("/login/");


    })
});