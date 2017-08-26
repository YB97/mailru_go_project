function showModal(stat) {
    if (stat){
        $('#suc_modal').modal('show')
    }else{
        $('#fail_modal').modal('show');
    }
}

$('#reg_btn').click(function () {
    let login = $('#login').val();
    let password = $('#password').val();
    $.post("/reg/", {"login": login, "password": password}, function (response, status) {
        if(status == "success") {
            showModal(true);
            window.location.replace("/");
        } else{
            console.log(response);
            showModal(false);
        }

    })
});