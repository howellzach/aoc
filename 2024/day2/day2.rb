#!/usr/bin/env ruby

levels_list = []

f = File.open('input.txt', 'r')
f.each_line do |line|
  s = line.split(' ')
  levels_list << s.map(&:to_i)
end
f.close

def check_safety(levels)
  output = {}
  dirs = []
diffs = []
  levels.each_with_index do |report, index|
    if levels.length != index + 1
      diff = report - levels[index + 1]
      if diff.abs > 3
        # return false, index
        output[index] = false
      end

      dir = ''
      if report > levels[index + 1]
        dir = 'inc'
      elsif report < levels[index + 1]
        dir = 'dec'
      elsif report == levels[index + 1]
        dir = 'same'
        # return false, index
        output[index] = false
      end
      dirs << dir
      diffs << diff.abs
    end
  end

  starting_dir = dirs[0]
  dirs.each_with_index do |d, i|
    if starting_dir != d
      # return false, i+1
      output[i] = false
    end
    output[i] = true if output[i] != false
  end

  output
end

level_safety_list = []
levels_list.each do |level|
  level_safety_list << check_safety(level)
end

res1 = 0
level_safety_list.each do |level|
  res1 += 1 if level.values.uniq.size == 1 && level.values[0] == true
end

puts "Result 1: #{res1}"

p level_safety_list
res2 = 0
level_safety_list.each_with_index do |level, index|
  if level.values.uniq.size == 1 && level.values[0] == true
    res2 += 1
  else
    level.each_with_index do |_l, i|
      p "starting over"
      level_copy = levels_list[index].clone
      p i
      p level_copy
      level_copy.delete_at(i)
      p level_copy
      check = check_safety(level_copy)
      p check
      if check.values.uniq.size == 1 && check.values[0] == true
        puts "=================== #{check} ======================"
        p "found safe"
        res2 += 1
        break
      end
    end
  end
end

puts "Result 2: #{res2}"
