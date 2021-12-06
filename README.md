# Golang Quiz multiple choice

Objectives
- REST API app
- Cobra CLI Framework
- Local database sqlite-database.db

## Target
- User should be able to get questions with a number of answers
- User should be able to select just one answer per question.
- User should be able to answer all the questions and then post his/hers answers and get back how many correct answers they had, displayed to the user.
- User should see how good he/she did compare to others that have taken the quiz, "You were better than 60% of all quizzers"

## Questions file and timeout settings
```go
fileName := flag.String("file", "problems.yaml", "The name of the file with the problems")
limit := flag.Int("limit", 100, "The time limit for the quiz in seconds")
```

## Layers
Packages description:
- game
    - model contains the actors involved, the domain
    - business the business logic and starter game quiz
    - definetions func and struct to parse the yaml file

- data
    - persistence layer and entities for DTO

- service
    - business logic between persistence and presentation layer

- controller
    - to handle REST call and dispatch to the proper service

- cmd
    - cobra definitions 

## Build and Run
```
 go build
./golang-quiz init

```
- it will drop and create tables from scratch

## Server port and usage
- HomePage http://localhost:3000/
- Port 3000

## REST descriptions
- Add Player : http://localhost:3000/addPlayer/?username=Zena POST
- Get Player : http://localhost:3000/player/{playerId} GET
- Start Shell Game : http://localhost:3000/play/ GET
- Score Player : http://localhost:3000/score/{playerId} GET
- Subscribe Game Quiz  : http://localhost:3000/subscribe/{playerId} GET
- Get Problems Player : http://localhost:3000/problems/{playerId} GET
- Send Answer : http://localhost:3000/sendAnswer/{playerId}?answer=your_answer POST
- Players : http://localhost:3000/players/ GET
- Problems : http://localhost:3000/problems/ GET

## Play Usage
- Add a numbers of Players using : http://localhost:3000/addPlayer/?username=Zena POST
    - 3 Players already exist (you can skip)
- Start Game : At the startup of Application, the system it will ask if you want to play or exit, run the server or rebuild DB
- Play the Quiz 
    - Choice a player (just select the number) 
    - Answer the question digiting 
    - Enjoy! (I hope)
- Check your Rank using #Get Player
- Check your Score using #Score Player
