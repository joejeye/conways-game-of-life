# Conway's Game of Life

A simple implementation of the [Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life).

## Build (Windows Only)

Open a Powershell command prompt and enter the following command
```ps1
.\build.ps1
```

## Quick Start

Run the following command in a Powershell prompt
```ps1
.\run.ps1 30 50 600
```
This command creates a 5-by-20 2D world, and initialized the world with 30 live cells randomly scattered within. The 2D world is updated evey 125 ms. And the output is something like the following

![Demo](./media/render1715062751484.gif)

## Usage

```ps1
.\run.ps1 [NumRows [NumCols [NumInitLiveCells]]]
```

where 
- `NumRows` is the number of rows in the 2D world
- `NumCols` is the number of columns in the 2D world
- `NumInitLiveCells` is the number of live cells randomly scattered in the initial 2D world