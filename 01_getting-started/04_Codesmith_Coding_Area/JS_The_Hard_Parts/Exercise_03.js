function outer() {
    this.counter = 0; // this variable is outside incrementCounter's scope
    this.incrementCounter = function(){
        this.counter ++;
        this.console.log('counter', this.counter);
    }
    return this.incrementCounter;
}

var willCounter = new outer();
var jasCounter = new outer();

willCounter.incrementCounter();
willCounter.incrementCounter();
jasCounter.incrementCounter();
jasCounter.incrementCounter();