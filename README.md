# Infinity Monkeys

This repo was suppose to be just a funny simple code to try an implementation of the [Infinite monkey theorem](https://en.wikipedia.org/wiki/Infinite_monkey_theorem). At the end, I learned a lot about go-routines and channels. 
A brief explanation about the packages:
- pubsub:
    this package implement a simple pubsub, the whole tutorial can be find [here](https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/)
- monkey:
    this package has three functions:
    - CharacterPicker: choose a random number between 32 and 5011(all the latin utf-8 simbols) and turn it into a string 
    - TrueInfinityMonkey: Only returns the word if it is found the whole sequence
    - EasyInfinityMonkey: Returns the word if found all the letters. The EasyInfinityMonkey function it's a lot faster, since it's not required to found randomly the whole sequence