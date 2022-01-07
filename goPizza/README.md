# goPizza

**goPizza** is a Pizzeria that contains a golang-workerpool to create Pizza.
There are several tasks for different workers that each work concurrent.
However, each step is depending on another task, which is why the workers have
to communicate with eachother.

## Workers

1. Order-Taker: Takes the order that is given via cli (probably a menu with
   menu-item numbers could be beneficial)
1. Dough-Maker: Prepares the dough and place it for the Sauce-Maker
1. Sauce-Maker: Prepares the sauce and place it on the dough
1. Topping-Placer: Takes the dough with sauce and places toppings on it
1. Oven-Placer: Takes the uncooked pizza and places it into the oven and takes
   it out if it is finished
1. Pizza-Seller: Takes the cash and gives out the finished pizza

## Tools

- cobra-cli (just for fun)
- workerpool with six workers
