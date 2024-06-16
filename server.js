const fs = require('fs');   
const express = require("express");
const bodyParser = require("body-parser");
const { exec } = require('child_process');
const path = require("path");
const __path = path.resolve("public");
const app = express();

const goScript = path.join(__dirname, 'ma.go');
const command = `go run ${goScript}`;


const PORT = 8000;

app.use(express.text());
// app.use(bodyParser.urlencoded({extended:true}))

app.get("/", (req, res) => {
  res.sendFile(`${__path}/index.html`);
});

app.post("/submit", (req, res) => {
  console.log(req.body);
  const data = req.body
  const filename = 'test.txt'
    fs.writeFile(filename, data, 'utf8', (err) => {
        if (err) {
            console.error('Error writing to file', err);
        } else {
            console.log('Data successfully written to file');
        }
    });
    exec(command, (error, stdout, stderr) => {
        if (error) {
            console.error(`Error executing Go script: ${error.message}`);
            return;
        }
        if (stderr) {
            console.error(`Go script stderr: ${stderr}`);
            return;
        }
        console.log(`Go script output: ${stdout}`);
    });

  res.send(200);
});
app.post("/logs", (req, res) => {
  const logs = { 34: "23904823049x32094", 328: "20394f3492" };
  res.send(JSON.stringify(logs));
});

app.use(express.static(__path));

app.listen(PORT, () => {
  console.log(`LISTENING ON ${PORT}`);
});
