import * as fs from "fs"

const data = fs.readFileSync('input.txt', 'utf-8')
const dataSplit: Array<string> = data.split('\n')

const legend = {
    "A": "Rock",
    "B": "Paper",
    "C": "Scissors",
    "X": "Lose",
    "Y": "Draw",
    "Z": "Win"
}

const scoring = {
    "Rock": 1,
    "Paper": 2,
    "Scissors": 3,
    "Win": 6,
    "Tie": 3
}

function pt1Convert(input: Array<String>) : Array<String> {
    let outArray = []
    var o: String
    for (let i of input) {
        const [op, me] = i.split(' ')
        if (me == "X") {
            o = "A"
        } else if (me == "Y") {
            o = "B"
        } else if (me == "Z") {
            o = "C"
        }
        outArray.push(op + " " + o)
    }
    return outArray
}

function calcScores(input: Array<String>) : Array<Number> {
    let myScore = 0
    let opScore = 0
    for (let i of input) {
        const [op, me] = i.split(' ')
        let opChoice = legend[op]
        let myChoice = legend[me]
        opScore = opScore + scoring[opChoice]
        myScore = myScore + scoring[myChoice]

        if (myChoice == "Rock" && opChoice == "Scissors") {
            myScore = myScore + scoring["Win"]
        } else if (myChoice == "Paper" && opChoice == "Rock") {
            myScore = myScore + scoring["Win"]
        } else if (myChoice == "Scissors" && opChoice == "Paper") {
            myScore = myScore + scoring["Win"]
        } else if (myChoice == opChoice) {
            myScore = myScore + scoring["Tie"]
            opScore = opScore + scoring["Tie"]
        } else {
            opScore = opScore + scoring["Win"]
        }
    }
    return [myScore, opScore]
}

function decryptStrat(input: Array<String>) : Array<String> {
    let outArray = []
    var o: String
    for (let i of input) {
        const [op, outcome] = i.split(' ')
        let opChoice = legend[op]
        let outcomeGoal = legend[outcome]
        if (outcomeGoal == "Lose") {
            if (opChoice == "Rock") {
                o = "C"
            } else if (opChoice == "Paper") {
                o = "A"
            } else if (opChoice == "Scissors") {
                o = "B"
            }
        } else if (outcomeGoal == "Win") {
            if (opChoice == "Rock") {
                o = "B"
            } else if (opChoice == "Paper") {
                o = "C"
            } else if (opChoice == "Scissors") {
                o = "A"
            }
        } else {
            o = op
        }
        outArray.push(op + " " + o)

    }
    return outArray
}

const pt1ConvertData = pt1Convert(dataSplit)
const [myScore1, opScore1] = calcScores(pt1ConvertData)
console.log("Result 1: %d", myScore1)
console.log("Opponent score: %d\n", opScore1)

const newData = decryptStrat(dataSplit)
const [myScore2, opScore2] = calcScores(newData)
console.log("Result 2: %s", myScore2)
console.log("Opponent score: %s", opScore2)