require('dotenv').config();

const express = require('express');
const bodyParser = require('body-parser');
const logger = require('morgan');
const nocache = require('nocache');
const path = require('path');

// Envirnments
const envProp = process.env;
console.log(process.env)
const port = envProp.PORT || 8080;

const app = express();

// Diable x-powered
app.disable('x-powered-by');

// Serving Static folder
app.use(express.static(path.join(__dirname, 'Public')));

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.use(nocache());

const checkAuthentication = (req, res, next) => {
  // Placehoder for Auth
  next();
};

app.get('/getEnv', (req, res) => {
  const properties = {
    baseUrl: process.env.baseUrl,
    ATLAS_URI:process.env.ATLAS_URI
  };
  res.json(properties);
});

app.get('/', checkAuthentication, (req, res) => {
  res.render(path.join(__dirname, 'Public', 'index'));
});

app.get('/*', (req, res) => {
  res.sendFile(__dirname + '/Public/index.html');
});

// Start listening on defined port
const server = app.listen(port, () => {
  console.log(`APP is running on port ${server.address().port}`);
});
