// Express
const express = require("express");
const app = express();
const port = process.env.PORT || 3000;

// Utils
const utils = require("./utils");

// DB
const db = require("./schedule");

// Templates
app.set('views', './views');
app.set('view engine', 'pug');

// Folder for static files
app.use(express.static('img'));
app.use(express.static('js'));
app.use(express.static('css'));

// UI
app.get("/", (req, res) => {
    let [starOfWeak, dayOfWeek] = utils.getDayOfWeek("yesterday");
    const group1yesterday = db["group1"][dayOfWeek][starOfWeak],
          group2yesterday = db["group2"][dayOfWeek][starOfWeak];

    [starOfWeak, dayOfWeek] = utils.getDayOfWeek("today");
    const group1today = db["group1"][dayOfWeek][starOfWeak],
          group2today = db["group2"][dayOfWeek][starOfWeak];

    [starOfWeak, dayOfWeek] = utils.getDayOfWeek("tomorrow");
    const group1tomorrow = db["group1"][dayOfWeek][starOfWeak],
          group2tomorrow = db["group2"][dayOfWeek][starOfWeak];

    res.render("index", {
	    group1yesterday: group1yesterday,
	    group1today: group1today,
	    group1tomorrow: group1tomorrow,
	    group2yesterday: group2yesterday,
	    group2today: group2today,
	    group2tomorrow: group2tomorrow,
    });
});

// Get all db json
app.get("/db", (req, res) => {
    res.json(db);
});

// Get schedule times, json or text
app.get("/time", (req, res) => {
    if (req.query.type == "json") {
      res.json(db.time);
    } else {
      let text = "";

      Object.entries(db.time).forEach(([key, value]) => {
        text += `${key}: ${value}\n`;
      });

      res.send(text);
    }
});

// Get lessons for group and week, json or text
app.get("/s/", (req, res) => {
    const group = `group${req.query.group}`,
          day = req.query.day,
          type = req.query.type;

    if (group && day && type) {
      if (!(["group1", "group2"].includes(group)))
        res.send("Error");

      if (!(["yesterday", "today", "tomorrow"].includes(day)))
        res.send("Error");

	    const [starOfWeak, dayOfWeek] = utils.getDayOfWeek(day);

      if (type == "json") {
        res.json(db[group][dayOfWeek][starOfWeak]);
      } else {
        let text = "";

        Object.entries(db[group][dayOfWeek][starOfWeak]).forEach(([key, value]) => {
          text += `${key}: ${value}\n`;
        });

        res.send(text);
      }
    } else {
      res.send("Error");
    }
});

// Start server
app.listen(port, () => console.log(`localhost:${port}`));
