
var width = 500;

function boardClick() {
	console.log("CLICKED");
}

function repaintBoard(size, moves) {
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


	for (var i = 0; i < moves.length; i++) {
		console.log(moves[i]);
		var x = (moves[i][0] + 1) * increment;
		var y = (moves[i][1] + 1) * increment;

		if (i % 2 == 0) {
			ctx.fillStyle = "#FFFFFF"
		}
		else {
			ctx.fillStyle = "#000000"
		}

		ctx.beginPath();
		ctx.arc(x, y, increment / 2, 0, 2 * Math.PI, false);
		ctx.fill();
	}
}
