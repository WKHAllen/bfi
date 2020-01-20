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

// Set the status text
function setStatus(status) {
    $('#status').text(status);
}

// Disable an element
function disableElement(selector) {
    $(selector).prop('disabled', true);
    $(selector).addClass('disabled');
}

// Enable an element
function enableElement(selector) {
    $(selector).prop('disabled', false);
    $(selector).removeClass('disabled');
}

// Get the server-provided session ID
function getSessionID() {
    return $('#sessionID').text();
}

// Add to the output area
function displayOutput(char) {
    $('#code-output').append(char);
}

// Highlight a section of the code
function highlightCode(start, end) {
    var codeBox = document.getElementById('code');
    codeBox.focus();
    codeBox.setSelectionRange(start, end + 1);
}

// Run the code
function runCode() {
    disableElement('#run-button');
    setStatus('Status: running');
    $.ajax({
        url: '/interpret',
        type: 'GET',
        data: {
            sessionID: getSessionID(),
            code: $('#code').val()
        },
        dataType: 'json',
        success: (data) => {
            if (data.error) {
                setStatus(`Error: ${data.error}, Index: ${data.index}`);
                highlightCode(data.index, data.index);
                enableElement('#run-button');
            } else {
                // TODO: check which code was returned
            }
        },
        error: (err) => {
            setStatus(`Unexpected error: ${err}`);
            enableElement('#run-button');
        }
    });
}
