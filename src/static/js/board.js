
var width = 700;
var size;
var moves;
var data;
var gameId;

var tentativeX;
var tentativeY;

$(function() {
	$("#board").click(boardClick);
	$("#board").mousemove(boardMouseover);
});


function setGameId(newGameId) {
	gameId = newGameId;
}

function freshDataCallback(freshData) {
	data = JSON.parse(freshData);
	size = data["Size"]
	moves = data["Moves"]
	repaintBoard();
}

function getFreshData() {
	$.get("/gamedata/" + gameId, freshDataCallback);
}

function boardClick(event) {
	var increment = width / (size + 1);
	var x = Math.round(event.offsetX / increment - 1);
	var y = Math.round(event.offsetY / increment - 1);

	$.post("/move/" + gameId + "/" + x + "/" + y, freshDataCallback);
}

function boardMouseover(event) {
	var increment = width / (size + 1);
	var x = Math.round(event.offsetX / increment - 1);
	var y = Math.round(event.offsetY / increment - 1);

	tentativeX = x;
	tentativeY = y;
	repaintBoard();
}

function repaintBoard() {
	var board = $("#board")[0];
	var ctx = board.getContext("2d");
	ctx.fillStyle = "#EDB809";
	ctx.fillRect(0, 0, width, width);

	var increment = width / (size + 1);
	for (var i = 0; i < size; i++) {
		ctx.beginPath();
		ctx.moveTo(increment, increment * (i + 1) );
		ctx.lineTo(increment * size, increment * (i + 1));
		ctx.stroke();

		ctx.beginPath();
		ctx.moveTo(increment * (i + 1), increment);
		ctx.lineTo(increment * (i + 1), increment * size);
		ctx.stroke();
	}


	for (var i = 0; i < moves.length; i++) {
		if (moves[i][0] == -1 && moves[i][1] == -1) {
			// This is how we signify a skipped move
			continue;
		}

		var x = (moves[i][0] + 1) * increment;
		var y = (moves[i][1] + 1) * increment;

		if (i % 2 == 0) {
			ctx.fillStyle = "#000000"
		}
		else {
			ctx.fillStyle = "#FFFFFF"
		}

		ctx.beginPath();
		ctx.arc(x, y, increment / 2, 0, 2 * Math.PI, false);
		ctx.fill();
	}

	// if (typeof possibleMove !== 'undefined') {
	// 	var x = possibleMove[0];
	// 	var y = possibleMove[1];
	// 	if (x >= 0 && y >= 0 && x < size && y < size) {
	//     	if (moves.length % 2 == 0) {
	// 			ctx.fillStyle = "rgba(255, 255, 255, 0.5)";
	// 		}
	// 		else {
	// 			ctx.fillStyle = "rgba(0, 0, 0, 0.5)";
	// 		}

	// 		var x = (possibleMove[0] + 1) * increment;
	// 		var y = (possibleMove[1] + 1) * increment;

	// 		ctx.beginPath();
	// 		ctx.arc(x, y, increment / 2, 0, 2 * Math.PI, false);
	// 		ctx.fill();
	// 	}
	// }
}
