
const WORLD_W = 50;
const WORLD_H = 50;
const WORLD_GAP = 2;
const WORLD_COLS = 16;
const WORLD_ROWS = 12;

var levelOne = [4, 4, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4,
                4, 4, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
                4, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
                1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1,
                1, 0, 0, 0, 1, 1, 1, 4, 4, 4, 4, 1, 1, 0, 0, 1,
                1, 0, 0, 1, 1, 0, 0, 1, 4, 4, 1, 1, 0, 0, 0, 1,
                1, 0, 0, 1, 0, 0, 0, 0, 1, 4, 1, 0, 0, 0, 0, 1,
                1, 0, 0, 1, 0, 3, 0, 0, 0, 1, 1, 0, 0, 5, 0, 1,
                1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1,
                1, 2, 0, 1, 0, 0, 5, 0, 0, 0, 5, 0, 0, 1, 0, 1,
                1, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1,
                1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1];

var   worldGrid = [];
const WORLD_ROAD = 0;
const WORLD_WALL = 1;
const WORLD_PLAYERSTART = 2;
const WORLD_GOAL = 3;
const WORLD_KEY = 4;
const WORLD_DOOR = 5;


function returnTileTypeAtColRow(col,row){
    if(col >= 0 && col < WORLD_COLS && row >= 0
        && row < WORLD_ROWS) {
        var worldIndexUnderCoord = rowColToArrayIndex(col, row);
        return worldGrid[worldIndexUnderCoord];
    } else {
        return WORLD_WALL;
    }
}

function getTileTypeAtPixelCoord(atX, atY){
    var warriorWorldCol = Math.floor(atX / WORLD_W);
    var warriorWorldRow = Math.floor(atY / WORLD_H);
    var worldIndexUnderWarrior = rowColToArrayIndex(warriorWorldCol, warriorWorldRow);

    if(warriorWorldCol >= 0 && warriorWorldCol < WORLD_COLS &&
        warriorWorldRow >= 0 && warriorWorldRow < WORLD_ROWS) {
        var tileHere = returnTileTypeAtColRow(warriorWorldCol,warriorWorldRow);

        return tileHere;
    } // end of valid col and row

    return WORLD_WALL; // treat outside of map as solid
} // end of getTileTypeAtPixelCoord func

function rowColToArrayIndex(col, row) {
    return col + WORLD_COLS * row;
}

function drawWorld() {
    var arrayIndex = 0;
    var drawTileX = 0;
    var drawTileY = 0;
    for(var eachRow=0;eachRow<WORLD_ROWS;eachRow++) {
        for (var eachCol = 0;eachCol<WORLD_COLS;eachCol++) {

            var arrayIndex = rowColToArrayIndex(eachCol, eachRow);
            var tileKindHere = worldGrid[arrayIndex];
            var useImg = worldPics[tileKindHere];

            canvasContext.drawImage(useImg, drawTileX, drawTileY);
            drawTileX += WORLD_W;
            arrayIndex++;
        } // end of for each col
        drawTileY += WORLD_H;
        drawTileX = 0;
    } // end of for each row
} // end of drawTrack func