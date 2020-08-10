const options = {
  buttonColorLight: '#14854f',
  label: '🌓',
  autoMatchOsTheme: true
}

const darkmode = new Darkmode(options);
darkmode.showWidget();

const CORS_PROXY = "https://cors-anywhere.herokuapp.com/"
let parser = new RSSParser();
let news = document.getElementById("news");

parser.parseURL(CORS_PROXY + `http://www.polytech21.ru/component/ninjarsssyndicator/?feed_id=1&format=raw`,
  (err, feed) => {
    if (err) throw err;

    news.innerHTML = "";

    for (let i = 0; i <= 4; i++) {
      if (i == 4) {
        news.innerHTML += `<a href="${feed.items[i].link}">${feed.items[i].title}</a><br>${feed.items[i].pubDate}`;
      } else {
        news.innerHTML += `<a href="${feed.items[i].link}">${feed.items[i].title}</a><br>${feed.items[i].pubDate}<hr>`;
      }
    }
});
