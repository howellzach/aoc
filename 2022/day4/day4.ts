import * as fs from "fs"

const data = fs.readFileSync('input.txt', 'utf-8')
const dataSplit: Array<string> = data.split('\n')

function range(start: number, end: number) {
    return [...Array(end - start + 1).keys()].map(x => x + start)
}

function createSections(data: Array<string>): Array<Array<Array<number>>> {
    let sections = []
    for (let i of data) {
        let x = i.split(",")[0].split("-")
        let y = i.split(",")[1].split("-")
        let range1 = range(Number(x[0]), Number(x[1]))
        let range2 = range(Number(y[0]), Number(y[1]))
        sections.push([range1, range2])
    }
    return sections
}

function sectionContainsEvery(s1: Array<number>, s2: Array<number>) {
    let check1: boolean = s2.every(v => s1.includes(v))
    let check2: boolean = s1.every(v => s2.includes(v))
    if (check1 || check2) {
        return true
    } else {
        return false
    }
}

function sectionContainsSome(s1: Array<number>, s2: Array<number>) {
    let check1: boolean = s2.some(v => s1.includes(v))
    let check2: boolean = s1.some(v => s2.includes(v))
    if (check1 || check2) {
        return true
    } else {
        return false
    }
}

let sections = createSections(dataSplit)
let pt1Res = 0
for (let i of sections) {
    if (sectionContainsEvery(i[0], i[1])) {
        pt1Res++
    }
}
console.log("Part 1 result: %s", pt1Res)

let pt2Res = 0
for (let i of sections) {
    if (sectionContainsSome(i[0], i[1])) {
        pt2Res++
    }
}
console.log("Part 2 result: %s", pt2Res)