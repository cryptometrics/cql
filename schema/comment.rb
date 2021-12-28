# frozen_string_literal: true

# Comment is a module for methods to maninipulate comments for a specific
# programming language
module Comment
  def split_comment_by_char(delimiter, comment, count)
    line =	delimiter
    split = []
    comment.split(' ').each do |word|
      if (line + word).length >= count
        split << (line += "\n").rstrip
        line = delimiter
      end
      line = line + ' ' + word
    end
    split << line.rstrip
  end

  def split_go_comment_by_char(comment, count = 80)
    split_comment_by_char('//', comment, count)
  end

  def split_graphql_comment_by_char(comment, count = 80)
    split_comment_by_char('#', comment, count)
  end

  def format_go_comment(comment, count = 80)
    split_go_comment_by_char(comment, count).join('')
  end

  def format_graphql_comment(comment, count = 80)
    split_graphql_comment_by_char(comment, count).join('')
  end
end
