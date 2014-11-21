

word_list = %w{ apple
                pair
                zephyr
                consul
                cat
                dog
                granada
                tree
                building
                window
                inhilipilification
               }

def getWord(level)
  @word_ptr = 0
  case level
    when 'l'
      low = 3
      high = 5
    when 'm'
      low =5
      high = 7
    when 'h'
      low = 7
      high = 99
  end

  while word_ptr < word_list.length do
    w = word_list[word_ptr]
    word_ptr += 1
    if w.size >= low and w.size <= high
      return w
    end
  end
end



puts 'Hangman game'
puts
print 'Number of allowed guesses: '
number_of_guesses = 6 #gets.chomp()
print 'Enter difficulty level (L/M/H): '
level = gets.chomp.downcase


while  do
  word = getWord(level) #word_list[word_ptr]
  if word == nil
    break
  end

  guess_str = word.length.times.collect {'-'}.join()
  guess_count = 0
  correct_guesses = 0

    while guess_count <= number_of_guesses and correct_guesses < word.length
      puts "Word to guess: #{guess_str}"
      print 'Enter guess: '
      c_guess = gets.chomp()
      guess_count += 1

      index = 0
      word.each_char do |c|
        if c == c_guess
          guess_str[index] = c
          correct_guesses += 1
        end
        index += 1
      end
    end

  if guess_count == number_of_guesses
    puts 'You are dead.'
    break
  end

  word_ptr += 1
  puts 'Try again? (y/n): '
  unless gets.chomp == 'y'
    break
  end
end