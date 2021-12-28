require 'fileutils'

module Protomodel
  def self.pkg
    'package protomodel'
  end

  def write_protomodel
    Dir.chdir(PARENT_DIR + '/protomodel') do
      File.open(go_model_filename, 'w') do |f|
        f.write(Protomodel.pkg)
      end
    end
  end
end
