

StartLine = '*** START OF THIS PROJECT GUTENBERG EBOOK BIRDS AND ALL NATURE, JULY 1898 ***'

def main()

  fname = String.new
  begin
    print 'Enter filename (default ./pg47498.txt):'
#    fname = gets.chomp
#    if fname.length == 0
      fname = './pg47498.txt'
#    end
  rescue Exception => e
    puts "Exception, "+e
  end
  puts

  onforward = true

  begin
    wordHash = Hash.new
    open(fname) do |file|
      file.each do |line|

        if onforward
          if line.include?(StartLine)
            onforward = false
          end
          next
        end
        wordsInLine = line.chomp.split(' ')
        wordsInLine.each do |w|
          # to lowercase and remove non alpha chars
          lcw = w.downcase.gsub(/[^a-z]/i, '')
          if lcw.length == 0
            next
          end
           if wordHash.has_key?(lcw)
            # exists in hash so just increment count
            wordHash[lcw] = wordHash[lcw] + 1
          else
            # new word
            wordHash.store(lcw, 1)
          end
        end
      end
    end
  rescue Exception => e
    puts "A problem occurred opening the file #{fname} " + e
    return
  end

#  print wordHash.to_s
  sortedWords = wordHash.sort_by {|word, count| count}

  puts "The sorted list of words, #{wordHash.size} in total"
  sortedWords.each { |wc| puts wc }

  puts
  puts

  puts 'Top Ten most frequent'
  for j in 1..10
    puts sortedWords[-1*j].to_s
  end

  puts
  puts

  puts 'Top 10 least frequent'
  10.times do |i|
    puts sortedWords[i].to_s
  end
end


main()

#eof