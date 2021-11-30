# Golang Quiz multiple choise

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

## Anatomy
Packages description:
- game
    - model contains the actors involved, the domain
    - business the business logic and starter game quiz
    - definetions func and struct to parse the yaml file
- data
    - persistence layer and entities for DTO
-service
    - business logic between persistence and presentation layer
-controller
    - to handle REST call and dispatch to the proper service
-cmd
    - cobra definitions 

## Build and Run
```
 go build
./golang-quiz init

```
- it will drop and create tables from scratch

## Server port and usage
- HomePage http://localhost:10000/
- Port 10000

## REST descriptions
- Add Player : http://localhost:10000/addPlayer/?username=Zena POST
- Get Player : http://localhost:10000/player/1 GET
- Start Game : http://localhost:10000/play/ GET
    - questions are show on terminal where did you build
    - to answer, just digit it on terminal
- Score Player : http://localhost:10000/score/1 GET
