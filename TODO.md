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
___
- [ ] add multi level states
- [ ] add obstacle to l2
- [ ] level 2

## Notes
#### March 24, 2026
Prioritizing other work today. Still need to add multi level states.

I think a refactoring is in order regarding the coordinates system which will affect:
- boundary logic
- tile sizing
- user-tile position

This will enhance the games ability to handle different level sizes more gracefully,
and will allow the UI to render reliably.
