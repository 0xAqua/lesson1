document.getElementById('upload-form').addEventListener('submit', uploadFile);

const uploadFile = (event) => {
    event.preventDefault();

    var form = document.getElementById('upload-form');
    var formData = new FormData(form);
    toggleLoading(true);

    fetch('/upload', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        toggleLoading(false);
        if(data.error) {
            alert(data.error);
        } else {
            document.getElementById('original-content').innerHTML = data.original_content;
            document.getElementById('specification').innerHTML = data.specification;

            var downloadLink = document.getElementById('download-link');
            downloadLink.href = data.download_link;
            downloadLink.hidden = false;
        }
    })
    .catch(error => {
        toggleLoading(false);
        alert('Error: ' + error);
    });
}

const toggleLoading = (isLoading) => {
    var loadingElement = document.getElementById('loading');
    if (isLoading) {
        loadingElement.hidden = false;
    } else {
        loadingElement.hidden = true;
    }
}
