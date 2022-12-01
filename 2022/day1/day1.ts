import * as fs from "fs";

const data = fs.readFileSync('input.txt', 'utf-8');
const dataSplit: Array<String> = data.split('\n')

let elves = []
let elfCal = 0
dataSplit.forEach(function (v, i) {
    if (i == 0) {
        elfCal = Number(v)
    }
    if (v != "") {
        elfCal = elfCal + Number(v)
    } else {
        elves.push(elfCal)
        elfCal = 0
    }
})

let mostElf: number = elves[0]
for (let e of elves) {
    if (e > mostElf) {
        mostElf = e
    }
}
console.log("Result 1: %d", mostElf)

function getTopElvesCal(n: number) {
    const sortedElves = elves.sort((a, b) => b - a)
    let topElvesCal = 0
    for (let i = 0; i <= n-1; i++) {
        topElvesCal = topElvesCal + sortedElves[i]
    }
    return topElvesCal
}

const result2 = getTopElvesCal(3)
console.log("Result 2: %d", result2)