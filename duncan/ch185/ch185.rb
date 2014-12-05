

class NamePairs
  attr_accessor :long
  attr_accessor :short
  @long = String.new
  @short = String.new

  def printNames
    print "#{@short} : #{@long} \n"
  end
end





def main()

  print 'Enter filename :'
  fname = './enable1.txt' #  gets.chomp
  puts

  names = Array.new
  open(fname) do |file|
    file.each do |line|
#      puts line

      fp = NamePairs.new
      fp.long = line.chomp
      s = String.new(fp.long)
      if s.start_with?('at')
        s[0..1] = '@'
        fp.short = s
        names.push(fp)
      end

    end
  end

  names.each do |n|
    n.printNames
  end

end

















main()