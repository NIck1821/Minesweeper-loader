# Minesweeper-loader

Дан файл с решеткой N x N (разделители пробел и \n), который представляет собой
минное поле из игры Сапёр (Minesweeper). O - означает что на поле нет мины, X -
что есть. Нужно заменить все O на количество смежных мин (всего смежных клеток
8 - по горизонтали, вертикали и диагонали).

Пример:

Входной файл:

```
O O O X O X O O
O O X O X O O X
X X O X X O X O
O X O X O X X X
O O X O X O O X
X O O O X O X O
O O O X O X O X
X O X X O X O X
```

Вывод в консоль:

```
0 1 2 X 3 X 2 1
2 3 X 5 X 4 3 X
X X 5 X X 5 X 4
3 X 5 X 5 X X X
2 3 X 4 X 5 5 X
X 2 2 4 X 4 X 3
2 3 3 X 5 X 5 X
X 2 X X 4 X 4 X
```
