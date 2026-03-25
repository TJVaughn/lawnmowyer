# TODO
- [x] render anything
- [x] add debug statement 
- [x] render a square
- [x] render a row with proper spacing 
- [x] render all rows 
- [x] register keypress, print debug msg on press 
- [x] keypress moves player
- [x] When user presses key, it only moves them 1 square
- [x] Add bounds to where they can move
- [x] as a user goes over a square, the color changes
- [x] if player goes over cut square - fail
- [x] game intro state
- [x] level end state - fail and success
- [x] if all the squares are changed, the level ends
- [x] level 1
- [x] add multi level states
- [x] add obstacle to l2
- [x] level 2
- [x] levels 3 - 8 
___

## Notes
#### March 24, 2026
Prioritizing other work today. Still need to add multi level states.

I think a refactoring is in order regarding the coordinates system which will affect:
- boundary logic
- tile sizing
- user-tile position

This will enhance the games ability to handle different level sizes more gracefully,
and will allow the UI to render reliably.

Update: spent some time tonight, yet to refactor tile sizing, but some refactoring was done to create the different level states.
I think the concept of this game has been fleshed out, and it was a good refresher in go. 
I could continue to add more challenging levels as it has become a memory game with hidden obstacles. 
Or perhaps moving pieces, but then you need collision logic. Adding art would make the game better.
At this point, this game would be difficult for a toddler, which was partly my initial goal. 
Simple game. Use go. Kid friendly.

Time to move on to the next project.
