var showdown  = require('showdown');
const { readFileSync } = require('fs');
const { join } = require('path');

module.exports = (req, res) => {
  let styleData = readFileSync(join(__dirname, 'style.css')).toString('utf8');

  converter = new showdown.Converter({
    ghCompatibleHeaderId: true,
    simpleLineBreaks: true,
    ghMentions: true,
  });

  let preContent = `
  <html>
    <head>
      <title> MarkdownWebNow </title>
      <meta name="viewport" content="width=device-width, initial-scale=1">
    </head>
    <body>
      <div id='content'>
  `

  let postContent = `

      </div>
      <style type='text/css'>` + styleData + `</style>
    </body>
  </html>`;

  html = preContent + converter.makeHtml(req.body) + postContent

  converter.setFlavor('github');
  res.status(200).send(html)
}
