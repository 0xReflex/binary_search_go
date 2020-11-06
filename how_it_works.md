# How it works???

## It works on a simple phenomena
## like~~ 
*You have 10 books and you want the 5th book. How will you find it
There are 2 possible ways*
1. You start from 1,2,3,4 and then you willreach 5 **It will take you 5 moves``And
	thats a lot ``**
2. You can just
	Check the middel term like in this it is 5 so you got it in just 1 move
	but if you want the "" 7 ""th book
  	 ________________________________________
	| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 |
	 ---------------------------------------- 
	  |										|
	  L 									R 
	.0 Take two terms one ""left"" as 0 starting of the list and other ""right"" at the end of the array

	.1 Check the middle term in this example it is 5

	.2 Here 5 is not is equal to 7 so put left to middle that is 5
  	 ________________________________________
	| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 |
	 ---------------------------------------- 
	 				  |				       |
	   				  L		        	   R 

	.3 Now find the middle again like Left + Right/2 == 5 + 10 /2 = 7.5 
	round 7.5 = 7
	 _______________________________________
	| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 |
	 ---------------------------------------- 
	 				  |		  |	           |
	   				  L		  M      	   R

	.4 Check if the middle term is the term we wanted if yes exit else continue

	.5 **Boom !!!!** We found the 7th book with only 2 moves and if it is a million more books then it is very useful because you can't count it one by one.  
