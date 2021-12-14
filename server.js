const express = require("express");
const nunjucks = require("nunjucks");

const { getRandomWord } = require("./utils.js");

const app = express();
const port = 3000;

const global = {
  level: "",
  words: [],
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
  getRandomWord(req.query.level, (words) => {
    global.words = words;
    res.render("game.html", { level: global.level, words: global.words });
  });
});

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});
