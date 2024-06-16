document.getElementById('journalForm').addEventListener('submit', function(event) {
    event.preventDefault();

    const journalEntry = document.getElementById('journalEntry').value;

    let logs = JSON.parse(localStorage.getItem('journalLogs')) || [];
    logs.push(journalEntry);
    localStorage.setItem('journalLogs', JSON.stringify(logs));

    console.log(journalEntry)

    
    
    document.getElementById('journalEntry').value = '';
    
    const successMessage = document.getElementById('successMessage');
    successMessage.style.display = 'flex';
    
    setTimeout(() => {
        successMessage.style.display = 'none';
        location.reload();
    }, 2000);

    const data = {log:journalEntry}
    console.log(data)
    fetch('/submit', {
        method: 'POST',
        headers: {
            'Content-Type': 'text/plain'
        },
        body: journalEntry
    })
});

document.getElementById('viewLogs').addEventListener('click', function() {
    const logsDiv = document.getElementById('logs');
    const closeButton = document.getElementById('closeLogs');

    // Toggle logs visibility
    if (logsDiv.style.display === 'block') {
        logsDiv.style.animation = 'slideOut 0.5s ease-in-out';

        setTimeout(() => {
            logsDiv.style.display = 'none';
            logsDiv.style.animation = '';
        }, 500);

        // Hide close button
        closeButton.style.display = 'none';
    } else {
        logsDiv.innerHTML = '';
        let logs = JSON.parse(localStorage.getItem('journalLogs')) || [];

        if (logs.length === 0) {
            logsDiv.innerHTML = '<p>No journal entries found.</p>';
        } else {
            logs.forEach(log => {
                const logEntry = document.createElement('div');
                logEntry.className = 'log-entry';
                logEntry.textContent = log;
                logsDiv.appendChild(logEntry);
            });
        }

        logsDiv.style.display = 'block';

        // Show close button
        closeButton.style.display = 'block';
    }
});

document.getElementById('closeLogs').addEventListener('click', function() {
    const logsDiv = document.getElementById('logs');
    logsDiv.style.animation = 'slideOut 0.5s ease-in-out';

    setTimeout(() => {
        logsDiv.style.display = 'none';
        logsDiv.style.animation = '';
    }, 500);

    // Hide close button
    this.style.display = 'none';
});


const submitBttn = document.getElementById('submission')

submitBttn.addEventListener("onclick", () =>{
    

})