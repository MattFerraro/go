
var size = 13;
var width = 500;

function boardClick() {
	console.log("CLICKED");
}

function repaintBoard() {
	console.log("loaded");
	var board = $("#board")[0];
	var ctx = board.getContext("2d");
	ctx.fillStyle = "#EDB809";
	ctx.fillRect(0, 0, 500, 500);

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
}
