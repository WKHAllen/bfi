const codeKey = 'code';

// On document ready
$(() => {
    var code = localStorage.getItem(codeKey);
    if (code !== null) {
        $('#code').text(code);
    }
});

// Save the code to local storage
function saveCode() {
    var code = $('#code').val();
    localStorage.setItem(codeKey, code);
}

// Run the code
function runCode() {
    // TODO: send code to server via AJAX
}
