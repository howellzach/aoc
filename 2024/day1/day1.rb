#!/usr/bin/env ruby

l1 = []
l2 = []

f = File.open("input.txt", "r")
f.each_line do |line|
  s = line.split(" ")
  l1 << Integer(s[0])
  l2 << Integer(s[1])
end
f.close

l1s = l1.sort_by(&:to_i)
l2s = l2.sort_by(&:to_i)

diffs = []
l1s.each_with_index do |item, index|
  n = item - l2s[index]
  diffs << n.abs
end

puts "Result 1: #{diffs.sum}"

def get_frequency_mult(l1, l2)
  l = []
  l1.each do |item|
    l << item * l2.count(item)
  end
  return l.sum
end


res2 = get_frequency_mult(l1, l2)

puts "Result 2: #{res2}"

