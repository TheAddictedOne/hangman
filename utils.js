const fs = require("fs");

function getRandomWord(level, callback) {
  let filename = process.cwd() + `/files/${level}.txt`;

  console.log(filename);

  fs.readFile(filename, (err, data) => {
    console.log("reading file...");
    console.log(data);
    callback(data.toString().split("\n"));
  });
}

module.exports = {
  getRandomWord,
};
