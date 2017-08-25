$('#fileupload').fileupload({
    // Send cross-domain cookies
    xhrFields: {withCredentials: true},
   // url: 'http://localhost:8080/recognition/',
    // Chunk size in bytes
    maxChunkSize: 1000000,
    // Enable file resume
    add: function (e, data) {
        let that = this;
        $.ajax({
            url: 'http://localhost:8080/recognition/load_file/',
            xhrFields: {withCredentials: true},
            data: {file: data.files[0].name}
        }).done(function(result) {
            let file = result.file;
            data.uploadedBytes = file && file.size;
            $.blueimp.fileupload.prototype.options.add.call(that, e, data);

        });

    }

});