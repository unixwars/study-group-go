
require 'open-uri'


happy = ['love' 'loved' 'like' 'liked' 'awesome' 'amazing' 'good' 'great' 'excellent']
sad = ['hate' 'hated' 'dislike' 'disliked' 'awful' 'terrible' 'bad' 'painful' 'worst']


def readPage(page)

#  eocMarker = "\ufeff"

  linearray = []

  open(page){|pageDoc|
    pageDoc.each_line{|line|
      i = line.index('33558957')
      if (i != nil and i > 0)
        puts line

        lineElements = line.split(',')

        commentsArray.push(lineElements[5])

#        strEnd = line.index(eocMarker)
#        if strEnd != nil
#          s = line.sub(29, strEnd)
#          puts s
#        end
      end
 #     if i == nil or i >= page.length
 #       break
 #     end
     }
  }


#  page = wpage.gets
#  while i = page.index('class=CT')
#    if i == nil or i >= page.length
#      break
#    end
#    puts i
 # while ('div class="CT">')
#    if line.include?('/"CT/"')
#      puts i
#    end
#  end



# stackoverflow http://stackoverflow.com/questions/1878891/how-to-load-a-web-page-and-search-for-a-word-in-ruby
# http://www.rubydoc.info/github/sparklemotion/nokogiri/Nokogiri/CSS

#  require 'nokogiri'
#  require 'open-uri'
#  doc = Nokogiri(open(page))
#  comments = doc.css('.CT').map do [l]
#    puts l
#  end


# read file now parse commentArray

# TODO

end #eof






def main ()
  puts 'Webpage scraper CH190'
  print 'Enter web page :'
#  page = gets.chomp

  page = 'https://plus.googleapis.com/u/0/_/widget/render/comments?first_party_property=YOUTUBE&href=https://www.youtube.com/watch?v=lMOQnJ-SLHI'

  if page.size < 6
    exit(0)
  end

  readPage(page)


end











main()