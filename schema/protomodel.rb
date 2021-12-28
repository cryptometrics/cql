require 'fileutils'

module Protomodel
  def self.msg
    "\n// * This is a generated file, do not edit"
  end

  def self.pkg
    'package protomodel'
  end

  def write_protomodel
    Dir.chdir(PARENT_DIR + '/protomodel') do
      File.open(go_model_filename, 'w') do |f|
        f.write(Protomodel.pkg)
        f.write(Protomodel.msg)
      end
      # In addition to fixing imports, goimports also formats your code in the
      # same style as gofmt so it can be used as a replacement for your editor's
      # gofmt-on-save hook.
      `/go/bin/goimports -w #{go_model_filename}`
    end
  end
end
