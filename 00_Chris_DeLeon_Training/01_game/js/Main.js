    var canvas, canvasContext;

    var blueWarrior = new warriorClass();

    window.onload = function() {
        canvas = document.getElementById('gameCanvas');
        canvasContext = canvas.getContext('2d');

        colorRect(0,0, canvas.width, canvas.height, 'black');
        colorText("LOADING IMAGES", canvas.width/2, canvas.height/2, 'white');

        loadImages();
    }

    function imageLoadingDoneSoStartGame() {

        var framesPerSecond = 30;
        setInterval(updateAll, 1000/framesPerSecond);

        setupInput();

        loadLevel(levelOne);
    }

    function loadLevel(whichLevel) {
        worldGrid = whichLevel.slice();
        blueWarrior.reset(warriorPic, "Blue Warrior");
    }

    function updateAll() {
        moveAll();
        drawAll();
    }

    function moveAll() {
        blueWarrior.move();

    }

    function drawAll() {
        drawWorld();
        blueWarrior.draw();
    }

    var myArray = ["Eric Andrade", "e_andrade_21@hotmail.com"];

    var myData = {fullName: " ", skype: " ", github:" "};

    function cutName(str){
        var newArray;
        newArray = str.split(" ");
        console.log(newArray);
        return newArray;
    }
    myData.fullName = cutName(myArray[0]);
    myData.skype = myArray[1];
    console.log(myData);

