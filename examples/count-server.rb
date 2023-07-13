require 'sinatra'

counter = 0

get '/' do
  "#{counter += 1}\n"
end
