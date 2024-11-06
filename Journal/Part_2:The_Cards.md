# The Goal
This part of the project was two fold. Make a presntable card layout, 
and make it ready to recieve database information when requested. 


## General Design
The card is relatively simple, but here is the gist. I want to have 25 of
the leagues best players from worlds on display, we are talking Caps, Faker, Knight, etc. These players will all have primary stats. To be clear these stats will be found in the next steps, we are still only in front-end world here. 

What I will then want to do, is using a card template have the user be able to search different types of queries via the cards. Let's say for example you want to see the win/loss ratio between Caps and Faker. Well you should be able to select Faker and see his stats. Then you should be able to select Caps card and a new screen showing a comparison of stats should occur. But I digress, this is getting ahead of ourselfs, for now we need the meat and potatoes. The cards!

So, the cards have within themselves dummy data, again, no DB no data. However you, the user should be able to select the card, and on hovering over the card you should be able to see the players general stats, (this will be in my db and loaded when my function for players runs. 

The end result of this chapter is cards that can spawn in by a JS function and then be interacted with from the user. 

### Example!
![Current Demo](Pics/output3.gif)

## Testing 
Ahh yes Testing in the Frontend...well for now, there is still no tests, I now have the card template for which I want to make cards with and find information. That stated, the next section will indeed have a filter bar where you can query how many players you want to view. For now, just trust me, the tests will come.

## Problems
The most obvious hurdle here is all the CSS. Because let me tell you it can be challenging knowing 0 CSS and going full on. However, chatgpt has defintely helped sort out a lot of the problems, though I am sure I will have to go back and remove some fields that are now worthless from the constant tinkering to make my particular boxes show appropriately.

# Epilogue
Finally...enough Fronend to do Backend, praise be! The next steps will be Setting up the DB making queries and testing things with the front end functionality, although I may need to add an event handler to make sure the cards can be input. Currently manually doing them in console is a bit much for my taste.


# TODOS
- [ ] Set up DB
- [ ] Make DB Queries
- [ ] Find Data 
- [ ] Make Testing Suites 
- [ ] Make DashBoard
- [ ] Make Landing Page
- [ ] Make Login/Auth
