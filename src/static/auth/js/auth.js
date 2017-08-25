$('#reg_btn').click(function () {
    window.location.replace("/registration/")
});

function showModal(stat) {
    if (stat){
        $('#suc_modal').modal('show')
    }else{
        $('#fail_modal').modal('show');
    }
}

$('#login_btn').click(function () {
    let login = $('#login').val();
    let password = $('#password').val();
    $.post("/log/?userData="+JSON.stringify({"LOGIN": login, "PASSWORD": password}), function (response) {

        if (status == "success"){
            showModal(true);
            window.location.replace("/recognition/");
        } else {
            showModal(false)
        }
    })
});