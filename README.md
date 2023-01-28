# Taxi

A common reinforcement learning environment is the taxi environment. A description of it is located here: https://gymnasium.farama.org/environments/toy_text/taxi/

It is a fairly straightforward environment and problem. The goal of this repository is to do something like this:

- implement the "game" via TinyGo, and ensure it is somewhat playable from CLI
- using the same implementation make it "playable" in a browser

```
+---------+
|R: | : :G|
| : | : : |
| : : : : |
| | : | : |
|Y| : |B: |
+---------+
```


We'll display the state of the game using openmoji (to avoid using sprites), and we'll use css grids to display them

- taxi: https://openmoji.org/library/emoji-1F696/
- passenger: https://openmoji.org/library/emoji-1F64B/
- red stop: https://openmoji.org/library/emoji-26E9/ or https://openmoji.org/library/emoji-1F7E5/
- green stop: https://openmoji.org/library/emoji-1F3DE or https://openmoji.org/library/emoji-1F7E9/
- yellow stop: https://openmoji.org/library/emoji-1F3D6 or https://openmoji.org/library/emoji-1F7E8/
- blue stop: https://openmoji.org/library/emoji-1F3A1/ or https://openmoji.org/library/emoji-1F7E6/
- tree: https://openmoji.org/library/emoji-1F333/

The game has a finite number of states being:

- location of taxi (25 locations, represented by `[][]int32`)
- destination (4 states, being `int32` representing which color stop, being `{0: red, 1: green, 2: yellow, 3: blue}`)
- location of passenger (same as destination, with `{4: in tax}`)

We will simplify the action space to movement only and assume the taxi will automatically pickup and drop off the passenger at the correct time

- 0: move down
- 1: move up
- 2: move right
- 3: move left

For simplicity, we'll implement the game purely using CLI, where we input the current state, and choose to show game board only, or to perform an action. The rough usage would be:

```sh
go run main.go -taxi-x 0 -taxi-y 1 -passenger 0 -destination 1
go run main.go -taxi-x 0 -taxi-y 1 -passenger 0 -destination 1 -act 0
go run main.go -taxi-x 0 -taxi-y 3 -act 1 -passenger 3 -goal 0
```

Example

```sh
$ go run main.go -taxi-x 0 -taxi-y 3 -act 1 -passenger 3 -goal 0
+---------+             
|R: | : :G|             
| : | : : |             
| : : : : |             
|T| : | : |             
|Y| : |B: |             
+---------+             
                        
~Taxi Move~             
+---------+             
|R: | : :G|             
| : | : : |             
|T: : : : |             
| | : | : |             
|Y| : |B: |             
+---------+             
Passegner location: B
```

