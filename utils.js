const fs = require("fs");

function getRandomWord(level, callback) {
  let filename = process.cwd() + `/files/${level}.txt`;

  console.log(filename);

  fs.readFile(filename, (err, data) => {
    const words = data.toString().split("\n");
    const num = Math.floor(Math.random() * words.length);
    callback(words[num]);
  });
}

module.exports = {
  getRandomWord,
};
