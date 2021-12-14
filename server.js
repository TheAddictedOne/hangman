const express = require("express");
const bodyParser = require("body-parser");
const nunjucks = require("nunjucks");

const { getRandomWord } = require("./utils.js");

const app = express();
const port = 3000;

const state = {
  level: "",
  word: "",
  letters: [],
};

nunjucks.configure("views", {
  autoescape: true,
  express: app,
});

app.use(bodyParser.urlencoded({ extended: true }));

app.get("/", function (req, res) {
  res.render("index.html");
});

app.get("/game", function (req, res) {
  state.level = req.query.level;
  getRandomWord(req.query.level, (word) => {
    state.word = word;
    res.render("game.html", { ...state });
  });
});

app.post("/game", function (req, res) {
  console.log("checking ", req.body.letter, " in ", state.letters);
  if (!state.letters.includes(req.body.letter)) {
    state.letters.push(req.body.letter);
  }
  res.render("game.html", { ...state });
});

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});
