import * as fs from "fs"

const data = fs.readFileSync('input.txt', 'utf-8')
const dataSplit: Array<string> = data.split('\n')

const stream = dataSplit[0]

function isUnique(str: string) {
  return new Set(str).size == str.length;
}

function commsSystem(data: string, chars: number): number {
    let count = 0
    let buffer = ""
    for (let i of data) {
        count++
        buffer += i
        if (buffer.length > chars) {
            buffer = buffer.substring(1)
        }
        if (buffer.length == chars) {
            if (isUnique(buffer)) {
                break
            }
        }
    }
    return count
}

let pt1Res = commsSystem(stream, 4)
console.log("First start-of-packet marker: %d", pt1Res)

let pt2Res = commsSystem(stream, 14)
console.log("First start-of-message marker: %d", pt2Res)