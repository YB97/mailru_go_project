$('#reg_btn').click(function () {
    let login = $('#login').val();
    let password = $('#password').val();
    $.post("/reg/?userData="+JSON.stringify({"login": login, "password": password}), function (response, status) {
        if(status === "success") {
            window.location.replace("/");
        } else{
            $.error("User already exist");
        }

    })
});