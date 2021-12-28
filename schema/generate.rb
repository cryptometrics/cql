#!/usr/local/bin/ruby -w
# frozen_string_literal: true

require_relative 'endpoint_state'
require_relative 'scheme'
require_relative 'protomodel'

PARENT_DIR = File.expand_path('..', Dir.pwd)

def generate_models
  Dir.glob("#{File.dirname(__FILE__)}/model/*.json").map do |filename|
    next if filename.include?('schema.json')

    scheme = Scheme.new(filename)
    scheme.write_protomodel
  end
end

generate_models
