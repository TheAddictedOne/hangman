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

function wordToFind(word, letters) {
  if (!letters) {
    return "_".repeat(word.length).split("");
  }

  return word.split("").map((letter) => {
    return letters.filter((l) => l === letter).used ? letter : "_";
  });
}

function getAllLetters() {
  const first = 97; // ASCII for "a"
  const total = 26; // Number of letters
  const letters = [];

  for (let i = first; i < first + total; i++) {
    letters.push({
      value: String.fromCharCode(i),
      used: false,
    });
  }

  return letters;
}

module.exports = {
  getRandomWord,
  wordToFind,
  getAllLetters,
};
