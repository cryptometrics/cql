require 'fileutils'

module Protomodel
  def self.msg
    "\n// * This is a generated file, do not edit\n"
  end

  def self.pkg
    'package protomodel'
  end

  def protomodel_comment
    "\n" + go_comment
  end

  def protomodel_struct
    literals = []
    fields.each do |field|
      literal = field.description.empty? ? '' : "\n// #{field.description}\n"
      literal += "#{field.go_protofield_name} #{field.go_type}"
      literal += "`json:\"#{field.identifier}\"`"
      literals << literal
    end
    literals = literals.sort.join("\n")
    "\ntype #{go_model_name} struct {#{literals}}"
  end

  def write_protomodel
    Dir.chdir(PARENT_DIR + '/protomodel') do
      File.open(go_model_filename, 'w') do |f|
        f.write(Protomodel.pkg)
        f.write(Protomodel.msg)
        f.write(protomodel_comment)
        f.write(protomodel_struct)
      end
      # In addition to fixing imports, goimports also formats your code in the
      # same style as gofmt so it can be used as a replacement for your editor's
      # gofmt-on-save hook.
      `/go/bin/goimports -w #{go_model_filename}`
    end
  end
end
