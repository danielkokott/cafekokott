var express = require("express");
var app = express();
app.use(express.logger());

app.use("/static", express.static("./static"));
app.use("/bootstrap", express.static("./bootstrap"));

app.get('/', function(request, response) {
	response.sendfile('./web/workinprogress.html');
});

app.get('/test', function(request, response) {
	response.sendfile('./web/index.html');
});

var port = process.env.PORT || 5000;
	app.listen(port, function() {
});