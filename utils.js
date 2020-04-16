
// Method for get week of year
Date.prototype.getWeek = function () {
  const onejan = new Date(this.getFullYear(), 0, 1);
  return Math.ceil((((this - onejan) / 86400000) + onejan.getDay() + 1) / 7); // Magic!
}

// Day -> Name of the week
function numToDay(num) {
    const days = {
      0: "sunday",
      1: "monday",
      2: "tuesday",
      3: "wednesday",
      4: "thursday",
      5: "friday",
      6: "saturday"
    }

    return days[num];
}

// Get star of week
function getStarWeek(num) {
    if ((num % 2) == 0) {
      return "*"
    } else {
      return "**"
    }
}

// Get day and star, for a particular day
function getDayOfWeek(day) {
    const todayDate = new Date();

    let yesterdayDate = new Date(todayDate);
    yesterdayDate.setDate(yesterdayDate.getDate() - 1);

    let tomorrowDate = new Date(todayDate);
    tomorrowDate.setDate(tomorrowDate.getDate() + 1);

    if (day == "yesterday") {
      const d = yesterdayDate.getDay(),
            w = yesterdayDate.getWeek();

      return [getStarWeek(w), numToDay(d)];

    } else if (day == "today") {
      const d = todayDate.getDay(),
            w = yesterdayDate.getWeek();

      return [getStarWeek(w), numToDay(d)];

    } else {
      const d = tomorrowDate.getDay(),
            w = yesterdayDate.getWeek();

      return [getStarWeek(w), numToDay(d)];
    }
}

// Export
exports.getDayOfWeek = getDayOfWeek;
