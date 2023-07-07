# TETRIS-OPTIMISER

## INTRODUCTION

In this project we have to create a program that receives only one argument, a path to a text file which will contain a list of [tetrominoes](https://en.wikipedia.org/wiki/Tetromino) and assemble them in order to create the smallest square possible.

## ALGO
We extract data from text file with *ioutil.ReadFile*.
Verify if :
-After 4 lines we have newline,
-number of dot = 4*numberOfDiez,
-Each diez is linked to at least one diez.
Resolution :
We'll get all tetrominoes indexs and put them in double int tab with corresponding charactere(A=65 to Z=90).
We'll adapt these indexs to the finalTab(tab we'll print a the end), and try to place each tetrominoe and the next tetrominoe at the same time if can't a tetrominoe and the next we delete the placed tetrominoe...
## USAGE
 ```shell
go run . sample.txt
```
## License & Copyright
**abdouladieng**  