const express = require('express');
const { exec } = require('child_process');

const app = express();
const port = process.env.PORT || 3000;

app.get('/', (req, res) => {
  res.json({
    message: 'GitHub CLI Sample App',
    endpoints: {
      '/version': 'Get GitHub CLI version',
      '/repos': 'List repositories',
      '/issues': 'List issues'
    }
  });
});

app.get('/version', (req, res) => {
  exec('gh --version', (error, stdout, stderr) => {
    if (error) {
      res.status(500).json({ error: error.message });
      return;
    }
    res.json({ version: stdout.trim() });
  });
});

app.get('/repos', (req, res) => {
  exec('gh repo list --json name,description,url --limit 5', (error, stdout, stderr) => {
    if (error) {
      res.status(500).json({ error: error.message });
      return;
    }
    try {
      const repos = JSON.parse(stdout);
      res.json({ repositories: repos });
    } catch (e) {
      res.status(500).json({ error: 'Failed to parse repository data' });
    }
  });
});

app.get('/issues', (req, res) => {
  exec('gh issue list --json title,number,state --limit 5', (error, stdout, stderr) => {
    if (error) {
      res.status(500).json({ error: error.message });
      return;
    }
    try {
      const issues = JSON.parse(stdout);
      res.json({ issues: issues });
    } catch (e) {
      res.status(500).json({ error: 'Failed to parse issue data' });
    }
  });
});

app.listen(port, () => {
  console.log(`GitHub CLI Sample App listening at http://localhost:${port}`);
}); 