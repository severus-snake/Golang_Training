function fib(n){

    var array = [n];

    if(array[n] == 0){
        array.push(1);
    } else {
        array.push(n);
    }

    for(i = 0; i <= 10; i++){
        var result = array[i] + array[i + 1];
        array.push(result);
        console.log(array);
    }
    return array
}

fib(1);

// this code works for both 0 && 1 and 1 && 1