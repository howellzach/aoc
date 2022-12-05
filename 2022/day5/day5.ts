import * as fs from "fs"

const data = fs.readFileSync('input.txt', 'utf-8')
const dataSplit: Array<string> = data.split('\n')

let procedureInput: string[] = []
let crateInput: string[] = []
let instructions = false
let stacks = true
for (let i of dataSplit) {
    if (i == "") {
        instructions = true
        stacks = false
    }
    if (stacks) {
        crateInput.push(i)
    }
    if (instructions) {
        procedureInput.push(i)
    }
}

procedureInput.shift()
let procedure: string[][] = []
for (let i of procedureInput) {
    let p = i.split(" ")
    procedure.push(p)
}

function createCargos(crateInput: string[]): [Object, Object] {
    let pt1Cargo = {}
    for (let i of crateInput.slice(-1)[0]) {
        if (i != " ") {
            pt1Cargo[i] = []
        }
    }

    let maxStorage = Object.keys(pt1Cargo).length
    let idx = 1
    for (let stack = 0; stack < maxStorage; stack++) {
        for (let i of crateInput.slice(0, -1)) {
            let crate = i[idx]
            if (crate != " ") {
                pt1Cargo[stack + 1].push(crate)
            }
        }
        idx = idx + 4
    }

    let pt2Cargo = structuredClone(pt1Cargo)

    return [pt1Cargo, pt2Cargo]
}

function moveItemsOnebyOne(cargo: Object, procedure: string[][]) {
    for (let i of procedure) {
        let q = Number(i[1])
        let sourceStack = i[3]
        let targetStack = i[5]
        while (q > 0) {
            cargo[targetStack].unshift(cargo[sourceStack][0])
            cargo[sourceStack].shift()
            q--
        }
    }
}

function moveItemsAsGroups(cargo: Object, procedure: string[][]) {
    for (let i of procedure) {
        let q = Number(i[1])
        let sourceStack = i[3]
        let targetStack = i[5]
        cargo[targetStack] = [...cargo[sourceStack].slice(0, q), ...cargo[targetStack]]
        cargo[sourceStack].splice(0, q)
    }
}

function main() {
    let [pt1Cargo, pt2Cargo] = createCargos(crateInput)

    moveItemsOnebyOne(pt1Cargo, procedure)
    let pt1Res = ""
    for (let i in pt1Cargo) {
        pt1Res += pt1Cargo[i][0]
    }
    console.log("Part 1 Result: %s", pt1Res)

    moveItemsAsGroups(pt2Cargo, procedure)
    let pt2Res = ""
    for (let i in pt2Cargo) {
        pt2Res += pt2Cargo[i][0]
    }
    console.log("Part 2 Result: %s", pt2Res)
}

main()
