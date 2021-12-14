const express = require("express");
const nunjucks = require("nunjucks");

const { getRandomWord } = require("./utils.js");

const app = express();
const port = 3000;

const global = {
  level: "",
  word: "",
};

nunjucks.configure("views", {
  autoescape: true,
  express: app,
});

app.get("/", function (req, res) {
  res.render("index.html");
});

app.get("/game", function (req, res) {
  global.level = req.query.level;
  getRandomWord(req.query.level, (word) => {
    global.word = word;
    res.render("game.html", { ...global });
  });
});

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});
