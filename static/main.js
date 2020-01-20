const codeKey = 'code';

const BFIDone = 0
const BFIInput = 1
const BFIOutput = 2

// On document ready populate the code textarea
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

function getSessionID() {
    return $('#sessionID').text();
}

function displayOutput(char) {
    $('#code-output').append(char);
}

// Run the code
function runCode() {
    $.ajax({
        url: '/interpret',
        type: 'GET',
        data: {
            sessionID: getSessionID(),
            code: $('#code').val()
        },
        dataType: 'json',
        success: (data) => {
            console.log(data);
            // TODO: check `data` to see why it returned
            displayOutput(data);
        },
        error: (err) => {
            // TODO: display the error
        }
    });
}
