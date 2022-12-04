import * as fs from "fs"

const data = fs.readFileSync('input.txt', 'utf-8')
const dataSplit: Array<string> = data.split('\n')

function makeRucks(data: Array<String>): Array<Object> {
    let rucksacks = []
    for (let i of data) {
        rucksacks.push({
            "bothCom": i,
            "com1": i.slice(0, (i.length / 2)),
            "com2": i.slice(-(i.length / 2)),
        })
    }
    for (let i of rucksacks) {
        for (let j of i["com1"]) {
            for (let k of i["com2"]) {
                if (j == k) {
                    i["repeat"] = j
                }
            }
        }
    }
    for (let i of rucksacks) {
        i["repPriority"] = getPriority(i["repeat"])
    }
    return rucksacks
}

function getPriority(item: string): Number {
    const priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    return priority.indexOf(item) + 1
}

function groupElves(rucksacks: Array<Object>, groupSize: number) {
    let count = 1
    let outerGroup = []
    let innerGroup = []
    for (let i of rucksacks) {
        if (count % groupSize != 0) {
            innerGroup.push(i)
            count++
        } else {
            innerGroup.push(i)
            outerGroup.push(innerGroup)
            innerGroup = []
            count++
        }
    }
    return outerGroup
}

function findBadges(group: Array<Array<Object>>) {
    for (let outer of group) {
        let elf1 = outer[0]
        let elf2 = outer[1]
        let elf3 = outer[2]
        for (let i of elf1['bothCom']) {
            for (let j of elf2['bothCom']) {
                for (let k of elf3['bothCom']) {
                    if (i == k && i == j) {
                        elf1['badge'] = i
                        elf2['badge'] = i
                        elf3['badge'] = i
                        elf1['badgePriority'] = getPriority(i)
                    }
                }
            }
        }
    }
}


let rucksacks = makeRucks(dataSplit)
let pt1Res = 0
rucksacks.forEach(i => { pt1Res = pt1Res + i["repPriority"] })
console.log("Part 1 result: %s", pt1Res)

let groupedElves = groupElves(rucksacks, 3)
findBadges(groupedElves)
let pt2Res = 0
groupedElves.forEach(i => { pt2Res = pt2Res + i[0]['badgePriority'] })
console.log("Part 2 result: %s", pt2Res)