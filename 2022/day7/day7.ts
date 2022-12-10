import * as fs from "fs"

const data = fs.readFileSync('input.txt', 'utf-8')
const dataSplit: string[] = data.split('\n')

let filesystem = {
    "/": {
        "type": "dir",
        "size": 0,
        "children": {}
    }
}

let current_path = []

function addToFilesystem(line: string[]) {
    let item: object
    let isFile: boolean
    let name = line[1]
    if (line[0] == "dir") {
        isFile = false
        item = {
            "type": "dir",
            "size": 0,
            "children": {}
        }
    } else {
        isFile = true
        item = {
            "type": "file",
            "size": Number(line[0]),
        }
    }

    let fsPath = ""
    for (let i of current_path) {
        fsPath += `.${i}.children`
    }

    let fsObj = objByString(filesystem, fsPath)
    fsObj[name] = item

    if (isFile) {
        let parentPath = fsPath.replace(/(.children$)/, "")
        let parObj = objByString(filesystem, parentPath)
        parObj["size"] += item["size"]
        while (parentPath != "./") {
            parentPath = parentPath.replace(/(.children.[A-Za-z]*$)/, "")
            parObj = objByString(filesystem, parentPath)
            parObj["size"] += item["size"]
        }
    }
}

function objByString(o: object, s: string) {
    s = s.replace(/\[(\w+)\]/g, '.$1')
    s = s.replace(/^\./, '')
    let a = s.split('.')
    for (let i in a) {
        let k = a[i]
        if (k in o) {
            o = o[k]
        } else {
            return
        }
    }
    return o
}

function checkSizes(files: object, target: number, destArr: number[]) {
    for (let i in files) {
        if (files[i]["type"] == "dir") {
            // console.log(files[i])
            if (files[i]["size"] <= target) {
                destArr.push(files[i]["size"])
            }
            checkSizes(files[i]["children"], target, destArr)
        }
    }
}

function getSmallestSolution(arr: number[], target: number) {
    arr.sort(function (a, b) { return a - b; })
    for (let i of arr) {
        if (i > target) {
            return i
        }
    }
}

for (let i of dataSplit) {
    let line = i.split(" ")
    if (line[0] == "$") {
        if (line[1] == "cd" && line[2] != "..") {
            current_path.push(line[2])
        }
        if (line[1] == "cd" && line[2] == "..") {
            current_path.pop()
        }
    } else {
        addToFilesystem(line)
    }
}

let pt1ResArr = []
checkSizes(filesystem, 100000, pt1ResArr)
let pt1Res = pt1ResArr.reduce((partialSum, a) => partialSum + a, 0)
console.log("Result 1: %s", pt1Res)

let unusedSpace = 70000000 - filesystem["/"]["size"]
let targetSpace = 30000000 - unusedSpace
let pt2ResArr = []
checkSizes(filesystem, 70000000, pt2ResArr)
let pt2Res = getSmallestSolution(pt2ResArr, targetSpace)
console.log("Result 2: %s", pt2Res)