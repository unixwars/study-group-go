

class NamePairs
  attr_accessor :long
  attr_accessor :short
  @long = String.new
  @short = String.new

  def printNames
    print "#{@short} : #{@long} \n"
  end

  # sorting function
  def <=> (l)
    if l.short.length > @short.length
      return -1
    elsif l.short.length < @short.length
      return 1
    else
      return 0
    end
  end
end





def main()

  print 'Enter filename (default ./enable1.txt):'
  fname = gets.chomp
  if fname.length == 0
    fname = './enable1.txt'
  end
  puts

  begin
  names = Array.new
  open(fname) do |file|
    file.each do |line|

      fp = NamePairs.new
      fp.long = line.chomp
      s = String.new(fp.long)
      if s.start_with?('at')
        s[0..1] = '@'
        if s.include?('at')
          i = s.index('at')
          s[i..i+1] = '@'
        end
        fp.short = s
        names.push(fp)
      end
    end
  end

  rescue
    puts "Bad filename #{fname}"
    return
  end

  sortedNames = names.sort

  for j in 1..10
    sortedNames[-1 *j].printNames
  end

  10.times do |i|
    sortedNames[i].printNames
  end


end




main()