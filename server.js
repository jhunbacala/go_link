const express = require('express');
const fs = require('fs');
const path = require('path');

const app = express();
const PORT = process.env.PORT || 8080;
const URLS_FILE = path.join(__dirname, 'urls.json');

// Load URLs from JSON
let urls = {};

function loadUrls() {
  try {
    const data = fs.readFileSync(URLS_FILE, 'utf8');
    urls = JSON.parse(data).shorts || {};
    console.log('Loaded URLs:', Object.keys(urls));
  } catch (err) {
    console.error('Error loading urls.json:', err);
    urls = {};
  }
}

loadUrls();

// Reload on file change (optional for dev)
fs.watchFile(URLS_FILE, (curr, prev) => {
  if (curr.mtime !== prev.mtime) {
    console.log('urls.json changed, reloading...');
    loadUrls();
  }
});

// Redirect route
app.get('/:short', (req, res) => {
  const short = req.params.short;
  const originalUrl = urls[short];

  if (originalUrl) {
    console.log(`Redirecting ${short} -> ${originalUrl}`);
    return res.redirect(301, originalUrl);
  }

  res.status(404).send('Short link not found');
});

// Home page
app.get('/', (req, res) => {
  res.send(`
    <h1>Go Link - Simple URL Shortener</h1>
    <p>Edit <code>urls.json</code> to add more short links.</p>
    <p>Example: <a href="/gh">/gh</a></p>
  `);
});

app.listen(PORT, () => {
  console.log(`🚀 Server running on http://localhost:${PORT}`);
});
