const express = require("express");
const bodyParser = require("body-parser");
const nunjucks = require("nunjucks");

const { getRandomWord, wordToFind, getAllLetters } = require("./utils.js");

const app = express();
const port = 3000;

const state = {
  level: "",
  word: "",
  wordToFind: "",
  letters: getAllLetters(),
  errors: 0,
};

nunjucks.configure("views", {
  autoescape: true,
  express: app,
});

app.use(express.static("public"));
app.use(bodyParser.urlencoded({ extended: true }));

app.get("/", function (req, res) {
  res.render("index.html");
});

app.get("/game", function (req, res) {
  const { level } = req.query;

  state.level = level;
  getRandomWord(level, (word) => {
    state.word = word;
    state.wordToFind = wordToFind(word);
    res.render("game.html", { ...state });
  });
});

app.post("/game", function (req, res) {
  const { letter } = req.body;

  if (!state.word.includes(letter)) {
    state.errors++;
  }

  state.letters.find((l) => l === letter).used = true;

  console.log("new:", state.letters);

  res.render("game.html", {
    ...state,
    wordToFind: wordToFind(state.word, state.letters),
  });
});

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});
