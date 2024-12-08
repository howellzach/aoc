#!/usr/bin/env ruby

s = ''
f = File.open('input.txt', 'r')
f.each_line do |line|
  s << line.strip
end
f.close

matches = s.scan(/mul\((\d{1,3}),(\d{1,3})\)/)

res1 = 0
matches.each do |match|
  nums = match.map(&:to_i)
  res = nums[0] * nums[1]
  res1 += res
end

puts "Result 1: #{res1}"

p s
s = s.gsub(/don't\(\).*?do\(\)/, '')
puts "\n\n==========\n\n\n"
p s

matches = s.scan(/mul\((\d{1,3}),(\d{1,3})\)/)

res2 = 0
matches.each do |match|
  nums = match.map(&:to_i)
  res = nums[0] * nums[1]
  res2 += res
end

puts "Result 2: #{res2}"
