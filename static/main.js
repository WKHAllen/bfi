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
    $('#code-output').append(String.fromCharCode(char));
}

// Highlight a section of the code
function highlightCode(start, end) {
    var codeBox = document.getElementById('code');
    codeBox.focus();
    codeBox.setSelectionRange(start, end + 1);
}

// Display an error
function showError(err) {
    setStatus('Unexpected error, see console for details');
    console.log(err);
    enableElement('#run-button');
}

// Run the code
function runCode() {
    disableElement('#run-button');
    $('#code-output').text('');
    setStatus('Status: running');
    $.ajax({
        url: '/interpret',
        type: 'GET',
        data: {
            sessionID: getSessionID(),
            code: $('#code').val()
        },
        dataType: 'json',
        success: mainReturn,
        error: showError
    });
}

// Handle data returned from initial AJAX call
function mainReturn(data) {
    if (data.error) {
        // Error or timeout
        if (data.index !== undefined) {
            setStatus(`Error: ${data.error}, Index: ${data.index}`);
            highlightCode(data.index, data.index);
        } else {
            setStatus(`Error: ${data.error}`);
        }
        enableElement('#run-button');
    } else {
        // Successful return
        if (data.returnCode == BFIDone) {
            // Done
            setStatus('Status: complete');
            enableElement('#run-button');
        }
        if (data.returnCode == BFIInput) {
            // Wait for input
            // TODO: wait for input
        } else if (data.returnCode == BFIOutput) {
            // Display the character and continue
            displayOutput(data.displayByte);
            $.ajax({
                url: '/returnOutput',
                type: 'GET',
                data: {
                    sessionID: getSessionID()
                },
                dataType: 'json',
                success: mainReturn,
                error: showError
            });
        }
    }
}
