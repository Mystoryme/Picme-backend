var express = require('express');
var router = express.Router();

/* GET users listing. */
router.get('/', function(req, res, next) {
  res.send('Hello Pic me');
});

router.get("/posts", function (req, res, next) {
  res.send("Post");
});

module.exports = router;
